//*
// This file is auto generated by tto v0.4.5 !
// If you want to modify this code, please read the guide
// to modify code template.
//
// Get started: https://github.com/ixre/tto
//
// Copyright (C) 2021 <no value>, All rights reserved.
//
// name : template_service.proto
// author : jarrysix
// date : 2021/12/02 10:37:45
// description : 条码服务
// history :

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: code_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SaveQrTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,json=id,proto3" json:"Id"`
	//* 模板标题
	Title string `protobuf:"bytes,2,opt,name=Title,json=title,proto3" json:"Title"`
	//* 背景图片
	BgImage string `protobuf:"bytes,3,opt,name=BgImage,json=bgImage,proto3" json:"BgImage"`
	//* 垂直偏离量
	OffsetX int32 `protobuf:"varint,4,opt,name=OffsetX,json=offsetX,proto3" json:"OffsetX"`
	//* 垂直偏移量
	OffsetY int32 `protobuf:"varint,5,opt,name=OffsetY,json=offsetY,proto3" json:"OffsetY"`
	//* 二维码模板文本
	Comment string `protobuf:"bytes,6,opt,name=Comment,json=comment,proto3" json:"Comment"`
	//* 回调地址
	CallbackUrl string `protobuf:"bytes,7,opt,name=CallbackUrl,json=callbackUrl,proto3" json:"CallbackUrl"`
	//* 是否启用
	Enabled int32 `protobuf:"varint,8,opt,name=Enabled,json=enabled,proto3" json:"Enabled"`
}

func (x *SaveQrTemplateRequest) Reset() {
	*x = SaveQrTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveQrTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveQrTemplateRequest) ProtoMessage() {}

func (x *SaveQrTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_code_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveQrTemplateRequest.ProtoReflect.Descriptor instead.
func (*SaveQrTemplateRequest) Descriptor() ([]byte, []int) {
	return file_code_service_proto_rawDescGZIP(), []int{0}
}

func (x *SaveQrTemplateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SaveQrTemplateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SaveQrTemplateRequest) GetBgImage() string {
	if x != nil {
		return x.BgImage
	}
	return ""
}

func (x *SaveQrTemplateRequest) GetOffsetX() int32 {
	if x != nil {
		return x.OffsetX
	}
	return 0
}

func (x *SaveQrTemplateRequest) GetOffsetY() int32 {
	if x != nil {
		return x.OffsetY
	}
	return 0
}

func (x *SaveQrTemplateRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *SaveQrTemplateRequest) GetCallbackUrl() string {
	if x != nil {
		return x.CallbackUrl
	}
	return ""
}

func (x *SaveQrTemplateRequest) GetEnabled() int32 {
	if x != nil {
		return x.Enabled
	}
	return 0
}

type SaveQrTemplateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode int64  `protobuf:"varint,1,opt,name=ErrCode,json=errCode,proto3" json:"ErrCode"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=ErrMsg,json=errMsg,proto3" json:"ErrMsg"`
	Id      int64  `protobuf:"varint,3,opt,name=Id,json=id,proto3" json:"Id"`
}

func (x *SaveQrTemplateResponse) Reset() {
	*x = SaveQrTemplateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveQrTemplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveQrTemplateResponse) ProtoMessage() {}

func (x *SaveQrTemplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_code_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveQrTemplateResponse.ProtoReflect.Descriptor instead.
func (*SaveQrTemplateResponse) Descriptor() ([]byte, []int) {
	return file_code_service_proto_rawDescGZIP(), []int{1}
}

func (x *SaveQrTemplateResponse) GetErrCode() int64 {
	if x != nil {
		return x.ErrCode
	}
	return 0
}

func (x *SaveQrTemplateResponse) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *SaveQrTemplateResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CommQrTemplateId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=Value,json=value,proto3" json:"Value"`
}

func (x *CommQrTemplateId) Reset() {
	*x = CommQrTemplateId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommQrTemplateId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommQrTemplateId) ProtoMessage() {}

