// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.0
// source: message/member_query_dto.proto

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

// * 会员
type SMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * 编号
	Id int64 `protobuf:"zigzag64,1,opt,name=id,proto3" json:"id"`
	// * 会员编码/邀请码
	UserCode string `protobuf:"bytes,2,opt,name=userCode,proto3" json:"userCode"`
	// * 会员标志
	UserFlag int32 `protobuf:"zigzag32,3,opt,name=userFlag,proto3" json:"userFlag"`
	// 角色标志,1: 用户, 2:商户雇员, 4:扩展1, 8:扩展2, 16:扩展3
	Role int32 `protobuf:"zigzag32,4,opt,name=role,proto3" json:"role"`
	// * 用户名
	Username string `protobuf:"bytes,5,opt,name=username,proto3" json:"username"`
	// * 经验值
	Exp int64 `protobuf:"zigzag64,6,opt,name=exp,proto3" json:"exp"`
	// * 等级
	Level int32 `protobuf:"zigzag32,7,opt,name=level,proto3" json:"level"`
	// * 会员头像
	Portrait string `protobuf:"bytes,8,opt,name=portrait,proto3" json:"portrait"`
	// * 昵称
	Nickname string `protobuf:"bytes,9,opt,name=nickname,proto3" json:"nickname"`
	// * 注册IP
	RegIp string `protobuf:"bytes,10,opt,name=regIp,proto3" json:"regIp"`
	// * 注册来源
	RegFrom string `protobuf:"bytes,11,opt,name=regFrom,proto3" json:"regFrom"`
	// * 手机号
	Phone string `protobuf:"bytes,12,opt,name=phone,proto3" json:"phone"`
	// * 电子邮箱
	Email string `protobuf:"bytes,13,opt,name=email,proto3" json:"email"`
	// * 真实姓名
	RealName string `protobuf:"bytes,14,opt,name=realName,proto3" json:"realName"`
	// * 高级用户级别
	PremiumUser int32 `protobuf:"zigzag32,15,opt,name=premiumUser,proto3" json:"premiumUser"`
	// * 高级用户过期时间
	PremiumExpires int64 `protobuf:"zigzag64,16,opt,name=premiumExpires,proto3" json:"premiumExpires"`
	// * 注册时间
	RegTime int64 `protobuf:"varint,17,opt,name=regTime,proto3" json:"regTime"`
	// * 最后登录时间
	LastLoginTime int64 `protobuf:"varint,18,opt,name=lastLoginTime,proto3" json:"lastLoginTime"`
}

func (x *SMember) Reset() {
	*x = SMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_member_query_dto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SMember) ProtoMessage() {}

func (x *SMember) ProtoReflect() protoreflect.Message {
	mi := &file_message_member_query_dto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SMember.ProtoReflect.Descriptor instead.
func (*SMember) Descriptor() ([]byte, []int) {
	return file_message_member_query_dto_proto_rawDescGZIP(), []int{0}
}

func (x *SMember) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SMember) GetUserCode() string {
	if x != nil {
		return x.UserCode
	}
	return ""
}

func (x *SMember) GetUserFlag() int32 {
	if x != nil {
		return x.UserFlag
	}
	return 0
}

func (x *SMember) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *SMember) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SMember) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *SMember) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *SMember) GetPortrait() string {
	if x != nil {
		return x.Portrait
	}
	return ""
}

func (x *SMember) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *SMember) GetRegIp() string {
	if x != nil {
		return x.RegIp
	}
	return ""
}

func (x *SMember) GetRegFrom() string {
	if x != nil {
		return x.RegFrom
	}
	return ""
}

func (x *SMember) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SMember) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SMember) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *SMember) GetPremiumUser() int32 {
	if x != nil {
		return x.PremiumUser
	}
	return 0
}

func (x *SMember) GetPremiumExpires() int64 {
	if x != nil {
		return x.PremiumExpires
	}
	return 0
}

func (x *SMember) GetRegTime() int64 {
	if x != nil {
		return x.RegTime
	}
	return 0
}

func (x *SMember) GetLastLoginTime() int64 {
	if x != nil {
		return x.LastLoginTime
	}
	return 0
}

