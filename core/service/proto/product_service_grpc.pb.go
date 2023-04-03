// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: product_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ProductService_GetProductModel_FullMethodName     = "/ProductService/GetProductModel"
	ProductService_GetModels_FullMethodName           = "/ProductService/GetModels"
	ProductService_GetAttr_FullMethodName             = "/ProductService/GetAttr"
	ProductService_GetAttrItem_FullMethodName         = "/ProductService/GetAttrItem"
	ProductService_SaveProductModel_FullMethodName    = "/ProductService/SaveProductModel"
	ProductService_DeleteModel__FullMethodName        = "/ProductService/DeleteModel_"
	ProductService_GetBrand_FullMethodName            = "/ProductService/GetBrand"
	ProductService_SaveBrand_FullMethodName           = "/ProductService/SaveBrand"
	ProductService_DeleteBrand_FullMethodName         = "/ProductService/DeleteBrand"
	ProductService_GetBrands_FullMethodName           = "/ProductService/GetBrands"
	ProductService_GetCategory_FullMethodName         = "/ProductService/GetCategory"
	ProductService_DeleteCategory_FullMethodName      = "/ProductService/DeleteCategory"
	ProductService_SaveCategory_FullMethodName        = "/ProductService/SaveCategory"
	ProductService_GetCategoryTreeNode_FullMethodName = "/ProductService/GetCategoryTreeNode"
	ProductService_FindParentCategory_FullMethodName  = "/ProductService/FindParentCategory"
	ProductService_GetProduct_FullMethodName          = "/ProductService/GetProduct"
	ProductService_SaveProduct_FullMethodName         = "/ProductService/SaveProduct"
	ProductService_DeleteProduct_FullMethodName       = "/ProductService/DeleteProduct"
	ProductService_SaveProductInfo_FullMethodName     = "/ProductService/SaveProductInfo"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	// 获取产品模型及模型的规格属性
	GetProductModel(ctx context.Context, in *ProductModelId, opts ...grpc.CallOption) (*SProductModel, error)
	// 获取产品模型
	GetModels(ctx context.Context, in *GetModelsRequest, opts ...grpc.CallOption) (*ProductModelListResponse, error)
	// 获取属性
	GetAttr(ctx context.Context, in *ProductAttrId, opts ...grpc.CallOption) (*SProductAttr, error)
	// 获取属性项
	GetAttrItem(ctx context.Context, in *ProductAttrItemId, opts ...grpc.CallOption) (*SProductAttrItem, error)
	// 保存产品模型
	SaveProductModel(ctx context.Context, in *SaveProductModelRequest, opts ...grpc.CallOption) (*Result, error)
	// 删除产品模型
	DeleteModel_(ctx context.Context, in *ProductModelId, opts ...grpc.CallOption) (*Result, error)
	// Get 产品品牌
	GetBrand(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SProductBrand, error)
	// Save 产品品牌
	SaveBrand(ctx context.Context, in *SProductBrand, opts ...grpc.CallOption) (*Result, error)
	// Delete 产品品牌
	DeleteBrand(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// 获取所有产品品牌
	GetBrands(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProductBrandListResponse, error)
	// 获取商品分类
	GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*SProductCategory, error)
	// 获取商品分类和选项
	//
	//	rpc GetCategoryAndOptions(mchId int64, id int32) (*product.Category,
	//	domain.IOptionStore)
	DeleteCategory(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// 保存产品分类
	SaveCategory(ctx context.Context, in *SaveProductCategoryRequest, opts ...grpc.CallOption) (*SaveProductCategoryResponse, error)
	// 获取分类树形数据
	GetCategoryTreeNode(ctx context.Context, in *CategoryTreeRequest, opts ...grpc.CallOption) (*CategoryTreeResponse, error)
	FindParentCategory(ctx context.Context, in *CategoryIdRequest, opts ...grpc.CallOption) (*CategoriesResponse, error)
	// 获取产品值
	GetProduct(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*SProduct, error)
	// 保存产品
	SaveProduct(ctx context.Context, in *SaveProductRequest, opts ...grpc.CallOption) (*SaveProductResponse, error)
	// 删除产品
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*Result, error)
	// 保存货品描述
	SaveProductInfo(ctx context.Context, in *ProductInfoRequest, opts ...grpc.CallOption) (*Result, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetProductModel(ctx context.Context, in *ProductModelId, opts ...grpc.CallOption) (*SProductModel, error) {
	out := new(SProductModel)
	err := c.cc.Invoke(ctx, ProductService_GetProductModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetModels(ctx context.Context, in *GetModelsRequest, opts ...grpc.CallOption) (*ProductModelListResponse, error) {
	out := new(ProductModelListResponse)
	err := c.cc.Invoke(ctx, ProductService_GetModels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetAttr(ctx context.Context, in *ProductAttrId, opts ...grpc.CallOption) (*SProductAttr, error) {
	out := new(SProductAttr)
	err := c.cc.Invoke(ctx, ProductService_GetAttr_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetAttrItem(ctx context.Context, in *ProductAttrItemId, opts ...grpc.CallOption) (*SProductAttrItem, error) {
	out := new(SProductAttrItem)
	err := c.cc.Invoke(ctx, ProductService_GetAttrItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SaveProductModel(ctx context.Context, in *SaveProductModelRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, ProductService_SaveProductModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteModel_(ctx context.Context, in *ProductModelId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, ProductService_DeleteModel__FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetBrand(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*SProductBrand, error) {
	out := new(SProductBrand)
	err := c.cc.Invoke(ctx, ProductService_GetBrand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SaveBrand(ctx context.Context, in *SProductBrand, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, ProductService_SaveBrand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteBrand(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, ProductService_DeleteBrand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetBrands(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProductBrandListResponse, error) {
	out := new(ProductBrandListResponse)
	err := c.cc.Invoke(ctx, ProductService_GetBrands_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*SProductCategory, error) {
	out := new(SProductCategory)
	err := c.cc.Invoke(ctx, ProductService_GetCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteCategory(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, ProductService_DeleteCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SaveCategory(ctx context.Context, in *SaveProductCategoryRequest, opts ...grpc.CallOption) (*SaveProductCategoryResponse, error) {
	out := new(SaveProductCategoryResponse)
	err := c.cc.Invoke(ctx, ProductService_SaveCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetCategoryTreeNode(ctx context.Context, in *CategoryTreeRequest, opts ...grpc.CallOption) (*CategoryTreeResponse, error) {
	out := new(CategoryTreeResponse)
	err := c.cc.Invoke(ctx, ProductService_GetCategoryTreeNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) FindParentCategory(ctx context.Context, in *CategoryIdRequest, opts ...grpc.CallOption) (*CategoriesResponse, error) {
	out := new(CategoriesResponse)
	err := c.cc.Invoke(ctx, ProductService_FindParentCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*SProduct, error) {
	out := new(SProduct)
	err := c.cc.Invoke(ctx, ProductService_GetProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SaveProduct(ctx context.Context, in *SaveProductRequest, opts ...grpc.CallOption) (*SaveProductResponse, error) {
	out := new(SaveProductResponse)
	err := c.cc.Invoke(ctx, ProductService_SaveProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, ProductService_DeleteProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) SaveProductInfo(ctx context.Context, in *ProductInfoRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, ProductService_SaveProductInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	// 获取产品模型及模型的规格属性
	GetProductModel(context.Context, *ProductModelId) (*SProductModel, error)
	// 获取产品模型
	GetModels(context.Context, *GetModelsRequest) (*ProductModelListResponse, error)
	// 获取属性
	GetAttr(context.Context, *ProductAttrId) (*SProductAttr, error)
	// 获取属性项
	GetAttrItem(context.Context, *ProductAttrItemId) (*SProductAttrItem, error)
	// 保存产品模型
	SaveProductModel(context.Context, *SaveProductModelRequest) (*Result, error)
	// 删除产品模型
	DeleteModel_(context.Context, *ProductModelId) (*Result, error)
	// Get 产品品牌
	GetBrand(context.Context, *Int64) (*SProductBrand, error)
	// Save 产品品牌
	SaveBrand(context.Context, *SProductBrand) (*Result, error)
	// Delete 产品品牌
	DeleteBrand(context.Context, *Int64) (*Result, error)
	// 获取所有产品品牌
	GetBrands(context.Context, *Empty) (*ProductBrandListResponse, error)
	// 获取商品分类
	GetCategory(context.Context, *GetCategoryRequest) (*SProductCategory, error)
	// 获取商品分类和选项
	//
	//	rpc GetCategoryAndOptions(mchId int64, id int32) (*product.Category,
	//	domain.IOptionStore)
	DeleteCategory(context.Context, *Int64) (*Result, error)
	// 保存产品分类
	SaveCategory(context.Context, *SaveProductCategoryRequest) (*SaveProductCategoryResponse, error)
	// 获取分类树形数据
	GetCategoryTreeNode(context.Context, *CategoryTreeRequest) (*CategoryTreeResponse, error)
	FindParentCategory(context.Context, *CategoryIdRequest) (*CategoriesResponse, error)
	// 获取产品值
	GetProduct(context.Context, *ProductId) (*SProduct, error)
	// 保存产品
	SaveProduct(context.Context, *SaveProductRequest) (*SaveProductResponse, error)
	// 删除产品
	DeleteProduct(context.Context, *DeleteProductRequest) (*Result, error)
	// 保存货品描述
	SaveProductInfo(context.Context, *ProductInfoRequest) (*Result, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) GetProductModel(context.Context, *ProductModelId) (*SProductModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductModel not implemented")
}
func (UnimplementedProductServiceServer) GetModels(context.Context, *GetModelsRequest) (*ProductModelListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModels not implemented")
}
func (UnimplementedProductServiceServer) GetAttr(context.Context, *ProductAttrId) (*SProductAttr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttr not implemented")
}
func (UnimplementedProductServiceServer) GetAttrItem(context.Context, *ProductAttrItemId) (*SProductAttrItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttrItem not implemented")
}
func (UnimplementedProductServiceServer) SaveProductModel(context.Context, *SaveProductModelRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveProductModel not implemented")
}
func (UnimplementedProductServiceServer) DeleteModel_(context.Context, *ProductModelId) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteModel_ not implemented")
}
func (UnimplementedProductServiceServer) GetBrand(context.Context, *Int64) (*SProductBrand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrand not implemented")
}
func (UnimplementedProductServiceServer) SaveBrand(context.Context, *SProductBrand) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBrand not implemented")
}
func (UnimplementedProductServiceServer) DeleteBrand(context.Context, *Int64) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBrand not implemented")
}
func (UnimplementedProductServiceServer) GetBrands(context.Context, *Empty) (*ProductBrandListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrands not implemented")
}
func (UnimplementedProductServiceServer) GetCategory(context.Context, *GetCategoryRequest) (*SProductCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedProductServiceServer) DeleteCategory(context.Context, *Int64) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedProductServiceServer) SaveCategory(context.Context, *SaveProductCategoryRequest) (*SaveProductCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveCategory not implemented")
}
func (UnimplementedProductServiceServer) GetCategoryTreeNode(context.Context, *CategoryTreeRequest) (*CategoryTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryTreeNode not implemented")
}
func (UnimplementedProductServiceServer) FindParentCategory(context.Context, *CategoryIdRequest) (*CategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindParentCategory not implemented")
}
func (UnimplementedProductServiceServer) GetProduct(context.Context, *ProductId) (*SProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductServiceServer) SaveProduct(context.Context, *SaveProductRequest) (*SaveProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveProduct not implemented")
}
func (UnimplementedProductServiceServer) DeleteProduct(context.Context, *DeleteProductRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServiceServer) SaveProductInfo(context.Context, *ProductInfoRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveProductInfo not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_GetProductModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductModelId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetProductModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductModel(ctx, req.(*ProductModelId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetModels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetModels(ctx, req.(*GetModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetAttr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductAttrId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetAttr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetAttr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetAttr(ctx, req.(*ProductAttrId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetAttrItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductAttrItemId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetAttrItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetAttrItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetAttrItem(ctx, req.(*ProductAttrItemId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SaveProductModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveProductModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SaveProductModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_SaveProductModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SaveProductModel(ctx, req.(*SaveProductModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteModel__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductModelId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteModel_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_DeleteModel__FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteModel_(ctx, req.(*ProductModelId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetBrand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetBrand(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SaveBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SProductBrand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SaveBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_SaveBrand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SaveBrand(ctx, req.(*SProductBrand))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_DeleteBrand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteBrand(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetBrands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetBrands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetBrands_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetBrands(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetCategory(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_DeleteCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteCategory(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SaveCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveProductCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SaveCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_SaveCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SaveCategory(ctx, req.(*SaveProductCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetCategoryTreeNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetCategoryTreeNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetCategoryTreeNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetCategoryTreeNode(ctx, req.(*CategoryTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_FindParentCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).FindParentCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_FindParentCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).FindParentCategory(ctx, req.(*CategoryIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SaveProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SaveProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_SaveProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SaveProduct(ctx, req.(*SaveProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_SaveProductInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).SaveProductInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_SaveProductInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).SaveProductInfo(ctx, req.(*ProductInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductModel",
			Handler:    _ProductService_GetProductModel_Handler,
		},
		{
			MethodName: "GetModels",
			Handler:    _ProductService_GetModels_Handler,
		},
		{
			MethodName: "GetAttr",
			Handler:    _ProductService_GetAttr_Handler,
		},
		{
			MethodName: "GetAttrItem",
			Handler:    _ProductService_GetAttrItem_Handler,
		},
		{
			MethodName: "SaveProductModel",
			Handler:    _ProductService_SaveProductModel_Handler,
		},
		{
			MethodName: "DeleteModel_",
			Handler:    _ProductService_DeleteModel__Handler,
		},
		{
			MethodName: "GetBrand",
			Handler:    _ProductService_GetBrand_Handler,
		},
		{
			MethodName: "SaveBrand",
			Handler:    _ProductService_SaveBrand_Handler,
		},
		{
			MethodName: "DeleteBrand",
			Handler:    _ProductService_DeleteBrand_Handler,
		},
		{
			MethodName: "GetBrands",
			Handler:    _ProductService_GetBrands_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _ProductService_GetCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _ProductService_DeleteCategory_Handler,
		},
		{
			MethodName: "SaveCategory",
			Handler:    _ProductService_SaveCategory_Handler,
		},
		{
			MethodName: "GetCategoryTreeNode",
			Handler:    _ProductService_GetCategoryTreeNode_Handler,
		},
		{
			MethodName: "FindParentCategory",
			Handler:    _ProductService_FindParentCategory_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
		{
			MethodName: "SaveProduct",
			Handler:    _ProductService_SaveProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
		{
			MethodName: "SaveProductInfo",
			Handler:    _ProductService_SaveProductInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product_service.proto",
}