func (x *CommQrTemplateId) ProtoReflect() protoreflect.Message {
	mi := &file_code_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommQrTemplateId.ProtoReflect.Descriptor instead.
func (*CommQrTemplateId) Descriptor() ([]byte, []int) {
	return file_code_service_proto_rawDescGZIP(), []int{2}
}

func (x *CommQrTemplateId) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type SQrTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,json=id,proto3" json:"Id"`
	//* 模板标题
	Title string `protobuf:"bytes,2,opt,name=Title,json=title,proto3" json:"Title"`
	//* 背景图片
	BgImage string `protobuf:"bytes,3,opt,name=BgImage,json=bgImage,proto3" json:"BgImage"`
	//* 垂直偏离量
	OffsetX int32 `protobuf:"varint,4,opt,name=OffsetX,json=offsetX,proto3" json:"OffsetX"`
	//* 垂直偏移量
	OffsetY int32 `protobuf:"varint,5,opt,name=OffsetY,json=offsetY,proto3" json:"OffsetY"`
	//* 二维码模板文本
	Comment string `protobuf:"bytes,6,opt,name=Comment,json=comment,proto3" json:"Comment"`
	//* 回调地址
	CallbackUrl string `protobuf:"bytes,7,opt,name=CallbackUrl,json=callbackUrl,proto3" json:"CallbackUrl"`
	//* 是否启用
	Enabled int32 `protobuf:"varint,8,opt,name=Enabled,json=enabled,proto3" json:"Enabled"`
}

func (x *SQrTemplate) Reset() {
	*x = SQrTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_code_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQrTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQrTemplate) ProtoMessage() {}

func (x *SQrTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_code_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQrTemplate.ProtoReflect.Descriptor instead.
func (*SQrTemplate) Descriptor() ([]byte, []int) {
	return file_code_service_proto_rawDescGZIP(), []int{3}
}

func (x *SQrTemplate) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SQrTemplate) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SQrTemplate) GetBgImage() string {
	if x != nil {
		return x.BgImage
	}
	return ""
}

func (x *SQrTemplate) GetOffsetX() int32 {
	if x != nil {
		return x.OffsetX
	}
	return 0
}

func (x *SQrTemplate) GetOffsetY() int32 {
	if x != nil {
		return x.OffsetY
	}
	return 0
}

func (x *SQrTemplate) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *SQrTemplate) GetCallbackUrl() string {
	if x != nil {
		return x.CallbackUrl
	}
	return ""
}

func (x *SQrTemplate) GetEnabled() int32 {
	if x != nil {
		return x.Enabled
	}
	return 0
}

var File_code_service_proto protoreflect.FileDescriptor

var file_code_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x51, 0x72, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x67, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x67, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x58, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x58, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x59, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x59,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x5a, 0x0a, 0x16, 0x53, 0x61, 0x76, 0x65, 0x51, 0x72,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x72,
	0x72, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x4d,
	0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x28, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x51, 0x72, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xd7, 0x01, 0x0a,
	0x0b, 0x53, 0x51, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x67, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x67, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x58, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x58, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x59, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x59,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x32, 0xb8, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x64, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65, 0x51, 0x72,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x51,
	0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x51, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x51, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x51, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x1a,
	0x0c, 0x2e, 0x53, 0x51, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x00, 0x12,
	0x30, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x51, 0x72, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x42, 0x1f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x67, 0x6f, 0x32, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_code_service_proto_rawDescOnce sync.Once
	file_code_service_proto_rawDescData = file_code_service_proto_rawDesc
)

func file_code_service_proto_rawDescGZIP() []byte {
	file_code_service_proto_rawDescOnce.Do(func() {
		file_code_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_code_service_proto_rawDescData)
	})
	return file_code_service_proto_rawDescData
}

