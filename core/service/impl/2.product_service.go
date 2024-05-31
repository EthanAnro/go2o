package impl

import (
	"errors"

	"github.com/ixre/go2o/core/domain/interface/item"
	promodel "github.com/ixre/go2o/core/domain/interface/pro_model"
	"github.com/ixre/go2o/core/domain/interface/product"
	"github.com/ixre/go2o/core/infrastructure/format"
	"github.com/ixre/go2o/core/service/proto"
	"github.com/ixre/gof/types"
	"golang.org/x/net/context"
)

var _ proto.ProductServiceServer = new(productService)

// 产品服务
type productService struct {
	pmRepo      promodel.IProductModelRepo
	catRepo     product.ICategoryRepo
	productRepo product.IProductRepo
	serviceUtil
	proto.UnimplementedProductServiceServer
}

func NewProductService(pmRepo promodel.IProductModelRepo,
	catRepo product.ICategoryRepo,
	proRepo product.IProductRepo) proto.ProductServiceServer {
	return &productService{
		pmRepo:      pmRepo,
		catRepo:     catRepo,
		productRepo: proRepo,
	}
}

// GetModel 获取产品模型
func (p *productService) GetProductModel(_ context.Context, id *proto.ProductModelId) (*proto.SProductModel, error) {
	im := p.pmRepo.GetModel(int(id.Value))
	if im != nil {
		return p.fromProductModel(im), nil
	}
	return nil, product.ErrNoSuchProduct
}

// 　转换产品模型数据
func (p *productService) fromProductModel(im promodel.IProductModel) *proto.SProductModel {
	if im == nil {
		return nil
	}
	ret := p.parseModelDto(im.Value())
	// 绑定属性
	attrs := im.Attrs()
	ret.Attrs = make([]*proto.SProductAttr, len(attrs))
	for i, v := range attrs {
		attr := p.appendAttrItems(p.parseProductAttrDto(v), v.Items)
		ret.Attrs[i] = attr
	}
	// 绑定规格
	specList := im.Specs()
	ret.Specs = make([]*proto.SProductSpec, len(specList))
	for i, v := range specList {
		spec := p.appendSpecItems(p.parseSpecDto(v), v.Items)
		ret.Specs[i] = spec
	}
	// 绑定品牌
	brands := im.Brands()
	ret.Brands = make([]int64, len(brands))
	for i, v := range brands {
		ret.Brands[i] = int64(v.Id)
	}
	return ret
}

func (p *productService) appendAttrItems(attr *proto.SProductAttr, items []*promodel.AttrItem) *proto.SProductAttr {
	attr.Items = make([]*proto.SProductAttrItem, len(items))
	for i1, v1 := range items {
		attr.Items[i1] = p.parseProductAttrItemDto(v1)
	}
	return attr
}

func (p *productService) appendSpecItems(spec *proto.SProductSpec, items promodel.SpecItemList) *proto.SProductSpec {
	spec.Items = make([]*proto.SProductSpecItem, len(items))
	for i1, v1 := range items {
		spec.Items[i1] = p.parseProductSpecItemDto(v1)
	}
	return spec
}

// GetModels 获取产品模型
func (p *productService) GetModels(_ context.Context, r *proto.GetModelsRequest) (*proto.ProductModelListResponse, error) {
	list := p.pmRepo.SelectProModel("")
	arr := make([]*proto.SProductModel, 0)
	for _, v := range list {
		if v.Enabled == 0 {
			continue
		}
		arr = append(arr, p.parseModelDto(v))
	}
	return &proto.ProductModelListResponse{
		Value: arr,
	}, nil
}

// GetAttr 获取属性
func (p *productService) GetAttr(_ context.Context, id *proto.ProductAttrId) (*proto.SProductAttr, error) {
	v := p.pmRepo.AttrService().GetAttr(int(id.Value))
	if v != nil {
		attr := p.parseProductAttrDto(v)
		attr = p.appendAttrItems(attr, v.Items)
		return attr, nil
	}
	return nil, product.ErrNoSuchAttr
}

