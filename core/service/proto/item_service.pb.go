// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: item_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ItemDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId   int64 `protobuf:"zigzag64,1,opt,name=ItemId,proto3" json:"ItemId"`
	ItemType int32 `protobuf:"zigzag32,2,opt,name=ItemType,proto3" json:"ItemType"`
}

func (x *ItemDetailRequest) Reset() {
	*x = ItemDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemDetailRequest) ProtoMessage() {}

func (x *ItemDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_item_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemDetailRequest.ProtoReflect.Descriptor instead.
func (*ItemDetailRequest) Descriptor() ([]byte, []int) {
	return file_item_service_proto_rawDescGZIP(), []int{0}
}

func (x *ItemDetailRequest) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *ItemDetailRequest) GetItemType() int32 {
	if x != nil {
		return x.ItemType
	}
	return 0
}

type PagingGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemType   EItemSalesType `protobuf:"varint,1,opt,name=ItemType,proto3,enum=EItemSalesType" json:"ItemType"`
	SellerId   int64          `protobuf:"varint,2,opt,name=sellerId,proto3" json:"sellerId"`
	CategoryId int64          `protobuf:"varint,3,opt,name=categoryId,proto3" json:"categoryId"`
	// 关键词
	Keyword string         `protobuf:"bytes,4,opt,name=keyword,proto3" json:"keyword"`
	Params  *SPagingParams `protobuf:"bytes,5,opt,name=params,proto3" json:"params"`
}

func (x *PagingGoodsRequest) Reset() {
	*x = PagingGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagingGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingGoodsRequest) ProtoMessage() {}

func (x *PagingGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_item_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingGoodsRequest.ProtoReflect.Descriptor instead.
func (*PagingGoodsRequest) Descriptor() ([]byte, []int) {
	return file_item_service_proto_rawDescGZIP(), []int{1}
}

func (x *PagingGoodsRequest) GetItemType() EItemSalesType {
	if x != nil {
		return x.ItemType
	}
	return EItemSalesType_IT_NORMAL
}

func (x *PagingGoodsRequest) GetSellerId() int64 {
	if x != nil {
		return x.SellerId
	}
	return 0
}

func (x *PagingGoodsRequest) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *PagingGoodsRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *PagingGoodsRequest) GetParams() *SPagingParams {
	if x != nil {
		return x.Params
	}
	return nil
}

type PagingShopGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId     int64          `protobuf:"varint,1,opt,name=shopId,proto3" json:"shopId"`
	CategoryId int64          `protobuf:"varint,2,opt,name=categoryId,proto3" json:"categoryId"`
	Params     *SPagingParams `protobuf:"bytes,3,opt,name=params,proto3" json:"params"`
}

func (x *PagingShopGoodsRequest) Reset() {
	*x = PagingShopGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagingShopGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingShopGoodsRequest) ProtoMessage() {}

func (x *PagingShopGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_item_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingShopGoodsRequest.ProtoReflect.Descriptor instead.
func (*PagingShopGoodsRequest) Descriptor() ([]byte, []int) {
	return file_item_service_proto_rawDescGZIP(), []int{2}
}

func (x *PagingShopGoodsRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *PagingShopGoodsRequest) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *PagingShopGoodsRequest) GetParams() *SPagingParams {
	if x != nil {
		return x.Params
	}
	return nil
}

type SaleLabelItemsRequest_ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId  int64          `protobuf:"varint,1,opt,name=shopId,proto3" json:"shopId"`
	LabelId int32          `protobuf:"varint,2,opt,name=labelId,proto3" json:"labelId"`
	Params  *SPagingParams `protobuf:"bytes,3,opt,name=params,proto3" json:"params"`
}

func (x *SaleLabelItemsRequest_) Reset() {
	*x = SaleLabelItemsRequest_{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleLabelItemsRequest_) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleLabelItemsRequest_) ProtoMessage() {}

