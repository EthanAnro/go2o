/**
 * Copyright 2015 @ 56x.net.
 * name : express_test.go
 * author : jarryliu
 * date : 2016-07-05 19:16
 * description :
 * history :
 */
package domain

import (
	"testing"

	expressImpl "github.com/ixre/go2o/core/domain/express"
	"github.com/ixre/go2o/core/domain/interface/express"
	"github.com/ixre/go2o/core/inject"
)

func TestExpressTemplateImpl_Save(t *testing.T) {
	valRepo := inject.GetValueRepo()
	rep := inject.GetExpressRepo()

	list := rep.GetExpressProviders()
	for _, v := range list {
		t.Log(v.Id, "=>", v.Name, v.Code)
	}

	// 用户的快递设置
	u := expressImpl.NewUserExpress(104, rep, valRepo)
	//创建快递模板
	tpl := u.CreateTemplate(&express.ExpressTemplate{
		Name:      "普通快递运费模板",
		IsFree:    0,
		Basis:     1,
		FirstUnit: 2,
		FirstFee:  10,
		AddUnit:   1,
		AddFee:    10,
		Enabled:   1,
	})
	// 保存模板
	id, err := tpl.Save()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	cul := u.CreateCalculator()
	cul.Add(id, 3)
	cul.Calculate("")

	// 计算运费
	t.Log("快递运费为:", cul.Total())
	// 删除模板
	err = u.DeleteTemplate(id)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
