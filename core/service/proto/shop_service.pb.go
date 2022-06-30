// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shop_service.proto

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

// * 店铺标志
type EShopFlag int32

const (
	EShopFlag__5 EShopFlag = 0
	// * 自营
	EShopFlag_SelfSale EShopFlag = 1
)

var EShopFlag_name = map[int32]string{
	0: "_5",
	1: "SelfSale",
}
var EShopFlag_value = map[string]int32{
	"_5":       0,
	"SelfSale": 1,
}

func (x EShopFlag) String() string {
	return proto.EnumName(EShopFlag_name, int32(x))
}
func (EShopFlag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_fbbfbc92aee29573, []int{0}
}

type TurnShopRequest struct {
	ShopId               int64    `protobuf:"zigzag64,1,opt,name=ShopId,json=shopId,proto3" json:"ShopId"`
	On                   bool     `protobuf:"varint,2,opt,name=On,json=on,proto3" json:"On"`
	Reason               string   `protobuf:"bytes,3,opt,name=Reason,json=reason,proto3" json:"Reason"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TurnShopRequest) Reset()         { *m = TurnShopRequest{} }
func (m *TurnShopRequest) String() string { return proto.CompactTextString(m) }
func (*TurnShopRequest) ProtoMessage()    {}
func (*TurnShopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_fbbfbc92aee29573, []int{0}
}
func (m *TurnShopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TurnShopRequest.Unmarshal(m, b)
}
func (m *TurnShopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TurnShopRequest.Marshal(b, m, deterministic)
}
func (dst *TurnShopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TurnShopRequest.Merge(dst, src)
}
func (m *TurnShopRequest) XXX_Size() int {
	return xxx_messageInfo_TurnShopRequest.Size(m)
}
func (m *TurnShopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TurnShopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TurnShopRequest proto.InternalMessageInfo

func (m *TurnShopRequest) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *TurnShopRequest) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

func (m *TurnShopRequest) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

// 店铺
type SShop struct {
	// * 店铺编号
	Id int64 `protobuf:"varint,1,opt,name=Id,json=id,proto3" json:"Id"`
	// * 商户编号
	MerchantId int64 `protobuf:"varint,2,opt,name=MerchantId,json=merchantId,proto3" json:"MerchantId"`
	// * 卖家会员编号
	SellerMid int64 `protobuf:"varint,3,opt,name=SellerMid,json=sellerMid,proto3" json:"SellerMid"`
	// * 店铺名称
	ShopName string `protobuf:"bytes,4,opt,name=ShopName,json=shopName,proto3" json:"ShopName"`
	// 店铺标题
	ShopTitle string `protobuf:"bytes,5,opt,name=ShopTitle,json=shopTitle,proto3" json:"ShopTitle"`
	// 店铺公告
	ShopNotice string `protobuf:"bytes,6,opt,name=ShopNotice,json=shopNotice,proto3" json:"ShopNotice"`
	// 标志
	Flag int32 `protobuf:"varint,7,opt,name=Flag,json=flag,proto3" json:"Flag"`
	// * 店铺标志
	Logo string `protobuf:"bytes,8,opt,name=Logo,json=logo,proto3" json:"Logo"`
	// * 个性化域名
	Alias string `protobuf:"bytes,9,opt,name=Alias,json=alias,proto3" json:"Alias"`
	// * 自定义 域名
	Host string `protobuf:"bytes,10,opt,name=Host,json=host,proto3" json:"Host"`
	// * 电话
	Telephone string `protobuf:"bytes,11,opt,name=Telephone,json=telephone,proto3" json:"Telephone"`
	// * 状态
	State                int32    `protobuf:"varint,12,opt,name=State,json=state,proto3" json:"State"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SShop) Reset()         { *m = SShop{} }
func (m *SShop) String() string { return proto.CompactTextString(m) }
func (*SShop) ProtoMessage()    {}
func (*SShop) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_fbbfbc92aee29573, []int{1}
}
func (m *SShop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SShop.Unmarshal(m, b)
}
func (m *SShop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SShop.Marshal(b, m, deterministic)
}
func (dst *SShop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SShop.Merge(dst, src)
}
func (m *SShop) XXX_Size() int {
	return xxx_messageInfo_SShop.Size(m)
}
func (m *SShop) XXX_DiscardUnknown() {
	xxx_messageInfo_SShop.DiscardUnknown(m)
}

