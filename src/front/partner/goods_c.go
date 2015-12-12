/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2014-02-05 21:53
 * description :
 * history :
 */
package partner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jsix/gof"
	gfmt "github.com/jsix/gof/util/fmt"
	"github.com/jsix/gof/web"
	"github.com/jsix/gof/web/mvc"
	"go2o/src/cache"
	"go2o/src/core/domain/interface/sale"
	"go2o/src/core/infrastructure/format"
	"go2o/src/core/service/dps"
	"go2o/src/core/variable"
	"html/template"
	"net/url"
	"strconv"
	"strings"
)

var _ mvc.Filter = new(goodsC)

type goodsC struct {
	*baseC
}

//商品列表
func (this *goodsC) Item_list(ctx *echox.Context)error{
	r, w := ctx.Request, ctx.Response
	r.ParseForm()

	cateOpts := cache.GetDropOptionsOfCategory(this.GetPartnerId(ctx))
	ctx.App.Template().Execute(w, gof.TemplateDataMap{
		"cate_opts":  template.HTML(cateOpts),
		"no_pic_url": format.GetGoodsImageUrl(""),
	}, "views/partner/goods/item_list.html")
}

//商品选择
func (this *goodsC) Goods_select(ctx *echox.Context)error{
	r, w := ctx.Request, ctx.Response
	r.ParseForm()
	cateOpts := cache.GetDropOptionsOfCategory(this.GetPartnerId(ctx))
	ctx.App.Template().Execute(w, gof.TemplateDataMap{
		"cate_opts":  template.HTML(cateOpts),
		"no_pic_url": format.GetGoodsImageUrl(""),
	}, "views/partner/goods/goods_select.html")
}

func (this *goodsC) Create(ctx *echox.Context)error{
	partnerId := this.GetPartnerId(ctx)
	shopChks := cache.GetShopCheckboxs(partnerId, "")
	cateOpts := cache.GetDropOptionsOfCategory(partnerId)

	e := &sale.ValueItem{
		Image: ctx.App.Config().GetString(variable.NoPicPath),
	}
	js, _ := json.Marshal(e)

	ctx.App.Template().Execute(ctx.Response, gof.TemplateDataMap{
		"entity":    template.JS(js),
		"shop_chk":  template.HTML(shopChks),
		"cate_opts": template.HTML(cateOpts),
	},
		"views/partner/goods/create_goods.html")
}

func (this *goodsC) Edit(ctx *echox.Context)error{
	partnerId := this.GetPartnerId(ctx)
	r, w := ctx.Request, ctx.Response
	var e *sale.ValueItem
	id, _ := strconv.Atoi(r.URL.Query().Get("item_id"))
	e = dps.SaleService.GetValueItem(partnerId, id)
	if e == nil {
		w.Write([]byte("商品不存在"))
		return
	}
	e.Description = ""
	js, _ := json.Marshal(e)

	shopChks := cache.GetShopCheckboxs(partnerId, e.ApplySubs)
	cateOpts := cache.GetDropOptionsOfCategory(partnerId)

	ctx.App.Template().Execute(w,
		gof.TemplateDataMap{
			"entity":    template.JS(js),
			"shop_chk":  template.HTML(shopChks),
			"cate_opts": template.HTML(cateOpts),
		},
		"views/partner/goods/update_goods.html")
}

// 保存商品描述
func (this *goodsC) Item_info(ctx *echox.Context)error{
	partnerId := this.GetPartnerId(ctx)
	r, w := ctx.Request, ctx.Response
	var e *sale.ValueItem
	id, _ := strconv.Atoi(r.URL.Query().Get("item_id"))
	e = dps.SaleService.GetValueItem(partnerId, id)
	if e == nil {
		w.Write([]byte("商品不存在"))
		return
	}
	ctx.App.Template().Execute(w, gof.TemplateDataMap{
		"item_id":   e.Id,
		"item_info": template.HTML(e.Description),
	}, "views/partner/goods/item_info.html")
}

func (this *goodsC) Save_item_info_post(ctx *echox.Context)error{
	partnerId := this.GetPartnerId(ctx)
	r := ctx.Request
	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("ItemId"))
	info := r.FormValue("Info")

	var result gof.Message
	err := dps.SaleService.SaveItemInfo(partnerId, id, info)

	if err != nil {
		result.Message = err.Error()
	} else {
		result.Result = true
	}

	return ctx.JSON(http.StatusOK,result)
}

func (this *goodsC) SaveItem_post(ctx *echox.Context)error{
	partnerId := this.GetPartnerId(ctx)
	r := ctx.Request
	var result gof.Message
	r.ParseForm()

	e := sale.ValueItem{}
	web.ParseFormToEntity(r.Form, &e)

	e.State = 1 //todo: 暂时使用

	id, err := dps.SaleService.SaveItem(partnerId, &e)

	if err != nil {
		result.Message = err.Error()
	} else {
		result.Result = true
		result.Data = id
	}
	return ctx.JSON(http.StatusOK,result)
}

func (this *goodsC) Del_post(ctx *echox.Context)error{
	partnerId := this.GetPartnerId(ctx)
	r := ctx.Request
	var result gof.Message

	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("id"))
	err := dps.SaleService.DeleteGoods(partnerId, id)

	if err != nil {
		result.Message = err.Error()
	} else {
		result.Result = true
	}
	return ctx.JSON(http.StatusOK,result)
}

func (this *goodsC) Del_item_post(ctx *echox.Context)error{
	partnerId := this.GetPartnerId(ctx)
	r := ctx.Request
	var result gof.Message

	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("id"))
	err := dps.SaleService.DeleteItem(partnerId, id)

	if err != nil {
		result.Message = err.Error()
	} else {
		result.Result = true
	}
	return ctx.JSON(http.StatusOK,result)
}

