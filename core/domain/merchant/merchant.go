/**
 * Copyright 2014 @ 56x.net.
 * name :
 * author : jarryliu
 * date : 2013-12-12 16:55
 * description :
 * history :
 */

package merchant

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ixre/go2o/core/domain/interface/domain/enum"
	"github.com/ixre/go2o/core/domain/interface/item"
	"github.com/ixre/go2o/core/domain/interface/member"
	"github.com/ixre/go2o/core/domain/interface/merchant"
	"github.com/ixre/go2o/core/domain/interface/merchant/shop"
	"github.com/ixre/go2o/core/domain/interface/merchant/user"
	"github.com/ixre/go2o/core/domain/interface/merchant/wholesaler"
	"github.com/ixre/go2o/core/domain/interface/registry"
	"github.com/ixre/go2o/core/domain/interface/valueobject"
	"github.com/ixre/go2o/core/domain/interface/wallet"
	si "github.com/ixre/go2o/core/domain/merchant/shop"
	userImpl "github.com/ixre/go2o/core/domain/merchant/user"
	wsImpl "github.com/ixre/go2o/core/domain/merchant/wholesale"
	"github.com/ixre/go2o/core/infrastructure"
	"github.com/ixre/go2o/core/infrastructure/domain"
	"github.com/ixre/go2o/core/infrastructure/domain/util"
	"github.com/ixre/go2o/core/variable"
)

var _ merchant.IMerchantManager = new(merchantManagerImpl)

type merchantManagerImpl struct {
	rep     merchant.IMerchantRepo
	valRepo valueobject.IValueRepo
}

func NewMerchantManager(rep merchant.IMerchantRepo,
	valRepo valueobject.IValueRepo) merchant.IMerchantManager {
	return &merchantManagerImpl{
		rep:     rep,
		valRepo: valRepo,
	}
}

// CreateSignUpToken 创建会员申请商户密钥
func (m *merchantManagerImpl) CreateSignUpToken(memberId int64) string {
	return m.rep.CreateSignUpToken(memberId)
}

// GetMemberFromSignUpToken 根据商户申请密钥获取会员编号
func (m *merchantManagerImpl) GetMemberFromSignUpToken(token string) int64 {
	return m.rep.GetMemberFromSignUpToken(token)
}

// RemoveSignUp 删除会员的商户申请资料
func (m *merchantManagerImpl) RemoveSignUp(memberId int) error {
	return m.rep.DeleteMerchantSignUpByMemberId(memberId)
}

// 检查商户注册信息是否正确
func (m *merchantManagerImpl) checkSignUpInfo(v *merchant.MchSignUp) error {
	if v.MemberId <= 0 {
		return errors.New("会员不存在")
	}
	//todo: validate and check merchant name exists
	if v.MchName == "" {
		return merchant.ErrMissingMerchantName
	}
	if v.CompanyName == "" {
		return merchant.ErrMissingCompanyName
	}
	if v.CompanyNo == "" {
		return merchant.ErrMissingCompanyNo
	}
	if v.Address == "" {
		return merchant.ErrMissingAddress
	}
	if v.PersonName == "" {
		return merchant.ErrMissingPersonName
	}
	if v.PersonId == "" {
		return merchant.ErrMissingPersonId
	}
	if util.CheckChineseCardID(v.PersonId) != nil {
		return merchant.ErrPersonCardId
	}
	if v.Phone == "" {
		return merchant.ErrMissingPhone
	}
	if v.CompanyImage == "" {
		return merchant.ErrMissingCompanyImage
	}
	if v.PersonImage == "" {
		return merchant.ErrMissingPersonImage
	}
	return nil
}

// CommitSignUpInfo 提交商户注册信息
func (m *merchantManagerImpl) CommitSignUpInfo(v *merchant.MchSignUp) (int, error) {
	err := m.checkSignUpInfo(v)
	if err != nil {
		return 0, err
	}
	v.Reviewed = enum.ReviewAwaiting
	v.SubmitTime = time.Now().Unix()
	v.UpdateTime = time.Now().Unix()
	return m.rep.SaveSignUpInfo(v)

}

