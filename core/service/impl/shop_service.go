/**
 * Copyright 2015 @ 56x.net.
 * name : content_service
 * author : jarryliu
 * date : -- :
 * description :
 * history :
 */
package impl

import (
	"context"
	"regexp"

	"github.com/ixre/go2o/core/domain/interface/merchant"
	"github.com/ixre/go2o/core/domain/interface/merchant/shop"
	"github.com/ixre/go2o/core/domain/interface/registry"
	"github.com/ixre/go2o/core/query"
	"github.com/ixre/go2o/core/service/proto"
	context2 "golang.org/x/net/context"
)

var _ proto.ShopServiceServer = new(shopServiceImpl)

type shopServiceImpl struct {
	repo         shop.IShopRepo
	mchRepo      merchant.IMerchantRepo
	shopRepo     shop.IShopRepo
	query        *query.ShopQuery
	registryRepo registry.IRegistryRepo
	serviceUtil
	proto.UnimplementedShopServiceServer
}

func NewShopService(rep shop.IShopRepo, mchRepo merchant.IMerchantRepo,
	shopRepo shop.IShopRepo, registryRepo registry.IRegistryRepo, query *query.ShopQuery) *shopServiceImpl {
	return &shopServiceImpl{
		repo:         rep,
		mchRepo:      mchRepo,
		registryRepo: registryRepo,
		shopRepo:     shopRepo,
		query:        query,
	}
}

// SaveOfflineShop 保存门店
func (si *shopServiceImpl) SaveOfflineShop(_ context.Context, r *proto.SStore) (*proto.Result, error) {
	mch := si.mchRepo.GetMerchant(int(r.MerchantId))
	var err error
	if mch == nil {
		err = merchant.ErrNoSuchMerchant
	} else {
		mgr := mch.ShopManager()
		store, v := si.parseOfflineShop(r)
		var sp shop.IShop
		if store.Id > 0 {
			// 保存商店
			sp = mgr.GetStore(int(store.Id))
		} else {
			//创建商店
			sp = mgr.CreateShop(store)
		}
		err = sp.SetValue(store)
		if err == nil {
			ofs := sp.(shop.IOfflineShop)
			err = ofs.SetShopValue(v)
			if err == nil {
				err = sp.Save()
			}
		}
	}
	return si.error(err), nil
}

// * 查询自营店铺
func (si *shopServiceImpl) GetSelfSupportShops(_ context.Context, r *proto.SelfSupportShopRequest) (*proto.ShopListResponse, error) {
	shops := si.shopRepo.QuerySelfSupportShops()
	rsp := &proto.ShopListResponse{
		List: []*proto.SShop{},
	}
	for _, v := range shops {
		rsp.List = append(rsp.List, &proto.SShop{
			Id:       v.Id,
			ShopName: v.ShopName,
			Flag:     int32(v.Flag),
			Logo:     v.Logo,
			State:    int32(v.State),
		})
	}
	return rsp, nil
}

func (si *shopServiceImpl) DeleteStore(_ context.Context, id *proto.StoreId) (*proto.Result, error) {
	panic("implement me")
}

func (si *shopServiceImpl) QueryShopId(c context2.Context, req *proto.ShopAliasRequest) (*proto.Int64, error) {
	shopId := si.shopRepo.GetShopIdByAlias(req.ShopAlias)
	return &proto.Int64{Value: shopId}, nil
}

func (si *shopServiceImpl) GetShop(_ context.Context, req *proto.GetShopIdRequest) (*proto.SShop, error) {
	is := si.shopRepo.GetShop(req.ShopId)
	if is != nil {
		iop := is.(shop.IOnlineShop)
		iv := iop.GetShopValue()
		ret := si.parseShopDto(iv)
		ret.ShopTitle = ret.ShopName
		ret.Host = iv.Host
		ret.Logo = iv.Logo
		ret.Alias = iv.Alias
		ret.Telephone = iv.Telephone
		// 返回SellerMid
		im := si.mchRepo.GetMerchant(int(iv.VendorId))
		if im != nil {
			ret.SellerMid = int64(im.GetValue().MemberId)
		}
		return ret, nil
	}
	return nil, shop.ErrNoSuchShop
}

// CheckMerchantShopState 检查商户是否开通店铺
func (si *shopServiceImpl) CheckMerchantShopState(_ context.Context, id *proto.MerchantId) (*proto.CheckShopResponse, error) {
	sp := si.shopRepo.GetOnlineShopOfMerchant(int(id.Value))
	ret := &proto.CheckShopResponse{}
	if sp != nil {
		ret.Status = 1
		ret.Remark = "已开通"
		ret.ShopId = sp.Id
	} else {
		//todo: 返回审核中状态
	}
	return ret, nil
}

var shopHostRegexp *regexp.Regexp

