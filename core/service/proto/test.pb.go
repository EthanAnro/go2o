// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package proto // import "../core/service/proto"

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

type EState int32

const (
	EState__       EState = 0
	EState_Normal  EState = 1
	EState_Stopped EState = 3
)

var EState_name = map[int32]string{
	0: "_",
	1: "Normal",
	3: "Stopped",
}
var EState_value = map[string]int32{
	"_":       0,
	"Normal":  1,
	"Stopped": 3,
}

func (x EState) String() string {
	return proto.EnumName(EState_name, int32(x))
}
func (EState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_test_e29543d9ddb6e4ea, []int{0}
}

type User struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	GroupId              int64             `protobuf:"zigzag64,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Extra                map[string]string `protobuf:"bytes,3,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_e29543d9ddb6e4ea, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *User) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

type UserResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State                EState   `protobuf:"varint,2,opt,name=state,proto3,enum=EState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_e29543d9ddb6e4ea, []int{1}
}
func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (dst *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(dst, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResponse) GetState() EState {
	if m != nil {
		return m.State
	}
	return EState__
}

func init() {
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterMapType((map[string]string)(nil), "User.ExtraEntry")
	proto.RegisterType((*UserResponse)(nil), "UserResponse")
	proto.RegisterEnum("EState", EState_name, EState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterServiceClient is the client API for GreeterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterServiceClient interface {
	Hello(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error)
}

type greeterServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreeterServiceClient(cc *grpc.ClientConn) GreeterServiceClient {
	return &greeterServiceClient{cc}
}

func (c *greeterServiceClient) Hello(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/GreeterService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServiceServer is the server API for GreeterService service.
type GreeterServiceServer interface {
	Hello(context.Context, *User) (*UserResponse, error)
}

func RegisterGreeterServiceServer(s *grpc.Server, srv GreeterServiceServer) {
	s.RegisterService(&_GreeterService_serviceDesc, srv)
}

func _GreeterService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GreeterService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).Hello(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreeterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GreeterService",
	HandlerType: (*GreeterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _GreeterService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_test_e29543d9ddb6e4ea) }

var fileDescriptor_test_e29543d9ddb6e4ea = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0x97, 0x75, 0xd9, 0xfe, 0x7b, 0xf7, 0x77, 0x94, 0x17, 0x0f, 0x73, 0x32, 0x2c, 0x3d,
	0x48, 0xf1, 0x90, 0x62, 0xbd, 0x0c, 0x3d, 0x29, 0x14, 0xf5, 0xe2, 0x21, 0xc5, 0x8b, 0x97, 0x51,
	0xd7, 0x17, 0x19, 0x76, 0x4d, 0x49, 0xb2, 0xe1, 0x3e, 0x89, 0x5f, 0x57, 0x9a, 0x08, 0x7a, 0xf0,
	0x92, 0xbc, 0xcf, 0x2f, 0x3c, 0x4f, 0x9e, 0x04, 0xc0, 0x92, 0xb1, 0xa2, 0xd5, 0xca, 0xaa, 0xf8,
	0x93, 0xc1, 0xe0, 0xd9, 0x90, 0x46, 0x84, 0x41, 0x53, 0x6e, 0x69, 0xc6, 0x22, 0x96, 0x8c, 0xa5,
	0x9b, 0xf1, 0x04, 0xfe, 0xbd, 0x69, 0xb5, 0x6b, 0x57, 0x9b, 0x6a, 0xd6, 0x8f, 0x58, 0x82, 0x72,
	0xe4, 0xf4, 0x63, 0x85, 0xe7, 0xc0, 0xe9, 0xc3, 0xea, 0x72, 0x16, 0x44, 0x41, 0x32, 0xc9, 0x42,
	0xd1, 0x85, 0x88, 0xbc, 0x43, 0x79, 0x63, 0xf5, 0x41, 0xfa, 0xe3, 0xf9, 0x12, 0xe0, 0x07, 0x62,
	0x08, 0xc1, 0x3b, 0x1d, 0xbe, 0xef, 0xe8, 0x46, 0x3c, 0x06, 0xbe, 0x2f, 0xeb, 0x1d, 0xb9, 0xfc,
	0xb1, 0xf4, 0xe2, 0xba, 0xbf, 0x64, 0xf1, 0x2d, 0xfc, 0xef, 0x32, 0x25, 0x99, 0x56, 0x35, 0x86,
	0xfe, 0x2c, 0xb8, 0x00, 0x6e, 0x6c, 0x69, 0xbd, 0x7b, 0x9a, 0x8d, 0x44, 0x5e, 0x74, 0x52, 0x7a,
	0x7a, 0x91, 0xc0, 0xd0, 0x03, 0xe4, 0xc0, 0x56, 0x61, 0x0f, 0x01, 0x86, 0x4f, 0x4a, 0x6f, 0xcb,
	0x3a, 0x64, 0x38, 0x81, 0x51, 0x61, 0x55, 0xdb, 0x52, 0x15, 0x06, 0xd9, 0x25, 0x4c, 0xef, 0x35,
	0x91, 0x25, 0x5d, 0x90, 0xde, 0x6f, 0xd6, 0x84, 0x67, 0xc0, 0x1f, 0xa8, 0xae, 0x15, 0x72, 0xf7,
	0xb4, 0xf9, 0x91, 0xf8, 0xdd, 0x26, 0xee, 0xdd, 0x2d, 0x5e, 0x4e, 0x85, 0x48, 0xd7, 0x4a, 0x53,
	0x6a, 0xbc, 0x27, 0x75, 0x5f, 0x7a, 0xe3, 0xd6, 0xd7, 0xa1, 0xdb, 0xae, 0xbe, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x4d, 0xeb, 0x16, 0xa3, 0x6d, 0x01, 0x00, 0x00,
}