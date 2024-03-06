package impl

/**
 * Copyright (C) 2007-2020 56X.NET,All rights reserved.
 *
 * name : 2.cart_service.go
 * author : jarrysix (jarrysix#gmail.com)
 * date : 2020-09-27 11:04
 * description :
 * history :
 */

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/ixre/go2o/core/domain/interface/cart"
	proItem "github.com/ixre/go2o/core/domain/interface/item"
	"github.com/ixre/go2o/core/domain/interface/merchant"
	"github.com/ixre/go2o/core/domain/interface/merchant/shop"
	"github.com/ixre/go2o/core/service/parser"
	"github.com/ixre/go2o/core/service/proto"
	"github.com/ixre/gof/types"
	"github.com/ixre/gof/util"
)

var _ proto.CartServiceServer = new(cartServiceImpl)

type cartServiceImpl struct {
	itemRepo proItem.IItemRepo
	cartRepo cart.ICartRepo
	mchRepo  merchant.IMerchantRepo
	shopRepo shop.IShopRepo
	serviceUtil
	proto.UnimplementedCartServiceServer
}

func NewCartService(cartRepo cart.ICartRepo,
	goodsRepo proItem.IItemRepo,
	mchRepo merchant.IMerchantRepo,
	shopRepo shop.IShopRepo) *cartServiceImpl {
	return &cartServiceImpl{
		cartRepo: cartRepo,
		itemRepo: goodsRepo,
		mchRepo:  mchRepo,
		shopRepo: shopRepo,
	}
}

/*---------------- 批发购物车 ----------------*/

// WholesaleCartV1 批发购物车接口
func (s *cartServiceImpl) WholesaleCartV1(_ context.Context, r *proto.WsCartRequest) (*proto.Result, error) {
	//todo: check member
	c := s.cartRepo.GetMyCart(r.MemberId, cart.KWholesale)
	if r.Data == nil {
		r.Data = map[string]string{}
	}
	switch r.Action {
	case "GET":
		return s.wsGetCart(c, r.Data)
	case "MINI":
		return s.wsGetSimpleCart(c, r.Data)
	case "PUT":
		return s.wsPutItem(c, r.Data)
	case "UPDATE":
		return s.wsUpdateItem(c, r.Data)
	case "CHECK":
		return s.wsCheckCart(c, r.Data)
	}
	return s.result(errors.New("unknown action")), nil
}

// 转换勾选字典,数据如：{"1":["10","11"],"2":["20","21"]}
func (s *cartServiceImpl) parseCheckedMap(data string) (m map[int64][]int64) {
	if data != "" && data != "{}" {
		src := map[string][]string{}
		err := json.Unmarshal([]byte(data), &src)
		if err == nil {
			m = map[int64][]int64{}
			for k, v := range src {
				itemId, _ := strconv.Atoi(k)
				var skuList []int64
				for _, v2 := range v {
					skuId, _ := strconv.Atoi(v2)
					skuList = append(skuList, int64(skuId))
				}
				m[int64(itemId)] = skuList
			}
			return m
		}
	}
	return nil
}

// 获取可结算的购物车
func (s *cartServiceImpl) wsGetCart(c cart.ICartAggregateRoot, data map[string]string) (*proto.Result, error) {
	//统计checked
	checked := s.parseCheckedMap(data["checked"])
	checkout := data["checkout"] == "true"
	v := c.(cart.IWholesaleCart).JdoData(checkout, checked)
	if v != nil {
		for _, v2 := range v.Seller {
			mch := s.mchRepo.GetMerchant(int(v2.SellerId))
			if mch != nil {
				v2.Data["SellerName"] = mch.GetValue().CompanyName
			}
		}
	}
	return s.success(nil), nil
}

// 获取简易的购物车
func (s *cartServiceImpl) wsGetSimpleCart(c cart.ICartAggregateRoot, data map[string]string) (*proto.Result, error) {
	size, err := strconv.Atoi(data["size"])
	if err != nil {
		size = 5
	}
	qd := c.(cart.IWholesaleCart).QuickJdoData(size)
	return s.success(map[string]string{"JsonData": qd}), nil
}

