package cart

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/ixre/go2o/core/domain/interface/cart"
	"github.com/ixre/go2o/core/domain/interface/item"
	"github.com/ixre/go2o/core/domain/interface/member"
	"github.com/ixre/go2o/core/domain/interface/merchant"
	"github.com/ixre/go2o/core/domain/interface/merchant/shop"
	"github.com/ixre/go2o/core/infrastructure/format"
	"github.com/ixre/gof/util"
)

var _ cart.ICartAggregateRoot = new(wholesaleCartImpl)
var _ cart.IWholesaleCart = new(wholesaleCartImpl)

type wCartQuickSkuJdo struct {
	// 商品编号
	ItemId int64
	// SKU编号
	SkuId int64
	// 商品名称
	ItemName string
	// SKU编码
	SkuCode string
	// SKU图片
	SkuImage string
	// 规格文本
	SpecWord string
	// 数量
	Quantity int32
	// 价格
	Price int64
}

type wholesaleCartImpl struct {
	value      *cart.WsCart
	repo       cart.ICartRepo
	itemRepo   item.IItemRepo
	memberRepo member.IMemberRepo
	mchRepo    merchant.IMerchantRepo
	summary    string
	shop       shop.IShop
	deliver    member.IDeliverAddress
	snapMap    map[int64]*item.Snapshot
}

func CreateWholesaleCart(val *cart.WsCart, rep cart.ICartRepo,
	memberRepo member.IMemberRepo, mchRepo merchant.IMerchantRepo,
	itemRepo item.IItemRepo) cart.ICartAggregateRoot {
	return (&wholesaleCartImpl{
		value:      val,
		repo:       rep,
		memberRepo: memberRepo,
		itemRepo:   itemRepo,
	}).init()
}

func (c *wholesaleCartImpl) init() cart.ICartAggregateRoot {
	// 获取购物车项
	if c.GetAggregateRootId() > 0 {
		if c.value.Items == nil {
			c.value.Items = c.repo.SelectWsCartItem("cart_id= $1",
				c.GetAggregateRootId())
		}
	}
	if c.value.Items == nil {
		c.value.Items = []*cart.WsCartItem{}
	}
	// 初始化购物车的信息
	if c.value != nil && c.value.Items != nil {
		c.setAttachGoodsInfo(c.value.Items)
	}
	return c
}

// 购物车种类
func (c *wholesaleCartImpl) Kind() cart.Kind {
	return cart.KWholesale
}

// 获取买家编号
func (c *wholesaleCartImpl) BuyerId() int64 {
	return c.value.BuyerId
}

// Bind 绑定买家
func (c *wholesaleCartImpl) Bind(buyerId int) error {
	if c.value.BuyerId > 0 {
		return errors.New("cart has already bind buyer")
	}
	c.value.BuyerId = int64(buyerId)
	_, err := c.Save()
	return err
}

func (c *wholesaleCartImpl) Clone() cart.ICartAggregateRoot {
	panic("implement me")
}

func (c *wholesaleCartImpl) Prepare() error {
	return c.check()
}

// 检查购物车(仅结算商品)
func (c *wholesaleCartImpl) check() error {
	if c.value == nil || len(c.value.Items) == 0 {
		return cart.ErrEmptyShoppingCart
	}
	for _, v := range c.value.Items {
		it := c.itemRepo.GetItem(v.ItemId)
		if it == nil {
			return item.ErrNoSuchItem // 没有商品
		}
		// 验证批发权限
		wsIt := it.Wholesale()
		if wsIt == nil || !wsIt.CanWholesale() {
			return item.ErrItemWholesaleOff
		}
		// 验证库存
		stock := it.GetValue().StockNum
		if v.SkuId > 0 {
			if sku := it.GetSku(v.SkuId); sku != nil {
				stock = sku.Stock
			}
		}
		if stock == 0 {
			return fmt.Errorf(item.ErrFullOfStock.Error(), it.GetValue().Title) // 已经卖完了
		}
		if stock < v.Quantity {
			return item.ErrOutOfStock // 超出库存
		}
	}
	return nil
}

// 获取商品的快照列表
func (c *wholesaleCartImpl) getSnapshotsMap(items []*cart.WsCartItem) map[int64]*item.Snapshot {
	if c.snapMap == nil && items != nil {
		l := len(items)
		c.snapMap = make(map[int64]*item.Snapshot, l)
		if l > 0 {
			ids := make([]int64, l)
			for i, v := range items {
				ids[i] = v.ItemId
			}
			snapList := c.itemRepo.GetSnapshots(ids)
			for _, v := range snapList {
				v2 := v
				c.snapMap[v.ItemId] = &v2
			}
		}
	}
	return c.snapMap
}

