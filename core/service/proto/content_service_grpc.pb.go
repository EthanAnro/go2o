// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: content_service.proto

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

// ContentServiceClient is the client API for ContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentServiceClient interface {
	// 获取页面
	GetPage(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SPage, error)
	// 保存页面
	SavePage(ctx context.Context, in *SPage, opts ...grpc.CallOption) (*Result, error)
	// 删除页面
	DeletePage(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// 获取所有栏目
	GetArticleCategories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ArticleCategoriesResponse, error)
	// 获取文章栏目,可传入ID或者别名
	GetArticleCategory(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SArticleCategory, error)
	// 保存文章栏目
	SaveArticleCategory(ctx context.Context, in *SArticleCategory, opts ...grpc.CallOption) (*Result, error)
	// 删除文章分类
	DeleteArticleCategory(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// 获取文章
	GetArticle(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SArticle, error)
	// 删除文章
	DeleteArticle(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// 保存文章
	SaveArticle(ctx context.Context, in *SArticle, opts ...grpc.CallOption) (*Result, error)
	// * 获取置顶的文章,cat
	QueryTopArticles(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*ArticleListResponse, error)
	// * 获取分页文章
	QueryPagingArticles(ctx context.Context, in *PagingArticleRequest, opts ...grpc.CallOption) (*ArticleListResponse, error)
}

type contentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContentServiceClient(cc grpc.ClientConnInterface) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) GetPage(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SPage, error) {
	out := new(SPage)
	err := c.cc.Invoke(ctx, "/ContentService/GetPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) SavePage(ctx context.Context, in *SPage, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ContentService/SavePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) DeletePage(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ContentService/DeletePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetArticleCategories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ArticleCategoriesResponse, error) {
	out := new(ArticleCategoriesResponse)
	err := c.cc.Invoke(ctx, "/ContentService/GetArticleCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetArticleCategory(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SArticleCategory, error) {
	out := new(SArticleCategory)
	err := c.cc.Invoke(ctx, "/ContentService/GetArticleCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) SaveArticleCategory(ctx context.Context, in *SArticleCategory, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ContentService/SaveArticleCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) DeleteArticleCategory(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ContentService/DeleteArticleCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetArticle(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SArticle, error) {
	out := new(SArticle)
	err := c.cc.Invoke(ctx, "/ContentService/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) DeleteArticle(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ContentService/DeleteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) SaveArticle(ctx context.Context, in *SArticle, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ContentService/SaveArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) QueryTopArticles(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*ArticleListResponse, error) {
	out := new(ArticleListResponse)
	err := c.cc.Invoke(ctx, "/ContentService/QueryTopArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) QueryPagingArticles(ctx context.Context, in *PagingArticleRequest, opts ...grpc.CallOption) (*ArticleListResponse, error) {
	out := new(ArticleListResponse)
	err := c.cc.Invoke(ctx, "/ContentService/QueryPagingArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServiceServer is the server API for ContentService service.
// All implementations must embed UnimplementedContentServiceServer
// for forward compatibility
type ContentServiceServer interface {
	// 获取页面
	GetPage(context.Context, *IdOrName) (*SPage, error)
	// 保存页面
	SavePage(context.Context, *SPage) (*Result, error)
	// 删除页面
	DeletePage(context.Context, *Int64) (*Result, error)
	// 获取所有栏目
	GetArticleCategories(context.Context, *Empty) (*ArticleCategoriesResponse, error)
	// 获取文章栏目,可传入ID或者别名
	GetArticleCategory(context.Context, *IdOrName) (*SArticleCategory, error)
	// 保存文章栏目
	SaveArticleCategory(context.Context, *SArticleCategory) (*Result, error)
	// 删除文章分类
	DeleteArticleCategory(context.Context, *Int64) (*Result, error)
	// 获取文章
	GetArticle(context.Context, *IdOrName) (*SArticle, error)
	// 删除文章
	DeleteArticle(context.Context, *Int64) (*Result, error)
	// 保存文章
	SaveArticle(context.Context, *SArticle) (*Result, error)
	// * 获取置顶的文章,cat
	QueryTopArticles(context.Context, *IdOrName) (*ArticleListResponse, error)
	// * 获取分页文章
	QueryPagingArticles(context.Context, *PagingArticleRequest) (*ArticleListResponse, error)
	mustEmbedUnimplementedContentServiceServer()
}

// UnimplementedContentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContentServiceServer struct {
}

func (UnimplementedContentServiceServer) GetPage(context.Context, *IdOrName) (*SPage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPage not implemented")
}
func (UnimplementedContentServiceServer) SavePage(context.Context, *SPage) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SavePage not implemented")
}
func (UnimplementedContentServiceServer) DeletePage(context.Context, *Int64) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePage not implemented")
}
func (UnimplementedContentServiceServer) GetArticleCategories(context.Context, *Empty) (*ArticleCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleCategories not implemented")
}
func (UnimplementedContentServiceServer) GetArticleCategory(context.Context, *IdOrName) (*SArticleCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleCategory not implemented")
}
func (UnimplementedContentServiceServer) SaveArticleCategory(context.Context, *SArticleCategory) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveArticleCategory not implemented")
}
func (UnimplementedContentServiceServer) DeleteArticleCategory(context.Context, *Int64) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticleCategory not implemented")
}
func (UnimplementedContentServiceServer) GetArticle(context.Context, *IdOrName) (*SArticle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedContentServiceServer) DeleteArticle(context.Context, *Int64) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedContentServiceServer) SaveArticle(context.Context, *SArticle) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveArticle not implemented")
}
func (UnimplementedContentServiceServer) QueryTopArticles(context.Context, *IdOrName) (*ArticleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTopArticles not implemented")
}
func (UnimplementedContentServiceServer) QueryPagingArticles(context.Context, *PagingArticleRequest) (*ArticleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPagingArticles not implemented")
}
func (UnimplementedContentServiceServer) mustEmbedUnimplementedContentServiceServer() {}

// UnsafeContentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServiceServer will
// result in compilation errors.
type UnsafeContentServiceServer interface {
	mustEmbedUnimplementedContentServiceServer()
}

func RegisterContentServiceServer(s grpc.ServiceRegistrar, srv ContentServiceServer) {
	s.RegisterService(&ContentService_ServiceDesc, srv)
}

func _ContentService_GetPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdOrName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/GetPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetPage(ctx, req.(*IdOrName))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_SavePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SPage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).SavePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/SavePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).SavePage(ctx, req.(*SPage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_DeletePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DeletePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/DeletePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeletePage(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetArticleCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetArticleCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/GetArticleCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetArticleCategories(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetArticleCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdOrName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetArticleCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/GetArticleCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetArticleCategory(ctx, req.(*IdOrName))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_SaveArticleCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SArticleCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).SaveArticleCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/SaveArticleCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).SaveArticleCategory(ctx, req.(*SArticleCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_DeleteArticleCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DeleteArticleCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/DeleteArticleCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeleteArticleCategory(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdOrName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetArticle(ctx, req.(*IdOrName))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/DeleteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeleteArticle(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_SaveArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SArticle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).SaveArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/SaveArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).SaveArticle(ctx, req.(*SArticle))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_QueryTopArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdOrName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).QueryTopArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/QueryTopArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).QueryTopArticles(ctx, req.(*IdOrName))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_QueryPagingArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PagingArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).QueryPagingArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/QueryPagingArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).QueryPagingArticles(ctx, req.(*PagingArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContentService_ServiceDesc is the grpc.ServiceDesc for ContentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPage",
			Handler:    _ContentService_GetPage_Handler,
		},
		{
			MethodName: "SavePage",
			Handler:    _ContentService_SavePage_Handler,
		},
		{
			MethodName: "DeletePage",
			Handler:    _ContentService_DeletePage_Handler,
		},
		{
			MethodName: "GetArticleCategories",
			Handler:    _ContentService_GetArticleCategories_Handler,
		},
		{
			MethodName: "GetArticleCategory",
			Handler:    _ContentService_GetArticleCategory_Handler,
		},
		{
			MethodName: "SaveArticleCategory",
			Handler:    _ContentService_SaveArticleCategory_Handler,
		},
		{
			MethodName: "DeleteArticleCategory",
			Handler:    _ContentService_DeleteArticleCategory_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _ContentService_GetArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _ContentService_DeleteArticle_Handler,
		},
		{
			MethodName: "SaveArticle",
			Handler:    _ContentService_SaveArticle_Handler,
		},
		{
			MethodName: "QueryTopArticles",
			Handler:    _ContentService_QueryTopArticles_Handler,
		},
		{
			MethodName: "QueryPagingArticles",
			Handler:    _ContentService_QueryPagingArticles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content_service.proto",
}
