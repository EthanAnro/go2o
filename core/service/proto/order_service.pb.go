// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: order_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_order_service_proto protoreflect.FileDescriptor

var file_order_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xed, 0x05, 0x0a,
	0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a,
	0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x56, 0x31, 0x12, 0x13,
	0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x14, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4f, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x12, 0x1b, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x56, 0x32, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x2d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x56, 0x32, 0x1a,
	0x0d, 0x2e, 0x53, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00,
	0x12, 0x27, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x56, 0x32, 0x1a, 0x0d, 0x2e, 0x53, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x10, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x18, 0x2e,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x12, 0x26, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x43, 0x61, 0x73, 0x68, 0x50, 0x61, 0x79, 0x12, 0x06, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x1a,
	0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x16, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x17, 0x50, 0x72, 0x65,
	0x70, 0x61, 0x72, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x5f, 0x12, 0x14, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e,
	0x6f, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x1d, 0x0a, 0x06,
	0x50, 0x69, 0x63, 0x6b, 0x55, 0x70, 0x12, 0x08, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f,
	0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x04, 0x53,
	0x68, 0x69, 0x70, 0x12, 0x15, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x68, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x0d, 0x42, 0x75, 0x79, 0x65, 0x72, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x08, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f,
	0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x08, 0x4c,
	0x6f, 0x67, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x08, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e,
	0x6f, 0x1a, 0x07, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x42, 0x1f, 0x0a, 0x13,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x32, 0x6f, 0x2e,
	0x72, 0x70, 0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_order_service_proto_goTypes = []interface{}{
	(*SubmitOrderRequest)(nil),         // 0: SubmitOrderRequest
	(*PrepareOrderRequest)(nil),        // 1: PrepareOrderRequest
	(*SubmitNormalOrderV2Request)(nil), // 2: SubmitNormalOrderV2Request
	(*OrderNoV2)(nil),                  // 3: OrderNoV2
	(*TradeOrderSubmitRequest)(nil),    // 4: TradeOrderSubmitRequest
	(*Int64)(nil),                      // 5: Int64
	(*TradeOrderTicketRequest)(nil),    // 6: TradeOrderTicketRequest
	(*CancelOrderRequest)(nil),         // 7: CancelOrderRequest
	(*OrderNo)(nil),                    // 8: OrderNo
	(*OrderShipmentRequest)(nil),       // 9: OrderShipmentRequest
	(*StringMap)(nil),                  // 10: StringMap
	(*PrepareOrderResponse)(nil),       // 11: PrepareOrderResponse
	(*NormalOrderSubmitResponse)(nil),  // 12: NormalOrderSubmitResponse
	(*SParentOrder)(nil),               // 13: SParentOrder
	(*SSingleOrder)(nil),               // 14: SSingleOrder
	(*Result)(nil),                     // 15: Result
	(*String)(nil),                     // 16: String
}
var file_order_service_proto_depIdxs = []int32{
	0,  // 0: OrderService.SubmitOrderV1:input_type -> SubmitOrderRequest
	1,  // 1: OrderService.PrepareOrder:input_type -> PrepareOrderRequest
	2,  // 2: OrderService.SubmitNormalOrder_:input_type -> SubmitNormalOrderV2Request
	3,  // 3: OrderService.GetParentOrder:input_type -> OrderNoV2
	3,  // 4: OrderService.GetOrder:input_type -> OrderNoV2
	4,  // 5: OrderService.SubmitTradeOrder:input_type -> TradeOrderSubmitRequest
	5,  // 6: OrderService.TradeOrderCashPay:input_type -> Int64
	6,  // 7: OrderService.TradeOrderUpdateTicket:input_type -> TradeOrderTicketRequest
	1,  // 8: OrderService.PrepareOrderWithCoupon_:input_type -> PrepareOrderRequest
	7,  // 9: OrderService.CancelOrder:input_type -> CancelOrderRequest
	8,  // 10: OrderService.ConfirmOrder:input_type -> OrderNo
	8,  // 11: OrderService.PickUp:input_type -> OrderNo
	9,  // 12: OrderService.Ship:input_type -> OrderShipmentRequest
	8,  // 13: OrderService.BuyerReceived:input_type -> OrderNo
	8,  // 14: OrderService.LogBytes:input_type -> OrderNo
	10, // 15: OrderService.SubmitOrderV1:output_type -> StringMap
	11, // 16: OrderService.PrepareOrder:output_type -> PrepareOrderResponse
	12, // 17: OrderService.SubmitNormalOrder_:output_type -> NormalOrderSubmitResponse
	13, // 18: OrderService.GetParentOrder:output_type -> SParentOrder
	14, // 19: OrderService.GetOrder:output_type -> SSingleOrder
	15, // 20: OrderService.SubmitTradeOrder:output_type -> Result
	15, // 21: OrderService.TradeOrderCashPay:output_type -> Result
	15, // 22: OrderService.TradeOrderUpdateTicket:output_type -> Result
	10, // 23: OrderService.PrepareOrderWithCoupon_:output_type -> StringMap
	15, // 24: OrderService.CancelOrder:output_type -> Result
	15, // 25: OrderService.ConfirmOrder:output_type -> Result
	15, // 26: OrderService.PickUp:output_type -> Result
	15, // 27: OrderService.Ship:output_type -> Result
	15, // 28: OrderService.BuyerReceived:output_type -> Result
	16, // 29: OrderService.LogBytes:output_type -> String
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_order_service_proto_init() }
func file_order_service_proto_init() {
	if File_order_service_proto != nil {
		return
	}
	file_global_proto_init()
	file_message_order_dto_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_service_proto_goTypes,
		DependencyIndexes: file_order_service_proto_depIdxs,
	}.Build()
	File_order_service_proto = out.File
	file_order_service_proto_rawDesc = nil
	file_order_service_proto_goTypes = nil
	file_order_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	// 提交订单
	SubmitOrderV1(ctx context.Context, in *SubmitOrderRequest, opts ...grpc.CallOption) (*StringMap, error)
	// 预生成订单
	PrepareOrder(ctx context.Context, in *PrepareOrderRequest, opts ...grpc.CallOption) (*PrepareOrderResponse, error)
	// 提交普通订单
	SubmitNormalOrder_(ctx context.Context, in *SubmitNormalOrderV2Request, opts ...grpc.CallOption) (*NormalOrderSubmitResponse, error)
	//
	// 获取订单信息
	GetParentOrder(ctx context.Context, in *OrderNoV2, opts ...grpc.CallOption) (*SParentOrder, error)
	// 获取子订单,orderId
	GetOrder(ctx context.Context, in *OrderNoV2, opts ...grpc.CallOption) (*SSingleOrder, error)
	// 提交交易订单
	SubmitTradeOrder(ctx context.Context, in *TradeOrderSubmitRequest, opts ...grpc.CallOption) (*Result, error)
	// 交易单现金支付,orderId
	TradeOrderCashPay(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error)
	// 上传交易单发票
	TradeOrderUpdateTicket(ctx context.Context, in *TradeOrderTicketRequest, opts ...grpc.CallOption) (*Result, error)
	// 预生成订单，使用优惠券
	PrepareOrderWithCoupon_(ctx context.Context, in *PrepareOrderRequest, opts ...grpc.CallOption) (*StringMap, error)
	// 取消订单
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*Result, error)
	// 确定订单
	ConfirmOrder(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error)
	// 备货完成
	PickUp(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error)
	// 订单发货,并记录配送服务商编号及单号
	Ship(ctx context.Context, in *OrderShipmentRequest, opts ...grpc.CallOption) (*Result, error)
	// 买家收货
	BuyerReceived(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error)
	// 获取订单日志
	LogBytes(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*String, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) SubmitOrderV1(ctx context.Context, in *SubmitOrderRequest, opts ...grpc.CallOption) (*StringMap, error) {
	out := new(StringMap)
	err := c.cc.Invoke(ctx, "/OrderService/SubmitOrderV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) PrepareOrder(ctx context.Context, in *PrepareOrderRequest, opts ...grpc.CallOption) (*PrepareOrderResponse, error) {
	out := new(PrepareOrderResponse)
	err := c.cc.Invoke(ctx, "/OrderService/PrepareOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) SubmitNormalOrder_(ctx context.Context, in *SubmitNormalOrderV2Request, opts ...grpc.CallOption) (*NormalOrderSubmitResponse, error) {
	out := new(NormalOrderSubmitResponse)
	err := c.cc.Invoke(ctx, "/OrderService/SubmitNormalOrder_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetParentOrder(ctx context.Context, in *OrderNoV2, opts ...grpc.CallOption) (*SParentOrder, error) {
	out := new(SParentOrder)
	err := c.cc.Invoke(ctx, "/OrderService/GetParentOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *OrderNoV2, opts ...grpc.CallOption) (*SSingleOrder, error) {
	out := new(SSingleOrder)
	err := c.cc.Invoke(ctx, "/OrderService/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) SubmitTradeOrder(ctx context.Context, in *TradeOrderSubmitRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/SubmitTradeOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) TradeOrderCashPay(ctx context.Context, in *Int64, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/TradeOrderCashPay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) TradeOrderUpdateTicket(ctx context.Context, in *TradeOrderTicketRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/TradeOrderUpdateTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) PrepareOrderWithCoupon_(ctx context.Context, in *PrepareOrderRequest, opts ...grpc.CallOption) (*StringMap, error) {
	out := new(StringMap)
	err := c.cc.Invoke(ctx, "/OrderService/PrepareOrderWithCoupon_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ConfirmOrder(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/ConfirmOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) PickUp(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/PickUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) Ship(ctx context.Context, in *OrderShipmentRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/Ship", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) BuyerReceived(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/OrderService/BuyerReceived", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) LogBytes(ctx context.Context, in *OrderNo, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/OrderService/LogBytes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	// 提交订单
	SubmitOrderV1(context.Context, *SubmitOrderRequest) (*StringMap, error)
	// 预生成订单
	PrepareOrder(context.Context, *PrepareOrderRequest) (*PrepareOrderResponse, error)
	// 提交普通订单
	SubmitNormalOrder_(context.Context, *SubmitNormalOrderV2Request) (*NormalOrderSubmitResponse, error)
	//
	// 获取订单信息
	GetParentOrder(context.Context, *OrderNoV2) (*SParentOrder, error)
	// 获取子订单,orderId
	GetOrder(context.Context, *OrderNoV2) (*SSingleOrder, error)
	// 提交交易订单
	SubmitTradeOrder(context.Context, *TradeOrderSubmitRequest) (*Result, error)
	// 交易单现金支付,orderId
	TradeOrderCashPay(context.Context, *Int64) (*Result, error)
	// 上传交易单发票
	TradeOrderUpdateTicket(context.Context, *TradeOrderTicketRequest) (*Result, error)
	// 预生成订单，使用优惠券
	PrepareOrderWithCoupon_(context.Context, *PrepareOrderRequest) (*StringMap, error)
	// 取消订单
	CancelOrder(context.Context, *CancelOrderRequest) (*Result, error)
	// 确定订单
	ConfirmOrder(context.Context, *OrderNo) (*Result, error)
	// 备货完成
	PickUp(context.Context, *OrderNo) (*Result, error)
	// 订单发货,并记录配送服务商编号及单号
	Ship(context.Context, *OrderShipmentRequest) (*Result, error)
	// 买家收货
	BuyerReceived(context.Context, *OrderNo) (*Result, error)
	// 获取订单日志
	LogBytes(context.Context, *OrderNo) (*String, error)
}

// UnimplementedOrderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (*UnimplementedOrderServiceServer) SubmitOrderV1(context.Context, *SubmitOrderRequest) (*StringMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitOrderV1 not implemented")
}
func (*UnimplementedOrderServiceServer) PrepareOrder(context.Context, *PrepareOrderRequest) (*PrepareOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareOrder not implemented")
}
func (*UnimplementedOrderServiceServer) SubmitNormalOrder_(context.Context, *SubmitNormalOrderV2Request) (*NormalOrderSubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitNormalOrder_ not implemented")
}
func (*UnimplementedOrderServiceServer) GetParentOrder(context.Context, *OrderNoV2) (*SParentOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParentOrder not implemented")
}
func (*UnimplementedOrderServiceServer) GetOrder(context.Context, *OrderNoV2) (*SSingleOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (*UnimplementedOrderServiceServer) SubmitTradeOrder(context.Context, *TradeOrderSubmitRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTradeOrder not implemented")
}
func (*UnimplementedOrderServiceServer) TradeOrderCashPay(context.Context, *Int64) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeOrderCashPay not implemented")
}
func (*UnimplementedOrderServiceServer) TradeOrderUpdateTicket(context.Context, *TradeOrderTicketRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeOrderUpdateTicket not implemented")
}
func (*UnimplementedOrderServiceServer) PrepareOrderWithCoupon_(context.Context, *PrepareOrderRequest) (*StringMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareOrderWithCoupon_ not implemented")
}
func (*UnimplementedOrderServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (*UnimplementedOrderServiceServer) ConfirmOrder(context.Context, *OrderNo) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmOrder not implemented")
}
func (*UnimplementedOrderServiceServer) PickUp(context.Context, *OrderNo) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PickUp not implemented")
}
func (*UnimplementedOrderServiceServer) Ship(context.Context, *OrderShipmentRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ship not implemented")
}
func (*UnimplementedOrderServiceServer) BuyerReceived(context.Context, *OrderNo) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyerReceived not implemented")
}
func (*UnimplementedOrderServiceServer) LogBytes(context.Context, *OrderNo) (*String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogBytes not implemented")
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_SubmitOrderV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).SubmitOrderV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/SubmitOrderV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).SubmitOrderV1(ctx, req.(*SubmitOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_PrepareOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).PrepareOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/PrepareOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).PrepareOrder(ctx, req.(*PrepareOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_SubmitNormalOrder__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitNormalOrderV2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).SubmitNormalOrder_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/SubmitNormalOrder_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).SubmitNormalOrder_(ctx, req.(*SubmitNormalOrderV2Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetParentOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNoV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetParentOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/GetParentOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetParentOrder(ctx, req.(*OrderNoV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNoV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*OrderNoV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_SubmitTradeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeOrderSubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).SubmitTradeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/SubmitTradeOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).SubmitTradeOrder(ctx, req.(*TradeOrderSubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_TradeOrderCashPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int64)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).TradeOrderCashPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/TradeOrderCashPay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).TradeOrderCashPay(ctx, req.(*Int64))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_TradeOrderUpdateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeOrderTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).TradeOrderUpdateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/TradeOrderUpdateTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).TradeOrderUpdateTicket(ctx, req.(*TradeOrderTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_PrepareOrderWithCoupon__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).PrepareOrderWithCoupon_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/PrepareOrderWithCoupon_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).PrepareOrderWithCoupon_(ctx, req.(*PrepareOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ConfirmOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ConfirmOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/ConfirmOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ConfirmOrder(ctx, req.(*OrderNo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_PickUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).PickUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/PickUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).PickUp(ctx, req.(*OrderNo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_Ship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderShipmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).Ship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/Ship",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).Ship(ctx, req.(*OrderShipmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_BuyerReceived_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).BuyerReceived(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/BuyerReceived",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).BuyerReceived(ctx, req.(*OrderNo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_LogBytes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderNo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).LogBytes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/LogBytes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).LogBytes(ctx, req.(*OrderNo))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitOrderV1",
			Handler:    _OrderService_SubmitOrderV1_Handler,
		},
		{
			MethodName: "PrepareOrder",
			Handler:    _OrderService_PrepareOrder_Handler,
		},
		{
			MethodName: "SubmitNormalOrder_",
			Handler:    _OrderService_SubmitNormalOrder__Handler,
		},
		{
			MethodName: "GetParentOrder",
			Handler:    _OrderService_GetParentOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
		{
			MethodName: "SubmitTradeOrder",
			Handler:    _OrderService_SubmitTradeOrder_Handler,
		},
		{
			MethodName: "TradeOrderCashPay",
			Handler:    _OrderService_TradeOrderCashPay_Handler,
		},
		{
			MethodName: "TradeOrderUpdateTicket",
			Handler:    _OrderService_TradeOrderUpdateTicket_Handler,
		},
		{
			MethodName: "PrepareOrderWithCoupon_",
			Handler:    _OrderService_PrepareOrderWithCoupon__Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _OrderService_CancelOrder_Handler,
		},
		{
			MethodName: "ConfirmOrder",
			Handler:    _OrderService_ConfirmOrder_Handler,
		},
		{
			MethodName: "PickUp",
			Handler:    _OrderService_PickUp_Handler,
		},
		{
			MethodName: "Ship",
			Handler:    _OrderService_Ship_Handler,
		},
		{
			MethodName: "BuyerReceived",
			Handler:    _OrderService_BuyerReceived_Handler,
		},
		{
			MethodName: "LogBytes",
			Handler:    _OrderService_LogBytes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order_service.proto",
}