func (c *wholesaleCartImpl) getBuyerLevelId() int {
	if c.value.BuyerId > 0 {
		m := c.memberRepo.GetMember(c.value.BuyerId)
		if m != nil {
			return m.GetValue().Level
		}
	}
	return 0
}

func (c *wholesaleCartImpl) setItemInfo(snap *item.GoodsItem, level int) {
	// 设置会员价
	if level > 0 {
		gds := c.itemRepo.CreateItem(snap)
		snap.Price = gds.GetPromotionPrice(level)
	}
}

// 设置附加的商品信息
func (c *wholesaleCartImpl) setAttachGoodsInfo(items []*cart.WsCartItem) {
	list := c.getSnapshotsMap(items)
	if list == nil {
		return
	}
	var sku *item.Sku
	for _, v := range items {
		it := c.itemRepo.GetItem(v.ItemId)
		if it == nil {
			continue
		}
		if v.SkuId > 0 {
			sku = it.GetSku(v.SkuId)
		} else {
			iv := it.GetValue()
			sku = &item.Sku{
				ProductId:   iv.ProductId,
				ItemId:      iv.Id,
				Title:       iv.Title,
				Image:       iv.Image,
				SpecData:    "",
				SpecWord:    "",
				Code:        iv.Code,
				OriginPrice: iv.OriginPrice,
				Price:       iv.Price,
				Cost:        iv.Cost,
				Weight:      iv.Weight,
				Bulk:        iv.Bulk,
				Stock:       iv.StockNum,
				SaleNum:     iv.SaleNum,
			}
		}
		v.Sku = item.ParseSkuMedia(it.GetValue(), sku)

		//  会员价
		//var level int32
		//if gv.LevelSales == 1 && level != -1 {
		//    if level == 0 {
		//        level = c.getBuyerLevelId()
		//    }
		//    //todo: ???
		//    //c.setGoodsInfo(gv, level)
		//}
		//
	}
}

// 获取聚合根编号
func (c *wholesaleCartImpl) GetAggregateRootId() int32 {
	return c.value.ID
}

func (c *wholesaleCartImpl) GetValue() cart.WsCart {
	return *c.value
}

// 获取商品集合
func (c *wholesaleCartImpl) Items() []*cart.WsCartItem {
	return c.getItems()
}

func (c *wholesaleCartImpl) getItems() []*cart.WsCartItem {
	return c.value.Items
}

// 根据SKU获取项
func (c *wholesaleCartImpl) getSkuItem(itemId, skuId int64) *cart.WsCartItem {
	for _, v := range c.value.Items {
		if v.ItemId == itemId && v.SkuId == skuId {
			return v
		}
	}
	return nil
}

// 检查是否有库存
func (c *wholesaleCartImpl) checkItemStock(itemId int64, skuId int64) (*item.GoodsItem, *item.Sku, int32, error) {
	var sku *item.Sku

	it := c.itemRepo.GetItem(itemId)
	if it == nil {
		return nil, sku, 0, item.ErrNoSuchItem // 没有商品
	}
	iv := it.GetValue()
	// 库存,如有SKU，则使用SKU的库存
	stock := iv.StockNum
	// 判断是否上架
	if iv.ShelveState != item.ShelvesOn {
		return nil, sku, stock, fmt.Errorf(item.ErrNotOnShelves.Error(), iv.Title) //未上架
	}
	// 验证批发权限
	wsIt := it.Wholesale()
	if wsIt == nil || !wsIt.CanWholesale() {
		return &iv, sku, stock, item.ErrItemWholesaleOff
	}
	// 判断商品SkuId
	if skuId > 0 {
		sku = it.GetSku(skuId)
		if sku == nil {
			return &iv, nil, stock, item.ErrNoSuchSku
		}
		stock = sku.Stock
		//todo: 如果SKU没有启用批发,或没有达到最低的数量
		//arr := wsIt.GetSkuPrice(skuId)
	} else if iv.SkuNum > 0 {
		return &iv, sku, stock, cart.ErrItemNoSku
	}
	// 检查是否已经卖完了
	if stock == 0 {
		return &iv, sku, stock, fmt.Errorf(item.ErrFullOfStock.Error(), it.GetValue().Title) // 已经卖完了
	}
	return &iv, sku, stock, nil
}

