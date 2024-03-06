/**
 * Copyright 2015 @ 56x.net.
 * name : goods
 * author : jarryliu
 * date : 2016-06-28 23:54
 * description :
 * history :
 */
package item

import (
	"sort"

	promodel "github.com/ixre/go2o/core/domain/interface/pro_model"
	"github.com/ixre/go2o/core/domain/interface/product"
	"github.com/ixre/go2o/core/domain/interface/promotion"
	"github.com/ixre/go2o/core/domain/interface/valueobject"
	"github.com/ixre/go2o/core/infrastructure/domain"
)

const (
	// 普通商品
	ItemNormal int32 = 1
	// 批发商品
	ItemWholesale int32 = 2
)

const (
	// 仓库中的商品
	ShelvesInWarehouse int32 = 0
	// 已下架
	ShelvesDown int32 = 1
	// 已上架
	ShelvesOn int32 = 2
	// 已拒绝上架 (不允许上架)
	ShelvesIncorrect int32 = 3
)

// 商品标志
const (
	// 自营商品
	FlagSelfSales = 1
	// 免邮
	FlagFreeDelivery = 2
	// 推荐商品
	FlagRecommend = 4
	// 积分兑换商品
	FlagExchange = 8
	// 赠品
	FlagGift = 16
	// 分销商品
	FlagAffiliate = 32
	// 新品
	FlagNewOnShelve = 64
	// 热销商品
	FlagHotSales = 128
	// 平台配送
	FlagSelfDelivery = 256
)

var (
	ErrNoSuchItem = domain.NewError(
		"no_such_goods", "商品不存在")

	ErrIncorrectProductCategory = domain.NewError(
		"err_item_incorrect_product_category", "非法的商品分类")

	ErrItemWholesaleOff = domain.NewError(
		"err_item_wholesale_off", "商品已下架或待审核!")

	ErrLatestSnapshot = domain.NewError(
		"latest_snapshot", "已经是最新的快照")

	ErrNoSuchSnapshot = domain.NewError(
		"no_such_snapshot", "商品快照不存在")

	ErrNotOnShelves = domain.NewError(
		"not_on_shelves", "商品\"%s\"已下架")

	ErrGoodsMinProfitRate = domain.NewError(
		"err_goods_min_profit_rate", "商品利润率不能低于%s")
)

