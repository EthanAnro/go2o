package cart

import (
	"errors"
	"fmt"
	"time"

	"log"

	"github.com/ixre/go2o/core/domain/interface/cart"
	"github.com/ixre/go2o/core/domain/interface/item"
	"github.com/ixre/go2o/core/domain/interface/member"
	"github.com/ixre/go2o/core/infrastructure/domain"
)

var _ cart.ICartAggregateRoot = new(cartImpl)
var _ cart.INormalCart = new(cartImpl)

type cartImpl struct {
	value      *cart.NormalCart
	rep        cart.ICartRepo
	goodsRepo  item.IItemRepo
	memberRepo member.IMemberRepo
	//summary    string
	//shop       shop.IShop
	//deliver    member.IDeliverAddress
	snapMap map[int64]*item.Snapshot
}

func NewNormalCart(val *cart.NormalCart, rep cart.ICartRepo,
	memberRepo member.IMemberRepo, goodsRepo item.IItemRepo) cart.ICartAggregateRoot {
	c := &cartImpl{
		value:      val,
		rep:        rep,
		memberRepo: memberRepo,
		goodsRepo:  goodsRepo,
	}
	return c.init()
}

// 创建新的临时购物车
func CreateTempNormalCart(buyerId int, cartCode string, rep cart.ICartRepo, memberRepo member.IMemberRepo,
	goodsRepo item.IItemRepo) cart.ICartAggregateRoot {
	unix := time.Now().Unix()
	if cartCode == "" {
		cartCode = domain.GenerateCartCode(unix, time.Now().Nanosecond())
	}
	value := &cart.NormalCart{
		CartCode:   cartCode,
		DeliverId:  0,
		BuyerId:    int64(buyerId),
		PaymentOpt: 1,
		CreateTime: unix,
		UpdateTime: unix,
		Items:      []*cart.NormalCartItem{},
	}
	return NewNormalCart(value, rep, memberRepo, goodsRepo)
}

func (c *cartImpl) init() cart.ICartAggregateRoot {
	// 获取购物车项
	if c.GetAggregateRootId() > 0 {
		if c.value.Items == nil {
			c.value.Items = c.rep.SelectNormalCartItem("cart_id= $1 ORDER BY id DESC",
				c.GetAggregateRootId())
		}
	}
	if c.value.Items == nil {
		c.value.Items = []*cart.NormalCartItem{}
	}
	// 初始化购物车的信息
	if c.value != nil && c.value.Items != nil {
		c.setAttachGoodsInfo(c.value.Items)
		arr := c.value.Items
		c.value.Items = make([]*cart.NormalCartItem, 0)
		for _, v := range arr {
			if v.Sku != nil {
				c.value.Items = append(c.value.Items, v)
			}
		}
	}
	return c
}

// 购物车种类
func (c *cartImpl) Kind() cart.Kind {
	return cart.KNormal
}

// 获取买家编号
func (c *cartImpl) BuyerId() int64 {
	return c.value.BuyerId
}

// Bind 绑定买家
func (c *cartImpl) Bind(buyerId int) error {
	if c.value.BuyerId > 0 {
		return errors.New("cart has already bind buyer")
	}
	c.value.BuyerId = int64(buyerId)
	_, err := c.Save()
	return err
}

func (c *cartImpl) Clone() cart.ICartAggregateRoot {
	panic("implement me")
}

func (c *cartImpl) Prepare() error {
	return c.check()
}