// 添加项
func (c *wholesaleCartImpl) put(itemId, skuId int64, quantity int32, checkOnly bool) (*cart.WsCartItem, error) {
	var err error
	if c.value.Items == nil {
		c.value.Items = []*cart.WsCartItem{}
	}
	// 检查库存
	iv, sku, stock, err := c.checkItemStock(itemId, skuId)
	// 添加数量
	var dst *cart.WsCartItem
	for _, v := range c.value.Items {
		if v.ItemId == itemId && v.SkuId == skuId {
			dst = v
			// 设置商品数量
			finalQuantity := v.Quantity + quantity
			if finalQuantity > stock {
				return v, item.ErrOutOfStock // 库存不足
			}
			v.Quantity = finalQuantity
		} else {
			if checkOnly {
				//v.
			}
		}
	}
	if dst != nil {
		// 对现有商品数量进行调整，直接返回
		return dst, nil
	}
	// 向购物车中添加新的商品项
	c.snapMap = nil

	// 设置商品的相关信息
	c.setItemInfo(iv, c.getBuyerLevelId())

	v := &cart.WsCartItem{
		CartId:   c.GetAggregateRootId(),
		SellerId: iv.VendorId,
		ShopId:   iv.ShopId,
		ItemId:   iv.Id,
		SkuId:    skuId,
		Quantity: quantity,
		Sku:      item.ParseSkuMedia(*iv, sku),
	}
	c.value.Items = append(c.value.Items, v)
	return v, err
}

// // 更新项
// func (c *wholesaleCartImpl) update(itemId, skuId int64, quantity int32) error {
// 	if c.value.Items == nil {
// 		return cart.ErrEmptyShoppingCart
// 	}
// 	ci := c.getSkuItem(itemId, skuId)
// 	if ci == nil {
// 		return cart.ErrItemNoSku
// 	}
// 	it := c.itemRepo.GetItem(itemId)
// 	if it == nil {
// 		return item.ErrNoSuchItem // 没有商品
// 	}
// 	iv := it.GetValue()
// 	// 库存,如有SKU，则使用SKU的库存
// 	stock := iv.StockNum
// 	if quantity > stock {
// 		return item.ErrOutOfStock
// 	}
// 	// 判断商品SkuId
// 	if skuId > 0 {
// 		var sku *item.Sku
// 		sku = it.GetSku(skuId)
// 		if sku == nil {
// 			return item.ErrNoSuchSku
// 		}
// 		stock = sku.Stock
// 	}
// 	// 检查是否已经卖完了
// 	if stock == 0 {
// 		return item.ErrFullOfStock
// 	}
// 	// 超出库存
// 	if quantity > stock {
// 		return item.ErrOutOfStock
// 	}
// 	ci.Quantity = quantity
// 	return nil
// }

// Put 添加项
func (c *wholesaleCartImpl) Put(itemId, skuId int64, num int32, checkOnly bool) error {
	_, err := c.put(itemId, skuId, num, checkOnly)
	return err
}

// ResetQuantity 重置数量
func (c *wholesaleCartImpl) ResetQuantity(itemId, skuId int64, quantity int32) error {
	if c.value.Items == nil {
		return cart.ErrEmptyShoppingCart
	}
	var dst *cart.WsCartItem
	// 查找商品项
	for _, v := range c.value.Items {
		if v.ItemId == itemId && v.SkuId == skuId {
			dst = v
			break
		}
	}
	if dst == nil {
		// 不包含该商品
		return cart.ErrNoMatchItem
	}
	if quantity > dst.Quantity {
		// 如果商品数量增加，需验证库存
		_, _, stock, err := c.checkItemStock(itemId, skuId)
		if err != nil {
			// 校验库存失败
			return err
		}
		if quantity > stock {
			// 库存不足
			return item.ErrOutOfStock
		}
	}
	// 更新数量,如果商品数量为0,则在Save时，删除商品项
	dst.Quantity = quantity
	// 清除商品缓存
	//c.snapMap = nil
	return nil
}

// Code 获取购物车编码
func (c *wholesaleCartImpl) Code() string {
	return c.value.Code
}

// SetBuyerAddress 设置购买会员收货地址
func (c *wholesaleCartImpl) SetBuyerAddress(addressId int64) error {
	if c.value.BuyerId < 0 {
		return cart.ErrCartNoBuyer
	}
	m := c.memberRepo.GetMember(c.value.BuyerId)
	if m == nil {
		return member.ErrNoSuchMember
	}
	addr := m.Profile().GetAddress(addressId)
	if addr == nil {
		return member.ErrNoSuchAddress
	}
	return c.setBuyerAddress(addressId)
}