// * 资料
type SProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId   int64  `protobuf:"zigzag64,1,opt,name=memberId,proto3" json:"memberId"`
	Nickname   string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname"`
	Portrait   string `protobuf:"bytes,3,opt,name=portrait,proto3" json:"portrait"`
	Gender     int32  `protobuf:"zigzag32,4,opt,name=gender,proto3" json:"gender"`
	BirthDay   string `protobuf:"bytes,5,opt,name=birthDay,proto3" json:"birthDay"`
	Phone      string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone"`
	Address    string `protobuf:"bytes,7,opt,name=address,proto3" json:"address"`
	Im         string `protobuf:"bytes,8,opt,name=im,proto3" json:"im"`
	Email      string `protobuf:"bytes,9,opt,name=email,proto3" json:"email"`
	Province   int32  `protobuf:"zigzag32,10,opt,name=province,proto3" json:"province"`
	City       int32  `protobuf:"zigzag32,11,opt,name=city,proto3" json:"city"`
	District   int32  `protobuf:"zigzag32,12,opt,name=district,proto3" json:"district"`
	Remark     string `protobuf:"bytes,13,opt,name=remark,proto3" json:"remark"`
	Ext1       string `protobuf:"bytes,14,opt,name=ext1,proto3" json:"ext1"`
	Ext2       string `protobuf:"bytes,15,opt,name=ext2,proto3" json:"ext2"`
	Ext3       string `protobuf:"bytes,16,opt,name=ext3,proto3" json:"ext3"`
	Ext4       string `protobuf:"bytes,17,opt,name=ext4,proto3" json:"ext4"`
	Ext5       string `protobuf:"bytes,18,opt,name=ext5,proto3" json:"ext5"`
	Ext6       string `protobuf:"bytes,19,opt,name=ext6,proto3" json:"ext6"`
	UpdateTime int64  `protobuf:"zigzag64,20,opt,name=updateTime,proto3" json:"updateTime"`
}

func (x *SProfile) Reset() {
	*x = SProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_member_query_dto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SProfile) ProtoMessage() {}

func (x *SProfile) ProtoReflect() protoreflect.Message {
	mi := &file_message_member_query_dto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SProfile.ProtoReflect.Descriptor instead.
func (*SProfile) Descriptor() ([]byte, []int) {
	return file_message_member_query_dto_proto_rawDescGZIP(), []int{1}
}

func (x *SProfile) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *SProfile) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *SProfile) GetPortrait() string {
	if x != nil {
		return x.Portrait
	}
	return ""
}

func (x *SProfile) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *SProfile) GetBirthDay() string {
	if x != nil {
		return x.BirthDay
	}
	return ""
}

func (x *SProfile) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SProfile) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SProfile) GetIm() string {
	if x != nil {
		return x.Im
	}
	return ""
}

func (x *SProfile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SProfile) GetProvince() int32 {
	if x != nil {
		return x.Province
	}
	return 0
}

func (x *SProfile) GetCity() int32 {
	if x != nil {
		return x.City
	}
	return 0
}

func (x *SProfile) GetDistrict() int32 {
	if x != nil {
		return x.District
	}
	return 0
}

func (x *SProfile) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *SProfile) GetExt1() string {
	if x != nil {
		return x.Ext1
	}
	return ""
}

func (x *SProfile) GetExt2() string {
	if x != nil {
		return x.Ext2
	}
	return ""
}

func (x *SProfile) GetExt3() string {
	if x != nil {
		return x.Ext3
	}
	return ""
}

func (x *SProfile) GetExt4() string {
	if x != nil {
		return x.Ext4
	}
	return ""
}

func (x *SProfile) GetExt5() string {
	if x != nil {
		return x.Ext5
	}
	return ""
}

func (x *SProfile) GetExt6() string {
	if x != nil {
		return x.Ext6
	}
	return ""
}

func (x *SProfile) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type SComplexMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname       string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname"`
	Portrait       string `protobuf:"bytes,2,opt,name=portrait,proto3" json:"portrait"`
	Phone          string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone"`
	Exp            int32  `protobuf:"zigzag32,4,opt,name=exp,proto3" json:"exp"`
	Level          int32  `protobuf:"zigzag32,5,opt,name=level,proto3" json:"level"`
	LevelName      string `protobuf:"bytes,6,opt,name=levelName,proto3" json:"levelName"`
	UserCode       string `protobuf:"bytes,7,opt,name=userCode,proto3" json:"userCode"`
	TrustAuthState int32  `protobuf:"zigzag32,8,opt,name=trustAuthState,proto3" json:"trustAuthState"`
	PremiumUser    int32  `protobuf:"zigzag32,9,opt,name=premiumUser,proto3" json:"premiumUser"`
	Flag           int32  `protobuf:"zigzag32,10,opt,name=flag,proto3" json:"flag"`
	UpdateTime     int64  `protobuf:"zigzag64,11,opt,name=updateTime,proto3" json:"updateTime"`
	// * 交易密码是否已设置
	TradePasswordHasSet bool `protobuf:"varint,12,opt,name=tradePasswordHasSet,proto3" json:"tradePasswordHasSet"`
}