// GetAttrItem 获取属性项
func (p *productService) GetAttrItem(_ context.Context, id *proto.ProductAttrItemId) (*proto.SProductAttrItem, error) {
	it := p.pmRepo.GetAttrItem(id.Value)
	if it != nil {
		return p.parseProductAttrItemDto(it), nil
	}
	return nil, product.ErrNoSuchAttr
}

// GetBrands 获取所有产品品牌
func (p *productService) GetBrands(_ context.Context, _ *proto.Empty) (*proto.ProductBrandListResponse, error) {
	list := p.pmRepo.BrandService().AllBrands()
	arr := make([]*proto.SProductBrand, len(list))
	for i, v := range list {
		arr[i] = p.parseBrandDto(v)
	}
	return &proto.ProductBrandListResponse{
		Value: arr,
	}, nil
}

// 为分类绑定品牌
func (p *productService) appendCategoryBrands(ic product.IGlobCatService, v product.ICategory, cat *proto.SProductCategory) {
	brands := ic.RelationBrands(v.GetDomainId())
	cat.Brands = make([]*proto.SProductBrand, len(brands))
	for i1, v1 := range brands {
		cat.Brands[i1] = p.parseBrandDto(v1)
	}
}

// GetCategory 获取商品分类
func (p *productService) GetCategory(_ context.Context, req *proto.GetCategoryRequest) (*proto.SProductCategory, error) {
	ic := p.catRepo.GlobCatService()
	v := ic.GetCategory(int(req.CategoryId))
	if v != nil {
		cat := p.parseCategoryDto(v.GetValue())
		if req.WithBrand {
			p.appendCategoryBrands(ic, v, cat)
		}
		if req.WithModel {
			cat.Model = p.fromProductModel(v.GetModel())
		}
		return cat, nil
	}
	return nil, product.ErrNoSuchCategory
}

// DeleteCategory 删除分类
func (p *productService) DeleteCategory(_ context.Context, id *proto.Int64) (*proto.Result, error) {
	err := p.catRepo.GlobCatService().DeleteCategory(int(id.Value))
	return p.error(err), nil
}

// SaveCategory 保存分类
func (p *productService) SaveCategory(_ context.Context, category *proto.SaveProductCategoryRequest) (*proto.SaveProductCategoryResponse, error) {
	sl := p.catRepo.GlobCatService()
	var ca product.ICategory
	v := p.parseCategory(category)
	if v.Id > 0 {
		ca = sl.GetCategory(v.Id)
	} else {
		ca = sl.CreateCategory(v)
	}
	err := ca.SetValue(v)
	if err == nil {
		_, err = ca.Save()
	}
	if err != nil {
		return &proto.SaveProductCategoryResponse{Error: err.Error()}, nil
	}
	return &proto.SaveProductCategoryResponse{CategoryId: int64(ca.GetDomainId())}, nil
}

// GetProduct 获取产品值
func (p *productService) GetProduct(_ context.Context, id *proto.ProductId) (*proto.SProduct, error) {
	pro := p.productRepo.GetProduct(id.Value)
	if pro != nil {
		ret := p.parseProductDto(pro.GetValue())
		attrs := pro.Attr()
		ret.Attrs = make([]*proto.SProductAttrValue, len(attrs))
		for i, v := range attrs {
			ret.Attrs[i] = p.parseProductAttrValueDto(v)
		}
		return ret, nil
	}
	return nil, product.ErrNoSuchProduct
}

