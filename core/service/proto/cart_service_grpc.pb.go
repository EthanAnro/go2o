// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: cart_service.proto

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
	CartService_WholesaleCartV1_FullMethodName = "/CartService/WholesaleCartV1"
	CartService_GetShoppingCart_FullMethodName = "/CartService/GetShoppingCart"
	CartService_ApplyItem_FullMethodName       = "/CartService/ApplyItem"
	CartService_CheckCart_FullMethodName       = "/CartService/CheckCart"
)

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	// 批发购物车接口
	WholesaleCartV1(ctx context.Context, in *WsCartRequest, opts ...grpc.CallOption) (*Result, error)
	// 获取购物车,当购物车编号不存在时,将返回一个新的购物车
	GetShoppingCart(ctx context.Context, in *ShoppingCartId, opts ...grpc.CallOption) (*SShoppingCart, error)
	// 购物车商品操作
	ApplyItem(ctx context.Context, in *CartItemOpRequest, opts ...grpc.CallOption) (*CartItemResponse, error)
	// 勾选商品结算
	CheckCart(ctx context.Context, in *CheckCartRequest, opts ...grpc.CallOption) (*Result, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) WholesaleCartV1(ctx context.Context, in *WsCartRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, CartService_WholesaleCartV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetShoppingCart(ctx context.Context, in *ShoppingCartId, opts ...grpc.CallOption) (*SShoppingCart, error) {
	out := new(SShoppingCart)
	err := c.cc.Invoke(ctx, CartService_GetShoppingCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) ApplyItem(ctx context.Context, in *CartItemOpRequest, opts ...grpc.CallOption) (*CartItemResponse, error) {
	out := new(CartItemResponse)
	err := c.cc.Invoke(ctx, CartService_ApplyItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) CheckCart(ctx context.Context, in *CheckCartRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, CartService_CheckCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility
type CartServiceServer interface {
	// 批发购物车接口
	WholesaleCartV1(context.Context, *WsCartRequest) (*Result, error)
	// 获取购物车,当购物车编号不存在时,将返回一个新的购物车
	GetShoppingCart(context.Context, *ShoppingCartId) (*SShoppingCart, error)
	// 购物车商品操作
	ApplyItem(context.Context, *CartItemOpRequest) (*CartItemResponse, error)
	// 勾选商品结算
	CheckCart(context.Context, *CheckCartRequest) (*Result, error)
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (UnimplementedCartServiceServer) WholesaleCartV1(context.Context, *WsCartRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WholesaleCartV1 not implemented")
}
func (UnimplementedCartServiceServer) GetShoppingCart(context.Context, *ShoppingCartId) (*SShoppingCart, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShoppingCart not implemented")
}
func (UnimplementedCartServiceServer) ApplyItem(context.Context, *CartItemOpRequest) (*CartItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyItem not implemented")
}
func (UnimplementedCartServiceServer) CheckCart(context.Context, *CheckCartRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCart not implemented")
}
func (UnimplementedCartServiceServer) mustEmbedUnimplementedCartServiceServer() {}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s grpc.ServiceRegistrar, srv CartServiceServer) {
	s.RegisterService(&CartService_ServiceDesc, srv)
}

func _CartService_WholesaleCartV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WsCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).WholesaleCartV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_WholesaleCartV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).WholesaleCartV1(ctx, req.(*WsCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetShoppingCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShoppingCartId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetShoppingCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_GetShoppingCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetShoppingCart(ctx, req.(*ShoppingCartId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_ApplyItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemOpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).ApplyItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_ApplyItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).ApplyItem(ctx, req.(*CartItemOpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_CheckCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).CheckCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_CheckCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).CheckCart(ctx, req.(*CheckCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CartService_ServiceDesc is the grpc.ServiceDesc for CartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WholesaleCartV1",
			Handler:    _CartService_WholesaleCartV1_Handler,
		},
		{
			MethodName: "GetShoppingCart",
			Handler:    _CartService_GetShoppingCart_Handler,
		},
		{
			MethodName: "ApplyItem",
			Handler:    _CartService_ApplyItem_Handler,
		},
		{
			MethodName: "CheckCart",
			Handler:    _CartService_CheckCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart_service.proto",
}