var file_code_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_code_service_proto_goTypes = []interface{}{
	(*SaveQrTemplateRequest)(nil),  // 0: SaveQrTemplateRequest
	(*SaveQrTemplateResponse)(nil), // 1: SaveQrTemplateResponse
	(*CommQrTemplateId)(nil),       // 2: CommQrTemplateId
	(*SQrTemplate)(nil),            // 3: SQrTemplate
	(*Result)(nil),                 // 4: Result
}
var file_code_service_proto_depIdxs = []int32{
	0, // 0: CodeService.SaveQrTemplate:input_type -> SaveQrTemplateRequest
	2, // 1: CodeService.GetQrTemplate:input_type -> CommQrTemplateId
	2, // 2: CodeService.DeleteQrTemplate:input_type -> CommQrTemplateId
	1, // 3: CodeService.SaveQrTemplate:output_type -> SaveQrTemplateResponse
	3, // 4: CodeService.GetQrTemplate:output_type -> SQrTemplate
	4, // 5: CodeService.DeleteQrTemplate:output_type -> Result
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_code_service_proto_init() }
func file_code_service_proto_init() {
	if File_code_service_proto != nil {
		return
	}
	file_global_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_code_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveQrTemplateRequest); i {
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
		file_code_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveQrTemplateResponse); i {
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
		file_code_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommQrTemplateId); i {
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
		file_code_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQrTemplate); i {
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
			RawDescriptor: file_code_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_code_service_proto_goTypes,
		DependencyIndexes: file_code_service_proto_depIdxs,
		MessageInfos:      file_code_service_proto_msgTypes,
	}.Build()
	File_code_service_proto = out.File
	file_code_service_proto_rawDesc = nil
	file_code_service_proto_goTypes = nil
	file_code_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CodeServiceClient is the client API for CodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CodeServiceClient interface {
	// 保存CommQrTemplate
	SaveQrTemplate(ctx context.Context, in *SaveQrTemplateRequest, opts ...grpc.CallOption) (*SaveQrTemplateResponse, error)
	// 获取CommQrTemplate
	GetQrTemplate(ctx context.Context, in *CommQrTemplateId, opts ...grpc.CallOption) (*SQrTemplate, error)
	// 删除CommQrTemplate
	DeleteQrTemplate(ctx context.Context, in *CommQrTemplateId, opts ...grpc.CallOption) (*Result, error)
}

type codeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeServiceClient(cc grpc.ClientConnInterface) CodeServiceClient {
	return &codeServiceClient{cc}
}

func (c *codeServiceClient) SaveQrTemplate(ctx context.Context, in *SaveQrTemplateRequest, opts ...grpc.CallOption) (*SaveQrTemplateResponse, error) {
	out := new(SaveQrTemplateResponse)
	err := c.cc.Invoke(ctx, "/CodeService/SaveQrTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServiceClient) GetQrTemplate(ctx context.Context, in *CommQrTemplateId, opts ...grpc.CallOption) (*SQrTemplate, error) {
	out := new(SQrTemplate)
	err := c.cc.Invoke(ctx, "/CodeService/GetQrTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeServiceClient) DeleteQrTemplate(ctx context.Context, in *CommQrTemplateId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/CodeService/DeleteQrTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeServiceServer is the server API for CodeService service.
type CodeServiceServer interface {
	// 保存CommQrTemplate
	SaveQrTemplate(context.Context, *SaveQrTemplateRequest) (*SaveQrTemplateResponse, error)
	// 获取CommQrTemplate
	GetQrTemplate(context.Context, *CommQrTemplateId) (*SQrTemplate, error)
	// 删除CommQrTemplate
	DeleteQrTemplate(context.Context, *CommQrTemplateId) (*Result, error)
}

// UnimplementedCodeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCodeServiceServer struct {
}

func (*UnimplementedCodeServiceServer) SaveQrTemplate(context.Context, *SaveQrTemplateRequest) (*SaveQrTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveQrTemplate not implemented")
}
func (*UnimplementedCodeServiceServer) GetQrTemplate(context.Context, *CommQrTemplateId) (*SQrTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQrTemplate not implemented")
}
func (*UnimplementedCodeServiceServer) DeleteQrTemplate(context.Context, *CommQrTemplateId) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQrTemplate not implemented")
}

func RegisterCodeServiceServer(s *grpc.Server, srv CodeServiceServer) {
	s.RegisterService(&_CodeService_serviceDesc, srv)
}

func _CodeService_SaveQrTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveQrTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServiceServer).SaveQrTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CodeService/SaveQrTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServiceServer).SaveQrTemplate(ctx, req.(*SaveQrTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeService_GetQrTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommQrTemplateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServiceServer).GetQrTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CodeService/GetQrTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServiceServer).GetQrTemplate(ctx, req.(*CommQrTemplateId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeService_DeleteQrTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommQrTemplateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeServiceServer).DeleteQrTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CodeService/DeleteQrTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeServiceServer).DeleteQrTemplate(ctx, req.(*CommQrTemplateId))
	}
	return interceptor(ctx, in, info, handler)
}

var _CodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CodeService",
	HandlerType: (*CodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveQrTemplate",
			Handler:    _CodeService_SaveQrTemplate_Handler,
		},
		{
			MethodName: "GetQrTemplate",
			Handler:    _CodeService_GetQrTemplate_Handler,
		},
		{
			MethodName: "DeleteQrTemplate",
			Handler:    _CodeService_DeleteQrTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "code_service.proto",
}
