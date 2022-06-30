// Code generated by protoc-gen-go. DO NOT EDIT.
// source: express_service.proto

package proto // import "./"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ExpressProviderListResponse struct {
	Value                []*SExpressProvider `protobuf:"bytes,1,rep,name=Value,json=value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ExpressProviderListResponse) Reset()         { *m = ExpressProviderListResponse{} }
func (m *ExpressProviderListResponse) String() string { return proto.CompactTextString(m) }
func (*ExpressProviderListResponse) ProtoMessage()    {}
func (*ExpressProviderListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_406b04ca6f4da78c, []int{0}
}
func (m *ExpressProviderListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpressProviderListResponse.Unmarshal(m, b)
}
func (m *ExpressProviderListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpressProviderListResponse.Marshal(b, m, deterministic)
}
func (dst *ExpressProviderListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpressProviderListResponse.Merge(dst, src)
}
func (m *ExpressProviderListResponse) XXX_Size() int {
	return xxx_messageInfo_ExpressProviderListResponse.Size(m)
}
func (m *ExpressProviderListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpressProviderListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExpressProviderListResponse proto.InternalMessageInfo

func (m *ExpressProviderListResponse) GetValue() []*SExpressProvider {
	if m != nil {
		return m.Value
	}
	return nil
}

// 快递服务商
type SExpressProvider struct {
	// 快递公司编号
	Id int64 `protobuf:"varint,1,opt,name=Id,json=id,proto3" json:"Id"`
	// 快递名称
	Name string `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name"`
	// 首字母，用于索引分组
	Letter string `protobuf:"bytes,3,opt,name=Letter,json=letter,proto3" json:"Letter"`
	// 分组,多个组,用","隔开
	GroupFlag string `protobuf:"bytes,4,opt,name=GroupFlag,json=groupFlag,proto3" json:"GroupFlag"`
	// 快递公司编码
	Code string `protobuf:"bytes,5,opt,name=Code,json=code,proto3" json:"Code"`
	// 接口编码
	ApiCode string `protobuf:"bytes,6,opt,name=ApiCode,json=apiCode,proto3" json:"ApiCode"`
	// 是否启用
	Enabled              bool     `protobuf:"varint,7,opt,name=Enabled,json=enabled,proto3" json:"Enabled"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SExpressProvider) Reset()         { *m = SExpressProvider{} }
func (m *SExpressProvider) String() string { return proto.CompactTextString(m) }
func (*SExpressProvider) ProtoMessage()    {}
func (*SExpressProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_406b04ca6f4da78c, []int{1}
}
func (m *SExpressProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SExpressProvider.Unmarshal(m, b)
}
func (m *SExpressProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SExpressProvider.Marshal(b, m, deterministic)
}
func (dst *SExpressProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SExpressProvider.Merge(dst, src)
}
func (m *SExpressProvider) XXX_Size() int {
	return xxx_messageInfo_SExpressProvider.Size(m)
}
func (m *SExpressProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_SExpressProvider.DiscardUnknown(m)
}

var xxx_messageInfo_SExpressProvider proto.InternalMessageInfo

func (m *SExpressProvider) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SExpressProvider) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SExpressProvider) GetLetter() string {
	if m != nil {
		return m.Letter
	}
	return ""
}

func (m *SExpressProvider) GetGroupFlag() string {
	if m != nil {
		return m.GroupFlag
	}
	return ""
}

func (m *SExpressProvider) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SExpressProvider) GetApiCode() string {
	if m != nil {
		return m.ApiCode
	}
	return ""
}

func (m *SExpressProvider) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

// 快递模板
type SExpressTemplate struct {
	// 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,json=id,proto3" json:"Id"`
	// 运营商编号
	SellerId int64 `protobuf:"varint,2,opt,name=SellerId,json=sellerId,proto3" json:"SellerId"`
	// 运费模板名称
	Name string `protobuf:"bytes,3,opt,name=Name,json=name,proto3" json:"Name"`
	// 是否卖价承担运费
	IsFree bool `protobuf:"varint,4,opt,name=IsFree,json=isFree,proto3" json:"IsFree"`
	// 运费计价依据
	Basis int32 `protobuf:"varint,5,opt,name=Basis,json=basis,proto3" json:"Basis"`
	// 首次计价单位,如首重为2kg
	FirstUnit int32 `protobuf:"varint,6,opt,name=FirstUnit,json=firstUnit,proto3" json:"FirstUnit"`
	// 首次计价单价,如续重1kg
	FirstFee int64 `protobuf:"varint,7,opt,name=FirstFee,json=firstFee,proto3" json:"FirstFee"`
	// 超过首次计价计算单位,如续重1kg
	AddUnit int32 `protobuf:"varint,8,opt,name=AddUnit,json=addUnit,proto3" json:"AddUnit"`
	// 超过首次计价单价，如续重1kg
	AddFee int64 `protobuf:"varint,9,opt,name=AddFee,json=addFee,proto3" json:"AddFee"`
	// 是否启用
	Enabled              bool     `protobuf:"varint,10,opt,name=Enabled,json=enabled,proto3" json:"Enabled"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SExpressTemplate) Reset()         { *m = SExpressTemplate{} }
func (m *SExpressTemplate) String() string { return proto.CompactTextString(m) }
func (*SExpressTemplate) ProtoMessage()    {}
func (*SExpressTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_406b04ca6f4da78c, []int{2}
}
func (m *SExpressTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SExpressTemplate.Unmarshal(m, b)
}
func (m *SExpressTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SExpressTemplate.Marshal(b, m, deterministic)
}
func (dst *SExpressTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SExpressTemplate.Merge(dst, src)
}
func (m *SExpressTemplate) XXX_Size() int {
	return xxx_messageInfo_SExpressTemplate.Size(m)
}
func (m *SExpressTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_SExpressTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_SExpressTemplate proto.InternalMessageInfo

func (m *SExpressTemplate) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SExpressTemplate) GetSellerId() int64 {
	if m != nil {
		return m.SellerId
	}
	return 0
}

func (m *SExpressTemplate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SExpressTemplate) GetIsFree() bool {
	if m != nil {
		return m.IsFree
	}
	return false
}

func (m *SExpressTemplate) GetBasis() int32 {
	if m != nil {
		return m.Basis
	}
	return 0
}

func (m *SExpressTemplate) GetFirstUnit() int32 {
	if m != nil {
		return m.FirstUnit
	}
	return 0
}

func (m *SExpressTemplate) GetFirstFee() int64 {
	if m != nil {
		return m.FirstFee
	}
	return 0
}

func (m *SExpressTemplate) GetAddUnit() int32 {
	if m != nil {
		return m.AddUnit
	}
	return 0
}

func (m *SExpressTemplate) GetAddFee() int64 {
	if m != nil {
		return m.AddFee
	}
	return 0
}

func (m *SExpressTemplate) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type ExpressTemplateId struct {
	SellerId             int64    `protobuf:"varint,1,opt,name=SellerId,json=sellerId,proto3" json:"SellerId"`
	TemplateId           int64    `protobuf:"varint,2,opt,name=TemplateId,json=templateId,proto3" json:"TemplateId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExpressTemplateId) Reset()         { *m = ExpressTemplateId{} }
