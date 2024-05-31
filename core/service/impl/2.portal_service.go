package impl

import (
	"errors"
	"fmt"

	"github.com/ixre/go2o/core/dao"
	"github.com/ixre/go2o/core/dao/impl"
	"github.com/ixre/go2o/core/dao/model"
	"github.com/ixre/go2o/core/service/proto"
	"github.com/ixre/gof/db/orm"
	"golang.org/x/net/context"
)

var _ proto.PortalServiceServer = new(portalService)

type portalService struct {
	repo          *impl.CommonDao
	dao           dao.IPortalDao
	searchWordDao dao.ISysSearchWordDao
	serviceUtil
	proto.UnimplementedPortalServiceServer
}

func NewPortalService(o orm.Orm, d *impl.CommonDao, portalDao dao.IPortalDao) proto.PortalServiceServer {
	return &portalService{
		repo:          d,
		dao:           portalDao,
		searchWordDao: impl.NewSysSearchWordDao(o),
	}
}

func (p *portalService) SaveNav(_ context.Context, r *proto.SaveNavRequest) (*proto.SaveNavResponse, error) {
	var dst *model.PortalNav
	if r.Id > 0 {
		if dst = p.dao.GetNav(r.Id); dst == nil {
			return &proto.SaveNavResponse{
				ErrCode: 2,
				ErrMsg:  "no such record",
			}, nil
		}
	} else {
		dst = &model.PortalNav{}

	}
	dst.Text = r.Text
	dst.Url = r.Url
	dst.Target = r.Target
	dst.Image = r.Image
	dst.NavType = r.NavType
	dst.NavGroup = r.NavGroup
	id, err := p.dao.SaveNav(dst)
	ret := &proto.SaveNavResponse{
		Id: int64(id),
	}
	if err != nil {
		ret.ErrCode = 1
		ret.ErrMsg = err.Error()
	}
	return ret, nil
}

func (p *portalService) GetNav(context context.Context, id *proto.PortalNavId) (*proto.SNav, error) {
	v := p.dao.GetNav(id.Value)
	if v == nil {
		return nil, fmt.Errorf("no such nav")
	}
	return p.parseNav(v), nil
}

func (p *portalService) QueryNavList(context context.Context, r *proto.QueryNavRequest) (*proto.QueryNavResponse, error) {
	arr := p.dao.SelectNav("nav_type= $1 AND nav_group = $2  order by id ASC LIMIT $3 OFFSET $4", r.NavType, r.Group, r.Size, r.Begin)
	ret := &proto.QueryNavResponse{
		List: make([]*proto.SNav, len(arr)),
	}
	for i, v := range arr {
		ret.List[i] = p.parseNav(v)
	}
	return ret, nil
}

func (p *portalService) DeleteNav(context context.Context, id *proto.PortalNavId) (*proto.Result, error) {
	err := p.dao.DeleteNav(id.Value)
	return p.error(err), nil
}

// 获取门户导航
func (p *portalService) GetPortalNav_(id int32) *model.PortalNav {
	return p.repo.GetPortalNav(id)
}

// 保存门户导航
func (p *portalService) SavePortalNav_(v *model.PortalNav) (*proto.Result, error) {
	_, err := p.repo.SavePortalNav(v)
	return p.result(err), nil
}

// 删除门户导航
func (p *portalService) DeletePortalNav_(id int32) (*proto.Result, error) {
	err := p.repo.DeletePortalNav(id)
	return p.result(err), nil
}

// 获取门户导航
func (p *portalService) SelectPortalNav(navType int32) []*model.PortalNav {
	return p.repo.SelectPortalNav("nav_type= $1", navType)
}

// 获取导航类型
func (p *portalService) GetPortalNavType_(id int32) *model.PortalNavType {
	return p.repo.GetPortalNavType(id)
}

// 保存导航类型
func (p *portalService) SavePortalNavType_(v *model.PortalNavType) (*proto.Result, error) {
	_, err := p.repo.SavePortalNavType(v)
	return p.result(err), nil
}

// 删除导航类型
func (p *portalService) DeletePortalNavType_(id int32) (*proto.Result, error) {
	err := p.repo.DeletePortalNavType(id)
	return p.result(err), nil
}