func (c *wholesaleCartImpl) setBuyerAddress(addressId int64) error {
	c.value.DeliverId = addressId
	_, err := c.Save()
	return err
}

// SignItemChecked 标记商品结算
func (c *wholesaleCartImpl) SignItemChecked(items []*cart.ItemPair) error {
	panic("not support")
}

// Save 保存购物车
func (c *wholesaleCartImpl) Save() (int32, error) {
	c.value.UpdateTime = time.Now().Unix()
	id, err := util.I32Err(c.repo.SaveWsCart(c.value))
	c.value.ID = id
	if c.value.Items != nil {
		for _, v := range c.value.Items {
			if v.Quantity <= 0 {
				//c.repo.RemoveCartItem(v.Id)
				_, _ = c.repo.BatchDeleteWsCartItem("id= $1", v.ID)
			} else {
				v.CartId = c.GetAggregateRootId()
				v.ID, err = util.I32Err(c.repo.SaveWsCartItem(v))
			}
		}
	}
	return id, err
}

// CheckedItems 获取勾选的商品
func (c *wholesaleCartImpl) CheckedItems(checked map[int64][]int64) []*cart.ItemPair {
	var items []*cart.ItemPair
	if checked != nil {
		for _, v := range c.value.Items {
			arr, ok := checked[v.ItemId]
			if !ok {
				continue
			}
			for _, skuId := range arr {
				if skuId == v.SkuId {
					items = append(items, &cart.ItemPair{
						ItemId:   int64(v.ItemId),
						SkuId:    skuId,
						SellerId: v.SellerId,
						Quantity: v.Quantity,
					})
				}
			}
		}
	}
	return items
}

// Release 释放购物车,如果购物车的商品全部结算,则返回true
func (c *wholesaleCartImpl) Release(checked map[int64][]int64) bool {
	if checked == nil {
		return true
	}
	//部分计算
	part := false
	for _, v := range c.value.Items {
		//判断sku是否被结算
		skuChecked := false
		//判断ItemId
		for itemId, skuList := range checked {
			if int64(v.ItemId) != itemId {
				continue
			}
			//判断SkuId
			for _, skuId := range skuList {
				if int64(v.SkuId) == skuId {
					skuChecked = true
					_ = c.ResetQuantity(v.ItemId, v.SkuId, 0)
				}
			}
		}
		if !part && !skuChecked {
			part = true
		}
	}
	_, _ = c.Save()
	return !part
}

// Destroy 销毁购物车
func (c *wholesaleCartImpl) Destroy() (err error) {
	c.snapMap = nil //clean
	if err = c.repo.EmptyCartItems(c.GetAggregateRootId()); err == nil {
		return c.repo.DeleteCart(c.GetAggregateRootId())
	}
	return err
}

// 获取购物车商品Jdo数据
func (c *wholesaleCartImpl) getItemJdoData(list []*cart.ItemPair,
	itemId int64) cart.WCartItemJdo {
	it := c.itemRepo.GetItem(itemId)
	v := it.GetValue()
	itw := it.Wholesale()
	itJdo := cart.WCartItemJdo{
		ItemId:    int64(itemId),
		ItemName:  v.Title,
		ItemImage: v.Image,
		Sku:       []cart.WCartSkuJdo{},
		Data:      map[string]string{},
	}
	skuSignMap := make(map[int64]bool)
	for _, v := range list {
		if v.ItemId != itemId || skuSignMap[int64(v.SkuId)] {
			continue
		}
		skuV := it.GetSku(v.SkuId)
		skuJdo := cart.WCartSkuJdo{
			SkuId:            v.SkuId,
			SkuCode:          skuV.Code,
			SkuImage:         skuV.Image,
			SpecWord:         skuV.SpecWord,
			Quantity:         v.Quantity,
			Price:            0,
			DiscountPrice:    0,
			CanSalesQuantity: skuV.Stock,
			JData:            "{}",
		}
		mp := map[string]interface{}{}
		mp["canSalesQuantity"] = skuV.Stock
		c.setSkuJdoData(itw, &skuJdo, mp)

		itJdo.Sku = append(itJdo.Sku, skuJdo)
		skuSignMap[skuJdo.SkuId] = true
	}
	return itJdo
}