// ReviewMchSignUp 审核商户注册信息
func (m *merchantManagerImpl) ReviewMchSignUp(id int, pass bool, remark string) error {
	var err error
	v := m.GetSignUpInfo(id)
	if v == nil {
		return merchant.ErrNoSuchSignUpInfo
	}
	if pass {
		v.Reviewed = enum.ReviewPass
		v.Remark = ""
		if err = m.createNewMerchant(v); err != nil {
			return err
		}
	} else {
		v.Reviewed = enum.ReviewReject
		v.Remark = remark
		if strings.TrimSpace(v.Remark) == "" {
			return merchant.ErrRequireRejectRemark
		}
	}
	v.UpdateTime = time.Now().Unix()
	_, err = m.rep.SaveSignUpInfo(v)
	return err
}

// 创建新商户
func (m *merchantManagerImpl) createNewMerchant(v *merchant.MchSignUp) error {
	unix := time.Now().Unix()
	mchVal := &merchant.Merchant{
		MemberId: v.MemberId,
		// 商户名称
		Name: v.MchName,
		// 是否自营
		SelfSales: 0,
		// 商户等级
		Level: 1,
		// 标志
		Logo: "",
		// 公司名称
		CompanyName: "",
		// 省
		Province: int(v.Province),
		// 市
		City: int(v.City),
		// 区
		District: int(v.District),
		// 是否启用
		Enabled: 1,
		Flag:    1,
		// 过期时间
		ExpiresTime: time.Now().Add(time.Hour * time.Duration(24*365)).Unix(),
		// 注册时间
		CreateTime: unix,
		// 更新时间
		UpdateTime: unix,
		// 登录时间
		LoginTime: 0,
		// 最后登录时间
		LastLoginTime: 0,
	}
	mch := m.rep.CreateMerchant(mchVal)
	err := mch.SetValue(mchVal)
	if err != nil {
		return err
	}
	mchId, err := mch.Save()
	if err == nil {
		names := m.valRepo.GetAreaNames([]int32{v.Province, v.City, v.District})
		location := strings.Join(names, "")
		ev := &merchant.EnterpriseInfo{
			MchId:        mchId,
			CompanyName:  v.CompanyName,
			CompanyNo:    v.CompanyNo,
			PersonName:   v.PersonName,
			PersonIdNo:   v.PersonId,
			PersonImage:  v.PersonImage,
			Tel:          v.Phone,
			Province:     v.Province,
			City:         v.City,
			District:     v.District,
			Location:     location,
			Address:      v.Address,
			CompanyImage: v.CompanyImage,
			AuthDoc:      v.AuthDoc,
			Reviewed:     v.Reviewed,
			ReviewTime:   unix,
			ReviewRemark: "",
			UpdateTime:   unix,
		}
		_, err = mch.ProfileManager().SaveEnterpriseInfo(ev)
		if err == nil {
			mch.ProfileManager().ReviewEnterpriseInfo(true, "")
		}
	}
	return err
}

// GetSignUpInfo 获取商户申请信息
func (m *merchantManagerImpl) GetSignUpInfo(id int) *merchant.MchSignUp {
	return m.rep.GetMerchantSignUpInfo(id)
}

// GetSignUpInfoByMemberId 获取会员申请的商户信息
func (m *merchantManagerImpl) GetSignUpInfoByMemberId(memberId int) *merchant.MchSignUp {
	return m.rep.GetMerchantSignUpByMemberId(memberId)
}

// GetMerchantByMemberId 获取会员关联的商户
func (m *merchantManagerImpl) GetMerchantByMemberId(memberId int) merchant.IMerchant {
	return m.rep.GetMerchantByMemberId(memberId)
}

var _ merchant.IMerchant = new(merchantImpl)

