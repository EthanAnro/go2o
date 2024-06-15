/**
 * Copyright 2015 @ 56x.net.
 * name : aftersales_query
 * author : jarryliu
 * date : 2016-07-18 19:27
 * description :
 * history :
 */
package query

import (
	"database/sql"
	"log"

	afterSales "github.com/ixre/go2o/core/domain/interface/aftersales"
	"github.com/ixre/go2o/core/dto"
	"github.com/ixre/gof/db"
)

type AfterSalesQuery struct {
	db.Connector
}

func NewAfterSalesQuery(db db.Connector) *AfterSalesQuery {
	return &AfterSalesQuery{
		Connector: db,
	}
}

// 获取分页售后单
func (a *AfterSalesQuery) QueryPagerAfterSalesOrderOfMember(memberId int64, begin,
	size int, where string) (int, []*dto.PagedMemberAfterSalesOrder) {
	list := []*dto.PagedMemberAfterSalesOrder{}
	total := 0
	if len(where) > 0 {
		where = " AND " + where
	}
	err := a.ExecScalar(`SELECT COUNT(1) FROM sale_after_order ao
	INNER JOIN item_trade_snapshot sn ON sn.id = ao.snapshot_id
	WHERE ao.buyer_id= $1 `+where, &total, memberId)
	if total > 0 {
		err = a.Query(`SELECT ao.id,ao.type,ao.order_no,ao.vendor_id,mch.name as vendor_name,
 ao.snapshot_id,ao.quantity,sn.sku_id,sn.goods_title,sn.img,ao.status,
 ao.create_time,ao.update_time FROM sale_after_order ao
INNER JOIN mch_merchant mch ON ao.vendor_id = mch.id
INNER JOIN item_trade_snapshot sn ON sn.id = ao.snapshot_id
WHERE ao.buyer_id= $1 ORDER BY ao.create_time DESC LIMIT $3 OFFSET $2`, func(rs *sql.Rows) {
			for rs.Next() {
				e := &dto.PagedMemberAfterSalesOrder{}
				rs.Scan(&e.Id, &e.Type, &e.OrderNo, &e.VendorId, &e.VendorName,
					&e.SnapshotId, &e.Quantity, &e.SkuId, &e.GoodsTitle,
					&e.GoodsImage, &e.Status, &e.CreateTime, &e.UpdateTime)
				e.StatusText = afterSales.Stat(e.Status).String()
				list = append(list, e)
			}
		}, memberId, begin, size)
	}
	if err != nil {
		log.Println("[ Query][ Error]:", err.Error())
	}
	return total, list
}

// 获取分页售后单
func (a *AfterSalesQuery) QueryPagerAfterSalesOrderOfVendor(vendorId int64, begin,
	size int, where string) (int, []*dto.PagedVendorAfterSalesOrder) {
	var list []*dto.PagedVendorAfterSalesOrder
	total := 0
	if len(where) > 0 {
		where = " AND " + where
	}
	a.ExecScalar(`SELECT COUNT(1) FROM sale_after_order ao
	INNER JOIN item_trade_snapshot sn ON sn.id = ao.snapshot_id
	WHERE ao.vendor_id= $1 `+where, &total, vendorId)

	if total > 0 {
		a.Query(`SELECT ao.id,ao.type,ao.order_no,ao.buyer_id,mp.name as buyer_name,
 ao.snapshot_id,ao.quantity,sn.sku_id,sn.goods_title,sn.img,ao.status,
 ao.create_time,ao.update_time FROM sale_after_order ao
INNER JOIN mm_profile mp ON mp.member_id = ao.buyer_id
INNER JOIN item_trade_snapshot sn ON sn.id = ao.snapshot_id
WHERE ao.vendor_id= $1 `+where+" ORDER BY id DESC LIMIT $3 OFFSET $2", func(rs *sql.Rows) {
			for rs.Next() {
				e := &dto.PagedVendorAfterSalesOrder{}
				rs.Scan(&e.Id, &e.Type, &e.OrderNo, &e.BuyerId, &e.BuyerName,
					&e.SnapshotId, &e.Quantity, &e.SkuId, &e.GoodsTitle,
					&e.GoodsImage, &e.Status, &e.CreateTime, &e.UpdateTime)
				e.StatusText = afterSales.Stat(e.Status).String()
				list = append(list, e)
			}
		}, vendorId, begin, size)
	}
	return total, list
}