func (p *portalService) parseNav(v *model.PortalNav) *proto.SNav {
	return &proto.SNav{
		Id:       v.Id,
		Text:     v.Text,
		Url:      v.Url,
		Target:   v.Target,
		Image:    v.Image,
		NavType:  v.NavType,
		NavGroup: v.NavGroup,
	}
}

func (p *portalService) SaveNavGroup(c context.Context, r *proto.SaveNavGroupRequest) (*proto.SaveNavGroupResponse, error) {
	var dst *model.NavGroup
	if r.Id > 0 {
		if dst = p.dao.GetNavGroup(r.Id); dst == nil {
			return &proto.SaveNavGroupResponse{
				ErrCode: 2,
				ErrMsg:  "no such record",
			}, nil
		}
	} else {
		dst = &model.NavGroup{}
	}

	dst.Name = r.Name
	id, err := p.dao.SaveNavGroup(dst)
	ret := &proto.SaveNavGroupResponse{
		Id: int64(id),
	}
	if err != nil {
		ret.ErrCode = 1
		ret.ErrMsg = err.Error()
	}
	return ret, nil
}

func (p *portalService) parseNavGroup(v *model.NavGroup) *proto.SNavGroup {
	return &proto.SNavGroup{
		Id:   int32(v.Id),
		Name: v.Name,
	}
}
func (p *portalService) QueryNavGroupList(c context.Context, r *proto.QueryNavGroupRequest) (*proto.QueryNavGroupResponse, error) {
	arr := p.dao.SelectNavGroup("")
	ret := &proto.QueryNavGroupResponse{
		Value: make([]*proto.SNavGroup, len(arr)),
	}
	for i, v := range arr {
		ret.Value[i] = p.parseNavGroup(v)
	}
	return ret, nil
}

func (p *portalService) DeleteNavGroup(c context.Context, id *proto.PortalNavGroupId) (*proto.Result, error) {
	err := p.dao.DeleteNavGroup(id.Value)
	return p.error(err), nil
}

// SaveSearchWord 保存热搜词
func (p *portalService) SaveSearchWord(_ context.Context, r *proto.SaveSearchWordRequest) (*proto.SaveSearchWordResponse, error) {
	var dst *model.SearchWord
	if r.Id > 0 {
		if dst = p.searchWordDao.GetSearchWord(r.Id); dst == nil {
			return &proto.SaveSearchWordResponse{
				ErrCode: 2,
				ErrMsg:  "no such record",
			}, nil
		}
	} else {
		dst = &model.SearchWord{}

	}

	dst.Word = r.Word
	dst.SearchCount = int(r.SearchCount)
	dst.Flag = int(r.Flag)

	id, err := p.searchWordDao.SaveSearchWord(dst)
	ret := &proto.SaveSearchWordResponse{
		Id: int64(id),
	}
	if err != nil {
		ret.ErrCode = 1
		ret.ErrMsg = err.Error()
	}
	return ret, nil
}

func (p *portalService) parseSearchWord(v *model.SearchWord) *proto.SSearchWord {
	return &proto.SSearchWord{
		Id:          v.Id,
		Word:        v.Word,
		SearchCount: int32(v.SearchCount),
		Flag:        int32(v.Flag),
	}
}

// GetSearchWord 获取热搜词
func (p *portalService) GetSearchWord(_ context.Context, id *proto.SysSearchWordId) (*proto.SSearchWord, error) {
	v := p.searchWordDao.GetSearchWord(id.Value)
	if v == nil {
		return nil, errors.New("no such search word")
	}
	return p.parseSearchWord(v), nil
}

// QuerySearchWordList 获取热搜词列表
func (p *portalService) QuerySearchWordList(_ context.Context, r *proto.QuerySearchWordRequest) (*proto.QuerySearchWordResponse, error) {
	arr := p.searchWordDao.SelectSearchWord(fmt.Sprintf(" 1=1 LIMIT %d", r.Size))
	ret := &proto.QuerySearchWordResponse{
		Value: make([]*proto.SSearchWord, len(arr)),
	}
	for i, v := range arr {
		ret.Value[i] = p.parseSearchWord(v)
	}
	return ret, nil
}

// DeleteSearchWord 删除热搜词
func (p *portalService) DeleteSearchWord(_ context.Context, id *proto.SysSearchWordId) (*proto.Result, error) {
	err := p.searchWordDao.DeleteSearchWord(id.Value)
	return p.error(err), nil
}