func (m *ExpressTemplateId) String() string { return proto.CompactTextString(m) }
func (*ExpressTemplateId) ProtoMessage()    {}
func (*ExpressTemplateId) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_406b04ca6f4da78c, []int{3}
}
func (m *ExpressTemplateId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpressTemplateId.Unmarshal(m, b)
}
func (m *ExpressTemplateId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpressTemplateId.Marshal(b, m, deterministic)
}
func (dst *ExpressTemplateId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpressTemplateId.Merge(dst, src)
}
func (m *ExpressTemplateId) XXX_Size() int {
	return xxx_messageInfo_ExpressTemplateId.Size(m)
}
func (m *ExpressTemplateId) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpressTemplateId.DiscardUnknown(m)
}

var xxx_messageInfo_ExpressTemplateId proto.InternalMessageInfo

func (m *ExpressTemplateId) GetSellerId() int64 {
	if m != nil {
		return m.SellerId
	}
	return 0
}

func (m *ExpressTemplateId) GetTemplateId() int64 {
	if m != nil {
		return m.TemplateId
	}
	return 0
}

type GetTemplatesRequest struct {
	SellerId int64 `protobuf:"varint,1,opt,name=SellerId,json=sellerId,proto3" json:"SellerId"`
	// 仅返回已启用的模板
	OnlyEnabled          bool     `protobuf:"varint,2,opt,name=OnlyEnabled,json=onlyEnabled,proto3" json:"OnlyEnabled"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTemplatesRequest) Reset()         { *m = GetTemplatesRequest{} }
func (m *GetTemplatesRequest) String() string { return proto.CompactTextString(m) }
func (*GetTemplatesRequest) ProtoMessage()    {}
func (*GetTemplatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_406b04ca6f4da78c, []int{4}
}
func (m *GetTemplatesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTemplatesRequest.Unmarshal(m, b)
}
func (m *GetTemplatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTemplatesRequest.Marshal(b, m, deterministic)
}
func (dst *GetTemplatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTemplatesRequest.Merge(dst, src)
}
func (m *GetTemplatesRequest) XXX_Size() int {
	return xxx_messageInfo_GetTemplatesRequest.Size(m)
}
func (m *GetTemplatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTemplatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTemplatesRequest proto.InternalMessageInfo

func (m *GetTemplatesRequest) GetSellerId() int64 {
	if m != nil {
		return m.SellerId
	}
	return 0
}

func (m *GetTemplatesRequest) GetOnlyEnabled() bool {
	if m != nil {
		return m.OnlyEnabled
	}
	return false
}

type ExpressTemplateListResponse struct {
	Value                []*SExpressTemplate `protobuf:"bytes,1,rep,name=Value,json=value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ExpressTemplateListResponse) Reset()         { *m = ExpressTemplateListResponse{} }
func (m *ExpressTemplateListResponse) String() string { return proto.CompactTextString(m) }
func (*ExpressTemplateListResponse) ProtoMessage()    {}
func (*ExpressTemplateListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_406b04ca6f4da78c, []int{5}
}
func (m *ExpressTemplateListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpressTemplateListResponse.Unmarshal(m, b)
}
func (m *ExpressTemplateListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpressTemplateListResponse.Marshal(b, m, deterministic)
}
func (dst *ExpressTemplateListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpressTemplateListResponse.Merge(dst, src)
}
func (m *ExpressTemplateListResponse) XXX_Size() int {
	return xxx_messageInfo_ExpressTemplateListResponse.Size(m)
}
func (m *ExpressTemplateListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpressTemplateListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExpressTemplateListResponse proto.InternalMessageInfo

func (m *ExpressTemplateListResponse) GetValue() []*SExpressTemplate {
	if m != nil {
		return m.Value
	}
	return nil
}

// 快递地区模板
type SExpressAreaTemplate struct {
	// 模板编号
	Id int64 `protobuf:"varint,1,opt,name=Id,json=id,proto3" json:"Id"`
	// 地区编号列表，通常精确到省即可
	CodeList string `protobuf:"bytes,2,opt,name=CodeList,json=codeList,proto3" json:"CodeList"`
	// 地区名称列表
	NameList string `protobuf:"bytes,3,opt,name=NameList,json=nameList,proto3" json:"NameList"`
	// 首次数值，如 首重为2kg
	FirstUnit int32 `protobuf:"varint,4,opt,name=FirstUnit,json=firstUnit,proto3" json:"FirstUnit"`
	// 首次金额，如首重10元
	FirstFee int64 `protobuf:"varint,5,opt,name=FirstFee,json=firstFee,proto3" json:"FirstFee"`
	// 增加数值，如续重1kg
	AddUnit int32 `protobuf:"varint,6,opt,name=AddUnit,json=addUnit,proto3" json:"AddUnit"`
	// 增加产生费用，如续重1kg 10元
	AddFee               int64    `protobuf:"varint,7,opt,name=AddFee,json=addFee,proto3" json:"AddFee"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SExpressAreaTemplate) Reset()         { *m = SExpressAreaTemplate{} }
func (m *SExpressAreaTemplate) String() string { return proto.CompactTextString(m) }
func (*SExpressAreaTemplate) ProtoMessage()    {}
func (*SExpressAreaTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_406b04ca6f4da78c, []int{6}
}
func (m *SExpressAreaTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SExpressAreaTemplate.Unmarshal(m, b)
}
func (m *SExpressAreaTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SExpressAreaTemplate.Marshal(b, m, deterministic)
}
func (dst *SExpressAreaTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SExpressAreaTemplate.Merge(dst, src)
}
func (m *SExpressAreaTemplate) XXX_Size() int {
	return xxx_messageInfo_SExpressAreaTemplate.Size(m)
}
func (m *SExpressAreaTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_SExpressAreaTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_SExpressAreaTemplate proto.InternalMessageInfo

func (m *SExpressAreaTemplate) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SExpressAreaTemplate) GetCodeList() string {
	if m != nil {
		return m.CodeList
	}
	return ""
}

func (m *SExpressAreaTemplate) GetNameList() string {
	if m != nil {
		return m.NameList
	}
	return ""
}

func (m *SExpressAreaTemplate) GetFirstUnit() int32 {
	if m != nil {
		return m.FirstUnit
	}
	return 0
}

func (m *SExpressAreaTemplate) GetFirstFee() int64 {
	if m != nil {
		return m.FirstFee
	}
	return 0
}

func (m *SExpressAreaTemplate) GetAddUnit() int32 {
	if m != nil {
		return m.AddUnit
	}
	return 0
}

func (m *SExpressAreaTemplate) GetAddFee() int64 {
	if m != nil {
		return m.AddFee
	}
	return 0
}

type SaveAreaExpTemplateRequest struct {
	SellerId             int64                 `protobuf:"varint,1,opt,name=SellerId,json=sellerId,proto3" json:"SellerId"`
	TemplateId           int64                 `protobuf:"varint,2,opt,name=TemplateId,json=templateId,proto3" json:"TemplateId"`
	Value                *SExpressAreaTemplate `protobuf:"bytes,3,opt,name=Value,json=value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SaveAreaExpTemplateRequest) Reset()         { *m = SaveAreaExpTemplateRequest{} }
func (m *SaveAreaExpTemplateRequest) String() string { return proto.CompactTextString(m) }
func (*SaveAreaExpTemplateRequest) ProtoMessage()    {}
func (*SaveAreaExpTemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_406b04ca6f4da78c, []int{7}
}
func (m *SaveAreaExpTemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveAreaExpTemplateRequest.Unmarshal(m, b)
}
func (m *SaveAreaExpTemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveAreaExpTemplateRequest.Marshal(b, m, deterministic)
}
func (dst *SaveAreaExpTemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveAreaExpTemplateRequest.Merge(dst, src)
}
func (m *SaveAreaExpTemplateRequest) XXX_Size() int {
	return xxx_messageInfo_SaveAreaExpTemplateRequest.Size(m)
}
func (m *SaveAreaExpTemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveAreaExpTemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveAreaExpTemplateRequest proto.InternalMessageInfo

func (m *SaveAreaExpTemplateRequest) GetSellerId() int64 {
	if m != nil {
		return m.SellerId
	}
	return 0
}

func (m *SaveAreaExpTemplateRequest) GetTemplateId() int64 {
	if m != nil {
		return m.TemplateId
	}
	return 0
}

func (m *SaveAreaExpTemplateRequest) GetValue() *SExpressAreaTemplate {
	if m != nil {
		return m.Value
	}
	return nil
}

type AreaTemplateId struct {
	SellerId             int64    `protobuf:"varint,1,opt,name=SellerId,json=sellerId,proto3" json:"SellerId"`
	TemplateId           int64    `protobuf:"varint,2,opt,name=TemplateId,json=templateId,proto3" json:"TemplateId"`
	AreaTemplateId       int64    `protobuf:"varint,3,opt,name=AreaTemplateId,json=areaTemplateId,proto3" json:"AreaTemplateId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AreaTemplateId) Reset()         { *m = AreaTemplateId{} }
func (m *AreaTemplateId) String() string { return proto.CompactTextString(m) }
func (*AreaTemplateId) ProtoMessage()    {}
func (*AreaTemplateId) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_406b04ca6f4da78c, []int{8}
}
func (m *AreaTemplateId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AreaTemplateId.Unmarshal(m, b)
}
func (m *AreaTemplateId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AreaTemplateId.Marshal(b, m, deterministic)
}
func (dst *AreaTemplateId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AreaTemplateId.Merge(dst, src)
}
func (m *AreaTemplateId) XXX_Size() int {
	return xxx_messageInfo_AreaTemplateId.Size(m)
}
func (m *AreaTemplateId) XXX_DiscardUnknown() {
	xxx_messageInfo_AreaTemplateId.DiscardUnknown(m)
}

var xxx_messageInfo_AreaTemplateId proto.InternalMessageInfo

func (m *AreaTemplateId) GetSellerId() int64 {
	if m != nil {
		return m.SellerId
	}
	return 0
}

func (m *AreaTemplateId) GetTemplateId() int64 {
	if m != nil {
		return m.TemplateId
	}
	return 0
}

func (m *AreaTemplateId) GetAreaTemplateId() int64 {
	if m != nil {
		return m.AreaTemplateId
	}
	return 0
}

type SaveTemplateResponse struct {
	ErrCode              int64    `protobuf:"varint,1,opt,name=ErrCode,json=errCode,proto3" json:"ErrCode"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,json=errMsg,proto3" json:"ErrMsg"`
	TemplateId           int64    `protobuf:"varint,3,opt,name=TemplateId,json=templateId,proto3" json:"TemplateId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveTemplateResponse) Reset()         { *m = SaveTemplateResponse{} }
func (m *SaveTemplateResponse) String() string { return proto.CompactTextString(m) }
func (*SaveTemplateResponse) ProtoMessage()    {}
func (*SaveTemplateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_express_service_406b04ca6f4da78c, []int{9}
}
func (m *SaveTemplateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveTemplateResponse.Unmarshal(m, b)
}
func (m *SaveTemplateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveTemplateResponse.Marshal(b, m, deterministic)
}
func (dst *SaveTemplateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveTemplateResponse.Merge(dst, src)
}
func (m *SaveTemplateResponse) XXX_Size() int {
	return xxx_messageInfo_SaveTemplateResponse.Size(m)
}
func (m *SaveTemplateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveTemplateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveTemplateResponse proto.InternalMessageInfo

func (m *SaveTemplateResponse) GetErrCode() int64 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *SaveTemplateResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *SaveTemplateResponse) GetTemplateId() int64 {
	if m != nil {
		return m.TemplateId
	}
	return 0
}

func init() {
	proto.RegisterType((*ExpressProviderListResponse)(nil), "ExpressProviderListResponse")
	proto.RegisterType((*SExpressProvider)(nil), "SExpressProvider")
	proto.RegisterType((*SExpressTemplate)(nil), "SExpressTemplate")
	proto.RegisterType((*ExpressTemplateId)(nil), "ExpressTemplateId")
	proto.RegisterType((*GetTemplatesRequest)(nil), "GetTemplatesRequest")
	proto.RegisterType((*ExpressTemplateListResponse)(nil), "ExpressTemplateListResponse")
	proto.RegisterType((*SExpressAreaTemplate)(nil), "SExpressAreaTemplate")
	proto.RegisterType((*SaveAreaExpTemplateRequest)(nil), "SaveAreaExpTemplateRequest")
	proto.RegisterType((*AreaTemplateId)(nil), "AreaTemplateId")
	proto.RegisterType((*SaveTemplateResponse)(nil), "SaveTemplateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExpressServiceClient is the client API for ExpressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExpressServiceClient interface {
	// 获取快递公司
	GetExpressProvider(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SExpressProvider, error)
	// 保存快递公司
	SaveExpressProvider(ctx context.Context, in *SExpressProvider, opts ...grpc.CallOption) (*Result, error)
	// 获取可用的快递公司
	GetProviders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ExpressProviderListResponse, error)
	// 保存快递模板
	SaveTemplate(ctx context.Context, in *SExpressTemplate, opts ...grpc.CallOption) (*SaveTemplateResponse, error)
	// 获取单个快递模板
	GetTemplate(ctx context.Context, in *ExpressTemplateId, opts ...grpc.CallOption) (*SExpressTemplate, error)
	// 获取卖家的快递模板
	GetTemplates(ctx context.Context, in *GetTemplatesRequest, opts ...grpc.CallOption) (*ExpressTemplateListResponse, error)
	// 删除模板
	DeleteTemplate(ctx context.Context, in *ExpressTemplateId, opts ...grpc.CallOption) (*Result, error)
	// 保存地区快递模板
	SaveAreaTemplate(ctx context.Context, in *SaveAreaExpTemplateRequest, opts ...grpc.CallOption) (*Result, error)
	// 删除模板地区设定
	DeleteAreaTemplate(ctx context.Context, in *AreaTemplateId, opts ...grpc.CallOption) (*Result, error)
}

type expressServiceClient struct {
	cc *grpc.ClientConn
}

func NewExpressServiceClient(cc *grpc.ClientConn) ExpressServiceClient {
	return &expressServiceClient{cc}
}

func (c *expressServiceClient) GetExpressProvider(ctx context.Context, in *IdOrName, opts ...grpc.CallOption) (*SExpressProvider, error) {
	out := new(SExpressProvider)
	err := c.cc.Invoke(ctx, "/ExpressService/GetExpressProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) SaveExpressProvider(ctx context.Context, in *SExpressProvider, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ExpressService/SaveExpressProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) GetProviders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ExpressProviderListResponse, error) {
	out := new(ExpressProviderListResponse)
	err := c.cc.Invoke(ctx, "/ExpressService/GetProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) SaveTemplate(ctx context.Context, in *SExpressTemplate, opts ...grpc.CallOption) (*SaveTemplateResponse, error) {
	out := new(SaveTemplateResponse)
	err := c.cc.Invoke(ctx, "/ExpressService/SaveTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) GetTemplate(ctx context.Context, in *ExpressTemplateId, opts ...grpc.CallOption) (*SExpressTemplate, error) {
	out := new(SExpressTemplate)
	err := c.cc.Invoke(ctx, "/ExpressService/GetTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) GetTemplates(ctx context.Context, in *GetTemplatesRequest, opts ...grpc.CallOption) (*ExpressTemplateListResponse, error) {
	out := new(ExpressTemplateListResponse)
	err := c.cc.Invoke(ctx, "/ExpressService/GetTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) DeleteTemplate(ctx context.Context, in *ExpressTemplateId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ExpressService/DeleteTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) SaveAreaTemplate(ctx context.Context, in *SaveAreaExpTemplateRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ExpressService/SaveAreaTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expressServiceClient) DeleteAreaTemplate(ctx context.Context, in *AreaTemplateId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ExpressService/DeleteAreaTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpressServiceServer is the server API for ExpressService service.
type ExpressServiceServer interface {
	// 获取快递公司
	GetExpressProvider(context.Context, *IdOrName) (*SExpressProvider, error)
	// 保存快递公司
	SaveExpressProvider(context.Context, *SExpressProvider) (*Result, error)
	// 获取可用的快递公司
	GetProviders(context.Context, *Empty) (*ExpressProviderListResponse, error)
	// 保存快递模板
	SaveTemplate(context.Context, *SExpressTemplate) (*SaveTemplateResponse, error)
	// 获取单个快递模板
	GetTemplate(context.Context, *ExpressTemplateId) (*SExpressTemplate, error)
	// 获取卖家的快递模板
	GetTemplates(context.Context, *GetTemplatesRequest) (*ExpressTemplateListResponse, error)
	// 删除模板
	DeleteTemplate(context.Context, *ExpressTemplateId) (*Result, error)
	// 保存地区快递模板
	SaveAreaTemplate(context.Context, *SaveAreaExpTemplateRequest) (*Result, error)
	// 删除模板地区设定
	DeleteAreaTemplate(context.Context, *AreaTemplateId) (*Result, error)
}

func RegisterExpressServiceServer(s *grpc.Server, srv ExpressServiceServer) {
	s.RegisterService(&_ExpressService_serviceDesc, srv)
}

func _ExpressService_GetExpressProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdOrName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).GetExpressProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/GetExpressProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).GetExpressProvider(ctx, req.(*IdOrName))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_SaveExpressProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SExpressProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).SaveExpressProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/SaveExpressProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).SaveExpressProvider(ctx, req.(*SExpressProvider))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_GetProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).GetProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/GetProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).GetProviders(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_SaveTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SExpressTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).SaveTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/SaveTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).SaveTemplate(ctx, req.(*SExpressTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_GetTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpressTemplateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).GetTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/GetTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).GetTemplate(ctx, req.(*ExpressTemplateId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_GetTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).GetTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/GetTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).GetTemplates(ctx, req.(*GetTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_DeleteTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpressTemplateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).DeleteTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/DeleteTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).DeleteTemplate(ctx, req.(*ExpressTemplateId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_SaveAreaTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveAreaExpTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).SaveAreaTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/SaveAreaTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).SaveAreaTemplate(ctx, req.(*SaveAreaExpTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpressService_DeleteAreaTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AreaTemplateId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpressServiceServer).DeleteAreaTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExpressService/DeleteAreaTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpressServiceServer).DeleteAreaTemplate(ctx, req.(*AreaTemplateId))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExpressService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ExpressService",
	HandlerType: (*ExpressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExpressProvider",
			Handler:    _ExpressService_GetExpressProvider_Handler,
		},
		{
			MethodName: "SaveExpressProvider",
			Handler:    _ExpressService_SaveExpressProvider_Handler,
		},
		{
			MethodName: "GetProviders",
			Handler:    _ExpressService_GetProviders_Handler,
		},
		{
			MethodName: "SaveTemplate",
			Handler:    _ExpressService_SaveTemplate_Handler,
		},
		{
			MethodName: "GetTemplate",
			Handler:    _ExpressService_GetTemplate_Handler,
		},
		{
			MethodName: "GetTemplates",
			Handler:    _ExpressService_GetTemplates_Handler,
		},
		{
			MethodName: "DeleteTemplate",
			Handler:    _ExpressService_DeleteTemplate_Handler,
		},
		{
			MethodName: "SaveAreaTemplate",
			Handler:    _ExpressService_SaveAreaTemplate_Handler,
		},
		{
			MethodName: "DeleteAreaTemplate",
			Handler:    _ExpressService_DeleteAreaTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "express_service.proto",
}

func init() {
	proto.RegisterFile("express_service.proto", fileDescriptor_express_service_406b04ca6f4da78c)
}

var fileDescriptor_express_service_406b04ca6f4da78c = []byte{
	// 744 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcf, 0x6e, 0xfb, 0x44,
	0x10, 0x8e, 0xe3, 0xf8, 0x4f, 0x26, 0x55, 0xf8, 0x75, 0x9b, 0x22, 0xcb, 0xad, 0x20, 0xf2, 0x01,
	0x22, 0x21, 0x6d, 0x51, 0x8a, 0x38, 0x94, 0x53, 0x4b, 0x93, 0x28, 0x52, 0xa1, 0xc8, 0x01, 0x0e,
	0x5c, 0x90, 0x13, 0x4f, 0x53, 0x4b, 0x4e, 0x6c, 0x76, 0x37, 0x51, 0xfb, 0x00, 0x70, 0xe2, 0x71,
	0x78, 0x0a, 0x9e, 0x0a, 0xed, 0xda, 0x4e, 0x1c, 0x27, 0x0d, 0x95, 0x7e, 0x27, 0xeb, 0x9b, 0xd9,
	0xd9, 0xfd, 0xe6, 0x9b, 0x3f, 0x86, 0x73, 0x7c, 0x49, 0x19, 0x72, 0xfe, 0x3b, 0x47, 0xb6, 0x8e,
	0x66, 0x48, 0x53, 0x96, 0x88, 0xc4, 0x3d, 0x99, 0xc7, 0xc9, 0x34, 0x88, 0x33, 0xe4, 0x0d, 0xe1,
	0x62, 0x90, 0x1d, 0xfb, 0x89, 0x25, 0xeb, 0x28, 0x44, 0xf6, 0x10, 0x71, 0xe1, 0x23, 0x4f, 0x93,
	0x25, 0x47, 0xf2, 0x25, 0x18, 0xbf, 0x06, 0xf1, 0x0a, 0x1d, 0xad, 0xab, 0xf7, 0x5a, 0xfd, 0x53,
	0x3a, 0xa9, 0x9c, 0xf6, 0x8d, 0xb5, 0xf4, 0x7b, 0xff, 0x68, 0xf0, 0xa1, 0xea, 0x23, 0x6d, 0xa8,
	0x8f, 0x43, 0x47, 0xeb, 0x6a, 0x3d, 0xdd, 0xaf, 0x47, 0x21, 0x21, 0xd0, 0xf8, 0x31, 0x58, 0xa0,
	0x53, 0xef, 0x6a, 0xbd, 0xa6, 0xdf, 0x58, 0x06, 0x0b, 0x24, 0x9f, 0x82, 0xf9, 0x80, 0x42, 0x20,
	0x73, 0x74, 0x65, 0x35, 0x63, 0x85, 0xc8, 0x25, 0x34, 0x47, 0x2c, 0x59, 0xa5, 0xc3, 0x38, 0x98,
	0x3b, 0x0d, 0xe5, 0x6a, 0xce, 0x0b, 0x83, 0xbc, 0xe9, 0xfb, 0x24, 0x44, 0xc7, 0xc8, 0x6e, 0x9a,
	0x25, 0x21, 0x12, 0x07, 0xac, 0xdb, 0x34, 0x52, 0x66, 0x53, 0x99, 0xad, 0x20, 0x83, 0xd2, 0x33,
	0x58, 0x06, 0xd3, 0x18, 0x43, 0xc7, 0xea, 0x6a, 0x3d, 0xdb, 0xb7, 0x30, 0x83, 0xde, 0x5f, 0xf5,
	0x2d, 0xed, 0x9f, 0x71, 0x91, 0xc6, 0x81, 0xc0, 0x3d, 0xda, 0x2e, 0xd8, 0x13, 0x8c, 0x63, 0x64,
	0xe3, 0x50, 0x51, 0xd7, 0x7d, 0x9b, 0xe7, 0x78, 0x93, 0x92, 0xbe, 0x9b, 0xd2, 0x98, 0x0f, 0x19,
	0xa2, 0xe2, 0x6d, 0xfb, 0x66, 0xa4, 0x10, 0xe9, 0x80, 0x71, 0x17, 0xf0, 0x88, 0x2b, 0xd6, 0x86,
	0x6f, 0x4c, 0x25, 0x90, 0x89, 0x0e, 0x23, 0xc6, 0xc5, 0x2f, 0xcb, 0x48, 0x28, 0xe2, 0x86, 0xdf,
	0x7c, 0x2a, 0x0c, 0xf2, 0x6d, 0xe5, 0x1d, 0x22, 0x2a, 0xee, 0xba, 0x6f, 0x3f, 0xe5, 0x58, 0x25,
	0x1c, 0x86, 0x2a, 0xce, 0x56, 0x71, 0x56, 0x90, 0x41, 0xc9, 0xe0, 0x36, 0x0c, 0x65, 0x4c, 0x53,
	0xc5, 0x98, 0x81, 0x42, 0x65, 0x21, 0x60, 0x57, 0x88, 0x47, 0x38, 0xad, 0xc8, 0x30, 0xde, 0x4d,
	0x5c, 0xab, 0x24, 0xfe, 0x19, 0xc0, 0xf6, 0x64, 0x2e, 0x0b, 0x88, 0x8d, 0xc5, 0x9b, 0xc0, 0xd9,
	0x08, 0x45, 0x71, 0x84, 0xfb, 0xf8, 0xc7, 0x0a, 0xb9, 0x38, 0x7a, 0x65, 0x17, 0x5a, 0x8f, 0xcb,
	0xf8, 0xb5, 0x60, 0x58, 0x57, 0x0c, 0x5b, 0xc9, 0xd6, 0x54, 0xea, 0xd6, 0xe2, 0xe2, 0xf7, 0x75,
	0x6b, 0x71, 0xba, 0xe8, 0xd6, 0x7f, 0x35, 0xe8, 0x14, 0xbe, 0x5b, 0x86, 0xc1, 0xb1, 0xd2, 0xcb,
	0x0e, 0x92, 0xaf, 0xe4, 0x5d, 0x6b, 0xcf, 0x72, 0x2c, 0x7d, 0xb2, 0xf4, 0xca, 0x97, 0x95, 0xdf,
	0x5e, 0xe6, 0x78, 0xb7, 0xa8, 0x8d, 0x63, 0x45, 0x35, 0xde, 0x2e, 0xaa, 0xf9, 0x56, 0x51, 0xad,
	0x72, 0x51, 0xbd, 0x3f, 0x35, 0x70, 0x27, 0xc1, 0x1a, 0x65, 0x22, 0x83, 0x97, 0x74, 0x93, 0xeb,
	0x3b, 0x14, 0xff, 0x9f, 0x22, 0x92, 0xaf, 0x0a, 0x41, 0x65, 0x7e, 0xad, 0xfe, 0x39, 0x3d, 0x24,
	0x5a, 0x21, 0xaa, 0x80, 0x76, 0xd9, 0xfc, 0x71, 0xfd, 0x43, 0xbe, 0xa8, 0xde, 0xa6, 0x38, 0xe8,
	0x7e, 0x3b, 0xd8, 0xb1, 0x7a, 0xcf, 0xd0, 0x91, 0xc9, 0x6f, 0xb3, 0xce, 0x7b, 0x41, 0xb6, 0x3a,
	0x63, 0x6a, 0x1b, 0x64, 0x4f, 0x5b, 0x98, 0x41, 0xa9, 0xe3, 0x80, 0xb1, 0x1f, 0xf8, 0x3c, 0xaf,
	0xa8, 0x89, 0x0a, 0x55, 0x18, 0xe9, 0x55, 0x46, 0xfd, 0xbf, 0x1b, 0xd0, 0xce, 0xd3, 0x9f, 0x64,
	0x1b, 0x95, 0x7c, 0x03, 0x64, 0x84, 0xa2, 0xba, 0xf6, 0x9a, 0x74, 0x1c, 0x3e, 0x32, 0xd9, 0x1b,
	0xee, 0xfe, 0xc2, 0xf4, 0x6a, 0xe4, 0x1a, 0xce, 0x24, 0xe5, 0x6a, 0xd8, 0xfe, 0x59, 0xd7, 0xa2,
	0x3e, 0xf2, 0x55, 0x2c, 0xbc, 0x1a, 0xf9, 0x16, 0x4e, 0x46, 0x28, 0x0a, 0x0f, 0x27, 0x26, 0x1d,
	0x2c, 0x52, 0xf1, 0xea, 0x5e, 0xd2, 0x23, 0xfb, 0xdb, 0xab, 0x91, 0x1b, 0x38, 0x29, 0xeb, 0x43,
	0xf6, 0x87, 0xc2, 0x3d, 0xa7, 0x87, 0x14, 0x54, 0x6f, 0xb6, 0x4a, 0x33, 0x4c, 0x08, 0xdd, 0x5b,
	0x11, 0xee, 0xfe, 0x75, 0x5e, 0x8d, 0xdc, 0x2b, 0xae, 0x9b, 0xd9, 0x27, 0x1d, 0x7a, 0x60, 0x15,
	0x6c, 0x99, 0x1f, 0x9a, 0x65, 0xaf, 0x46, 0xae, 0xa0, 0x7d, 0x8f, 0x31, 0x0a, 0x3c, 0x4a, 0xa0,
	0x24, 0xd1, 0x0d, 0x7c, 0x28, 0xe6, 0x60, 0x13, 0x72, 0x41, 0xdf, 0x1e, 0x8d, 0x72, 0xec, 0xd7,
	0x40, 0xb2, 0xc7, 0x76, 0xa2, 0x3f, 0xa1, 0xbb, 0x3d, 0x58, 0x8a, 0xb8, 0xfb, 0x1c, 0xce, 0x66,
	0xc9, 0x82, 0xce, 0x23, 0xf1, 0xbc, 0x9a, 0xd2, 0x79, 0xd2, 0x4f, 0x28, 0x4b, 0x67, 0xbf, 0xd9,
	0xf4, 0xea, 0x3b, 0xf5, 0x6b, 0x9d, 0x9a, 0xea, 0x73, 0xfd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xfc, 0x47, 0x0b, 0x5e, 0x88, 0x07, 0x00, 0x00,
}
