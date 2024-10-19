// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
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

type RelationStatus int32

const (
	RelationStatus_RELATION_STATUS_UNUSED       RelationStatus = 0
	RelationStatus_RELATION_STATUS_NOT_IN_GROUP RelationStatus = 1
	RelationStatus_RELATION_STATUS_OWNER        RelationStatus = 2
	RelationStatus_RELATION_STATUS_MEMBER       RelationStatus = 3
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

type RelationChangeType int32

const (
	RelationChangeType_RELATION_CHANGE_TYPE_UNUSED       RelationChangeType = 0
	RelationChangeType_RELATION_CHNAGE_TYPE_NONE         RelationChangeType = 1
	RelationChangeType_RELATION_CHANGE_TYPE_OWNER_ACCEPT RelationChangeType = 2
	RelationChangeType_RELATION_CHANGE_TYPE_MEMBER_QUIT  RelationChangeType = 3
)

// Enum value maps for RelationChangeType.
var (
	RelationChangeType_name = map[int32]string{
		0: "RELATION_CHANGE_TYPE_UNUSED",
		1: "RELATION_CHNAGE_TYPE_NONE",
		2: "RELATION_CHANGE_TYPE_OWNER_ACCEPT",
		3: "RELATION_CHANGE_TYPE_MEMBER_QUIT",
	}
	RelationChangeType_value = map[string]int32{
		"RELATION_CHANGE_TYPE_UNUSED":       0,
		"RELATION_CHNAGE_TYPE_NONE":         1,
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
	ApplyStatus_APPLY_STATUS_NONE     ApplyStatus = 1
	ApplyStatus_APPLY_STATUS_PENDING  ApplyStatus = 2
	ApplyStatus_APPLY_STATUS_ACCEPTED ApplyStatus = 3
	ApplyStatus_APPLY_STATUS_REJECTED ApplyStatus = 4
	// 如果群已经呗解散, 用户再次操作apply, 状态会变成group disbanded
	ApplyStatus_APPLY_STATUS_GROUP_DISBANDED ApplyStatus = 5
)

// Enum value maps for ApplyStatus.
var (
	ApplyStatus_name = map[int32]string{
		0: "APPLY_STATUS_UNUSED",
		1: "APPLY_STATUS_NONE",
		2: "APPLY_STATUS_PENDING",
		3: "APPLY_STATUS_ACCEPTED",
		4: "APPLY_STATUS_REJECTED",
		5: "APPLY_STATUS_GROUP_DISBANDED",
	}
	ApplyStatus_value = map[string]int32{
		"APPLY_STATUS_UNUSED":          0,
		"APPLY_STATUS_NONE":            1,
		"APPLY_STATUS_PENDING":         2,
		"APPLY_STATUS_ACCEPTED":        3,
		"APPLY_STATUS_REJECTED":        4,
		"APPLY_STATUS_GROUP_DISBANDED": 5,
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

type SessionEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginAt             int64  `protobuf:"varint,1,opt,name=login_at,json=loginAt,proto3" json:"login_at,omitempty"`
	Username            string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	SessionId           string `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	DeviceDesc          string `protobuf:"bytes,4,opt,name=device_desc,json=deviceDesc,proto3" json:"device_desc,omitempty"`
	GwAdvertiseAddrPort string `protobuf:"bytes,5,opt,name=gwAdvertiseAddrPort,proto3" json:"gwAdvertiseAddrPort,omitempty"`
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

func (x *SessionEntry) GetGwAdvertiseAddrPort() string {
	if x != nil {
		return x.GwAdvertiseAddrPort
	}
	return ""
}

var File_base_base_proto protoreflect.FileDescriptor

var file_base_base_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0xb7, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x12, 0x30, 0x0a, 0x13, 0x67, 0x77, 0x41, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x41, 0x64, 0x64, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x67, 0x77, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69,
	0x73, 0x65, 0x41, 0x64, 0x64, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x2a, 0x85, 0x01, 0x0a, 0x0e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a,
	0x16, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x55, 0x53, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x52, 0x45, 0x4c,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x49, 0x4e, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x52,
	0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f,
	0x57, 0x4e, 0x45, 0x52, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52,
	0x10, 0x03, 0x2a, 0xa1, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x4c,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x55, 0x53, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45,
	0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x4e, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x52, 0x45, 0x4c,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x10, 0x02,
	0x12, 0x24, 0x0a, 0x20, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f,
	0x51, 0x55, 0x49, 0x54, 0x10, 0x03, 0x2a, 0xaf, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x55, 0x53, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x15, 0x0a, 0x11, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02,
	0x12, 0x19, 0x0a, 0x15, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x41,
	0x50, 0x50, 0x4c, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x4a, 0x45,
	0x43, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x44, 0x49, 0x53,
	0x42, 0x41, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x05, 0x42, 0x1f, 0x5a, 0x1d, 0x70, 0x69, 0x67, 0x65,
	0x6f, 0x6e, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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
var file_base_base_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_base_base_proto_goTypes = []interface{}{
	(RelationStatus)(0),     // 0: base.RelationStatus
	(RelationChangeType)(0), // 1: base.RelationChangeType
	(ApplyStatus)(0),        // 2: base.ApplyStatus
	(*Empty)(nil),           // 3: base.Empty
	(*SessionEntry)(nil),    // 4: base.SessionEntry
}
var file_base_base_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_base_base_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
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