// SaveProduct 保存产品
func (p *productService) SaveProduct(_ context.Context, r *proto.SaveProductRequest) (*proto.SaveProductResponse, error) {
	var pro product.IProduct
	v := p.parseProduct(r)
	ret := &proto.SaveProductResponse{}
	if v.Id > 0 {
		pro = p.productRepo.GetProduct(v.Id)
		if pro == nil || pro.GetValue().VendorId != v.VendorId {
			ret.ErrCode = 1
			ret.ErrMsg = product.ErrNoSuchProduct.Error()
			return ret, nil
		}
	} else {
		pro = p.productRepo.CreateProduct(v)
	}
	// 保存
	err := pro.SetValue(v)
	if err == nil {
		// 保存属性
		if v.Attrs != nil {
			err = pro.SetAttr(v.Attrs)
		}
		if r.UpdateDescription {
			err = pro.SetDescribe(v.Description)
		}
		if err == nil {
			v.Id, err = pro.Save()
		}
	}
	if err != nil {
		ret.ErrMsg = err.Error()
		ret.ErrCode = 2
	}
	ret.ProductId = v.Id
	return ret, nil
}

// SaveProductInfo 保存货品描述
func (p *productService) SaveProductInfo(_ context.Context, r *proto.ProductInfoRequest) (*proto.Result, error) {
	pro := p.productRepo.GetProduct(r.ProductId)
	var err error
	if pro == nil {
		err = product.ErrNoSuchProduct
	} else {
		err = pro.SetDescribe(r.Info)
	}
	return p.error(err), nil
}

// SaveModel 保存产品模型
func (p *productService) SaveProductModel(_ context.Context, r *proto.SaveProductModelRequest) (*proto.Result, error) {
	var pm promodel.IProductModel
	v := p.parseProductModel(r)
	if v.Id > 0 {
		pm = p.pmRepo.GetModel(int(r.Id))
		if pm == nil {
			return p.error(errors.New("模型不存在")), nil
		}
	} else {
		pm = p.pmRepo.CreateModel(v)
	}
	err := pm.SetValue(v)
	if err == nil {
		// 保存属性
		if err == nil && r.UpdateAttrSpec {
			err = pm.SetAttrs(v.Attrs)
		}
		// 保存规格
		if err == nil && r.UpdateAttrSpec {
			err = pm.SetSpecs(v.Specs)
		}
		// 保存品牌
		if err == nil && !r.UpdateAttrSpec {
			err = pm.SetBrands(v.BrandArray)
		}
	}
	// 保存模型
	if err == nil {
		v.Id, err = pm.Save()
	}
	return p.result(err), nil
}

// DeleteModel_ 删除产品模型
func (p *productService) DeleteModel_(_ context.Context, id *proto.ProductModelId) (*proto.Result, error) {
	model := p.pmRepo.GetModel(int(id.Value))
	if model == nil {
		return p.result(errors.New("商品模型不存在")), nil
	}
	err := model.Destroy()
	return p.result(err), nil
}

// GetBrand 获取产品品牌
func (p *productService) GetBrand(_ context.Context, id *proto.Int64) (*proto.SProductBrand, error) {
	brand := p.pmRepo.BrandService().Get(int(id.Value))
	if brand != nil {
		return p.parseBrandDto(brand), nil
	}
	return nil, product.ErrNoSuchBrand
}

// SaveBrand Save 产品品牌
func (p *productService) SaveBrand(_ context.Context, brand *proto.SProductBrand) (*proto.Result, error) {
	v := p.parseBrand(brand)
	_, err := p.pmRepo.BrandService().SaveBrand(v)
	return p.result(err), nil
}

// DeleteBrand 删除产品品牌
func (p *productService) DeleteBrand(_ context.Context, id *proto.Int64) (*proto.Result, error) {
	err := p.pmRepo.BrandService().DeleteBrand(int(id.Value))
	return p.result(err), nil
}

// DeleteProduct 删除产品
func (p *productService) DeleteProduct(_ context.Context, r *proto.DeleteProductRequest) (*proto.Result, error) {
	var err error
	prod := p.productRepo.GetProduct(r.ProductId)
	if prod == nil || prod.GetValue().VendorId != r.SellerId {
		err = product.ErrNoSuchProduct
	} else {
		err = p.productRepo.DeleteProduct(r.ProductId)
		//todo: 删除商品
	}
	return p.error(err), nil
}