// 检查购物车(仅结算商品)
func (c *cartImpl) check() error {
	if c.value == nil || len(c.value.Items) == 0 {
		return cart.ErrEmptyShoppingCart
	}
	for _, v := range c.value.Items {
		if v.Checked == 1 {
			it := c.goodsRepo.GetItem(v.ItemId)
			if it == nil {
				return item.ErrNoSuchItem // 没有商品
			}
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
	}
	return nil
}

// 获取商品的快照列表
func (c *cartImpl) getSnapshotsMap(items []*cart.NormalCartItem) map[int64]*item.Snapshot {
	if c.snapMap == nil && items != nil {
		l := len(items)
		c.snapMap = make(map[int64]*item.Snapshot, l)
		if l > 0 {
			var ids = make([]int64, l)
			for i, v := range items {
				ids[i] = v.ItemId
			}
			snapList := c.goodsRepo.GetSnapshots(ids)
			for _, v := range snapList {
				v2 := v
				c.snapMap[v.ItemId] = &v2
			}
		}
	}
	return c.snapMap
}

func (c *cartImpl) getBuyerLevelId() int {
	if c.value.BuyerId > 0 {
		m := c.memberRepo.GetMember(c.value.BuyerId)
		if m != nil {
			return m.GetValue().Level
		}
	}
	return 0
}

func (c *cartImpl) setItemInfo(snap *item.GoodsItem, level int) {
	// 设置会员价
	if level > 0 {
		gds := c.goodsRepo.CreateItem(snap)
		snap.Price = gds.GetPromotionPrice(level)
	}
}

// 设置附加的商品信息
func (c *cartImpl) setAttachGoodsInfo(items []*cart.NormalCartItem) {
	list := c.getSnapshotsMap(items)
	if list == nil {
		return
	}
	var sku *item.Sku
	for _, v := range items {
		in := list[v.ItemId]
		if in == nil {
			continue
		}
		it := c.goodsRepo.GetItem(v.ItemId)
		if it == nil {
			continue
		}
		if v.SkuId > 0 {
			sku = it.GetSku(v.SkuId)
		} else {
			iv := it.GetValue()
			sku = &item.Sku{
				ProductId:   in.ProductId,
				ItemId:      v.ItemId,
				Title:       in.Title,
				Image:       in.Image,
				SpecData:    "",
				SpecWord:    "",
				Code:        in.Code,
				OriginPrice: in.OriginPrice,
				Price:       in.Price,
				Cost:        in.Cost,
				Weight:      in.Weight,
				Bulk:        in.Bulk,
				Stock:       iv.StockNum,
				SaleNum:     iv.SaleNum,
			}
		}
		v.Sku = item.ParseSkuMedia(it.GetValue(), sku)
	}
}

// 获取聚合根编号
func (c *cartImpl) GetAggregateRootId() int32 {
	return c.value.Id
}

// 获取购物车数据
func (c *cartImpl) Value() cart.NormalCart {
	return *c.value
}

// 获取商品集合
func (c *cartImpl) Items() []*cart.NormalCartItem {
	return c.getItems()
}

func (c *cartImpl) getItems() []*cart.NormalCartItem {
	return c.value.Items
}

// Put 添加项
func (c *cartImpl) Put(itemId, skuId int64, num int32, checkOnly bool) error {
	_, err := c.put(itemId, skuId, num, checkOnly)
	return err
}

// 检查是否有库存
func (c *cartImpl) checkItemStock(itemId int64, skuId int64) (*item.GoodsItem, *item.Sku, int32, error) {
	var sku *item.Sku

	it := c.goodsRepo.GetItem(itemId)
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
	// 判断商品SkuId
	if skuId > 0 {
		sku = it.GetSku(skuId)
		if sku == nil {
			return &iv, nil, stock, item.ErrNoSuchSku
		}
		stock = sku.Stock
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
func (c *cartImpl) put(itemId, skuId int64, num int32, checkOnly bool) (*cart.NormalCartItem, error) {
	var err error
	if c.value.Items == nil {
		c.value.Items = []*cart.NormalCartItem{}
	}
	// 检查库存
	iv, sku, stock, err := c.checkItemStock(itemId, skuId)
	if err != nil {
		return nil, err
	}
	var dst *cart.NormalCartItem
	// 添加数量
	for _, v := range c.value.Items {
		if checkOnly {
			// 取消掉其他要结算的商品
			v.Checked = 0
		}
		if v.ItemId == itemId && v.SkuId == skuId {
			dst = v
			finalQuantity := v.Quantity + num
			if checkOnly {
				// 立即购买
				finalQuantity = num
				v.Checked = 1
			}
			if finalQuantity > stock {
				return v, item.ErrOutOfStock // 库存不足
			}
			v.Quantity = finalQuantity
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
	v := &cart.NormalCartItem{
		CartId:   c.GetAggregateRootId(),
		VendorId: iv.VendorId,
		ShopId:   iv.ShopId,
		ItemId:   iv.Id,
		SkuId:    skuId,
		Quantity: num,
		Sku:      item.ParseSkuMedia(*iv, sku),
		Checked:  1,
	}
	c.value.Items = append(c.value.Items, v)
	return v, err
}

// ResetQuantity 重置数量
func (c *cartImpl) ResetQuantity(itemId, skuId int64, quantity int32) error {
	if c.value.Items == nil {
		return cart.ErrEmptyShoppingCart
	}

	var dst *cart.NormalCartItem
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

// GetItem 获取项
func (c *cartImpl) GetItem(itemId, skuId int64) *cart.NormalCartItem {
	if c.value != nil && c.value.Items != nil {
		for _, v := range c.value.Items {
			if v.ItemId == itemId && v.SkuId == skuId {
				return v
			}
		}
	}
	return nil
}

// Code 获取购物车编码
func (c *cartImpl) Code() string {
	return c.value.CartCode
}

// Combine 合并购物车，并返回新的购物车
func (c *cartImpl) Combine(ic cart.ICartAggregateRoot) cart.ICartAggregateRoot {
	if ic.Kind() != cart.KNormal {
		panic("only retail cart can be combine!")
	}
	if id := ic.GetAggregateRootId(); id != c.GetAggregateRootId() {
		rc := ic.(cart.INormalCart)
		for _, v := range rc.Items() {
			if it, err := c.put(v.ItemId,
				v.SkuId, v.Quantity, false); err == nil {
				if v.Checked == 1 {
					it.Checked = 1
				}
			}
		}
		err := ic.Destroy() //合并后,需销毁购物车
		if err != nil {
			log.Println("[ GO2O][ ERROR]: combine cart failed: ", err.Error())
		}
	}
	c.snapMap = nil //clean
	return c
}

// SignItemChecked 标记商品结算
func (c *cartImpl) SignItemChecked(items []*cart.ItemPair) error {
	mp := c.getItems()
	// 遍历购物车商品，默认不结算。
	for _, item := range mp {
		item.Checked = 0
		// 如果传入结算商品信息，则标记购物车项结算状态
		for _, v := range items {
			if v.SkuId == item.SkuId && v.ItemId == item.ItemId {
				item.Checked = v.Checked
				break
			}
		}
	}
	return c.check()
}

// Save 保存购物车
func (c *cartImpl) Save() (int32, error) {
	c.value.UpdateTime = time.Now().Unix()
	id, err := c.rep.SaveShoppingCart(c.value)
	c.value.Id = id
	if c.value.Items != nil {
		for _, v := range c.value.Items {
			if v.Quantity <= 0 {
				c.rep.RemoveCartItem(v.Id)
			} else {
				v.CartId = c.GetAggregateRootId()
				v.Id, err = c.rep.SaveCartItem(v)
			}
		}
	}
	return id, err
}

// 获取勾选的商品
func (c *cartImpl) CheckedItems(checked map[int64][]int64) []*cart.ItemPair {
	items := []*cart.ItemPair{}
	if checked != nil {
		for _, v := range c.value.Items {
			arr, ok := checked[int64(v.ItemId)]
			if ok {
				continue
			}
			for _, skuId := range arr {
				if skuId == int64(v.SkuId) {
					items = append(items, &cart.ItemPair{
						ItemId:   int64(v.ItemId),
						SkuId:    skuId,
						SellerId: v.VendorId,
						Quantity: v.Quantity,
					})
				}
			}
		}
	}
	return items
}

// 释放购物车,如果购物车的商品全部结算,则返回true
func (c *cartImpl) Release(_checked map[int64][]int64) bool {
	var checked []int
	for i, v := range c.value.Items {
		if v.Checked == 1 {
			checked = append(checked, i)
		}
	}
	// 如果为部分结算,则移除商品(更新数量为0)并返回false
	if len(checked) < len(c.value.Items) {
		for _, i := range checked {
			v := c.value.Items[i]
			c.ResetQuantity(v.ItemId, v.SkuId, 0)
		}
		c.Save()
		return false
	}
	return true
}

// 销毁购物车
func (c *cartImpl) Destroy() (err error) {
	c.snapMap = nil //clean
	if err = c.rep.EmptyCartItems(c.GetAggregateRootId()); err == nil {
		return c.rep.DeleteCart(c.GetAggregateRootId())
	}
	return err
}