func (this *goodsC) SetSaleTag(ctx *echox.Context)error{
	r := ctx.Request
	r.ParseForm()
	partnerId := this.GetPartnerId(ctx)
	goodsId, _ := strconv.Atoi(r.URL.Query().Get("id"))

	var tags []*sale.ValueSaleTag = dps.SaleService.GetAllSaleTags(partnerId)
	tagsHtml := getSaleTagsCheckBoxHtml(tags)

	var chkTags []*sale.ValueSaleTag = dps.SaleService.GetItemSaleTags(partnerId, goodsId)
	strArr := make([]string, len(chkTags))
	for i, v := range chkTags {
		strArr[i] = strconv.Itoa(v.Id)
	}

	tagVal := strings.Join(strArr, ",")

	ctx.App.Template().Execute(ctx.Response, gof.TemplateDataMap{
		"goodsId":  goodsId,
		"tagsHtml": template.HTML(tagsHtml),
		"tagValue": tagVal,
	}, "views/partner/goods/set_sale_tag.html")
}

func (this *goodsC) SaveGoodsSTag_post(ctx *echox.Context)error{
	r := ctx.Request
	var result gof.Message
	goodsId, err := strconv.Atoi(r.FormValue("GoodsId"))
	if err == nil {
		tags := strings.Split(r.FormValue("SaleTags"), ",")
		var ids []int = []int{}
		for _, v := range tags {
			if i, err := strconv.Atoi(v); err == nil {
				ids = append(ids, i)
			}
		}

		partnerId := this.GetPartnerId(ctx)
		err = dps.SaleService.SaveItemSaleTags(partnerId, goodsId, ids)
	}

	if err != nil {
		result.Message = err.Error()
	} else {
		result.Result = true
	}
	return ctx.JSON(http.StatusOK,result)
}

func (this *goodsC) ItemCtrl(ctx *echox.Context)error{

	itemId, _ := strconv.Atoi(ctx.Request.URL.Query().Get("item_id"))
	ctx.App.Template().Execute(ctx.Response, gof.TemplateDataMap{
		"item_id": itemId,
	}, "views/partner/goods/item_ctrl.html")
}

func (this *goodsC) LvPrice(ctx *echox.Context)error{
	partnerId := this.GetPartnerId(ctx)
	//todo: should be goodsId
	itemId, _ := strconv.Atoi(ctx.Request.URL.Query().Get("item_id"))
	goods := dps.SaleService.GetGoodsBySku(partnerId, itemId, 0)
	lvs := dps.PartnerService.GetMemberLevels(partnerId)
	var prices []*sale.MemberPrice = dps.SaleService.GetGoodsLevelPrices(partnerId, goods.GoodsId)

	var buf *bytes.Buffer = bytes.NewBufferString("")

	var fmtFunc = func(level int, levelName string, id int, price float32, enabled int) {
		buf.WriteString(fmt.Sprintf(`
		<tr>
                <td><input type="hidden" field="Id_%d" value="%d"/>
                    %s</td>
                <td align="center"><input type="number" field="Price_%d" value="%s"/></td>
                <td align="center"><input type="checkbox" field="Enabled_%d" %s/></td>
            </tr>
		`, level, id, levelName, level, format.FormatFloat(price), level,
			gfmt.BoolString(enabled == 1, "checked=\"checked\"", "")))
	}

	var b bool
	for _, v := range lvs {
		b = false
		for _, v1 := range prices {
			if v.Value == v1.Level {
				fmtFunc(v.Value, v.Name, v1.Id, v1.Price, v1.Enabled)
				b = true
				break
			}
		}
		if !b {
			fmtFunc(v.Value, v.Name, 0, goods.SalePrice, 0)
		}
	}

	ctx.App.Template().Execute(ctx.Response, gof.TemplateDataMap{
		"goods":   goods,
		"setHtml": template.HTML(buf.String()),
	}, "views/partner/goods/level_price.html")
}

func (this *goodsC) LvPrice_post(ctx *echox.Context)error{
	ctx.Request.ParseForm()
	var form url.Values = ctx.Request.Form
	goodsId, err := strconv.Atoi(form.Get("goodsId"))
	if err != nil {
		return ctx.JSON(http.StatusOK,gof.Message{Message: err.Error()})
		return
	}

	var priceSet []*sale.MemberPrice = []*sale.MemberPrice{}
	var id int
	var price float64
	var lv int
	var enabled int

	for k, _ := range form {
		if strings.HasPrefix(k, "Id_") {
			if lv, err = strconv.Atoi(k[3:]); err == nil {
				id, _ = strconv.Atoi(form.Get(k))
				price, _ = strconv.ParseFloat(ctx.Request.FormValue(fmt.Sprintf("Price_%d", lv)), 32)
				if ctx.Request.FormValue(fmt.Sprintf("Enabled_%d", lv)) == "on" {
					enabled = 1
				} else {
					enabled = 0
				}

				priceSet = append(priceSet, &sale.MemberPrice{
					Id:      id,
					Level:   lv,
					GoodsId: goodsId,
					Price:   float32(price),
					Enabled: enabled,
				})
			} else {
				return ctx.JSON(http.StatusOK,gof.Message{Message: err.Error()})
				return
			}
		}
	}

	partnerId := this.GetPartnerId(ctx)
	err = dps.SaleService.SaveMemberPrices(partnerId, goodsId, priceSet)
	if err != nil {
		return ctx.JSON(http.StatusOK,gof.Message{Message: err.Error()})
	} else {
		return ctx.JSON(http.StatusOK,gof.Message{Result: true})
	}
}