// FindParentCategory 获取分类包括所有的上级
func (p *productService) FindParentCategory(c context.Context, request *proto.CategoryIdRequest) (*proto.CategoriesResponse, error) {
	s := p.catRepo.GlobCatService()
	list := s.GetCategories()
	cat := s.GetCategory(int(request.CategoryId))
	arr := make([]*proto.SProductCategory, 0)
	if cat != nil {
		findParent := func(pid int64, arr []product.ICategory) int64 {
			for _, it := range arr {
				v := it.GetValue()
				if v.Id == int(pid) && v.ParentId > 0 {
					return int64(v.ParentId)
				}
			}
			return pid
		}

		for pid := request.CategoryId; pid > 0; {
			id := findParent(pid, list)
			if id == pid {
				break
			}
			arr = append([]*proto.SProductCategory{
				p.parseCategoryDto(s.GetCategory(int(id)).GetValue()),
			}, arr...)
			pid = id
		}
		arr = append(arr, p.parseCategoryDto(cat.GetValue()))
	}
	return &proto.CategoriesResponse{List: arr}, nil
}

// GetCategoryTreeNode 分类
func (p *productService) GetCategoryTreeNode(_ context.Context, r *proto.CategoryTreeRequest) (*proto.CategoryTreeResponse, error) {
	arr := p.catRepo.GlobCatService().GetCategories()
	// 懒加载,只加载一级,并设置IsLeaf
	if r.Lazy {
		roots := p.lazyLoadChildren(int(r.ParentId), arr, r)
		return &proto.CategoryTreeResponse{
			Value: roots,
		}, nil
	}
	root := &proto.SCategoryTree{
		Id: r.ParentId,
	}
	p.walkLoadCategory(root, arr, r)
	return &proto.CategoryTreeResponse{
		Value: root.Children,
	}, nil
}

// 是否为叶子节点
func (p *productService) testIsLeaf(parentId int, categories []product.ICategory, req *proto.CategoryTreeRequest) bool {
	// 遍历子分类
	for _, v := range categories {
		cat := v.GetValue()
		if cat.ParentId == parentId && p.testWalkCondition(req, cat) {
			return false
		}
	}
	return true
}

// 懒加载下级分类
func (p *productService) lazyLoadChildren(parentId int, categories []product.ICategory,
	req *proto.CategoryTreeRequest) []*proto.SCategoryTree {
	var arr []*proto.SCategoryTree
	// 遍历子分类
	for _, v := range categories {
		// if v.GetValue().Name == "礼品鲜花" {
		// 	print(fmt.Sprintf("%#v", v.GetValue()))
		// }
		if cat := v.GetValue(); cat.ParentId == parentId &&
			p.testWalkCondition(req, cat) {
			if req.OnlyEnabled && cat.Enabled == 0 {
				continue
			}
			cNode := &proto.SCategoryTree{
				Id:     int64(cat.Id),
				Name:   cat.Name,
				Url:    cat.CatUrl,
				Image:  cat.Icon,
				IsLeaf: p.testIsLeaf(cat.Id, categories, req),
			}
			arr = append(arr, cNode)
		}
	}
	return arr
}

// 排除分类
func (p *productService) testWalkCondition(req *proto.CategoryTreeRequest, cat *product.Category) bool {
	if req.OnlyEnabled && cat.Enabled == 0 {
		return false
	}
	if req.ExcludeIdList == nil {
		return true
	}
	for _, v := range req.ExcludeIdList {
		if v == int64(cat.Id) {
			return false
		}
	}
	return true
}

// 遍历加载下级分类
func (p *productService) walkLoadCategory(node *proto.SCategoryTree, categories []product.ICategory,
	req *proto.CategoryTreeRequest) {
	node.Children = []*proto.SCategoryTree{}
	// 遍历子分类
	for _, v := range categories {
		cat := v.GetValue()
		if cat.ParentId == int(node.Id) &&
			p.testWalkCondition(req, cat) {
			cNode := &proto.SCategoryTree{
				Id:    int64(cat.Id),
				Name:  cat.Name,
				Url:   cat.CatUrl,
				Image: cat.Icon,
			}
			node.Children = append(node.Children, cNode)
			p.walkLoadCategory(cNode, categories, req)
		}
	}
}