// 转换提交到购物车的数据(PUT和UPDATE), 数据如：91:1;92:1
func (s *cartServiceImpl) wsParseCartPostedData(skuData string) (arr []*cart.ItemPair) {
	arr = []*cart.ItemPair{}
	splitArr := strings.Split(skuData, ";")
	for _, str := range splitArr {
		i := strings.Index(str, ":")
		if i == -1 {
			continue
		}
		skuId, err := util.I64Err(strconv.Atoi(str[:i]))
		quantity, err1 := util.I32Err(strconv.Atoi(str[i+1:]))
		if err == nil && err1 == nil {
			arr = append(arr, &cart.ItemPair{
				SkuId:    skuId,
				Quantity: quantity,
			})
		}
	}
	return arr
}

// 生成结算提交的数据(立即购买),skuData参考数据：skuId:num;skuId2;num2
func (s *cartServiceImpl) createCheckedData(itemId int64, arr []*cart.ItemPair) string {
	buf := bytes.NewBufferString("{\"")
	buf.WriteString(strconv.Itoa(int(itemId)))
	buf.WriteString("\":[")
	for i, v := range arr {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString("\"")
		buf.WriteString(strconv.Itoa(int(v.SkuId)))
		buf.WriteString("\"")
	}
	buf.WriteString("]}")
	return buf.String()
}

// 放入商品，data["Data"]
func (s *cartServiceImpl) wsPutItem(c cart.ICartAggregateRoot, data map[string]string) (*proto.Result, error) {
	aId := c.GetAggregateRootId()
	itemId, err := util.I64Err(strconv.Atoi(data["ItemId"]))
	arr := s.wsParseCartPostedData(data["Data"])
	for _, v := range arr {
		err = c.Put(itemId, v.SkuId, v.Quantity, false)
		if err != nil {
			break
		}
	}
	if err == nil {
		_, err = c.Save()
		if err == nil {
			mp := make(map[string]string)
			mp["cartId"] = strconv.Itoa(int(aId))
			mp["checked"] = s.createCheckedData(itemId, arr)
			return s.success(mp), nil
		}
	}
	return s.result(err), nil
}

func (s *cartServiceImpl) wsUpdateItem(c cart.ICartAggregateRoot, data map[string]string) (*proto.Result, error) {
	itemId, err := util.I64Err(strconv.Atoi(data["ItemId"]))
	arr := s.wsParseCartPostedData(data["Data"])
	for _, v := range arr {
		err = c.ResetQuantity(itemId, v.SkuId, v.Quantity)
		if err != nil {
			break
		}
	}
	if err == nil {
		_, err = c.Save()
	}
	return s.result(err), nil
}

// 勾选购物车，格式如：1:2;1:5
func (s *cartServiceImpl) wsCheckCart(c cart.ICartAggregateRoot, data map[string]string) (*proto.Result, error) {
	checked := data["Checked"]
	var arr []*cart.ItemPair
	splitArr := strings.Split(checked, ";")
	for _, str := range splitArr {
		i := strings.Index(str, ":")
		if i == -1 {
			continue
		}
		itemId, err := util.I64Err(strconv.Atoi(str[:i]))
		skuId, err1 := util.I64Err(strconv.Atoi(str[i+1:]))
		if err == nil && err1 == nil {
			arr = append(arr, &cart.ItemPair{
				ItemId: itemId,
				SkuId:  skuId,
			})
		}
	}
	err := c.SignItemChecked(arr)
	return s.result(err), nil
}

/*---------------- 普通购物车 ----------------*/