func (x *SComplexMember) Reset() {
	*x = SComplexMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_member_query_dto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SComplexMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SComplexMember) ProtoMessage() {}

func (x *SComplexMember) ProtoReflect() protoreflect.Message {
	mi := &file_message_member_query_dto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SComplexMember.ProtoReflect.Descriptor instead.
func (*SComplexMember) Descriptor() ([]byte, []int) {
	return file_message_member_query_dto_proto_rawDescGZIP(), []int{2}
}

func (x *SComplexMember) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *SComplexMember) GetPortrait() string {
	if x != nil {
		return x.Portrait
	}
	return ""
}

func (x *SComplexMember) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SComplexMember) GetExp() int32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *SComplexMember) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *SComplexMember) GetLevelName() string {
	if x != nil {
		return x.LevelName
	}
	return ""
}

func (x *SComplexMember) GetUserCode() string {
	if x != nil {
		return x.UserCode
	}
	return ""
}

func (x *SComplexMember) GetTrustAuthState() int32 {
	if x != nil {
		return x.TrustAuthState
	}
	return 0
}

func (x *SComplexMember) GetPremiumUser() int32 {
	if x != nil {
		return x.PremiumUser
	}
	return 0
}

func (x *SComplexMember) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *SComplexMember) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *SComplexMember) GetTradePasswordHasSet() bool {
	if x != nil {
		return x.TradePasswordHasSet
	}
	return false
}

var File_message_member_query_dto_proto protoreflect.FileDescriptor

var file_message_member_query_dto_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe3, 0x03, 0x0a, 0x07, 0x53, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x46, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x46, 0x6c, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x12, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x11, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6f, 0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6f, 0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x49, 0x70, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x67, 0x49, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x67, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x0e, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x67, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xe4, 0x03, 0x0a, 0x08, 0x53, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6f, 0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6f, 0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x11, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x11, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x74, 0x31, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x74, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x74,
	0x32, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x74, 0x32, 0x12, 0x12, 0x0a,
	0x04, 0x65, 0x78, 0x74, 0x33, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x74,
	0x33, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x74, 0x34, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x65, 0x78, 0x74, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x74, 0x35, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x74, 0x35, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x74,
	0x36, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x74, 0x36, 0x12, 0x1e, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x12, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xf0, 0x02,
	0x0a, 0x0e, 0x53, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6f, 0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6f, 0x72, 0x74, 0x72, 0x61, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x11, 0x52, 0x03, 0x65, 0x78, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x6d,
	0x69, 0x75, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0b, 0x70,
	0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c,
	0x61, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x11, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x1e,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x12, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30,
	0x0a, 0x13, 0x74, 0x72, 0x61, 0x64, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48,
	0x61, 0x73, 0x53, 0x65, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x74, 0x72, 0x61,
	0x64, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x53, 0x65, 0x74,
	0x42, 0x1f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x67,
	0x6f, 0x32, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_member_query_dto_proto_rawDescOnce sync.Once
	file_message_member_query_dto_proto_rawDescData = file_message_member_query_dto_proto_rawDesc
)

func file_message_member_query_dto_proto_rawDescGZIP() []byte {
	file_message_member_query_dto_proto_rawDescOnce.Do(func() {
		file_message_member_query_dto_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_member_query_dto_proto_rawDescData)
	})
	return file_message_member_query_dto_proto_rawDescData
}

var file_message_member_query_dto_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_message_member_query_dto_proto_goTypes = []interface{}{
	(*SMember)(nil),        // 0: SMember
	(*SProfile)(nil),       // 1: SProfile
	(*SComplexMember)(nil), // 2: SComplexMember
}
var file_message_member_query_dto_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_member_query_dto_proto_init() }
func file_message_member_query_dto_proto_init() {
	if File_message_member_query_dto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_member_query_dto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SMember); i {
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
		file_message_member_query_dto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SProfile); i {
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
		file_message_member_query_dto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SComplexMember); i {
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
			RawDescriptor: file_message_member_query_dto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_member_query_dto_proto_goTypes,
		DependencyIndexes: file_message_member_query_dto_proto_depIdxs,
		MessageInfos:      file_message_member_query_dto_proto_msgTypes,
	}.Build()
	File_message_member_query_dto_proto = out.File
	file_message_member_query_dto_proto_rawDesc = nil
	file_message_member_query_dto_proto_goTypes = nil
	file_message_member_query_dto_proto_depIdxs = nil
}