func (x *SaleLabelItemsRequest_) ProtoReflect() protoreflect.Message {
	mi := &file_item_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleLabelItemsRequest_.ProtoReflect.Descriptor instead.
func (*SaleLabelItemsRequest_) Descriptor() ([]byte, []int) {
	return file_item_service_proto_rawDescGZIP(), []int{3}
}

func (x *SaleLabelItemsRequest_) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *SaleLabelItemsRequest_) GetLabelId() int32 {
	if x != nil {
		return x.LabelId
	}
	return 0
}

func (x *SaleLabelItemsRequest_) GetParams() *SPagingParams {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_item_service_proto protoreflect.FileDescriptor

var file_item_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x11, 0x49, 0x74,
	0x65, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52,
	0x06, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x11, 0x52, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x22, 0xbf, 0x01, 0x0a, 0x12, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x49, 0x74,
	0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x45,
	0x49, 0x74, 0x65, 0x6d, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x49,
	0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x26, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x53, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x78, 0x0a, 0x16, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x53,
	0x68, 0x6f, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22,
	0x72, 0x0a, 0x16, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x6f,
	0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x32, 0xda, 0x0a, 0x0a, 0x0b, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x06,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x1a, 0x12, 0x2e, 0x53, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08,
	0x53, 0x61, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x53, 0x61, 0x76,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x36, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x53, 0x6b, 0x75, 0x12,
	0x11, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x53, 0x6b, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x53, 0x55, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x56, 0x69, 0x65,
	0x77, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x41, 0x6e, 0x64, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x1a, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x6e, 0x64, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x2d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x06, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x1a, 0x0f,
	0x2e, 0x53, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x22,
	0x00, 0x12, 0x19, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x53, 0x6b, 0x75, 0x12, 0x06, 0x2e, 0x53, 0x6b,
	0x75, 0x49, 0x64, 0x1a, 0x05, 0x2e, 0x53, 0x53, 0x6b, 0x75, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0a,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0f, 0x53, 0x61, 0x76,
	0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x41, 0x73, 0x49, 0x6c, 0x6c, 0x65, 0x67, 0x61, 0x6c,
	0x12, 0x13, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6c, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x12, 0x30, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x13, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x12, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x79, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x53, 0x68, 0x6f, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x6f, 0x70, 0x50, 0x61, 0x67, 0x65, 0x64, 0x4f, 0x6e, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x17, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x68,
	0x6f, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x68, 0x6f, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x64, 0x4f, 0x6e, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x13, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x34, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x10, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x57, 0x68, 0x6f,
	0x6c, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x12, 0x06, 0x2e, 0x53, 0x6b, 0x75, 0x49, 0x64, 0x1a, 0x18, 0x2e, 0x53, 0x57, 0x73, 0x53, 0x6b,
	0x75, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x57, 0x68, 0x6f, 0x6c,
	0x65, 0x73, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x15, 0x2e, 0x53, 0x61, 0x76,
	0x65, 0x53, 0x6b, 0x75, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x57, 0x68, 0x6f, 0x6c, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x57,
	0x73, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x53, 0x57, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x57, 0x68, 0x6f, 0x6c, 0x65, 0x73, 0x61, 0x6c,
	0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x53, 0x61, 0x76, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x09, 0x2e, 0x49, 0x64, 0x4f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a,
	0x0b, 0x2e, 0x53, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x00, 0x12, 0x27,
	0x0a, 0x0d, 0x53, 0x61, 0x76, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x0b, 0x2e, 0x53, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x1a, 0x07, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x06, 0x2e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x51, 0x0a,
	0x1e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x42, 0x79, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x12,
	0x17, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x1a, 0x14, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x1f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x67,
	0x6f, 0x32, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_item_service_proto_rawDescOnce sync.Once
	file_item_service_proto_rawDescData = file_item_service_proto_rawDesc
)

func file_item_service_proto_rawDescGZIP() []byte {
	file_item_service_proto_rawDescOnce.Do(func() {
		file_item_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_item_service_proto_rawDescData)
	})
	return file_item_service_proto_rawDescData
}