var xxx_messageInfo_SShop proto.InternalMessageInfo

func (m *SShop) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SShop) GetMerchantId() int64 {
	if m != nil {
		return m.MerchantId
	}
	return 0
}

func (m *SShop) GetSellerMid() int64 {
	if m != nil {
		return m.SellerMid
	}
	return 0
}

func (m *SShop) GetShopName() string {
	if m != nil {
		return m.ShopName
	}
	return ""
}

func (m *SShop) GetShopTitle() string {
	if m != nil {
		return m.ShopTitle
	}
	return ""
}

func (m *SShop) GetShopNotice() string {
	if m != nil {
		return m.ShopNotice
	}
	return ""
}

func (m *SShop) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *SShop) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *SShop) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *SShop) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *SShop) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

func (m *SShop) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

// 店铺设置
type SShopConfig struct {
	// * 店铺标志
	Logo string `protobuf:"bytes,4,opt,name=Logo,json=logo,proto3" json:"Logo"`
	// * 自定义 域名
	Host string `protobuf:"bytes,5,opt,name=Host,json=host,proto3" json:"Host"`
	// * 个性化域名
	Alias string `protobuf:"bytes,6,opt,name=Alias,json=alias,proto3" json:"Alias"`
	// * 电话
	Tel                  string   `protobuf:"bytes,7,opt,name=Tel,json=tel,proto3" json:"Tel"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SShopConfig) Reset()         { *m = SShopConfig{} }
func (m *SShopConfig) String() string { return proto.CompactTextString(m) }
func (*SShopConfig) ProtoMessage()    {}
func (*SShopConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_fbbfbc92aee29573, []int{2}
}
func (m *SShopConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SShopConfig.Unmarshal(m, b)
}
func (m *SShopConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SShopConfig.Marshal(b, m, deterministic)
}
func (dst *SShopConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SShopConfig.Merge(dst, src)
}
func (m *SShopConfig) XXX_Size() int {
	return xxx_messageInfo_SShopConfig.Size(m)
}
func (m *SShopConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SShopConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SShopConfig proto.InternalMessageInfo

func (m *SShopConfig) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *SShopConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *SShopConfig) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *SShopConfig) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

// 获取店铺
type GetShopIdRequest struct {
	// 店铺编号
	ShopId               int64    `protobuf:"varint,1,opt,name=ShopId,json=shopId,proto3" json:"ShopId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShopIdRequest) Reset()         { *m = GetShopIdRequest{} }
func (m *GetShopIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetShopIdRequest) ProtoMessage()    {}
func (*GetShopIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_fbbfbc92aee29573, []int{3}
}
func (m *GetShopIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShopIdRequest.Unmarshal(m, b)
}
func (m *GetShopIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShopIdRequest.Marshal(b, m, deterministic)
}
func (dst *GetShopIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShopIdRequest.Merge(dst, src)
}
func (m *GetShopIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetShopIdRequest.Size(m)
}
func (m *GetShopIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShopIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShopIdRequest proto.InternalMessageInfo

func (m *GetShopIdRequest) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

// * 店铺别名请求
type ShopAliasRequest struct {
	// 店铺别名
	ShopAlias            string   `protobuf:"bytes,1,opt,name=ShopAlias,json=shopAlias,proto3" json:"ShopAlias"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShopAliasRequest) Reset()         { *m = ShopAliasRequest{} }
func (m *ShopAliasRequest) String() string { return proto.CompactTextString(m) }
func (*ShopAliasRequest) ProtoMessage()    {}
func (*ShopAliasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_fbbfbc92aee29573, []int{4}
}
func (m *ShopAliasRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShopAliasRequest.Unmarshal(m, b)
}
func (m *ShopAliasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShopAliasRequest.Marshal(b, m, deterministic)
}
func (dst *ShopAliasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShopAliasRequest.Merge(dst, src)
}
func (m *ShopAliasRequest) XXX_Size() int {
	return xxx_messageInfo_ShopAliasRequest.Size(m)
}
func (m *ShopAliasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShopAliasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShopAliasRequest proto.InternalMessageInfo

func (m *ShopAliasRequest) GetShopAlias() string {
	if m != nil {
		return m.ShopAlias
	}
	return ""
}

// 门店编号
type StoreId struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,json=value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreId) Reset()         { *m = StoreId{} }
func (m *StoreId) String() string { return proto.CompactTextString(m) }
func (*StoreId) ProtoMessage()    {}
func (*StoreId) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_fbbfbc92aee29573, []int{5}
}
func (m *StoreId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreId.Unmarshal(m, b)
}
func (m *StoreId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreId.Marshal(b, m, deterministic)
}
func (dst *StoreId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreId.Merge(dst, src)
}
func (m *StoreId) XXX_Size() int {
	return xxx_messageInfo_StoreId.Size(m)
}
func (m *StoreId) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreId.DiscardUnknown(m)
}

var xxx_messageInfo_StoreId proto.InternalMessageInfo

func (m *StoreId) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// 店铺
type SStore struct {
	Id           int64  `protobuf:"zigzag64,1,opt,name=Id,json=id,proto3" json:"Id"`
	MerchantId   int64  `protobuf:"zigzag64,2,opt,name=MerchantId,json=merchantId,proto3" json:"MerchantId"`
	Name         string `protobuf:"bytes,3,opt,name=Name,json=name,proto3" json:"Name"`
	Alias        string `protobuf:"bytes,4,opt,name=Alias,json=alias,proto3" json:"Alias"`
	State        int32  `protobuf:"zigzag32,5,opt,name=State,json=state,proto3" json:"State"`
	OpeningState int32  `protobuf:"zigzag32,8,opt,name=OpeningState,json=openingState,proto3" json:"OpeningState"`
	StorePhone   string `protobuf:"bytes,9,opt,name=StorePhone,json=storePhone,proto3" json:"StorePhone"`
	StoreNotice  string `protobuf:"bytes,11,opt,name=StoreNotice,json=storeNotice,proto3" json:"StoreNotice"`
	Province     int32  `protobuf:"varint,12,opt,name=Province,json=province,proto3" json:"Province"`
	City         int32  `protobuf:"varint,13,opt,name=City,json=city,proto3" json:"City"`
	District     int32  `protobuf:"varint,14,opt,name=District,json=district,proto3" json:"District"`
	// 地区名称
	Address string `protobuf:"bytes,15,opt,name=Address,json=address,proto3" json:"Address"`
	// 详细地址
	DetailAddress string `protobuf:"bytes,16,opt,name=DetailAddress,json=detailAddress,proto3" json:"DetailAddress"`
	// 纬度
	Lat float64 `protobuf:"fixed64,17,opt,name=Lat,json=lat,proto3" json:"Lat"`
	// 经度
	Lng float64 `protobuf:"fixed64,18,opt,name=Lng,json=lng,proto3" json:"Lng"`
	// 覆盖范围(公里)
	CoverRadius int32 `protobuf:"varint,19,opt,name=CoverRadius,json=coverRadius,proto3" json:"CoverRadius"`
	// 序号
	SortNum              int32    `protobuf:"varint,20,opt,name=SortNum,json=sortNum,proto3" json:"SortNum"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SStore) Reset()         { *m = SStore{} }
func (m *SStore) String() string { return proto.CompactTextString(m) }
func (*SStore) ProtoMessage()    {}
func (*SStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_fbbfbc92aee29573, []int{6}
}
func (m *SStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SStore.Unmarshal(m, b)
}
func (m *SStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SStore.Marshal(b, m, deterministic)
}
func (dst *SStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SStore.Merge(dst, src)
}
func (m *SStore) XXX_Size() int {
	return xxx_messageInfo_SStore.Size(m)
}
func (m *SStore) XXX_DiscardUnknown() {
	xxx_messageInfo_SStore.DiscardUnknown(m)
}

var xxx_messageInfo_SStore proto.InternalMessageInfo

func (m *SStore) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SStore) GetMerchantId() int64 {
	if m != nil {
		return m.MerchantId
	}
	return 0
}

func (m *SStore) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SStore) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *SStore) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *SStore) GetOpeningState() int32 {
	if m != nil {
		return m.OpeningState
	}
	return 0
}