type (
	// IGoodsItemAggregateRoot 商品
	IGoodsItemAggregateRoot interface {
		// GetAggregateRootId 获取聚合根编号
		GetAggregateRootId() int64
		// GetValue 设置值
		GetValue() GoodsItem
		// GetPackedValue 获取包装过的商品信息
		GetPackedValue() *valueobject.Goods
		// SetValue 设置值
		SetValue(*GoodsItem) error
		// GrantFlag 标志赋值, 如果flag小于零, 则异或运算(去除)
		GrantFlag(flag int) error
		// SetSku 设置SKU
		SetSku(arr []*Sku) error
		// Save 保存
		Save() (int64, error)
		// Product 获取产品
		Product() product.IProduct
		// Images 获取商品图片
		Images() []string
		// SetImages 保存商品图片
		SetImages(images []string) error
		// Snapshot 商品快照
		Snapshot() *Snapshot
		// Wholesale 批发
		Wholesale() IWholesaleItem
		// SkuArray 获取SKU数组
		SkuArray() []*Sku
		// SpecArray 获取商品的规格
		SpecArray() promodel.SpecList
		// GetSku 获取SKU
		GetSku(skuId int64) *Sku
		// GetPromotions 获取促销信息
		GetPromotions() []promotion.IPromotion
		// GetPromotionPrice 获取促销价
		GetPromotionPrice(level int) int64
		// GetLevelPrice 获取会员价销价,返回是否有会原价及价格
		GetLevelPrice(level int) (bool, int64)
		// GetPromotionDescribe 获取促销描述
		GetPromotionDescribe() map[string]string
		// GetLevelPrices 获取会员价
		GetLevelPrices() []*MemberPrice
		// SaveLevelPrice 保存会员价
		SaveLevelPrice(*MemberPrice) (int32, error)
		// IsOnShelves 是否上架
		IsOnShelves() bool
		// SetShelve 设置上架
		SetShelve(state int32, remark string) error
		// Review 审核
		Review(pass bool, remark string) error
		// Incorrect 标记为违规
		Incorrect(remark string) error
		// AddSalesNum 更新销售数量,扣减库存
		AddSalesNum(skuId int64, quantity int32) error
		// CancelSale 取消销售
		CancelSale(skuId int64, quantity int32, orderNo string) error
		// TakeStock 占用库存
		TakeStock(skuId int64, quantity int32) error
		// ReleaseStock 释放库存
		ReleaseStock(skuId int64, quantity int32) error
		//// 生成快照
		//GenerateSnapshot() (int64, error)
		//
		//// 获取最新的快照
		//GetLatestSnapshot() *goods.GoodsSnapshot
		// 回收商品
		Recycle() error
		// 从回收站中撤回
		RecycleRevert() error
		// Destroy 删除商品
		Destroy() error
	}

	// IWholesaleItem 商品批发
	IWholesaleItem interface {
		// GetDomainId 获取领域编号
		GetDomainId() int64
		// CanWholesale 是否允许批发
		CanWholesale() bool
		// Save 保存
		Save() (int32, error)

		// GetJsonDetailData 获取详细信息
		GetJsonDetailData() []byte

		// IsOnShelves 是否上架
		IsOnShelves() bool
		// SetShelve 设置上架
		SetShelve(state int32, remark string) error
		// Review 审核
		Review(pass bool, remark string) error
		// Incorrect 标记为违规
		Incorrect(remark string) error

		// GetWholesaleDiscount 根据商品金额获取折扣
		GetWholesaleDiscount(groupId int32, amount int32) float64
		// GetItemDiscount 获取全部批发折扣
		GetItemDiscount(groupId int32) []*WsItemDiscount
		// SaveItemDiscount 保存批发折扣
		SaveItemDiscount(groupId int32, arr []*WsItemDiscount) error

		// GetWholesalePrice 获取批发价格
		GetWholesalePrice(skuId int64, quantity int32) int64
		// GetSkuPrice 根据SKU获取价格设置
		GetSkuPrice(skuId int64) []*WsSkuPrice
		// SaveSkuPrice 保存批发SKU价格设置
		SaveSkuPrice(skuId int64, arr []*WsSkuPrice) error
	}

	IItemWholesaleRepo interface {
		// GetWsItem Get WsItem
		GetWsItem(primary interface{}) *WsItem
		// SelectWsItem Select WsItem
		SelectWsItem(where string, v ...interface{}) []*WsItem
		// SaveWsItem Save WsItem
		SaveWsItem(v *WsItem, create bool) (int, error)
		// DeleteWsItem Delete WsItem
		DeleteWsItem(primary interface{}) error
		// BatchDeleteWsItem Batch Delete WsItem
		BatchDeleteWsItem(where string, v ...interface{}) (int64, error)

		// GetWsSkuPrice Get WsSkuPrice
		GetWsSkuPrice(primary interface{}) *WsSkuPrice
		// SelectWsSkuPrice Select WsSkuPrice
		SelectWsSkuPrice(where string, v ...interface{}) []*WsSkuPrice
		// SaveWsSkuPrice Save WsSkuPrice
		SaveWsSkuPrice(v *WsSkuPrice) (int, error)
		// DeleteWsSkuPrice Delete WsSkuPrice
		DeleteWsSkuPrice(primary interface{}) error
		// BatchDeleteWsSkuPrice Batch Delete WsSkuPrice
		BatchDeleteWsSkuPrice(where string, v ...interface{}) (int64, error)

		// GetWsItemDiscount Get WsItemDiscount
		GetWsItemDiscount(primary interface{}) *WsItemDiscount
		// SelectWsItemDiscount Select WsItemDiscount
		SelectWsItemDiscount(where string, v ...interface{}) []*WsItemDiscount
		// SaveWsItemDiscount Save WsItemDiscount
		SaveWsItemDiscount(v *WsItemDiscount) (int, error)
		// DeleteWsItemDiscount Delete WsItemDiscount
		DeleteWsItemDiscount(primary interface{}) error
		// BatchDeleteWsItemDiscount Batch Delete WsItemDiscount
		BatchDeleteWsItemDiscount(where string, v ...interface{}) (int64, error)
	}

	// WsItem 批发商品
	WsItem struct {
		// 编号
		ID int64 `db:"id" pk:"yes" auto:"yes"`
		// 运营商编号
		VendorId int64 `db:"vendor_id"`
		// 商品编号
		ItemId int64 `db:"item_id"`
		// 价格
		Price int64 `db:"price"`
		// 价格区间
		PriceRange string `db:"price_range"`
		// 上架状态
		ShelveState int32 `db:"shelve_state"`
		// 是否审核通过
		ReviewStatus int32 `db:"review_status"`
		// 审核备注
		ReviewRemark string `db:"review_remark"`
	}

	// 商品批发价
	WsSkuPrice struct {
		// 编号
		ID int32 `db:"id" pk:"yes" auto:"yes"`
		// 商品编号
		ItemId int64 `db:"item_id"`
		// SKU编号
		SkuId int64 `db:"sku_id"`
		// 需要数量以上
		RequireQuantity int32 `db:"require_quantity"`
		// 批发价
		WholesalePrice int64 `db:"wholesale_price"`
	}

	// 批发商品折扣
	WsItemDiscount struct {
		// 编号
		ID int32 `db:"id" pk:"yes" auto:"yes"`
		// 商品编号
		ItemId int64 `db:"item_id"`
		// 客户分组
		BuyerGid int32 `db:"buyer_gid"`
		// 要求金额，默认为0
		RequireAmount int32 `db:"require_amount"`
		// 折扣率
		DiscountRate float64 `db:"discount_rate"`
	}
)