// 获取购物车
func (s *cartServiceImpl) getShoppingCart(buyerId int64, cartCode string) cart.ICartAggregateRoot {
	// 本地的购物车
	var ic cart.ICartAggregateRoot
	if len(cartCode) > 0 {
		ic = s.cartRepo.GetShoppingCartByKey(cartCode)
	}
	// 如果传入会员编号，则合并购物车,并返回会员的购物车
	if buyerId > 0 {
		// 获取用户购物车
		mc := s.cartRepo.GetMyCart(buyerId, cart.KNormal)
		if ic != nil {
			// 绑定临时购物车为会员购物车
			if mc == nil {
				ic.Bind(int(buyerId))
				return ic
			}
			// 会员购物车合并临时购物车
			if mc.GetAggregateRootId() != ic.GetAggregateRootId() {
				mc.(cart.INormalCart).Combine(ic)
				_, _ = mc.Save()
			}
		}
		// 如果会员购物车存在则返回, 不存在也需创建一个临时的购物车
		if mc != nil {
			return mc
		}
	}
	// 为其他会员的购物车, 则新建购物车
	if ic != nil && ic.BuyerId() > 0 {
		cartCode = ""
		ic = nil
	}
	// 生成一个新的code和购物车
	if ic == nil {
		return s.cartRepo.NewTempNormalCart(int(buyerId), cartCode)
	}
	// 如果只传入code,且购物车存在，直接返回。
	return ic
}

// GetShoppingCart 获取购物车,当购物车编号不存在时,将返回一个新的购物车
func (s *cartServiceImpl) GetShoppingCart(_ context.Context, r *proto.ShoppingCartId) (*proto.SShoppingCart, error) {
	if r.IsWholesale {
		return nil, errors.New("not implement")
	}
	c := s.getShoppingCart(r.GetUserId(), r.CartCode)
	return s.parseCart(c), nil
}

// 转换购物车数据
func (s *cartServiceImpl) parseCart(c cart.ICartAggregateRoot) *proto.SShoppingCart {
	dto := parser.ParseToDtoCart(c)
	for _, v := range dto.Sellers {
		is := s.shopRepo.GetShop(v.ShopId)
		if is != nil {
			io := is.(shop.IOnlineShop)
			v.ShopName = io.GetShopValue().ShopName
		} else {
			for _, it := range v.Items {
				_ = c.ResetQuantity(it.ItemId, it.SkuId, 0)
			}
		}
	}
	return dto
}

// ApplyItem 购物车商品操作
func (s *cartServiceImpl) ApplyItem(_ context.Context, r *proto.CartItemOpRequest) (*proto.CartItemResponse, error) {
	c := s.getShoppingCart(r.CartId.UserId, r.CartId.CartCode)
	if c == nil {
		return nil, cart.ErrNoSuchCart
	}
	it := r.Item
	if it == nil {
		return &proto.CartItemResponse{ErrCode: 1, ErrMsg: "no any item to put cart"}, nil
	}
	rc := c.(cart.INormalCart)
	items := make([]*proto.SShoppingCartItem, 0)
	var err error
	if r.Op == proto.ECartItemOp_PUT {
		// 加入购物车
		err = c.Put(it.ItemId, it.SkuId, it.Quantity, false)
	} else if r.Op == proto.ECartItemOp_CHECKOUT {
		// 立即购买，清空数量并重新加入
		c.ResetQuantity(it.ItemId, it.SkuId, 0)
		err = c.Put(it.ItemId, it.SkuId, it.Quantity, true)
	} else if r.Op == proto.ECartItemOp_CHANGE {
		// 更新数量
		err = c.ResetQuantity(it.ItemId, it.SkuId, it.Quantity)
	}
	if err != nil {
		return &proto.CartItemResponse{ErrCode: 1, ErrMsg: err.Error()}, nil
	}
	item := rc.GetItem(it.ItemId, it.SkuId)
	items = append(items, parser.ParseCartItem(item))
	if _, err := c.Save(); err != nil {
		return &proto.CartItemResponse{ErrCode: 1, ErrMsg: err.Error()}, nil
	}
	return &proto.CartItemResponse{
		Items: items,
	}, nil
}

// CheckCart 勾选商品结算
func (s *cartServiceImpl) CheckCart(_ context.Context, r *proto.CheckCartRequest) (*proto.Result, error) {
	c := s.getShoppingCart(r.Id.UserId, r.Id.CartCode)
	items := make([]*cart.ItemPair, len(r.Items))
	for i, v := range r.Items {
		items[i] = &cart.ItemPair{
			ItemId:  v.ItemId,
			SkuId:   v.SkuId,
			Checked: int32(types.ElseInt(v.Checked, 1, 0)),
		}
	}
	err := c.SignItemChecked(items)
	if err == nil {
		_, err = c.Save()
	}
	return s.error(err), nil
}