func (c *wholesaleCartImpl) setSkuJdoData(itw item.IWholesaleItem,
	sku *cart.WCartSkuJdo, mp map[string]interface{}) {
	prArr := itw.GetSkuPrice(sku.SkuId)
	price := itw.GetWholesalePrice(sku.SkuId, sku.Quantity)
	var priceRange [][]string
	for _, v := range prArr {
		priceRange = append(priceRange, []string{
			strconv.Itoa(int(v.RequireQuantity)),
			format.FormatIntMoney(v.WholesalePrice),
		})
	}
	sku.Price = price
	sku.DiscountPrice = price
	mp["priceRange"] = priceRange
	data, _ := json.Marshal(mp)
	sku.JData = string(data)

}

// Jdo数据
func (c *wholesaleCartImpl) JdoData(checkout bool, checked map[int64][]int64) *cart.WCartJdo {
	var items []*cart.ItemPair
	if checked != nil {
		items = c.CheckedItems(checked)
	} else {
		for _, v := range c.value.Items {
			items = append(items, &cart.ItemPair{
				ItemId:   v.ItemId,
				SkuId:    v.SkuId,
				SellerId: v.SellerId,
				Quantity: v.Quantity,
			})
		}
	}
	jdo := &cart.WCartJdo{
		Seller: []cart.WCartSellerJdo{},
		Data:   map[string]string{},
	}
	sellerMap := make(map[int64]int)
	itemSignMap := make(map[int64]bool)
	for _, v := range items {
		// 如果已处理过商品，则跳过
		if v.SellerId <= 0 || itemSignMap[int64(v.ItemId)] {
			continue
		}
		vi, ok := sellerMap[v.SellerId]
		//初始化SellerJdo
		if !ok {
			vJdo := cart.WCartSellerJdo{
				SellerId: v.SellerId,
				Item:     []cart.WCartItemJdo{},
				Data:     map[string]string{},
			}
			jdo.Seller = append(jdo.Seller, vJdo)
			vi = len(jdo.Seller) - 1
			sellerMap[v.SellerId] = vi
		}
		// 设置商品信息
		itJdo := c.getItemJdoData(items, v.ItemId)
		jdo.Seller[vi].Item = append(jdo.Seller[vi].Item, itJdo)
		itemSignMap[int64(v.ItemId)] = true
	}
	if checkout {
		c.checkoutJdoData(jdo)
	}
	return jdo
}

// 附加结算数据
func (c *wholesaleCartImpl) checkoutJdoData(jdo *cart.WCartJdo) {
	var totalAmount int64
	sellerAmountMap := map[int64]int64{}
	for _, s := range jdo.Seller {
		for _, i := range s.Item {
			for _, sku := range i.Sku {
				v := sellerAmountMap[s.SellerId]
				v += sku.Price * int64(sku.Quantity)
				sellerAmountMap[s.SellerId] = v
			}
		}
		sellerAmount := sellerAmountMap[s.SellerId]
		totalAmount += sellerAmount
		//卖家汇总
		s.Data["ItemAmount"] = format.FormatIntMoney(sellerAmount)
		s.Data["ExpressAmount"] = format.DecimalToString(0)
	}
	//总计
	jdo.Data["TotalExpressAmount"] = format.DecimalToString(0)
	jdo.Data["TotalItemAmount"] = format.FormatIntMoney(totalAmount)
	jdo.Data["FinalAmount"] = format.FormatIntMoney(totalAmount)
}

// 简单Jdo数据,max为最多数量
func (c *wholesaleCartImpl) QuickJdoData(max int) string {
	items := c.value.Items
	l := len(items)
	if max > l {
		max = l
	}
	skuList := []wCartQuickSkuJdo{}
	for _, v := range items[:max] {
		// 如果已处理过商品，则跳过
		if v.SellerId <= 0 {
			continue
		}
		it := c.itemRepo.GetItem(v.ItemId)
		itw := it.Wholesale()
		wPrice := itw.GetWholesalePrice(v.SkuId, v.Quantity)
		skuV := it.GetSku(v.SkuId)
		skuJdo := wCartQuickSkuJdo{
			ItemId:   int64(v.ItemId),
			SkuId:    int64(v.SkuId),
			ItemName: it.GetValue().Title,
			SkuCode:  skuV.Code,
			SkuImage: skuV.Image,
			SpecWord: skuV.SpecWord,
			Price:    wPrice,
			Quantity: v.Quantity,
		}
		skuList = append(skuList, skuJdo)
	}
	mp := map[string]interface{}{
		"len":  l,
		"item": skuList,
	}
	d, err := json.Marshal(mp)
	if err == nil {
		return string(d)
	}
	return "{\"error\":\"" + err.Error() + "\"}"
}
