// Code generated by protoc-gen-go. DO NOT EDIT.
// source: foundation_service.proto

package proto // import "."

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

// * 替换敏感词请求
type ReplaceSensitiveRequest struct {
	Word                 string   `protobuf:"bytes,1,opt,name=Word,proto3" json:"Word"`
	Replacement          string   `protobuf:"bytes,2,opt,name=Replacement,proto3" json:"Replacement"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplaceSensitiveRequest) Reset()         { *m = ReplaceSensitiveRequest{} }
func (m *ReplaceSensitiveRequest) String() string { return proto.CompactTextString(m) }
func (*ReplaceSensitiveRequest) ProtoMessage()    {}
func (*ReplaceSensitiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_foundation_service_eaa6caa00d874cdc, []int{0}
}
func (m *ReplaceSensitiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplaceSensitiveRequest.Unmarshal(m, b)
}
func (m *ReplaceSensitiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplaceSensitiveRequest.Marshal(b, m, deterministic)
}
func (dst *ReplaceSensitiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceSensitiveRequest.Merge(dst, src)
}
func (m *ReplaceSensitiveRequest) XXX_Size() int {
	return xxx_messageInfo_ReplaceSensitiveRequest.Size(m)
}
func (m *ReplaceSensitiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceSensitiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceSensitiveRequest proto.InternalMessageInfo

func (m *ReplaceSensitiveRequest) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func (m *ReplaceSensitiveRequest) GetReplacement() string {
	if m != nil {
		return m.Replacement
	}
	return ""
}

func init() {
	proto.RegisterType((*ReplaceSensitiveRequest)(nil), "ReplaceSensitiveRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FoundationServiceClient is the client API for FoundationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FoundationServiceClient interface {
	// * 检测是否包含敏感词
	CheckSensitive(ctx context.Context, in *String, opts ...grpc.CallOption) (*Bool, error)
	// * 替换敏感词
	ReplaceSensitive(ctx context.Context, in *ReplaceSensitiveRequest, opts ...grpc.CallOption) (*String, error)
	// * 获取短信API凭据, provider 短信服务商, 默认:http
	GetSmsApi(ctx context.Context, in *String, opts ...grpc.CallOption) (*SSmsApi, error)
	// * 保存短信API凭据,@provider 短信服务商, 默认:http
	SaveSmsApi(ctx context.Context, in *SmsApiSaveRequest, opts ...grpc.CallOption) (*Result, error)
	// * 保存面板HOOK数据,这通常是在第三方应用中初始化或调用,参见文档：BoardHooks
	SaveBoardHook(ctx context.Context, in *BoardHookSaveRequest, opts ...grpc.CallOption) (*Result, error)
	// 格式化资源地址并返回
	ResourceUrl(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	// 设置键值
	// rpc SetValue (Pair) returns (Result){}
	// 删除值,key
	// rpc DeleteValue (String) returns (Result){}
	// 根据前缀获取值,prefix
	// rpc GetValuesByPrefix (String) returns (StringMap){}
	// 注册单点登录应用,返回值：
	//   -  1. 成功，并返回token
	//   - -1. 接口地址不正确
	//   - -2. 已经注册
	RegisterApp(ctx context.Context, in *SSsoApp, opts ...grpc.CallOption) (*String, error)
	// 获取应用信息,name
	GetApp(ctx context.Context, in *String, opts ...grpc.CallOption) (*SSsoApp, error)
	// 获取单点登录应用
	GetAllSsoApp(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringListResponse, error)
	// 验证超级用户账号和密码
	SuperValidate(ctx context.Context, in *UserPwd, opts ...grpc.CallOption) (*SuperLoginResponse, error)
	// 保存超级用户账号和密码
	FlushSuperPwd(ctx context.Context, in *UserPwd, opts ...grpc.CallOption) (*Result, error)
	// 创建同步登录的地址,returnUrl
	GetSyncLoginUrl(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	// 获取地区名称
	GetAreaNames(ctx context.Context, in *GetAreaNamesRequest, opts ...grpc.CallOption) (*StringListResponse, error)
	// 获取省市区字符串
	GetAreaString(ctx context.Context, in *AreaStringRequest, opts ...grpc.CallOption) (*String, error)
	// 获取下级区域,code
	GetChildAreas(ctx context.Context, in *Int32, opts ...grpc.CallOption) (*AreaListResponse, error)
	// 获取移动应用设置
	GetMoAppConf(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SMobileAppConfig, error)
	// 保存移动应用设置
	SaveMoAppConf(ctx context.Context, in *SMobileAppConfig, opts ...grpc.CallOption) (*Result, error)
	// 获取微信接口配置
	GetWxApiConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SWxApiConfig, error)
	// 保存微信接口配置
	SaveWxApiConfig(ctx context.Context, in *SWxApiConfig, opts ...grpc.CallOption) (*Result, error)
	// 获取支付平台
	GetPayPlatform(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PaymentPlatformResponse, error)
	// 获取全局商户销售设置
	GetGlobMchSaleConf_(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SGlobMchSaleConf, error)
	// 保存全局商户销售设置
	SaveGlobMchSaleConf_(ctx context.Context, in *SGlobMchSaleConf, opts ...grpc.CallOption) (*Result, error)
}

type foundationServiceClient struct {
	cc *grpc.ClientConn
}

func NewFoundationServiceClient(cc *grpc.ClientConn) FoundationServiceClient {
	return &foundationServiceClient{cc}
}

func (c *foundationServiceClient) CheckSensitive(ctx context.Context, in *String, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/FoundationService/CheckSensitive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) ReplaceSensitive(ctx context.Context, in *ReplaceSensitiveRequest, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/ReplaceSensitive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetSmsApi(ctx context.Context, in *String, opts ...grpc.CallOption) (*SSmsApi, error) {
	out := new(SSmsApi)
	err := c.cc.Invoke(ctx, "/FoundationService/GetSmsApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveSmsApi(ctx context.Context, in *SmsApiSaveRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveSmsApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveBoardHook(ctx context.Context, in *BoardHookSaveRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveBoardHook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) ResourceUrl(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/ResourceUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) RegisterApp(ctx context.Context, in *SSsoApp, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/RegisterApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetApp(ctx context.Context, in *String, opts ...grpc.CallOption) (*SSsoApp, error) {
	out := new(SSsoApp)
	err := c.cc.Invoke(ctx, "/FoundationService/GetApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetAllSsoApp(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringListResponse, error) {
	out := new(StringListResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/GetAllSsoApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SuperValidate(ctx context.Context, in *UserPwd, opts ...grpc.CallOption) (*SuperLoginResponse, error) {
	out := new(SuperLoginResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/SuperValidate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) FlushSuperPwd(ctx context.Context, in *UserPwd, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/FlushSuperPwd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetSyncLoginUrl(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/GetSyncLoginUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetAreaNames(ctx context.Context, in *GetAreaNamesRequest, opts ...grpc.CallOption) (*StringListResponse, error) {
	out := new(StringListResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/GetAreaNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetAreaString(ctx context.Context, in *AreaStringRequest, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/FoundationService/GetAreaString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetChildAreas(ctx context.Context, in *Int32, opts ...grpc.CallOption) (*AreaListResponse, error) {
	out := new(AreaListResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/GetChildAreas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetMoAppConf(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SMobileAppConfig, error) {
	out := new(SMobileAppConfig)
	err := c.cc.Invoke(ctx, "/FoundationService/GetMoAppConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveMoAppConf(ctx context.Context, in *SMobileAppConfig, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveMoAppConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetWxApiConfig(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SWxApiConfig, error) {
	out := new(SWxApiConfig)
	err := c.cc.Invoke(ctx, "/FoundationService/GetWxApiConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveWxApiConfig(ctx context.Context, in *SWxApiConfig, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveWxApiConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetPayPlatform(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PaymentPlatformResponse, error) {
	out := new(PaymentPlatformResponse)
	err := c.cc.Invoke(ctx, "/FoundationService/GetPayPlatform", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) GetGlobMchSaleConf_(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SGlobMchSaleConf, error) {
	out := new(SGlobMchSaleConf)
	err := c.cc.Invoke(ctx, "/FoundationService/GetGlobMchSaleConf_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foundationServiceClient) SaveGlobMchSaleConf_(ctx context.Context, in *SGlobMchSaleConf, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/FoundationService/SaveGlobMchSaleConf_", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoundationServiceServer is the server API for FoundationService service.
type FoundationServiceServer interface {
	// * 检测是否包含敏感词
	CheckSensitive(context.Context, *String) (*Bool, error)
	// * 替换敏感词
	ReplaceSensitive(context.Context, *ReplaceSensitiveRequest) (*String, error)
	// * 获取短信API凭据, provider 短信服务商, 默认:http
	GetSmsApi(context.Context, *String) (*SSmsApi, error)
	// * 保存短信API凭据,@provider 短信服务商, 默认:http
	SaveSmsApi(context.Context, *SmsApiSaveRequest) (*Result, error)
	// * 保存面板HOOK数据,这通常是在第三方应用中初始化或调用,参见文档：BoardHooks
	SaveBoardHook(context.Context, *BoardHookSaveRequest) (*Result, error)
	// 格式化资源地址并返回
	ResourceUrl(context.Context, *String) (*String, error)
	// 设置键值
	// rpc SetValue (Pair) returns (Result){}
	// 删除值,key
	// rpc DeleteValue (String) returns (Result){}
	// 根据前缀获取值,prefix
	// rpc GetValuesByPrefix (String) returns (StringMap){}
	// 注册单点登录应用,返回值：
	//   -  1. 成功，并返回token
	//   - -1. 接口地址不正确
	//   - -2. 已经注册
	RegisterApp(context.Context, *SSsoApp) (*String, error)
	// 获取应用信息,name
	GetApp(context.Context, *String) (*SSsoApp, error)
	// 获取单点登录应用
	GetAllSsoApp(context.Context, *Empty) (*StringListResponse, error)
	// 验证超级用户账号和密码
	SuperValidate(context.Context, *UserPwd) (*SuperLoginResponse, error)
	// 保存超级用户账号和密码
	FlushSuperPwd(context.Context, *UserPwd) (*Result, error)
	// 创建同步登录的地址,returnUrl
	GetSyncLoginUrl(context.Context, *String) (*String, error)
	// 获取地区名称
	GetAreaNames(context.Context, *GetAreaNamesRequest) (*StringListResponse, error)
	// 获取省市区字符串
	GetAreaString(context.Context, *AreaStringRequest) (*String, error)
	// 获取下级区域,code
	GetChildAreas(context.Context, *Int32) (*AreaListResponse, error)
	// 获取移动应用设置
	GetMoAppConf(context.Context, *Empty) (*SMobileAppConfig, error)
	// 保存移动应用设置
	SaveMoAppConf(context.Context, *SMobileAppConfig) (*Result, error)
	// 获取微信接口配置
	GetWxApiConfig(context.Context, *Empty) (*SWxApiConfig, error)
	// 保存微信接口配置
	SaveWxApiConfig(context.Context, *SWxApiConfig) (*Result, error)
	// 获取支付平台
	GetPayPlatform(context.Context, *Empty) (*PaymentPlatformResponse, error)
	// 获取全局商户销售设置
	GetGlobMchSaleConf_(context.Context, *Empty) (*SGlobMchSaleConf, error)
	// 保存全局商户销售设置
	SaveGlobMchSaleConf_(context.Context, *SGlobMchSaleConf) (*Result, error)
}

func RegisterFoundationServiceServer(s *grpc.Server, srv FoundationServiceServer) {
	s.RegisterService(&_FoundationService_serviceDesc, srv)
}

func _FoundationService_CheckSensitive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).CheckSensitive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/CheckSensitive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).CheckSensitive(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_ReplaceSensitive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplaceSensitiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).ReplaceSensitive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/ReplaceSensitive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).ReplaceSensitive(ctx, req.(*ReplaceSensitiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetSmsApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetSmsApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetSmsApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetSmsApi(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveSmsApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsApiSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveSmsApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveSmsApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveSmsApi(ctx, req.(*SmsApiSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveBoardHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoardHookSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveBoardHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveBoardHook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveBoardHook(ctx, req.(*BoardHookSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_ResourceUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).ResourceUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/ResourceUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).ResourceUrl(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_RegisterApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SSsoApp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).RegisterApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/RegisterApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).RegisterApp(ctx, req.(*SSsoApp))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetApp(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetAllSsoApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetAllSsoApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetAllSsoApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetAllSsoApp(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SuperValidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPwd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SuperValidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SuperValidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SuperValidate(ctx, req.(*UserPwd))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_FlushSuperPwd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPwd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).FlushSuperPwd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/FlushSuperPwd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).FlushSuperPwd(ctx, req.(*UserPwd))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetSyncLoginUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetSyncLoginUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetSyncLoginUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetSyncLoginUrl(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetAreaNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAreaNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetAreaNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetAreaNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetAreaNames(ctx, req.(*GetAreaNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetAreaString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AreaStringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetAreaString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetAreaString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetAreaString(ctx, req.(*AreaStringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetChildAreas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int32)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetChildAreas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetChildAreas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetChildAreas(ctx, req.(*Int32))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetMoAppConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetMoAppConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetMoAppConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetMoAppConf(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveMoAppConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SMobileAppConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveMoAppConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveMoAppConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveMoAppConf(ctx, req.(*SMobileAppConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetWxApiConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetWxApiConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetWxApiConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetWxApiConfig(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveWxApiConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SWxApiConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveWxApiConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveWxApiConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveWxApiConfig(ctx, req.(*SWxApiConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetPayPlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetPayPlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetPayPlatform",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetPayPlatform(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_GetGlobMchSaleConf__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).GetGlobMchSaleConf_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/GetGlobMchSaleConf_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).GetGlobMchSaleConf_(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FoundationService_SaveGlobMchSaleConf__Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SGlobMchSaleConf)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoundationServiceServer).SaveGlobMchSaleConf_(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FoundationService/SaveGlobMchSaleConf_",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoundationServiceServer).SaveGlobMchSaleConf_(ctx, req.(*SGlobMchSaleConf))
	}
	return interceptor(ctx, in, info, handler)
}

var _FoundationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FoundationService",
	HandlerType: (*FoundationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckSensitive",
			Handler:    _FoundationService_CheckSensitive_Handler,
		},
		{
			MethodName: "ReplaceSensitive",
			Handler:    _FoundationService_ReplaceSensitive_Handler,
		},
		{
			MethodName: "GetSmsApi",
			Handler:    _FoundationService_GetSmsApi_Handler,
		},
		{
			MethodName: "SaveSmsApi",
			Handler:    _FoundationService_SaveSmsApi_Handler,
		},
		{
			MethodName: "SaveBoardHook",
			Handler:    _FoundationService_SaveBoardHook_Handler,
		},
		{
			MethodName: "ResourceUrl",
			Handler:    _FoundationService_ResourceUrl_Handler,
		},
		{
			MethodName: "RegisterApp",
			Handler:    _FoundationService_RegisterApp_Handler,
		},
		{
			MethodName: "GetApp",
			Handler:    _FoundationService_GetApp_Handler,
		},
		{
			MethodName: "GetAllSsoApp",
			Handler:    _FoundationService_GetAllSsoApp_Handler,
		},
		{
			MethodName: "SuperValidate",
			Handler:    _FoundationService_SuperValidate_Handler,
		},
		{
			MethodName: "FlushSuperPwd",
			Handler:    _FoundationService_FlushSuperPwd_Handler,
		},
		{
			MethodName: "GetSyncLoginUrl",
			Handler:    _FoundationService_GetSyncLoginUrl_Handler,
		},
		{
			MethodName: "GetAreaNames",
			Handler:    _FoundationService_GetAreaNames_Handler,
		},
		{
			MethodName: "GetAreaString",
			Handler:    _FoundationService_GetAreaString_Handler,
		},
		{
			MethodName: "GetChildAreas",
			Handler:    _FoundationService_GetChildAreas_Handler,
		},
		{
			MethodName: "GetMoAppConf",
			Handler:    _FoundationService_GetMoAppConf_Handler,
		},
		{
			MethodName: "SaveMoAppConf",
			Handler:    _FoundationService_SaveMoAppConf_Handler,
		},
		{
			MethodName: "GetWxApiConfig",
			Handler:    _FoundationService_GetWxApiConfig_Handler,
		},
		{
			MethodName: "SaveWxApiConfig",
			Handler:    _FoundationService_SaveWxApiConfig_Handler,
		},
		{
			MethodName: "GetPayPlatform",
			Handler:    _FoundationService_GetPayPlatform_Handler,
		},
		{
			MethodName: "GetGlobMchSaleConf_",
			Handler:    _FoundationService_GetGlobMchSaleConf__Handler,
		},
		{
			MethodName: "SaveGlobMchSaleConf_",
			Handler:    _FoundationService_SaveGlobMchSaleConf__Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foundation_service.proto",
}

func init() {
	proto.RegisterFile("foundation_service.proto", fileDescriptor_foundation_service_eaa6caa00d874cdc)
}

var fileDescriptor_foundation_service_eaa6caa00d874cdc = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xce, 0x10, 0x74, 0x9b, 0x59, 0x37, 0xe6, 0x0d, 0x51, 0x55, 0x80, 0x46, 0x04, 0x12, 0x68,
	0x60, 0xa0, 0x43, 0xe2, 0xb0, 0x53, 0x37, 0xb1, 0x82, 0xb4, 0x41, 0xd5, 0x68, 0x4c, 0xe2, 0x32,
	0xb9, 0xc9, 0x6b, 0x6a, 0xcd, 0x89, 0x8d, 0xed, 0x74, 0xf4, 0x0f, 0xe7, 0x8e, 0xec, 0xa4, 0xa9,
	0xd7, 0xd2, 0x53, 0xed, 0xf7, 0xfd, 0x78, 0xdf, 0x7b, 0x72, 0x83, 0x5a, 0x23, 0x51, 0xe4, 0x09,
	0x35, 0x4c, 0xe4, 0xd7, 0x1a, 0xd4, 0x84, 0xc5, 0x40, 0xa4, 0x12, 0x46, 0xb4, 0xb7, 0x52, 0x2e,
	0x86, 0x94, 0x57, 0xb7, 0xa7, 0x19, 0x68, 0x4d, 0x53, 0x78, 0xef, 0xf1, 0x13, 0x23, 0x4a, 0x34,
	0xfc, 0x81, 0x9e, 0x0c, 0x40, 0x72, 0x1a, 0x43, 0x04, 0xb9, 0x66, 0x86, 0x4d, 0x60, 0x00, 0xbf,
	0x0b, 0xd0, 0x06, 0x63, 0x74, 0xff, 0x4a, 0xa8, 0xa4, 0xb5, 0x76, 0xb0, 0xf6, 0x7a, 0x73, 0xe0,
	0xce, 0xf8, 0x00, 0x3d, 0xac, 0xe8, 0x19, 0xe4, 0xa6, 0x75, 0xcf, 0x41, 0x7e, 0xa9, 0xf3, 0x77,
	0x1d, 0xed, 0x9e, 0xd5, 0x9d, 0xa2, 0x32, 0x18, 0x0e, 0xd1, 0xf6, 0xe9, 0x18, 0xe2, 0x9b, 0xba,
	0x09, 0x5e, 0x27, 0x91, 0x51, 0x2c, 0x4f, 0xdb, 0x0f, 0xc8, 0x89, 0x10, 0x3c, 0x0c, 0xf0, 0x67,
	0xf4, 0x68, 0x31, 0x0a, 0x6e, 0x91, 0x15, 0xe9, 0xda, 0x33, 0x7d, 0x18, 0xe0, 0x03, 0xb4, 0xd9,
	0x03, 0x13, 0x65, 0xba, 0x2b, 0xd9, 0xdc, 0x77, 0x83, 0x44, 0x65, 0x29, 0x0c, 0xf0, 0x21, 0x42,
	0x11, 0x9d, 0x40, 0x45, 0xc1, 0xa4, 0x3c, 0xd8, 0xd2, 0xdc, 0x6e, 0x00, 0xba, 0xe0, 0x26, 0x0c,
	0xf0, 0x47, 0xd4, 0xb4, 0xc8, 0x89, 0xa0, 0x2a, 0xf9, 0x2a, 0xc4, 0x0d, 0x7e, 0x4c, 0xea, 0xf3,
	0x0a, 0xc9, 0x0b, 0xbb, 0x16, 0x2d, 0x0a, 0x15, 0xc3, 0xa5, 0xe2, 0xf3, 0x0c, 0x5e, 0xc8, 0xd0,
	0x52, 0x52, 0xa6, 0x0d, 0xa8, 0xae, 0x94, 0xd8, 0xa6, 0xd3, 0xa2, 0x2b, 0xa5, 0xcf, 0x79, 0x86,
	0x1a, 0x3d, 0x30, 0x16, 0xbe, 0x33, 0x85, 0xe3, 0x85, 0x01, 0x7e, 0x87, 0xb6, 0x2c, 0xcc, 0x79,
	0x59, 0xc1, 0x0d, 0xf2, 0x25, 0x93, 0x66, 0xda, 0xde, 0xab, 0xc8, 0xe7, 0x4c, 0x9b, 0x01, 0x68,
	0x29, 0x72, 0x0d, 0x61, 0x80, 0x3f, 0xa0, 0x66, 0x54, 0x48, 0x50, 0x3f, 0x29, 0x67, 0x09, 0x35,
	0x80, 0x37, 0xc8, 0xa5, 0x06, 0xd5, 0xbf, 0x4d, 0xac, 0xc2, 0x22, 0xe7, 0x22, 0x65, 0xb9, 0xa7,
	0x78, 0x89, 0x9a, 0x67, 0xbc, 0xd0, 0x63, 0x07, 0xf6, 0x6f, 0x13, 0x4f, 0xe1, 0x0d, 0xfb, 0x0a,
	0xed, 0xd8, 0x75, 0x4f, 0xf3, 0xd8, 0xe9, 0x57, 0x0d, 0x7c, 0x5c, 0xa6, 0x55, 0x40, 0xbf, 0xd3,
	0x0c, 0x34, 0xde, 0x27, 0xfe, 0x75, 0xb6, 0xc4, 0x15, 0xd9, 0x09, 0x6a, 0x56, 0xec, 0x12, 0xc6,
	0x98, 0xcc, 0x2f, 0xff, 0x79, 0x02, 0x6f, 0x1d, 0xff, 0x74, 0xcc, 0x78, 0x62, 0x79, 0x1a, 0x37,
	0xc8, 0xb7, 0xdc, 0x1c, 0x75, 0xda, 0xbb, 0x4e, 0xb7, 0xe0, 0x7e, 0xe8, 0xa2, 0x5d, 0xd8, 0x25,
	0x9e, 0x8a, 0x7c, 0x54, 0x2f, 0x72, 0x97, 0x44, 0x17, 0x62, 0xc8, 0x38, 0x54, 0x00, 0x4b, 0xdd,
	0xd6, 0xdd, 0x73, 0x98, 0xb3, 0x97, 0x59, 0xfe, 0x76, 0xde, 0xa0, 0xed, 0x1e, 0x98, 0xab, 0x3f,
	0x5d, 0xc9, 0x4a, 0xb0, 0x76, 0x6f, 0x92, 0xc8, 0x2b, 0xbb, 0x18, 0x3b, 0xd6, 0xd9, 0xe7, 0xde,
	0xe5, 0xf8, 0xbe, 0x9f, 0x9c, 0x6f, 0x9f, 0x4e, 0xfb, 0x9c, 0x9a, 0x91, 0x50, 0x59, 0xed, 0xdb,
	0x22, 0x7d, 0x3a, 0xb5, 0xff, 0xbd, 0x19, 0xe2, 0x4d, 0xda, 0x41, 0x7b, 0x3d, 0x30, 0x3d, 0x2e,
	0x86, 0x17, 0xf1, 0x38, 0xa2, 0x1c, 0xac, 0xf1, 0xb5, 0x3f, 0xf0, 0x02, 0xe6, 0x3a, 0xed, 0xdb,
	0x58, 0x4b, 0xa2, 0x65, 0xb2, 0x97, 0xef, 0xe4, 0x39, 0xda, 0x8b, 0x45, 0x46, 0x52, 0x66, 0xc6,
	0xc5, 0x90, 0xa4, 0xa2, 0x23, 0x88, 0x92, 0xf1, 0xaf, 0x75, 0x72, 0xec, 0x3e, 0x34, 0xc3, 0x86,
	0xfb, 0x39, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0x58, 0x5c, 0xf7, 0x13, 0xb7, 0x04, 0x00, 0x00,
}
