// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.2
// source: base/base.proto

package base

import (
	context "context"
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

// 关系状态, 一个人关于一个群的状态, 目前有非群成员/群主/群员三种关系
type RelationStatus int32

const (
	RelationStatus_RELATION_STATUS_UNUSED       RelationStatus = 0
	RelationStatus_RELATION_STATUS_NOT_IN_GROUP RelationStatus = 1 // 非成员
	RelationStatus_RELATION_STATUS_OWNER        RelationStatus = 2 // 群主
	RelationStatus_RELATION_STATUS_MEMBER       RelationStatus = 3 // 群员
)

// Enum value maps for RelationStatus.
var (
	RelationStatus_name = map[int32]string{
		0: "RELATION_STATUS_UNUSED",
		1: "RELATION_STATUS_NOT_IN_GROUP",
		2: "RELATION_STATUS_OWNER",
		3: "RELATION_STATUS_MEMBER",
	}
	RelationStatus_value = map[string]int32{
		"RELATION_STATUS_UNUSED":       0,
		"RELATION_STATUS_NOT_IN_GROUP": 1,
		"RELATION_STATUS_OWNER":        2,
		"RELATION_STATUS_MEMBER":       3,
	}
)

func (x RelationStatus) Enum() *RelationStatus {
	p := new(RelationStatus)
	*p = x
	return p
}

func (x RelationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RelationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_base_base_proto_enumTypes[0].Descriptor()
}

func (RelationStatus) Type() protoreflect.EnumType {
	return &file_base_base_proto_enumTypes[0]
}

func (x RelationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RelationStatus.Descriptor instead.
func (RelationStatus) EnumDescriptor() ([]byte, []int) {
	return file_base_base_proto_rawDescGZIP(), []int{0}
}

// 关系改变原因, 表示上次关系修改的原因
type RelationChangeType int32

const (
	RelationChangeType_RELATION_CHANGE_TYPE_UNUSED       RelationChangeType = 0
	RelationChangeType_RELATION_CHANGE_TYPE_CREATE_GROUP RelationChangeType = 1 // 因为创建群聊, 群主关于这个群的change type只能是这个, 且关系version为1
	RelationChangeType_RELATION_CHANGE_TYPE_OWNER_ACCEPT RelationChangeType = 2 // 因为群聊owner接受申请, 成为群员
	RelationChangeType_RELATION_CHANGE_TYPE_MEMBER_QUIT  RelationChangeType = 3 // 因为群员主动退出, 成为非群员
)

// Enum value maps for RelationChangeType.
var (
	RelationChangeType_name = map[int32]string{
		0: "RELATION_CHANGE_TYPE_UNUSED",
		1: "RELATION_CHANGE_TYPE_CREATE_GROUP",
		2: "RELATION_CHANGE_TYPE_OWNER_ACCEPT",
		3: "RELATION_CHANGE_TYPE_MEMBER_QUIT",
	}
	RelationChangeType_value = map[string]int32{
		"RELATION_CHANGE_TYPE_UNUSED":       0,
		"RELATION_CHANGE_TYPE_CREATE_GROUP": 1,
		"RELATION_CHANGE_TYPE_OWNER_ACCEPT": 2,
		"RELATION_CHANGE_TYPE_MEMBER_QUIT":  3,
	}
)

func (x RelationChangeType) Enum() *RelationChangeType {
	p := new(RelationChangeType)
	*p = x
	return p
}

func (x RelationChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RelationChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_base_base_proto_enumTypes[1].Descriptor()
}

func (RelationChangeType) Type() protoreflect.EnumType {
	return &file_base_base_proto_enumTypes[1]
}

func (x RelationChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RelationChangeType.Descriptor instead.
func (RelationChangeType) EnumDescriptor() ([]byte, []int) {
	return file_base_base_proto_rawDescGZIP(), []int{1}
}

type ApplyStatus int32

const (
	ApplyStatus_APPLY_STATUS_UNUSED   ApplyStatus = 0
	ApplyStatus_APPLY_STATUS_PENDING  ApplyStatus = 1 // 等待处理
	ApplyStatus_APPLY_STATUS_ACCEPTED ApplyStatus = 2 // 已被处理, 被accept
	ApplyStatus_APPLY_STATUS_REJECTED ApplyStatus = 3 // 已被处理, 被reject
	// 如果群已经被解散, 用户再次操作apply, 状态会变成group disbanded
	ApplyStatus_APPLY_STATUS_GROUP_DISBANDED ApplyStatus = 4
)

// Enum value maps for ApplyStatus.
var (
	ApplyStatus_name = map[int32]string{
		0: "APPLY_STATUS_UNUSED",
		1: "APPLY_STATUS_PENDING",
		2: "APPLY_STATUS_ACCEPTED",
		3: "APPLY_STATUS_REJECTED",
		4: "APPLY_STATUS_GROUP_DISBANDED",
	}
	ApplyStatus_value = map[string]int32{
		"APPLY_STATUS_UNUSED":          0,
		"APPLY_STATUS_PENDING":         1,
		"APPLY_STATUS_ACCEPTED":        2,
		"APPLY_STATUS_REJECTED":        3,
		"APPLY_STATUS_GROUP_DISBANDED": 4,
	}
)

func (x ApplyStatus) Enum() *ApplyStatus {
	p := new(ApplyStatus)
	*p = x
	return p
}

func (x ApplyStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApplyStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_base_base_proto_enumTypes[2].Descriptor()
}

func (ApplyStatus) Type() protoreflect.EnumType {
	return &file_base_base_proto_enumTypes[2]
}

func (x ApplyStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApplyStatus.Descriptor instead.
func (ApplyStatus) EnumDescriptor() ([]byte, []int) {
	return file_base_base_proto_rawDescGZIP(), []int{2}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_base_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_base_base_proto_rawDescGZIP(), []int{0}
}

// session信息, 一个session代表一个已经登录的长链接
type SessionEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginAt             int64  `protobuf:"varint,1,opt,name=login_at,json=loginAt,proto3" json:"login_at,omitempty"`
	Username            string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	SessionId           string `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	DeviceDesc          string `protobuf:"bytes,4,opt,name=device_desc,json=deviceDesc,proto3" json:"device_desc,omitempty"`
	GwAdvertiseAddrport string `protobuf:"bytes,5,opt,name=gw_advertise_addrport,json=gwAdvertiseAddrport,proto3" json:"gw_advertise_addrport,omitempty"`
}

func (x *SessionEntry) Reset() {
	*x = SessionEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionEntry) ProtoMessage() {}

func (x *SessionEntry) ProtoReflect() protoreflect.Message {
	mi := &file_base_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionEntry.ProtoReflect.Descriptor instead.
func (*SessionEntry) Descriptor() ([]byte, []int) {
	return file_base_base_proto_rawDescGZIP(), []int{1}
}

func (x *SessionEntry) GetLoginAt() int64 {
	if x != nil {
		return x.LoginAt
	}
	return 0
}

func (x *SessionEntry) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SessionEntry) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *SessionEntry) GetDeviceDesc() string {
	if x != nil {
		return x.DeviceDesc
	}
	return ""
}

func (x *SessionEntry) GetGwAdvertiseAddrport() string {
	if x != nil {
		return x.GwAdvertiseAddrport
	}
	return ""
}

// 某个人关于某个群的关系
type RelationEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId         string             `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	UserId          string             `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status          RelationStatus     `protobuf:"varint,3,opt,name=status,proto3,enum=base.RelationStatus" json:"status,omitempty"`
	ChangeType      RelationChangeType `protobuf:"varint,4,opt,name=change_type,json=changeType,proto3,enum=base.RelationChangeType" json:"change_type,omitempty"`
	RelationVersion int64              `protobuf:"varint,5,opt,name=relation_version,json=relationVersion,proto3" json:"relation_version,omitempty"`
	ChangeAt        int64              `protobuf:"varint,6,opt,name=change_at,json=changeAt,proto3" json:"change_at,omitempty"`
}

func (x *RelationEntry) Reset() {
	*x = RelationEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationEntry) ProtoMessage() {}

func (x *RelationEntry) ProtoReflect() protoreflect.Message {
	mi := &file_base_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationEntry.ProtoReflect.Descriptor instead.
func (*RelationEntry) Descriptor() ([]byte, []int) {
	return file_base_base_proto_rawDescGZIP(), []int{2}
}

func (x *RelationEntry) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *RelationEntry) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RelationEntry) GetStatus() RelationStatus {
	if x != nil {
		return x.Status
	}
	return RelationStatus_RELATION_STATUS_UNUSED
}

func (x *RelationEntry) GetChangeType() RelationChangeType {
	if x != nil {
		return x.ChangeType
	}
	return RelationChangeType_RELATION_CHANGE_TYPE_UNUSED
}

func (x *RelationEntry) GetRelationVersion() int64 {
	if x != nil {
		return x.RelationVersion
	}
	return 0
}

func (x *RelationEntry) GetChangeAt() int64 {
	if x != nil {
		return x.ChangeAt
	}
	return 0
}

var File_base_base_proto protoreflect.FileDescriptor

var file_base_base_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0xb9, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x12, 0x32, 0x0a, 0x15, 0x67, 0x77, 0x5f, 0x61,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x67, 0x77, 0x41, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x73, 0x65, 0x41, 0x64, 0x64, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x22, 0xf4, 0x01, 0x0a,
	0x0d, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x39, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x41, 0x74, 0x2a, 0x85, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x55, 0x53, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x02, 0x12,
	0x1a, 0x0a, 0x16, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x03, 0x2a, 0xa9, 0x01, 0x0a, 0x12,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
	0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x55, 0x53, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x52, 0x45,
	0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x10,
	0x02, 0x12, 0x24, 0x0a, 0x20, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x48,
	0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52,
	0x5f, 0x51, 0x55, 0x49, 0x54, 0x10, 0x03, 0x2a, 0x98, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x50, 0x50, 0x4c, 0x59,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x55, 0x53, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x18, 0x0a, 0x14, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x50,
	0x50, 0x4c, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50,
	0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x20, 0x0a, 0x1c, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x44, 0x49, 0x53, 0x42, 0x41, 0x4e, 0x44, 0x45, 0x44,
	0x10, 0x04, 0x42, 0x1f, 0x5a, 0x1d, 0x70, 0x69, 0x67, 0x65, 0x6f, 0x6e, 0x2f, 0x6b, 0x69, 0x74,
	0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_base_proto_rawDescOnce sync.Once
	file_base_base_proto_rawDescData = file_base_base_proto_rawDesc
)

func file_base_base_proto_rawDescGZIP() []byte {
	file_base_base_proto_rawDescOnce.Do(func() {
		file_base_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_base_proto_rawDescData)
	})
	return file_base_base_proto_rawDescData
}

var file_base_base_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_base_base_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_base_base_proto_goTypes = []interface{}{
	(RelationStatus)(0),     // 0: base.RelationStatus
	(RelationChangeType)(0), // 1: base.RelationChangeType
	(ApplyStatus)(0),        // 2: base.ApplyStatus
	(*Empty)(nil),           // 3: base.Empty
	(*SessionEntry)(nil),    // 4: base.SessionEntry
	(*RelationEntry)(nil),   // 5: base.RelationEntry
}
var file_base_base_proto_depIdxs = []int32{
	0, // 0: base.RelationEntry.status:type_name -> base.RelationStatus
	1, // 1: base.RelationEntry.change_type:type_name -> base.RelationChangeType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_base_base_proto_init() }
func file_base_base_proto_init() {
	if File_base_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_base_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionEntry); i {
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
		file_base_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationEntry); i {
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
			RawDescriptor: file_base_base_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_base_proto_goTypes,
		DependencyIndexes: file_base_base_proto_depIdxs,
		EnumInfos:         file_base_base_proto_enumTypes,
		MessageInfos:      file_base_base_proto_msgTypes,
	}.Build()
	File_base_base_proto = out.File
	file_base_base_proto_rawDesc = nil
	file_base_base_proto_goTypes = nil
	file_base_base_proto_depIdxs = nil
}

var _ context.Context