func (m *SStore) GetStorePhone() string {
	if m != nil {
		return m.StorePhone
	}
	return ""
}

func (m *SStore) GetStoreNotice() string {
	if m != nil {
		return m.StoreNotice
	}
	return ""
}

func (m *SStore) GetProvince() int32 {
	if m != nil {
		return m.Province
	}
	return 0
}

func (m *SStore) GetCity() int32 {
	if m != nil {
		return m.City
	}
	return 0
}

func (m *SStore) GetDistrict() int32 {
	if m != nil {
		return m.District
	}
	return 0
}

func (m *SStore) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SStore) GetDetailAddress() string {
	if m != nil {
		return m.DetailAddress
	}
	return ""
}

func (m *SStore) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *SStore) GetLng() float64 {
	if m != nil {
		return m.Lng
	}
	return 0
}

func (m *SStore) GetCoverRadius() int32 {
	if m != nil {
		return m.CoverRadius
	}
	return 0
}

func (m *SStore) GetSortNum() int32 {
	if m != nil {
		return m.SortNum
	}
	return 0
}

// 检查店铺结果
type CheckShopResponse struct {
	ShopId int64 `protobuf:"varint,1,opt,name=ShopId,json=shopId,proto3" json:"ShopId"`
	// 店铺开通状态,0:未开通 1:已开通 2:待审核 3:审核未通过
	Status               int32    `protobuf:"varint,2,opt,name=Status,json=status,proto3" json:"Status"`
	Remark               string   `protobuf:"bytes,3,opt,name=Remark,json=remark,proto3" json:"Remark"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckShopResponse) Reset()         { *m = CheckShopResponse{} }
func (m *CheckShopResponse) String() string { return proto.CompactTextString(m) }
func (*CheckShopResponse) ProtoMessage()    {}
func (*CheckShopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_shop_service_fbbfbc92aee29573, []int{7}
}
func (m *CheckShopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckShopResponse.Unmarshal(m, b)
}
func (m *CheckShopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckShopResponse.Marshal(b, m, deterministic)
}
func (dst *CheckShopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckShopResponse.Merge(dst, src)
}
func (m *CheckShopResponse) XXX_Size() int {
	return xxx_messageInfo_CheckShopResponse.Size(m)
}
func (m *CheckShopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckShopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckShopResponse proto.InternalMessageInfo

func (m *CheckShopResponse) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *CheckShopResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CheckShopResponse) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func init() {
	proto.RegisterType((*TurnShopRequest)(nil), "TurnShopRequest")
	proto.RegisterType((*SShop)(nil), "SShop")
	proto.RegisterType((*SShopConfig)(nil), "SShopConfig")
	proto.RegisterType((*GetShopIdRequest)(nil), "GetShopIdRequest")
	proto.RegisterType((*ShopAliasRequest)(nil), "ShopAliasRequest")
	proto.RegisterType((*StoreId)(nil), "StoreId")
	proto.RegisterType((*SStore)(nil), "SStore")
	proto.RegisterType((*CheckShopResponse)(nil), "CheckShopResponse")
	proto.RegisterEnum("EShopFlag", EShopFlag_name, EShopFlag_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShopServiceClient is the client API for ShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShopServiceClient interface {
	// * 获取店铺,shopId
	GetShop(ctx context.Context, in *GetShopIdRequest, opts ...grpc.CallOption) (*SShop, error)
	// * 根据别名查询店铺编号
	QueryShopId(ctx context.Context, in *ShopAliasRequest, opts ...grpc.CallOption) (*Int64, error)
	// rpc GetVendorShop_ (Int64) returns (SShop) {}
	// 检查商户是否开通店铺
	CheckMerchantShopState(ctx context.Context, in *MerchantId, opts ...grpc.CallOption) (*CheckShopResponse, error)
	// * 获取门店,storeId
	GetStore(ctx context.Context, in *StoreId, opts ...grpc.CallOption) (*SStore, error)
	// * 根据主机头获取店铺编号,host
	QueryShopByHost(ctx context.Context, in *String, opts ...grpc.CallOption) (*Int64, error)
	// 获取门店
	// rpc GetOfflineShop(1:sint32 shopId)returns(Shop)
	// 打开或关闭商店
	TurnShop(ctx context.Context, in *TurnShopRequest, opts ...grpc.CallOption) (*Result, error)
	// 保存门店
	SaveShop(ctx context.Context, in *SShop, opts ...grpc.CallOption) (*Result, error)
	// 保存门店
	SaveOfflineShop(ctx context.Context, in *SStore, opts ...grpc.CallOption) (*Result, error)
	// 删除商店
	DeleteStore(ctx context.Context, in *StoreId, opts ...grpc.CallOption) (*Result, error)
}

type shopServiceClient struct {
	cc *grpc.ClientConn
}

func NewShopServiceClient(cc *grpc.ClientConn) ShopServiceClient {
	return &shopServiceClient{cc}
}

func (c *shopServiceClient) GetShop(ctx context.Context, in *GetShopIdRequest, opts ...grpc.CallOption) (*SShop, error) {
	out := new(SShop)
	err := c.cc.Invoke(ctx, "/ShopService/GetShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) QueryShopId(ctx context.Context, in *ShopAliasRequest, opts ...grpc.CallOption) (*Int64, error) {
	out := new(Int64)
	err := c.cc.Invoke(ctx, "/ShopService/QueryShopId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) CheckMerchantShopState(ctx context.Context, in *MerchantId, opts ...grpc.CallOption) (*CheckShopResponse, error) {
	out := new(CheckShopResponse)
	err := c.cc.Invoke(ctx, "/ShopService/CheckMerchantShopState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) GetStore(ctx context.Context, in *StoreId, opts ...grpc.CallOption) (*SStore, error) {
	out := new(SStore)
	err := c.cc.Invoke(ctx, "/ShopService/GetStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) QueryShopByHost(ctx context.Context, in *String, opts ...grpc.CallOption) (*Int64, error) {
	out := new(Int64)
	err := c.cc.Invoke(ctx, "/ShopService/QueryShopByHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) TurnShop(ctx context.Context, in *TurnShopRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/TurnShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) SaveShop(ctx context.Context, in *SShop, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/SaveShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) SaveOfflineShop(ctx context.Context, in *SStore, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/SaveOfflineShop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopServiceClient) DeleteStore(ctx context.Context, in *StoreId, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/ShopService/DeleteStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServiceServer is the server API for ShopService service.
type ShopServiceServer interface {
	// * 获取店铺,shopId
	GetShop(context.Context, *GetShopIdRequest) (*SShop, error)
	// * 根据别名查询店铺编号
	QueryShopId(context.Context, *ShopAliasRequest) (*Int64, error)
	// rpc GetVendorShop_ (Int64) returns (SShop) {}
	// 检查商户是否开通店铺
	CheckMerchantShopState(context.Context, *MerchantId) (*CheckShopResponse, error)
	// * 获取门店,storeId
	GetStore(context.Context, *StoreId) (*SStore, error)
	// * 根据主机头获取店铺编号,host
	QueryShopByHost(context.Context, *String) (*Int64, error)
	// 获取门店
	// rpc GetOfflineShop(1:sint32 shopId)returns(Shop)
	// 打开或关闭商店
	TurnShop(context.Context, *TurnShopRequest) (*Result, error)
	// 保存门店
	SaveShop(context.Context, *SShop) (*Result, error)
	// 保存门店
	SaveOfflineShop(context.Context, *SStore) (*Result, error)
	// 删除商店
	DeleteStore(context.Context, *StoreId) (*Result, error)
}

func RegisterShopServiceServer(s *grpc.Server, srv ShopServiceServer) {
	s.RegisterService(&_ShopService_serviceDesc, srv)
}

func _ShopService_GetShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShopIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/GetShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetShop(ctx, req.(*GetShopIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_QueryShopId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShopAliasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).QueryShopId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/QueryShopId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).QueryShopId(ctx, req.(*ShopAliasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_CheckMerchantShopState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).CheckMerchantShopState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/CheckMerchantShopState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).CheckMerchantShopState(ctx, req.(*MerchantId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_GetStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).GetStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/GetStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).GetStore(ctx, req.(*StoreId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_QueryShopByHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).QueryShopByHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/QueryShopByHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).QueryShopByHost(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_TurnShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TurnShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).TurnShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/TurnShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).TurnShop(ctx, req.(*TurnShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_SaveShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SShop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).SaveShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/SaveShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).SaveShop(ctx, req.(*SShop))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_SaveOfflineShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).SaveOfflineShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/SaveOfflineShop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).SaveOfflineShop(ctx, req.(*SStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShopService_DeleteStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServiceServer).DeleteStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShopService/DeleteStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServiceServer).DeleteStore(ctx, req.(*StoreId))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShopService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShopService",
	HandlerType: (*ShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShop",
			Handler:    _ShopService_GetShop_Handler,
		},
		{
			MethodName: "QueryShopId",
			Handler:    _ShopService_QueryShopId_Handler,
		},
		{
			MethodName: "CheckMerchantShopState",
			Handler:    _ShopService_CheckMerchantShopState_Handler,
		},
		{
			MethodName: "GetStore",
			Handler:    _ShopService_GetStore_Handler,
		},
		{
			MethodName: "QueryShopByHost",
			Handler:    _ShopService_QueryShopByHost_Handler,
		},
		{
			MethodName: "TurnShop",
			Handler:    _ShopService_TurnShop_Handler,
		},
		{
			MethodName: "SaveShop",
			Handler:    _ShopService_SaveShop_Handler,
		},
		{
			MethodName: "SaveOfflineShop",
			Handler:    _ShopService_SaveOfflineShop_Handler,
		},
		{
			MethodName: "DeleteStore",
			Handler:    _ShopService_DeleteStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop_service.proto",
}

func init() { proto.RegisterFile("shop_service.proto", fileDescriptor_shop_service_fbbfbc92aee29573) }

var fileDescriptor_shop_service_fbbfbc92aee29573 = []byte{
	// 858 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0x4d, 0x8f, 0xdb, 0x36,
	0x10, 0xf5, 0xa7, 0x2c, 0x8f, 0xbc, 0x59, 0x9b, 0x09, 0x02, 0xc2, 0x68, 0xbb, 0xae, 0x90, 0xb6,
	0xc6, 0x1e, 0xd4, 0x22, 0xfd, 0xb8, 0xe4, 0x94, 0xec, 0xf6, 0x63, 0x81, 0x24, 0x9b, 0x48, 0x8b,
	0x1e, 0x5a, 0x14, 0x01, 0x57, 0xa2, 0x65, 0x22, 0xb4, 0xa8, 0x8a, 0x94, 0x81, 0xfd, 0x61, 0xfd,
	0x13, 0xfd, 0x45, 0x3d, 0x16, 0x1c, 0x4a, 0xb6, 0xba, 0x69, 0x4e, 0xc6, 0x7b, 0x33, 0x9a, 0x79,
	0x9c, 0x79, 0xa4, 0x81, 0xe8, 0xad, 0x2a, 0xdf, 0x69, 0x5e, 0xed, 0x45, 0xca, 0xa3, 0xb2, 0x52,
	0x46, 0x2d, 0x67, 0xb9, 0x54, 0xb7, 0x4c, 0x3a, 0x14, 0xbe, 0x85, 0xd3, 0x9b, 0xba, 0x2a, 0x92,
	0xad, 0x2a, 0x63, 0xfe, 0x67, 0xcd, 0xb5, 0x21, 0x8f, 0xc1, 0xb3, 0xf0, 0x2a, 0xa3, 0xfd, 0x55,
	0x7f, 0x4d, 0x62, 0x4f, 0x23, 0x22, 0x0f, 0x60, 0x70, 0x5d, 0xd0, 0xc1, 0xaa, 0xbf, 0xf6, 0xe3,
	0x81, 0x2a, 0x6c, 0x5e, 0xcc, 0x99, 0x56, 0x05, 0x1d, 0xae, 0xfa, 0xeb, 0x69, 0xec, 0x55, 0x88,
	0xc2, 0xbf, 0x06, 0x30, 0x4e, 0x6c, 0x05, 0xfb, 0x45, 0x53, 0x65, 0x18, 0x0f, 0x44, 0x46, 0x3e,
	0x03, 0x78, 0xc5, 0xab, 0x74, 0xcb, 0x0a, 0x73, 0x95, 0x61, 0xa5, 0x61, 0x0c, 0xbb, 0x03, 0x43,
	0x3e, 0x81, 0x69, 0xc2, 0xa5, 0xe4, 0xd5, 0x2b, 0x91, 0x61, 0xd1, 0x61, 0x3c, 0xd5, 0x2d, 0x41,
	0x96, 0xe0, 0xdb, 0xaa, 0xaf, 0xd9, 0x8e, 0xd3, 0x11, 0x76, 0xf4, 0x75, 0x83, 0xf1, 0xcb, 0xad,
	0x2a, 0x6f, 0x84, 0x91, 0x9c, 0x8e, 0x31, 0x38, 0xd5, 0x2d, 0x61, 0xfb, 0xe2, 0x97, 0xca, 0x88,
	0x94, 0x53, 0x0f, 0xc3, 0xa0, 0x0f, 0x0c, 0x21, 0x30, 0xfa, 0x49, 0xb2, 0x9c, 0x4e, 0x56, 0xfd,
	0xf5, 0x38, 0x1e, 0x6d, 0x24, 0xcb, 0x2d, 0xf7, 0x52, 0xe5, 0x8a, 0xfa, 0x98, 0x3d, 0x92, 0x2a,
	0x57, 0xe4, 0x11, 0x8c, 0x9f, 0x4b, 0xc1, 0x34, 0x9d, 0x22, 0x39, 0x66, 0x16, 0xd8, 0xcc, 0x5f,
	0x94, 0x36, 0x14, 0x5c, 0xe6, 0x56, 0x69, 0x63, 0xf5, 0xdc, 0x70, 0xc9, 0xcb, 0xad, 0x2a, 0x38,
	0x0d, 0x9c, 0x1e, 0xd3, 0x12, 0xb6, 0x4e, 0x62, 0x98, 0xe1, 0x74, 0x86, 0x0d, 0xc7, 0xda, 0x82,
	0xf0, 0x0f, 0x08, 0x70, 0x6c, 0x17, 0xaa, 0xd8, 0x88, 0xa3, 0x80, 0x51, 0x47, 0x40, 0xdb, 0x6a,
	0xdc, 0x69, 0x75, 0x10, 0xe5, 0x75, 0x45, 0xcd, 0x61, 0x78, 0xc3, 0x25, 0x9e, 0x68, 0x1a, 0x0f,
	0x0d, 0x97, 0xe1, 0x39, 0xcc, 0x7f, 0xe6, 0xc6, 0x6d, 0xf6, 0xff, 0x57, 0x3d, 0x6c, 0x57, 0x1d,
	0x7e, 0x03, 0x73, 0xcb, 0x63, 0xdd, 0x36, 0xb7, 0x19, 0xb1, 0xeb, 0xd5, 0x3f, 0x8e, 0x18, 0x89,
	0xf0, 0x0c, 0x26, 0x89, 0x51, 0x15, 0xbf, 0xca, 0xac, 0xa0, 0x5f, 0x99, 0xac, 0x79, 0x53, 0x73,
	0xbc, 0xb7, 0x20, 0xfc, 0x7b, 0x08, 0x5e, 0x82, 0x29, 0x1d, 0x5b, 0x90, 0x8f, 0xd8, 0x82, 0xfc,
	0xc7, 0x16, 0x04, 0x46, 0xb8, 0x74, 0x67, 0xb3, 0x51, 0x61, 0x17, 0x7e, 0x38, 0xf5, 0xa8, 0x7b,
	0xea, 0xc3, 0x60, 0xed, 0x80, 0x16, 0xcd, 0x60, 0x49, 0x08, 0xb3, 0xeb, 0x92, 0x17, 0xa2, 0xc8,
	0x5d, 0xd0, 0xc7, 0xe0, 0x4c, 0x75, 0x38, 0xb4, 0x88, 0x15, 0xf7, 0x06, 0x37, 0x36, 0x6d, 0x2c,
	0x72, 0x60, 0xc8, 0x0a, 0x02, 0x8c, 0x37, 0x1e, 0x72, 0x2b, 0x0d, 0xf4, 0x91, 0xb2, 0xf6, 0x7c,
	0x53, 0xa9, 0xbd, 0x28, 0xd2, 0x76, 0xaf, 0x7e, 0xd9, 0x60, 0x7b, 0x82, 0x0b, 0x61, 0xee, 0xe8,
	0x89, 0x33, 0x58, 0x2a, 0xcc, 0x9d, 0xcd, 0xbf, 0x14, 0xda, 0x54, 0x22, 0x35, 0xf4, 0x81, 0xcb,
	0xcf, 0x1a, 0x4c, 0x28, 0x4c, 0x9e, 0x67, 0x59, 0xc5, 0xb5, 0xa6, 0xa7, 0xd8, 0x69, 0xc2, 0x1c,
	0x24, 0x4f, 0xe0, 0xe4, 0x92, 0x1b, 0x26, 0x64, 0x1b, 0x9f, 0x63, 0xfc, 0x24, 0xeb, 0x92, 0x76,
	0xfb, 0x2f, 0x99, 0xa1, 0x8b, 0x55, 0x7f, 0xdd, 0x8f, 0x87, 0x92, 0x19, 0x64, 0x8a, 0x9c, 0x92,
	0x86, 0x29, 0x72, 0x7b, 0xa2, 0x0b, 0xb5, 0xe7, 0x55, 0xcc, 0x32, 0x51, 0x6b, 0xfa, 0x10, 0x25,
	0x04, 0xe9, 0x91, 0xb2, 0x2a, 0x12, 0x55, 0x99, 0xd7, 0xf5, 0x8e, 0x3e, 0xc2, 0xe8, 0x44, 0x3b,
	0x18, 0xfe, 0x0e, 0x8b, 0x8b, 0x2d, 0x4f, 0xdf, 0xbb, 0x67, 0x43, 0x97, 0xaa, 0xd0, 0xfc, 0x63,
	0x66, 0x42, 0xde, 0x30, 0x53, 0x6b, 0x5c, 0xed, 0x38, 0xf6, 0x34, 0x22, 0xf7, 0x7e, 0xec, 0x58,
	0xf5, 0xfe, 0xf8, 0x7e, 0x58, 0x74, 0xfe, 0x39, 0x4c, 0x7f, 0xb4, 0x85, 0xec, 0x95, 0x24, 0x1e,
	0x0c, 0xde, 0x7d, 0x3f, 0xef, 0x91, 0x19, 0xf8, 0x09, 0x97, 0x9b, 0x84, 0x49, 0x3e, 0xef, 0x3f,
	0xfd, 0x67, 0x00, 0x81, 0x4d, 0x49, 0xdc, 0xcb, 0x46, 0xbe, 0x84, 0x49, 0xe3, 0x6d, 0xb2, 0x88,
	0xee, 0xbb, 0x7c, 0xe9, 0x45, 0x78, 0xaf, 0xc2, 0x1e, 0x39, 0x87, 0xe0, 0x6d, 0xcd, 0xab, 0x3b,
	0x17, 0x27, 0x8b, 0xe8, 0xbe, 0xcb, 0x97, 0x5e, 0x74, 0x55, 0x98, 0x1f, 0xbe, 0x0b, 0x7b, 0xe4,
	0x19, 0x3c, 0xc6, 0x33, 0xb6, 0xd6, 0xc4, 0x7e, 0xe8, 0x95, 0x20, 0x3a, 0xda, 0x75, 0x49, 0xa2,
	0x0f, 0x26, 0x11, 0xf6, 0xc8, 0x19, 0xf8, 0x56, 0x06, 0xda, 0xdd, 0x8f, 0x9a, 0x9b, 0xb1, 0x9c,
	0x44, 0xee, 0x06, 0x84, 0x3d, 0xf2, 0x04, 0x4e, 0x0f, 0x4a, 0x5e, 0xdc, 0xd9, 0x4b, 0x4d, 0x26,
	0x51, 0x62, 0x2a, 0x51, 0xe4, 0x1d, 0x0d, 0x5f, 0x81, 0xdf, 0xbe, 0xce, 0x64, 0x1e, 0xdd, 0x7b,
	0xa8, 0x97, 0x93, 0x28, 0xe6, 0xba, 0x96, 0x26, 0xec, 0x91, 0x4f, 0xc1, 0x4f, 0xd8, 0x9e, 0x63,
	0x62, 0x73, 0xdc, 0x6e, 0xf8, 0x0b, 0x38, 0xb5, 0xe1, 0xeb, 0xcd, 0x46, 0x8a, 0xc2, 0x65, 0xb5,
	0x5a, 0xba, 0x69, 0x21, 0x04, 0x97, 0x5c, 0x72, 0xc3, 0x3f, 0x14, 0xde, 0xe6, 0xbc, 0x38, 0x83,
	0x87, 0xa9, 0xda, 0x45, 0xb9, 0x30, 0xdb, 0xfa, 0x36, 0xca, 0xd5, 0x53, 0x15, 0x55, 0x65, 0xfa,
	0x9b, 0x1f, 0x7d, 0xfd, 0x0c, 0xff, 0x51, 0x6e, 0x3d, 0xfc, 0xf9, 0xf6, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x71, 0x33, 0x71, 0x49, 0x7c, 0x06, 0x00, 0x00,
}