func (p *productService) GetBigCategories(mchId int64) []*proto.SProductCategory {
	cats := p.catRepo.GlobCatService().GetCategories()
	var list []*proto.SProductCategory
	for _, v := range cats {
		if v2 := v.GetValue(); v2.ParentId == 0 && v2.Enabled == 1 {
			v2.Icon = format.GetFileFullUrl(v2.Icon)
			list = append(list, p.parseCategoryDto(v2))
		}
	}
	return list
}

func (p *productService) GetChildCategories(mchId int64, parentId int64) []*proto.SProductCategory {
	cats := p.catRepo.GlobCatService().GetCategories()
	var list []*proto.SProductCategory
	for _, v := range cats {
		if vv := v.GetValue(); vv.ParentId == int(parentId) && vv.Enabled == 1 {
			vv.Icon = format.GetFileFullUrl(vv.Icon)
			p.setChild(cats, vv)
			list = append(list, p.parseCategoryDto(vv))
		}
	}
	return list
}

func (p *productService) setChild(list []product.ICategory, dst *product.Category) {
	for _, v := range list {
		if vv := v.GetValue(); vv.ParentId == dst.Id && vv.Enabled == 1 {
			if dst.Children == nil {
				dst.Children = []*product.Category{}
			}
			vv.Icon = format.GetFileFullUrl(vv.Icon)
			dst.Children = append(dst.Children, vv)
		}
	}
}

/***** 产品 *****/

// DeleteItem 删除货品
func (p *productService) DeleteItem(supplierId int64, productId int64) error {
	pro := p.productRepo.GetProduct(productId)
	if pro == nil || pro.GetValue().VendorId != supplierId {
		return product.ErrNoSuchProduct
	}
	return pro.Destroy()
}

// GetItemSaleLabels 获取商品的销售标签
func (p *productService) GetItemSaleLabels(mchId int64, itemId int64) []*item.Label {
	var list = make([]*item.Label, 0)
	//todo: refactor

	//sl := _s._rep.GetSale(mchId)
	//if goods := sl.ItemManager().GetItem(itemId); goods != nil {
	//	list = goods.GetSaleLabels()
	//}
	return list
}

// SaveItemSaleLabels 保存商品的销售标签
func (p *productService) SaveItemSaleLabels(mchId, itemId int64, tagIds []int) error {
	var err error

	//todo: refactor

	//sl := _s._rep.GetSale(mchId)
	//if goods := sl.ItemManager().GetItem(itemId); goods != nil {
	//	err = goods.SaveSaleLabels(tagIds)
	//} else {
	//	err = errors.New("商品不存在")
	//}
	return err
}

func (p *productService) parseModelDto(v *promodel.ProductModel) *proto.SProductModel {
	if v == nil {
		return nil
	}
	ret := &proto.SProductModel{
		Id:      int64(v.Id),
		Name:    v.Name,
		AttrStr: v.AttrStr,
		SpecStr: v.SpecStr,
		Attrs:   nil,
		Specs:   nil,
		Brands:  nil,
		Enabled: int32(v.Enabled),
	}
	return ret
}

func (p *productService) parseProductAttrDto(v *promodel.Attr) *proto.SProductAttr {
	return &proto.SProductAttr{
		Id:         int64(v.Id),
		Name:       v.Name,
		IsFilter:   int32(v.IsFilter),
		MultiCheck: int32(v.MultiCheck),
		SortNum:    int32(v.SortNum),
		ItemValues: v.ItemValues,
		Items:      nil,
	}
}

func (p *productService) parseProductAttrItemDto(v *promodel.AttrItem) *proto.SProductAttrItem {
	return &proto.SProductAttrItem{
		Id:      int64(v.Id),
		Value:   v.Value,
		SortNum: int32(v.SortNum),
	}
}