// QueryShopByHost 根据主机头获取店铺编号
func (si *shopServiceImpl) QueryShopByHost(_ context.Context, host *proto.String) (*proto.Int64, error) {
	if shopHostRegexp == nil {
		domain, _ := si.registryRepo.GetValue(registry.Domain)
		shopHostRegexp = regexp.MustCompile("([^\\.]+)." + domain)
	}
	userHost := host.Value
	if shopHostRegexp.MatchString(userHost) {
		matches := shopHostRegexp.FindAllStringSubmatch(host.Value, 1)
		userHost = matches[0][1]
	}
	shopId := si.query.QueryShopIdByHost(userHost)
	return &proto.Int64{Value: shopId}, nil
}

// GetStore 获取门店
func (si *shopServiceImpl) GetStore(_ context.Context, storeId *proto.StoreId) (*proto.SStore, error) {
	sp := si.shopRepo.GetStore(storeId.Value)
	if sp != nil {
		v := sp.GetValue()
		ifv := sp.(shop.IOfflineShop)
		iv := ifv.GetShopValue()
		ret := &proto.SStore{
			Id:            storeId.Value,
			MerchantId:    v.VendorId,
			Name:          v.Name,
			State:         v.State,
			OpeningState:  v.OpeningState,
			StorePhone:    iv.Tel,
			StoreNotice:   "",
			Province:      iv.Province,
			City:          iv.City,
			District:      iv.District,
			Address:       "",
			DetailAddress: iv.Address,
			Lat:           float64(iv.Lat),
			Lng:           float64(iv.Lng),
			CoverRadius:   int32(iv.CoverRadius),
			SortNum:       v.SortNum,
		}
		return ret, nil
	}
	return nil, shop.ErrNoSuchShop
}

// TurnShop 打开或关闭商店
func (si *shopServiceImpl) TurnShop(_ context.Context, r *proto.TurnShopRequest) (*proto.Result, error) {
	var err error
	sp := si.repo.GetShop(r.ShopId)
	if sp == nil {
		err = shop.ErrNoSuchShop
	} else {
		if r.On {
			err = sp.TurnOn()
		} else {
			err = sp.TurnOff(r.Reason)
		}
	}
	return si.result(err), nil
}

// OpenShop 设置商店是否营业
func (si *shopServiceImpl) OpenShop(_ context.Context, shopId int32, on bool, reason string) (*proto.Result, error) {
	var err error
	sp := si.repo.GetShop(int64(shopId))
	if sp == nil {
		err = shop.ErrNoSuchShop
	} else {
		if on {
			err = sp.Opening()
		} else {
			err = sp.Pause()
		}
	}
	return si.result(err), nil
}

// SaveShop 保存线上商店
func (si *shopServiceImpl) SaveShop(_ context.Context, s *proto.SShop) (*proto.Result, error) {
	mch := si.mchRepo.GetMerchant(int(s.MerchantId))
	var err error
	if mch == nil {
		err = merchant.ErrNoSuchMerchant
	} else {
		v1 := si.parse2OnlineShop(s)
		mgr := mch.ShopManager()
		sp := mgr.GetOnlineShop()
		if sp == nil {
			err = merchant.ErrNoSuchShop
		} else {
			ofs := sp.(shop.IOnlineShop)
			err = ofs.SetShopValue(v1)
			if err == nil {
				err = sp.Save()
			}
		}
	}
	return si.error(err), nil
}

func (si *shopServiceImpl) parse2OnlineShop(s *proto.SShop) *shop.OnlineShop {
	return &shop.OnlineShop{
		Id:         s.Id,
		VendorId:   s.MerchantId,
		ShopName:   s.ShopName,
		Logo:       s.Logo,
		Host:       s.Host,
		Telephone:  s.Telephone,
		ShopTitle:  s.ShopTitle,
		ShopNotice: s.ShopNotice,
		State:      int16(s.State),
	}
}

func (si *shopServiceImpl) parseOfflineShop(r *proto.SStore) (*shop.Shop, *shop.OfflineShop) {
	return &shop.Shop{
			Id:           r.Id,
			VendorId:     r.MerchantId,
			ShopType:     shop.TypeOfflineShop,
			Name:         r.Name,
			State:        r.State,
			OpeningState: r.OpeningState,
			SortNum:      r.SortNum,
		}, &shop.OfflineShop{
			ShopId:      int(r.Id),
			Tel:         r.StorePhone,
			Province:    r.Province,
			City:        r.City,
			District:    r.District,
			Address:     r.DetailAddress,
			Lng:         float32(r.Lng),
			Lat:         float32(r.Lat),
			CoverRadius: int(r.CoverRadius),
		}
}

func (si *shopServiceImpl) parseShopDto(v shop.OnlineShop) *proto.SShop {
	return &proto.SShop{
		Id:         v.Id,
		MerchantId: v.VendorId,
		SellerMid:  0,
		ShopName:   v.ShopName,
		ShopTitle:  v.ShopTitle,
		ShopNotice: v.ShopNotice,
		Flag:       int32(v.Flag),
		Logo:       v.Logo,
		Alias:      v.Alias,
		Host:       v.Host,
		Telephone:  v.Telephone,
		State:      int32(v.State),
	}
}