type (
	// IItemRepo 商品仓储
	IItemRepo interface {
		// SkuService 获取SKU服务
		SkuService() ISkuService
		// SnapshotService 获取快照服务
		SnapshotService() ISnapshotService

		// CreateItem 创建商品
		CreateItem(v *GoodsItem) IGoodsItemAggregateRoot

		// GetItem 获取商品
		GetItem(itemId int64) IGoodsItemAggregateRoot

		// GetValueGoods 获取商品
		GetValueGoods(itemId, skuId int64) *GoodsItem

		// GetItemImages  获取商品图片
		GetItemImages(itemId int64) []*Image
		// SaveItemImage 保存商品图片
		SaveItemImage(v *Image) (int, error)
		// DeleteItemImage 删除商品图片
		DeleteItemImage(id int64) error
		// 根据SKU-ID获取商品,SKU-ID为商品ID
		//todo: 循环引有,故为interface{}
		GetItemBySkuId(skuId int64) interface{}

		// GetValueGoodsById 获取商品
		GetValueGoodsById(goodsId int64) *GoodsItem

		// GetValueGoodsBySku 根据产品编号和SKU获取商品
		GetValueGoodsBySku(productId, skuId int64) *GoodsItem

		// SaveValueGoods 保存商品
		SaveValueGoods(*GoodsItem) (int64, error)

		// GetOnShelvesGoods 获取在货架上的商品
		GetOnShelvesGoods(mchId int64, start, end int,
			sortBy string) []*valueobject.Goods

		// GetGoodsByIds 根据编号获取商品
		GetGoodsByIds(ids ...int64) ([]*valueobject.Goods, error)

		// GetGoodSMemberLevelPrice 获取会员价
		GetGoodSMemberLevelPrice(goodsId int64) []*MemberPrice

		// SaveGoodSMemberLevelPrice 保存会员价
		SaveGoodSMemberLevelPrice(*MemberPrice) (int32, error)

		// RemoveGoodSMemberLevelPrice 移除会员价
		RemoveGoodSMemberLevelPrice(id int) error

		// SaveSnapshot 保存
		SaveSnapshot(*Snapshot) (int64, error)
		// DeleteSnapshot 删除商品快照
		DeleteSnapshot(itemId int64) error
		// GetSnapshots 根据指定商品快照
		GetSnapshots(skuIdArr []int64) []Snapshot

		// GetLatestSnapshot 获取最新的商品快照
		GetLatestSnapshot(itemId int64) *Snapshot

		// GetSalesSnapshot 获取指定的商品快照
		GetSalesSnapshot(id int64) *TradeSnapshot

		// GetSaleSnapshotByKey 根据Key获取商品快照
		GetSaleSnapshotByKey(key string) *TradeSnapshot

		// GetLatestSalesSnapshot 获取最新的商品销售快照
		GetLatestSalesSnapshot(itemId int64, skuId int64) *TradeSnapshot

		// DeleteItem 删除商品
		DeleteItem(itemId int) error
		// SaveSalesSnapshot 保存商品销售快照
		SaveSalesSnapshot(*TradeSnapshot) (int64, error)

		// GetItemSku Get ItemSku
		GetItemSku(primary interface{}) *Sku
		// SelectItemSku Select ItemSku
		SelectItemSku(where string, v ...interface{}) []*Sku
		// SaveItemSku Save ItemSku
		SaveItemSku(v *Sku) (int, error)
		// DeleteItemSku Delete ItemSku
		DeleteItemSku(primary interface{}) error
		// BatchDeleteItemSku Batch Delete ItemSku
		BatchDeleteItemSku(where string, v ...interface{}) (int64, error)
	}

	// GoodsItem 商品,临时改方便辨别
	GoodsItem struct {
		// 商品编号
		Id int64 `db:"id" pk:"yes" auto:"yes"`
		// 产品编号
		ProductId int64 `db:"product_id"`
		// 商品标志
		ItemFlag int `db:"item_flag"`
		// 分类编号
		CategoryId int32 `db:"cat_id"`
		// 供货商编号
		VendorId int64 `db:"vendor_id"`
		// 品牌编号(冗余)
		BrandId int32 `db:"brand_id"`
		// 店铺编号
		ShopId int64 `db:"shop_id"`
		// 店铺分类编号
		ShopCatId int32 `db:"shop_cat_id"`
		// 快递模板编号
		ExpressTid int32 `db:"express_tid"`
		// 商品标题
		Title string `db:"title"`
		// 短标题
		ShortTitle string `db:"short_title"`
		// 供货商编码
		Code string `db:"code"`
		// 主图
		Image string `db:"image"`
		// 视频介绍
		IntroVideo string `db:"intro_video"`
		// 销售价格区间
		PriceRange string `db:"price_range"`
		// 总库存
		StockNum int32 `db:"stock_num"`
		// 销售数量
		SaleNum int32 `db:"sale_num"`
		// SKU数量
		SkuNum int32 `db:"sku_num"`
		// 默认SKU编号
		SkuId int64 `db:"sku_id"`
		// 成本价
		Cost int64 `db:"cost"`
		// 销售价
		Price int64 `db:"price"`
		// 零售价
		OriginPrice int64 `db:"origin_price"`
		// 重量:克(g)
		Weight int32 `db:"weight"`
		// 体积:毫升(ml)
		Bulk int32 `db:"bulk"`
		// 商品购物保障
		SafeguardFlag int `db:"safeguard_flag"`
		// 是否上架
		ShelveState int32 `db:"shelve_state"`
		// 审核状态
		ReviewStatus int32 `db:"review_status"`
		// 审核备注
		ReviewRemark string `db:"review_remark"`
		// 排序序号
		SortNum int32 `db:"sort_num"`
		// 是否已被回收
		IsRecycle int `db:"is_recycle"`
		// 创建时间
		CreateTime int64 `db:"create_time"`
		// 更新时间
		UpdateTime int64 `db:"update_time"`
		// 促销价
		PromPrice int64 `db:"-"`
		// 规格项
		SkuArray []*Sku `db:"-"`
		// 图片
		Images []string `db:"-"`
	}

	// Image 产品图片
	Image struct {
		// 图片编号
		Id int64 `db:"id" pk:"yes" auto:"yes"`
		// 商品编号
		ItemId int64 `db:"item_id"`
		// 图片地址
		ImageUrl string `db:"image_url"`
		// 排列序号
		SortNum int `db:"sort_num"`
		// 创建时间
		CreateTime int64 `db:"create_time"`
	}

	// MemberPrice 会员价
	MemberPrice struct {
		// 编号
		Id int `db:"id" pk:"yes" auto:"yes"`
		// 商品编号
		GoodsId int64 `db:"goods_id"`
		// 等级
		LevelId int `db:"level"`
		// 价格
		Price int64 `db:"price"`
		// 限购数量
		MaxQuota int `db:"max_quota"`
		// 是否启用
		Enabled int `db:"enabled"`
	}

	AffiliateRate struct {
		// 编号
		Id int `db:"id" pk:"yes" auto:"yes"`
		// 商品编号
		ItemId int `db:"item_id"`
		// 上一级比例
		RateR1 int `db:"rate_r1"`
		// 上二级比例
		RateR2 int `db:"rate_r2"`
		// 自定义比例
		RateC int `db:"rate_c"`
		// 历史上一级比例
		OriginRateR1 int `db:"origin_rate_r1"`
		// 历史上二级比例
		OriginRateR2 int `db:"origin_rate_r2"`
		// 历史自定义比例
		OriginRateC int `db:"origin_rate_c"`
	}
)

var _ sort.Interface = SkuList{}

type SkuList []*Sku

func (s SkuList) Len() int {
	return len(s)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s SkuList) Less(i, j int) bool {
	b := s[i].Price - s[j].Price
	return b < 0 || b == 0 && s[i].Id < s[j].Id
}

// Swap swaps the elements with indexes i and j.
func (s SkuList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