func (p *productService) parseBrandDto(v *promodel.ProductBrand) *proto.SProductBrand {
	return &proto.SProductBrand{
		Id:           int64(v.Id),
		Name:         v.Name,
		Image:        v.Image,
		SiteUrl:      v.SiteUrl,
		Introduce:    v.Introduce,
		ReviewStatus: v.ReviewStatus,
		ReviewRemark: v.ReviewRemark,
		Enabled:      int32(v.Enabled),
		CreateTime:   v.CreateTime,
	}
}

func (p *productService) parseCategoryDto(v *product.Category) *proto.SProductCategory {
	return &proto.SProductCategory{
		Id:          int64(v.Id),
		ParentId:    int64(v.ParentId),
		ModelId:     int64(v.ModelId),
		Priority:    int32(v.Priority),
		Name:        v.Name,
		IsVirtual:   v.VirtualCat == 1,
		CategoryUrl: v.CatUrl,
		// 虚拟分类跳转地址
		RedirectUrl: v.RedirectUrl,
		Icon:        v.Icon,
		IconPoint:   v.IconPoint,
		Level:       int32(v.Level),
		SortNum:     int32(v.SortNum),
		FloorShow:   v.FloorShow == 1,
		Enabled:     v.Enabled == 1,
		CreateTime:  v.CreateTime,
		Options:     map[string]string{},
		Brands:      nil,
		Children:    nil,
	}
}

func (p *productService) parseCategory(v *proto.SaveProductCategoryRequest) *product.Category {
	return &product.Category{
		Id:          int(v.Id),
		ParentId:    int(v.ParentId),
		ModelId:     int(v.ModelId),
		Priority:    int(v.Priority),
		Name:        v.Name,
		VirtualCat:  types.ElseInt(v.IsVirtual, 1, 0),
		CatUrl:      v.CategoryUrl,
		RedirectUrl: v.RedirectUrl,
		Icon:        v.Icon,
		IconPoint:   v.IconPoint,
		SortNum:     int(v.SortNum),
		FloorShow:   types.ElseInt(v.FloorShow, 1, 0),
		Enabled:     types.ElseInt(v.Enabled, 1, 0),
	}
}

func (p *productService) parseProductDto(v product.Product) *proto.SProduct {
	return &proto.SProduct{
		Id:          v.Id,
		CategoryId:  int64(v.CatId),
		Name:        v.Name,
		VendorId:    v.VendorId,
		BrandId:     int64(v.BrandId),
		Code:        v.Code,
		Image:       v.Image,
		Description: v.Description,
		Remark:      v.Remark,
		State:       v.State,
		SortNum:     v.SortNum,
		CreateTime:  v.CreateTime,
		UpdateTime:  v.UpdateTime,
	}
}

func (p *productService) parseProduct(v *proto.SaveProductRequest) *product.Product {
	ret := &product.Product{
		Id:          v.Id,
		CatId:       int32(v.CategoryId),
		Name:        v.Name,
		VendorId:    v.VendorId,
		BrandId:     int32(v.BrandId),
		Code:        v.Code,
		Image:       v.Image,
		Remark:      v.Remark,
		State:       v.State,
		SortNum:     v.SortNum,
		Description: v.Description,
	}
	if v.Attrs != nil {
		ret.Attrs = make([]*product.AttrValue, len(v.Attrs))
		for i, v := range v.Attrs {
			ret.Attrs[i] = p.parseProductAttrValue(v)
		}
	}
	return ret
}