var file_item_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_item_service_proto_goTypes = []interface{}{
	(*ItemDetailRequest)(nil),           // 0: ItemDetailRequest
	(*PagingGoodsRequest)(nil),          // 1: PagingGoodsRequest
	(*PagingShopGoodsRequest)(nil),      // 2: PagingShopGoodsRequest
	(*SaleLabelItemsRequest_)(nil),      // 3: SaleLabelItemsRequest_
	(EItemSalesType)(0),                 // 4: EItemSalesType
	(*SPagingParams)(nil),               // 5: SPagingParams
	(*Int64)(nil),                       // 6: Int64
	(*SaveItemRequest)(nil),             // 7: SaveItemRequest
	(*ItemBySkuRequest)(nil),            // 8: ItemBySkuRequest
	(*GetItemAndSnapshotRequest)(nil),   // 9: GetItemAndSnapshotRequest
	(*SkuId)(nil),                       // 10: SkuId
	(*ItemReviewRequest)(nil),           // 11: ItemReviewRequest
	(*SaveLevelPriceRequest)(nil),       // 12: SaveLevelPriceRequest
	(*ItemIllegalRequest)(nil),          // 13: ItemIllegalRequest
	(*ShelveStateRequest)(nil),          // 14: ShelveStateRequest
	(*GetItemsByLabelRequest)(nil),      // 15: GetItemsByLabelRequest
	(*GetItemsRequest)(nil),             // 16: GetItemsRequest
	(*SaveSkuPricesRequest)(nil),        // 17: SaveSkuPricesRequest
	(*GetWsDiscountRequest)(nil),        // 18: GetWsDiscountRequest
	(*SaveItemDiscountRequest)(nil),     // 19: SaveItemDiscountRequest
	(*Empty)(nil),                       // 20: Empty
	(*IdOrName)(nil),                    // 21: IdOrName
	(*SItemLabel)(nil),                  // 22: SItemLabel
	(*SItemDataResponse)(nil),           // 23: SItemDataResponse
	(*SaveItemResponse)(nil),            // 24: SaveItemResponse
	(*SUnifiedViewItem)(nil),            // 25: SUnifiedViewItem
	(*ItemSnapshotResponse)(nil),        // 26: ItemSnapshotResponse
	(*STradeSnapshot)(nil),              // 27: STradeSnapshot
	(*SSku)(nil),                        // 28: SSku
	(*Result)(nil),                      // 29: Result
	(*String)(nil),                      // 30: String
	(*PagingShopGoodsResponse)(nil),     // 31: PagingShopGoodsResponse
	(*PagingGoodsResponse)(nil),         // 32: PagingGoodsResponse
	(*SWsSkuPriceListResponse)(nil),     // 33: SWsSkuPriceListResponse
	(*SWsItemDiscountListResponse)(nil), // 34: SWsItemDiscountListResponse
	(*ItemLabelListResponse)(nil),       // 35: ItemLabelListResponse
}
var file_item_service_proto_depIdxs = []int32{
	4,  // 0: PagingGoodsRequest.ItemType:type_name -> EItemSalesType
	5,  // 1: PagingGoodsRequest.params:type_name -> SPagingParams
	5,  // 2: PagingShopGoodsRequest.params:type_name -> SPagingParams
	5,  // 3: SaleLabelItemsRequest_.params:type_name -> SPagingParams
	6,  // 4: ItemService.GetItem:input_type -> Int64
	7,  // 5: ItemService.SaveItem:input_type -> SaveItemRequest
	8,  // 6: ItemService.GetItemBySku:input_type -> ItemBySkuRequest
	9,  // 7: ItemService.GetItemAndSnapshot:input_type -> GetItemAndSnapshotRequest
	6,  // 8: ItemService.GetTradeSnapshot:input_type -> Int64
	10, // 9: ItemService.GetSku:input_type -> SkuId
	11, // 10: ItemService.ReviewItem:input_type -> ItemReviewRequest
	12, // 11: ItemService.SaveLevelPrices:input_type -> SaveLevelPriceRequest
	13, // 12: ItemService.SignAsIllegal:input_type -> ItemIllegalRequest
	14, // 13: ItemService.SetShelveState:input_type -> ShelveStateRequest
	0,  // 14: ItemService.GetItemDetailData:input_type -> ItemDetailRequest
	15, // 15: ItemService.GetValueGoodsBySaleLabel:input_type -> GetItemsByLabelRequest
	2,  // 16: ItemService.GetShopPagedOnShelvesGoods:input_type -> PagingShopGoodsRequest
	1,  // 17: ItemService.GetPagedOnShelvesItem:input_type -> PagingGoodsRequest
	16, // 18: ItemService.GetItems:input_type -> GetItemsRequest
	10, // 19: ItemService.GetWholesalePriceArray:input_type -> SkuId
	17, // 20: ItemService.SaveWholesalePrice:input_type -> SaveSkuPricesRequest
	18, // 21: ItemService.GetWholesaleDiscountArray:input_type -> GetWsDiscountRequest
	19, // 22: ItemService.SaveWholesaleDiscount:input_type -> SaveItemDiscountRequest
	20, // 23: ItemService.GetAllSaleLabels:input_type -> Empty
	21, // 24: ItemService.GetSaleLabel:input_type -> IdOrName
	22, // 25: ItemService.SaveSaleLabel:input_type -> SItemLabel
	6,  // 26: ItemService.DeleteSaleLabel:input_type -> Int64
	3,  // 27: ItemService.GetPagedValueGoodsBySaleLabel_:input_type -> SaleLabelItemsRequest_
	23, // 28: ItemService.GetItem:output_type -> SItemDataResponse
	24, // 29: ItemService.SaveItem:output_type -> SaveItemResponse
	25, // 30: ItemService.GetItemBySku:output_type -> SUnifiedViewItem
	26, // 31: ItemService.GetItemAndSnapshot:output_type -> ItemSnapshotResponse
	27, // 32: ItemService.GetTradeSnapshot:output_type -> STradeSnapshot
	28, // 33: ItemService.GetSku:output_type -> SSku
	29, // 34: ItemService.ReviewItem:output_type -> Result
	29, // 35: ItemService.SaveLevelPrices:output_type -> Result
	29, // 36: ItemService.SignAsIllegal:output_type -> Result
	29, // 37: ItemService.SetShelveState:output_type -> Result
	30, // 38: ItemService.GetItemDetailData:output_type -> String
	31, // 39: ItemService.GetValueGoodsBySaleLabel:output_type -> PagingShopGoodsResponse
	31, // 40: ItemService.GetShopPagedOnShelvesGoods:output_type -> PagingShopGoodsResponse
	32, // 41: ItemService.GetPagedOnShelvesItem:output_type -> PagingGoodsResponse
	32, // 42: ItemService.GetItems:output_type -> PagingGoodsResponse
	33, // 43: ItemService.GetWholesalePriceArray:output_type -> SWsSkuPriceListResponse
	29, // 44: ItemService.SaveWholesalePrice:output_type -> Result
	34, // 45: ItemService.GetWholesaleDiscountArray:output_type -> SWsItemDiscountListResponse
	29, // 46: ItemService.SaveWholesaleDiscount:output_type -> Result
	35, // 47: ItemService.GetAllSaleLabels:output_type -> ItemLabelListResponse
	22, // 48: ItemService.GetSaleLabel:output_type -> SItemLabel
	29, // 49: ItemService.SaveSaleLabel:output_type -> Result
	29, // 50: ItemService.DeleteSaleLabel:output_type -> Result
	32, // 51: ItemService.GetPagedValueGoodsBySaleLabel_:output_type -> PagingGoodsResponse
	28, // [28:52] is the sub-list for method output_type
	4,  // [4:28] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_item_service_proto_init() }
func file_item_service_proto_init() {
	if File_item_service_proto != nil {
		return
	}
	file_global_proto_init()
	file_message_item_dto_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_item_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemDetailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_item_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagingGoodsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_item_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagingShopGoodsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_item_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleLabelItemsRequest_); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_item_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_item_service_proto_goTypes,
		DependencyIndexes: file_item_service_proto_depIdxs,
		MessageInfos:      file_item_service_proto_msgTypes,
	}.Build()
	File_item_service_proto = out.File
	file_item_service_proto_rawDesc = nil
	file_item_service_proto_goTypes = nil
	file_item_service_proto_depIdxs = nil
}