type merchantImpl struct {
	_value           *merchant.Merchant
	_account         merchant.IAccount
	_wholesaler      wholesaler.IWholesaler
	_host            string
	_repo            merchant.IMerchantRepo
	_wsRepo          wholesaler.IWholesaleRepo
	_itemRepo        item.IItemRepo
	_shopRepo        shop.IShopRepo
	_userRepo        user.IUserRepo
	_valRepo         valueobject.IValueRepo
	_memberRepo      member.IMemberRepo
	_userManager     user.IUserManager
	_confManager     merchant.IConfManager
	_saleManager     merchant.ISaleManager
	_levelManager    merchant.ILevelManager
	_kvManager       merchant.IKvManager
	_memberKvManager merchant.IKvManager
	//_mssManager      mss.IMssManager
	//_mssRepo          mss.IMssRepo
	_profileManager   merchant.IProfileManager
	_apiManager       merchant.IApiManager
	_shopManager      shop.IShopManager
	_walletRepo       wallet.IWalletRepo
	_registryRepo     registry.IRegistryRepo
	_lastBindMemberId int64 //  之前绑定的会员编号
}

func NewMerchant(v *merchant.Merchant, rep merchant.IMerchantRepo,
	wsRepo wholesaler.IWholesaleRepo, itemRepo item.IItemRepo,
	shopRepo shop.IShopRepo, userRepo user.IUserRepo, memberRepo member.IMemberRepo,
	walletRepo wallet.IWalletRepo, valRepo valueobject.IValueRepo, registryRepo registry.IRegistryRepo) merchant.IMerchant {
	mch := &merchantImpl{
		_value:        v,
		_repo:         rep,
		_wsRepo:       wsRepo,
		_itemRepo:     itemRepo,
		_shopRepo:     shopRepo,
		_userRepo:     userRepo,
		_valRepo:      valRepo,
		_memberRepo:   memberRepo,
		_walletRepo:   walletRepo,
		_registryRepo: registryRepo,
	}
	return mch
}

func (m *merchantImpl) GetRepo() merchant.IMerchantRepo {
	return m._repo
}

func (m *merchantImpl) GetAggregateRootId() int64 {
	return m._value.Id
}

// Complex 获取符合的商家信息
func (m *merchantImpl) Complex() *merchant.ComplexMerchant {
	src := m.GetValue()
	return &merchant.ComplexMerchant{
		Id:            int32(src.Id),
		MemberId:      src.MemberId,
		Usr:           src.LoginUser,
		Pwd:           src.LoginPwd,
		Name:          src.Name,
		SelfSales:     int32(src.SelfSales),
		Level:         int32(src.Level),
		Logo:          src.Logo,
		CompanyName:   src.CompanyName,
		Province:      int32(src.Province),
		City:          int32(src.City),
		District:      int32(src.District),
		Enabled:       int32(src.Enabled),
		ExpiresTime:   src.ExpiresTime,
		JoinTime:      src.CreateTime,
		UpdateTime:    src.UpdateTime,
		LoginTime:     src.LoginTime,
		LastLoginTime: src.LastLoginTime,
	}
}

func (m *merchantImpl) GetValue() merchant.Merchant {
	return *m._value
}

func (m *merchantImpl) SetValue(v *merchant.Merchant) error {
	tv := m._value
	if v.Id == tv.Id {
		tv.Name = v.Name
		tv.Province = v.Province
		tv.City = v.City
		tv.District = v.District
		if v.LastLoginTime > 0 {
			tv.LastLoginTime = v.LastLoginTime
		}
		if v.LoginTime > 0 {
			tv.LoginTime = v.LoginTime
		}

		if len(v.Logo) != 0 {
			tv.Logo = v.Logo
		}
		if len(v.CompanyName) != 0 {
			tv.CompanyName = v.CompanyName
		}
	}
	if v.LoginPwd != "" {
		tv.LoginPwd = v.LoginPwd
	}
	if m.GetAggregateRootId() <= 0 {
		m._value.MemberId = v.MemberId
	}
	tv.SelfSales = v.SelfSales
	tv.ExpiresTime = v.ExpiresTime
	tv.UpdateTime = time.Now().Unix()
	return nil
}