func (p *productService) parseProductModel(v *proto.SaveProductModelRequest) *promodel.ProductModel {
	ret := &promodel.ProductModel{
		Id:      int(v.Id),
		Name:    v.Name,
		Enabled: int(v.Enabled),
	}
	if v.Attrs != nil {
		ret.Attrs = make([]*promodel.Attr, len(v.Attrs))
		for i, v := range v.Attrs {
			ret.Attrs[i] = p.parseProductAttr(v)
		}
	}
	if v.Specs != nil {
		ret.Specs = make([]*promodel.Spec, len(v.Specs))
		for i, v := range v.Specs {
			ret.Specs[i] = p.parseProductSpec(v)
		}
	}
	if v.Brands != nil {
		ret.BrandArray = make([]int, len(v.Brands))
		for i, v := range v.Brands {
			ret.BrandArray[i] = int(v)
		}
	}
	return ret
}

func (p *productService) parseBrand(v *proto.SProductBrand) *promodel.ProductBrand {
	return &promodel.ProductBrand{
		Id:           int32(v.Id),
		Name:         v.Name,
		Image:        v.Image,
		SiteUrl:      v.SiteUrl,
		Introduce:    v.Introduce,
		ReviewStatus: v.ReviewStatus,
		ReviewRemark: v.ReviewRemark,
		Enabled:      int(v.Enabled),
		CreateTime:   v.CreateTime,
	}
}

func (p *productService) parseProductAttr(v *proto.SProductAttr) *promodel.Attr {
	ret := &promodel.Attr{
		Id:         int(v.Id),
		Name:       v.Name,
		IsFilter:   int(v.IsFilter),
		MultiCheck: int(v.MultiCheck),
		ItemValues: "",
		SortNum:    int(v.SortNum),
	}
	if v.Items != nil {
		ret.Items = make([]*promodel.AttrItem, len(v.Items))
		for i, v := range v.Items {
			ret.Items[i] = p.parseProductAttrItem(v)
		}
	}
	return ret
}

func (p *productService) parseProductSpec(v *proto.SProductSpec) *promodel.Spec {
	ret := &promodel.Spec{
		Id:         int(v.Id),
		Name:       v.Name,
		ItemValues: v.ItemValues,
		SortNum:    int(v.SortNum),
	}
	if v.Items != nil {
		ret.Items = make([]*promodel.SpecItem, len(v.Items))
		for i, v := range v.Items {
			ret.Items[i] = p.parseProductSpecItem(v)
		}
	}
	return ret
}

func (p *productService) parseProductAttrItem(v *proto.SProductAttrItem) *promodel.AttrItem {
	return &promodel.AttrItem{
		Id:      int(v.Id),
		Value:   v.Value,
		SortNum: int(v.SortNum),
	}
}

func (p *productService) parseProductSpecItem(v *proto.SProductSpecItem) *promodel.SpecItem {
	return &promodel.SpecItem{
		Id:      int(v.Id),
		Value:   v.Value,
		Color:   v.Color,
		SortNum: int(v.SortNum),
	}
}

func (p *productService) parseProductAttrValueDto(v *product.AttrValue) *proto.SProductAttrValue {
	return &proto.SProductAttrValue{
		Id:       v.Id,
		AttrId:   v.AttrId,
		AttrName: v.AttrName,
		AttrData: v.AttrData,
		AttrWord: v.AttrWord,
	}
}

func (p *productService) parseSpecDto(v *promodel.Spec) *proto.SProductSpec {
	return &proto.SProductSpec{
		Id:         int64(v.Id),
		Name:       v.Name,
		SortNum:    int32(v.SortNum),
		ItemValues: v.ItemValues,
		Items:      nil,
	}
}

func (p *productService) parseProductSpecItemDto(v *promodel.SpecItem) *proto.SProductSpecItem {
	return &proto.SProductSpecItem{
		Id:      int64(v.Id),
		Value:   v.Value,
		Color:   v.Color,
		SortNum: int32(v.SortNum),
	}
}

func (p *productService) parseProductAttrValue(v *proto.SProductAttrValue) *product.AttrValue {
	return &product.AttrValue{
		Id:       v.Id,
		AttrName: v.AttrName,
		AttrId:   v.AttrId,
		AttrData: v.AttrData,
		AttrWord: v.AttrWord,
	}
}