func (m *merchantImpl) BindMember(memberId int) error {
	if m._value.MemberId == int64(memberId) {
		return merchant.ErrMemberBindExists
	}
	exist := m._repo.CheckMemberBind(int64(memberId), m.GetAggregateRootId())
	if exist {
		return merchant.ErrBindAnotherMerchant
	}
	m._lastBindMemberId = m._value.MemberId
	m._value.MemberId = int64(memberId)
	if m.GetAggregateRootId() > 0 {
		err := m.applyBindMember()
		if err == nil {
			_, err = m.Save()
		}
		return err
	}
	return nil
}

func (m *merchantImpl) applyBindMember() error {
	// 解绑
	if m._lastBindMemberId > 0 {
		origin := m._memberRepo.GetMember(m._lastBindMemberId)
		if origin != nil {
			_ = origin.GrantFlag(-member.FlagSeller)
		}
	}
	// 添加商户标志
	im := m._memberRepo.GetMember(m._value.MemberId)
	if im == nil {
		return member.ErrNoSuchMember
	}
	err := im.GrantFlag(member.FlagSeller)
	if err == nil {
		m._lastBindMemberId = m._value.MemberId
	}
	return err
}

// Save 保存
func (m *merchantImpl) Save() (int64, error) {
	id := m.GetAggregateRootId()
	if id > 0 {
		id, err := m._repo.SaveMerchant(m._value)
		return int64(id), err
	}
	return m.createMerchant()
}

// SelfSales 是否自营
func (m *merchantImpl) SelfSales() bool {
	return m._value.SelfSales == 1 ||
		m.GetAggregateRootId() == 1
}

// Stat 获取商户的状态,判断 过期时间、判断是否停用。
// 过期时间通常按: 试合作期,即1个月, 后面每年延长一次。或者会员付费开通。
func (m *merchantImpl) Stat() error {
	if m._value == nil {
		return merchant.ErrNoSuchMerchant
	}
	if m._value.Enabled == 0 {
		return merchant.ErrMerchantDisabled
	}
	if m._value.ExpiresTime > 0 && m._value.ExpiresTime < time.Now().Unix() {
		return merchant.ErrMerchantExpires
	}
	return nil
}

// SetEnabled 设置商户启用或停用
func (m *merchantImpl) SetEnabled(enabled bool) error {
	if enabled {
		m._value.Enabled = 1
	} else {
		m._value.Enabled = 0
	}
	_, err := m.Save()
	return err
}

// Member 返回对应的会员编号
func (m *merchantImpl) Member() int64 {
	return m.GetValue().MemberId
}

// Account 获取商户账户
func (m *merchantImpl) Account() merchant.IAccount {
	if m._account == nil {
		v := m._repo.GetAccount(int(m.GetAggregateRootId()))
		m._account = newAccountImpl(m, v, m._memberRepo, m._walletRepo)
	}
	return m._account
}

// Wholesaler 获取批发商实例
func (m *merchantImpl) Wholesaler() wholesaler.IWholesaler {
	if m._wholesaler == nil {
		mchId := m.GetAggregateRootId()
		v := m._wsRepo.GetWsWholesaler(mchId)
		if v == nil {
			v, _ = m.createWholesaler()
		}
		m._wholesaler = wsImpl.NewWholesaler(mchId, v, m._wsRepo, m._itemRepo)
	}
	return m._wholesaler
}

// EnableWholesale 启用批发
func (m *merchantImpl) EnableWholesale() error {
	if m.Wholesaler() != nil {
		return errors.New("wholesale for merchant enabled!")
	}
	_, err := m.createWholesaler()
	return err
}

func (m *merchantImpl) createWholesaler() (*wholesaler.WsWholesaler, error) {
	v := &wholesaler.WsWholesaler{
		MchId:        m.GetAggregateRootId(),
		Rate:         1,
		ReviewStatus: enum.ReviewPass,
		//ReviewStatus: enum.ReviewAwaiting,
	}
	_, err := m._wsRepo.SaveWsWholesaler(v, true)
	return v, err
}

// 创建商户
func (m *merchantImpl) createMerchant() (int64, error) {
	unix := time.Now().Unix()
	m._value.ExpiresTime = unix + 3600*24*365
	m._value.UpdateTime = unix
	m._value.CreateTime = unix
	m._value.Enabled = 1
	m._value.Flag = 1
	m._value.LoginTime = 0
	m._value.LastLoginTime = 0
	if m._value.MemberId == 0 {
		if len(m._value.LoginUser) == 0 {
			return 0, merchant.ErrMissingMerchantUser
		}
		if m._repo.CheckUserExists(m._value.LoginUser, 0) {
			return 0, merchant.ErrMerchantUserExists
		}
	}

	id, err := m._repo.SaveMerchant(m._value)
	if err != nil {
		return int64(id), err
	}
	m._value.Id = int64(id)
	// 绑定会员
	if m._value.MemberId > 0 {
		err = m.applyBindMember()
	}
	// 创建API
	api := &merchant.ApiInfo{
		ApiId:     domain.NewApiId(int(id)),
		ApiSecret: domain.NewSecret(int(id)),
		WhiteList: "*",
		Enabled:   1,
	}
	err = m.ApiManager().SaveApiInfo(api)
	return int64(id), err
}

// GetMajorHost 获取商户的域名
func (m *merchantImpl) GetMajorHost() string {
	if len(m._host) == 0 {
		host := m._repo.GetMerchantMajorHost(int(m.GetAggregateRootId()))
		if len(host) == 0 {
			host = fmt.Sprintf("%s.%s", m._value.LoginUser, infrastructure.GetApp().
				Config().GetString(variable.ServerDomain))
		}
		m._host = host
	}
	return m._host
}

// UserManager 返回用户服务
func (m *merchantImpl) UserManager() user.IUserManager {
	if m._userManager == nil {
		m._userManager = userImpl.NewUserManager(m.GetAggregateRootId(), m._userRepo)
	}
	return m._userManager
}

// LevelManager 获取会员管理服务
func (m *merchantImpl) LevelManager() merchant.ILevelManager {
	if m._levelManager == nil {
		m._levelManager = NewLevelManager(int64(m.GetAggregateRootId()), m._repo)
	}
	return m._levelManager
}

// KvManager 获取键值管理器
func (m *merchantImpl) KvManager() merchant.IKvManager {
	if m._kvManager == nil {
		m._kvManager = newKvManager(m, "kvset")
	}
	return m._kvManager
}

// MemberKvManager 获取用户键值管理器
func (m *merchantImpl) MemberKvManager() merchant.IKvManager {
	if m._memberKvManager == nil {
		m._memberKvManager = newKvManager(m, "kvset_member")
	}
	return m._memberKvManager
}

// 消息系统管理器
//func (m *MerchantImpl) MssManager() mss.IMssManager {
//	if m._mssManager == nil {
//		m._mssManager = mssImpl.NewMssManager(m, m._mssRepo, m._repo)
//	}
//	return m._mssManager
//}

// ConfManager 返回设置服务
func (m *merchantImpl) ConfManager() merchant.IConfManager {
	if m._confManager == nil {
		m._confManager = newConfigManagerImpl(int(m.GetAggregateRootId()),
			m._repo, m._memberRepo, m._valRepo)
	}
	return m._confManager
}

// SaleManager 销售服务
func (m *merchantImpl) SaleManager() merchant.ISaleManager {
	if m._saleManager == nil {
		m._saleManager = newSaleManagerImpl(int(m.GetAggregateRootId()), m)
	}
	return m._saleManager
}

// ProfileManager 企业资料管理器
func (m *merchantImpl) ProfileManager() merchant.IProfileManager {
	if m._profileManager == nil {
		m._profileManager = newProfileManager(m, m._valRepo)
	}
	return m._profileManager
}

// ApiManager API服务
func (m *merchantImpl) ApiManager() merchant.IApiManager {
	if m._apiManager == nil {
		m._apiManager = newApiManagerImpl(m)
	}
	return m._apiManager
}

// ShopManager 商店服务
func (m *merchantImpl) ShopManager() shop.IShopManager {
	if m._shopManager == nil {
		m._shopManager = si.NewShopManagerImpl(m, m._shopRepo, m._valRepo, m._registryRepo)
	}
	return m._shopManager
}
