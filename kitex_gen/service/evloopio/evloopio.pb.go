// 这个文件是eventloop的请求响应定义

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: base/evloopio.proto

package evloopio

import (
	context "context"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	base "pigeon/kitex_gen/service/base"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AlterGroupMemberResponse_AlterGroupMemberResponseCode int32

const (
	AlterGroupMemberResponse_OK AlterGroupMemberResponse_AlterGroupMemberResponseCode = 0 // GROUP_DISBANDED = 1;
)

// Enum value maps for AlterGroupMemberResponse_AlterGroupMemberResponseCode.
var (
	AlterGroupMemberResponse_AlterGroupMemberResponseCode_name = map[int32]string{
		0: "OK",
	}
	AlterGroupMemberResponse_AlterGroupMemberResponseCode_value = map[string]int32{
		"OK": 0,
	}
)

func (x AlterGroupMemberResponse_AlterGroupMemberResponseCode) Enum() *AlterGroupMemberResponse_AlterGroupMemberResponseCode {
	p := new(AlterGroupMemberResponse_AlterGroupMemberResponseCode)
	*p = x
	return p
}

func (x AlterGroupMemberResponse_AlterGroupMemberResponseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlterGroupMemberResponse_AlterGroupMemberResponseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_base_evloopio_proto_enumTypes[0].Descriptor()
}

func (AlterGroupMemberResponse_AlterGroupMemberResponseCode) Type() protoreflect.EnumType {
	return &file_base_evloopio_proto_enumTypes[0]
}

func (x AlterGroupMemberResponse_AlterGroupMemberResponseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AlterGroupMemberResponse_AlterGroupMemberResponseCode.Descriptor instead.
func (AlterGroupMemberResponse_AlterGroupMemberResponseCode) EnumDescriptor() ([]byte, []int) {
	return file_base_evloopio_proto_rawDescGZIP(), []int{1, 0}
}

type SubscribeGroupResponse_SubscribeGroupCode int32

const (
	SubscribeGroupResponse_OK            SubscribeGroupResponse_SubscribeGroupCode = 0
	SubscribeGroupResponse_VERSION_OLD   SubscribeGroupResponse_SubscribeGroupCode = 1 // 这个错误, version有效调用方可以等待version更新后重试
	SubscribeGroupResponse_NO_PERMISSION SubscribeGroupResponse_SubscribeGroupCode = 2 // 这个错误, 没有权限, relation_id是有效的
)

// Enum value maps for SubscribeGroupResponse_SubscribeGroupCode.
var (
	SubscribeGroupResponse_SubscribeGroupCode_name = map[int32]string{
		0: "OK",
		1: "VERSION_OLD",
		2: "NO_PERMISSION",
	}
	SubscribeGroupResponse_SubscribeGroupCode_value = map[string]int32{
		"OK":            0,
		"VERSION_OLD":   1,
		"NO_PERMISSION": 2,
	}
)

func (x SubscribeGroupResponse_SubscribeGroupCode) Enum() *SubscribeGroupResponse_SubscribeGroupCode {
	p := new(SubscribeGroupResponse_SubscribeGroupCode)
	*p = x
	return p
}

func (x SubscribeGroupResponse_SubscribeGroupCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubscribeGroupResponse_SubscribeGroupCode) Descriptor() protoreflect.EnumDescriptor {
	return file_base_evloopio_proto_enumTypes[1].Descriptor()
}

func (SubscribeGroupResponse_SubscribeGroupCode) Type() protoreflect.EnumType {
	return &file_base_evloopio_proto_enumTypes[1]
}

func (x SubscribeGroupResponse_SubscribeGroupCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubscribeGroupResponse_SubscribeGroupCode.Descriptor instead.
func (SubscribeGroupResponse_SubscribeGroupCode) EnumDescriptor() ([]byte, []int) {
	return file_base_evloopio_proto_rawDescGZIP(), []int{3, 0}
}

type SendMessageResponse_SendMessageCode int32

const (
	SendMessageResponse_OK            SendMessageResponse_SendMessageCode = 0
	SendMessageResponse_VERSION_OLD   SendMessageResponse_SendMessageCode = 1 // 这个错误, version有效调用方可以等待version更新后重试
	SendMessageResponse_NO_PERMISSION SendMessageResponse_SendMessageCode = 2 // 这个错误, 没有权限, relation_id是有效的
)

// Enum value maps for SendMessageResponse_SendMessageCode.
var (
	SendMessageResponse_SendMessageCode_name = map[int32]string{
		0: "OK",
		1: "VERSION_OLD",
		2: "NO_PERMISSION",
	}
	SendMessageResponse_SendMessageCode_value = map[string]int32{
		"OK":            0,
		"VERSION_OLD":   1,
		"NO_PERMISSION": 2,
	}
)

func (x SendMessageResponse_SendMessageCode) Enum() *SendMessageResponse_SendMessageCode {
	p := new(SendMessageResponse_SendMessageCode)
	*p = x
	return p
}

func (x SendMessageResponse_SendMessageCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SendMessageResponse_SendMessageCode) Descriptor() protoreflect.EnumDescriptor {
	return file_base_evloopio_proto_enumTypes[2].Descriptor()
}

func (SendMessageResponse_SendMessageCode) Type() protoreflect.EnumType {
	return &file_base_evloopio_proto_enumTypes[2]
}

func (x SendMessageResponse_SendMessageCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SendMessageResponse_SendMessageCode.Descriptor instead.
func (SendMessageResponse_SendMessageCode) EnumDescriptor() ([]byte, []int) {
	return file_base_evloopio_proto_rawDescGZIP(), []int{5, 0}
}

// 更改群组成员, 由im-relation调用
type AlterGroupMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relation *base.RelationEntry `protobuf:"bytes,1,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *AlterGroupMemberRequest) Reset() {
	*x = AlterGroupMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_evloopio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlterGroupMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterGroupMemberRequest) ProtoMessage() {}

func (x *AlterGroupMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_evloopio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterGroupMemberRequest.ProtoReflect.Descriptor instead.
func (*AlterGroupMemberRequest) Descriptor() ([]byte, []int) {
	return file_base_evloopio_proto_rawDescGZIP(), []int{0}
}

func (x *AlterGroupMemberRequest) GetRelation() *base.RelationEntry {
	if x != nil {
		return x.Relation
	}
	return nil
}

type AlterGroupMemberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code            AlterGroupMemberResponse_AlterGroupMemberResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=evloopio.AlterGroupMemberResponse_AlterGroupMemberResponseCode" json:"code,omitempty"`
	RelationVersion int64                                                 `protobuf:"varint,2,opt,name=relation_version,json=relationVersion,proto3" json:"relation_version,omitempty"`
	CurrentSeqId    int64                                                 `protobuf:"varint,3,opt,name=current_seq_id,json=currentSeqId,proto3" json:"current_seq_id,omitempty"`
	ChangeAt        int64                                                 `protobuf:"varint,4,opt,name=changeAt,proto3" json:"changeAt,omitempty"`
}

func (x *AlterGroupMemberResponse) Reset() {
	*x = AlterGroupMemberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_evloopio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlterGroupMemberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterGroupMemberResponse) ProtoMessage() {}

func (x *AlterGroupMemberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_base_evloopio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterGroupMemberResponse.ProtoReflect.Descriptor instead.
func (*AlterGroupMemberResponse) Descriptor() ([]byte, []int) {
	return file_base_evloopio_proto_rawDescGZIP(), []int{1}
}

func (x *AlterGroupMemberResponse) GetCode() AlterGroupMemberResponse_AlterGroupMemberResponseCode {
	if x != nil {
		return x.Code
	}
	return AlterGroupMemberResponse_OK
}

func (x *AlterGroupMemberResponse) GetRelationVersion() int64 {
	if x != nil {
		return x.RelationVersion
	}
	return 0
}

func (x *AlterGroupMemberResponse) GetCurrentSeqId() int64 {
	if x != nil {
		return x.CurrentSeqId
	}
	return 0
}

func (x *AlterGroupMemberResponse) GetChangeAt() int64 {
	if x != nil {
		return x.ChangeAt
	}
	return 0
}

// 订阅群组, 由im-relay调用, 某个人订阅群组, im-relay先确认是否存在这个群(请求im-relation), 然后一致性hash
type SubscribeGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session *base.SessionEntry `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *SubscribeGroupRequest) Reset() {
	*x = SubscribeGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_evloopio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeGroupRequest) ProtoMessage() {}

func (x *SubscribeGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_evloopio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeGroupRequest.ProtoReflect.Descriptor instead.
func (*SubscribeGroupRequest) Descriptor() ([]byte, []int) {
	return file_base_evloopio_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeGroupRequest) GetSession() *base.SessionEntry {
	if x != nil {
		return x.Session
	}
	return nil
}

type SubscribeGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       SubscribeGroupResponse_SubscribeGroupCode `protobuf:"varint,1,opt,name=code,proto3,enum=evloopio.SubscribeGroupResponse_SubscribeGroupCode" json:"code,omitempty"`
	RelationId int64                                     `protobuf:"varint,2,opt,name=relation_id,json=relationId,proto3" json:"relation_id,omitempty"`
	MaxSeqId   int64                                     `protobuf:"varint,3,opt,name=max_seq_id,json=maxSeqId,proto3" json:"max_seq_id,omitempty"`
}

func (x *SubscribeGroupResponse) Reset() {
	*x = SubscribeGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_evloopio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeGroupResponse) ProtoMessage() {}

func (x *SubscribeGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_base_evloopio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeGroupResponse.ProtoReflect.Descriptor instead.
func (*SubscribeGroupResponse) Descriptor() ([]byte, []int) {
	return file_base_evloopio_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribeGroupResponse) GetCode() SubscribeGroupResponse_SubscribeGroupCode {
	if x != nil {
		return x.Code
	}
	return SubscribeGroupResponse_OK
}

func (x *SubscribeGroupResponse) GetRelationId() int64 {
	if x != nil {
		return x.RelationId
	}
	return 0
}

func (x *SubscribeGroupResponse) GetMaxSeqId() int64 {
	if x != nil {
		return x.MaxSeqId
	}
	return 0
}

// 发送消息
type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageData     []byte `protobuf:"bytes,1,opt,name=message_data,json=messageData,proto3" json:"message_data,omitempty"`
	CheckIdempotent bool   `protobuf:"varint,2,opt,name=check_idempotent,json=checkIdempotent,proto3" json:"check_idempotent,omitempty"` // 是否需要检查幂等
	IdempotentKey   string `protobuf:"bytes,3,opt,name=idempotent_key,json=idempotentKey,proto3" json:"idempotent_key,omitempty"`        // 幂等key
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_evloopio_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_evloopio_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_base_evloopio_proto_rawDescGZIP(), []int{4}
}

func (x *SendMessageRequest) GetMessageData() []byte {
	if x != nil {
		return x.MessageData
	}
	return nil
}

func (x *SendMessageRequest) GetCheckIdempotent() bool {
	if x != nil {
		return x.CheckIdempotent
	}
	return false
}

func (x *SendMessageRequest) GetIdempotentKey() string {
	if x != nil {
		return x.IdempotentKey
	}
	return ""
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       SendMessageResponse_SendMessageCode `protobuf:"varint,1,opt,name=code,proto3,enum=evloopio.SendMessageResponse_SendMessageCode" json:"code,omitempty"`
	RelationId int64                               `protobuf:"varint,2,opt,name=relation_id,json=relationId,proto3" json:"relation_id,omitempty"`
	MaxSeqId   int64                               `protobuf:"varint,3,opt,name=max_seq_id,json=maxSeqId,proto3" json:"max_seq_id,omitempty"`
	MessageSeq int64                               `protobuf:"varint,4,opt,name=message_seq,json=messageSeq,proto3" json:"message_seq,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_evloopio_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_base_evloopio_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_base_evloopio_proto_rawDescGZIP(), []int{5}
}

func (x *SendMessageResponse) GetCode() SendMessageResponse_SendMessageCode {
	if x != nil {
		return x.Code
	}
	return SendMessageResponse_OK
}

func (x *SendMessageResponse) GetRelationId() int64 {
	if x != nil {
		return x.RelationId
	}
	return 0
}

func (x *SendMessageResponse) GetMaxSeqId() int64 {
	if x != nil {
		return x.MaxSeqId
	}
	return 0
}

func (x *SendMessageResponse) GetMessageSeq() int64 {
	if x != nil {
		return x.MessageSeq
	}
	return 0
}

type UniversalGroupEvloopInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Input:
	//	*UniversalGroupEvloopInput_AlterGroupMember
	//	*UniversalGroupEvloopInput_SubscribeGroup
	//	*UniversalGroupEvloopInput_SendMessage
	Input isUniversalGroupEvloopInput_Input `protobuf_oneof:"input"`
}

func (x *UniversalGroupEvloopInput) Reset() {
	*x = UniversalGroupEvloopInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_evloopio_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniversalGroupEvloopInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniversalGroupEvloopInput) ProtoMessage() {}

func (x *UniversalGroupEvloopInput) ProtoReflect() protoreflect.Message {
	mi := &file_base_evloopio_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniversalGroupEvloopInput.ProtoReflect.Descriptor instead.
func (*UniversalGroupEvloopInput) Descriptor() ([]byte, []int) {
	return file_base_evloopio_proto_rawDescGZIP(), []int{6}
}

func (m *UniversalGroupEvloopInput) GetInput() isUniversalGroupEvloopInput_Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (x *UniversalGroupEvloopInput) GetAlterGroupMember() *AlterGroupMemberRequest {
	if x, ok := x.GetInput().(*UniversalGroupEvloopInput_AlterGroupMember); ok {
		return x.AlterGroupMember
	}
	return nil
}

func (x *UniversalGroupEvloopInput) GetSubscribeGroup() *SubscribeGroupRequest {
	if x, ok := x.GetInput().(*UniversalGroupEvloopInput_SubscribeGroup); ok {
		return x.SubscribeGroup
	}
	return nil
}

func (x *UniversalGroupEvloopInput) GetSendMessage() *SendMessageRequest {
	if x, ok := x.GetInput().(*UniversalGroupEvloopInput_SendMessage); ok {
		return x.SendMessage
	}
	return nil
}

type isUniversalGroupEvloopInput_Input interface {
	isUniversalGroupEvloopInput_Input()
}

type UniversalGroupEvloopInput_AlterGroupMember struct {
	AlterGroupMember *AlterGroupMemberRequest `protobuf:"bytes,1,opt,name=alter_group_member,json=alterGroupMember,proto3,oneof"`
}

type UniversalGroupEvloopInput_SubscribeGroup struct {
	SubscribeGroup *SubscribeGroupRequest `protobuf:"bytes,2,opt,name=subscribe_group,json=subscribeGroup,proto3,oneof"`
}

type UniversalGroupEvloopInput_SendMessage struct {
	SendMessage *SendMessageRequest `protobuf:"bytes,3,opt,name=send_message,json=sendMessage,proto3,oneof"` // GroupRequest disband_group = 4;
}

func (*UniversalGroupEvloopInput_AlterGroupMember) isUniversalGroupEvloopInput_Input() {}

func (*UniversalGroupEvloopInput_SubscribeGroup) isUniversalGroupEvloopInput_Input() {}

func (*UniversalGroupEvloopInput_SendMessage) isUniversalGroupEvloopInput_Input() {}

type UniversalGroupEvloopOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Output:
	//	*UniversalGroupEvloopOutput_AlterGroupMember
	//	*UniversalGroupEvloopOutput_SubscribeGroup
	//	*UniversalGroupEvloopOutput_SendMessage
	Output isUniversalGroupEvloopOutput_Output `protobuf_oneof:"output"`
}

func (x *UniversalGroupEvloopOutput) Reset() {
	*x = UniversalGroupEvloopOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_evloopio_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniversalGroupEvloopOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniversalGroupEvloopOutput) ProtoMessage() {}

func (x *UniversalGroupEvloopOutput) ProtoReflect() protoreflect.Message {
	mi := &file_base_evloopio_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniversalGroupEvloopOutput.ProtoReflect.Descriptor instead.
func (*UniversalGroupEvloopOutput) Descriptor() ([]byte, []int) {
	return file_base_evloopio_proto_rawDescGZIP(), []int{7}
}

func (m *UniversalGroupEvloopOutput) GetOutput() isUniversalGroupEvloopOutput_Output {
	if m != nil {
		return m.Output
	}
	return nil
}

func (x *UniversalGroupEvloopOutput) GetAlterGroupMember() *AlterGroupMemberResponse {
	if x, ok := x.GetOutput().(*UniversalGroupEvloopOutput_AlterGroupMember); ok {
		return x.AlterGroupMember
	}
	return nil
}

func (x *UniversalGroupEvloopOutput) GetSubscribeGroup() *SubscribeGroupResponse {
	if x, ok := x.GetOutput().(*UniversalGroupEvloopOutput_SubscribeGroup); ok {
		return x.SubscribeGroup
	}
	return nil
}

func (x *UniversalGroupEvloopOutput) GetSendMessage() *SendMessageResponse {
	if x, ok := x.GetOutput().(*UniversalGroupEvloopOutput_SendMessage); ok {
		return x.SendMessage
	}
	return nil
}

type isUniversalGroupEvloopOutput_Output interface {
	isUniversalGroupEvloopOutput_Output()
}

type UniversalGroupEvloopOutput_AlterGroupMember struct {
	AlterGroupMember *AlterGroupMemberResponse `protobuf:"bytes,1,opt,name=alter_group_member,json=alterGroupMember,proto3,oneof"`
}

type UniversalGroupEvloopOutput_SubscribeGroup struct {
	SubscribeGroup *SubscribeGroupResponse `protobuf:"bytes,2,opt,name=subscribe_group,json=subscribeGroup,proto3,oneof"`
}

type UniversalGroupEvloopOutput_SendMessage struct {
	SendMessage *SendMessageResponse `protobuf:"bytes,3,opt,name=send_message,json=sendMessage,proto3,oneof"` // DisbandGroupResponse disband_group = 4;
}

func (*UniversalGroupEvloopOutput_AlterGroupMember) isUniversalGroupEvloopOutput_Output() {}

func (*UniversalGroupEvloopOutput_SubscribeGroup) isUniversalGroupEvloopOutput_Output() {}

func (*UniversalGroupEvloopOutput_SendMessage) isUniversalGroupEvloopOutput_Output() {}

var File_base_evloopio_proto protoreflect.FileDescriptor

var file_base_evloopio_proto_rawDesc = []byte{
	0x0a, 0x13, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x6f, 0x1a,
	0x0f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4a, 0x0a, 0x17, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x84, 0x02, 0x0a,
	0x18, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70,
	0x69, 0x6f, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x6c, 0x74, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x71, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x74, 0x22, 0x26, 0x0a, 0x1c, 0x41,
	0x6c, 0x74, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x00, 0x22, 0x45, 0x0a, 0x15, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x07,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xe2, 0x01, 0x0a, 0x16, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x6f, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x65, 0x71, 0x49, 0x64, 0x22, 0x40, 0x0a,
	0x12, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x56,
	0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x4c, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d,
	0x4e, 0x4f, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x22,
	0x89, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x5f, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x65, 0x6d, 0x70, 0x6f,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x64,
	0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x22, 0xf7, 0x01, 0x0a, 0x13,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2d, 0x2e, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x73,
	0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x78,
	0x53, 0x65, 0x71, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x71, 0x22, 0x3d, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x4c, 0x44,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4f, 0x4e, 0x10, 0x02, 0x22, 0x86, 0x02, 0x0a, 0x19, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x51, 0x0a, 0x12, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x6f, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x10, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x41, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x76, 0x6c, 0x6f, 0x6f,
	0x70, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x8b,
	0x02, 0x0a, 0x1a, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x45, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x52, 0x0a,
	0x12, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x76, 0x6c, 0x6f,
	0x6f, 0x70, 0x69, 0x6f, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52,
	0x10, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x4b, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x76, 0x6c,
	0x6f, 0x6f, 0x70, 0x69, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x42,
	0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x6f, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x23, 0x5a, 0x21,
	0x70, 0x69, 0x67, 0x65, 0x6f, 0x6e, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x69,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_evloopio_proto_rawDescOnce sync.Once
	file_base_evloopio_proto_rawDescData = file_base_evloopio_proto_rawDesc
)

func file_base_evloopio_proto_rawDescGZIP() []byte {
	file_base_evloopio_proto_rawDescOnce.Do(func() {
		file_base_evloopio_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_evloopio_proto_rawDescData)
	})
	return file_base_evloopio_proto_rawDescData
}

var file_base_evloopio_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_base_evloopio_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_base_evloopio_proto_goTypes = []interface{}{
	(AlterGroupMemberResponse_AlterGroupMemberResponseCode)(0), // 0: evloopio.AlterGroupMemberResponse.AlterGroupMemberResponseCode
	(SubscribeGroupResponse_SubscribeGroupCode)(0),             // 1: evloopio.SubscribeGroupResponse.SubscribeGroupCode
	(SendMessageResponse_SendMessageCode)(0),                   // 2: evloopio.SendMessageResponse.SendMessageCode
	(*AlterGroupMemberRequest)(nil),                            // 3: evloopio.AlterGroupMemberRequest
	(*AlterGroupMemberResponse)(nil),                           // 4: evloopio.AlterGroupMemberResponse
	(*SubscribeGroupRequest)(nil),                              // 5: evloopio.SubscribeGroupRequest
	(*SubscribeGroupResponse)(nil),                             // 6: evloopio.SubscribeGroupResponse
	(*SendMessageRequest)(nil),                                 // 7: evloopio.SendMessageRequest
	(*SendMessageResponse)(nil),                                // 8: evloopio.SendMessageResponse
	(*UniversalGroupEvloopInput)(nil),                          // 9: evloopio.UniversalGroupEvloopInput
	(*UniversalGroupEvloopOutput)(nil),                         // 10: evloopio.UniversalGroupEvloopOutput
	(*base.RelationEntry)(nil),                                 // 11: base.RelationEntry
	(*base.SessionEntry)(nil),                                  // 12: base.SessionEntry
}
var file_base_evloopio_proto_depIdxs = []int32{
	11, // 0: evloopio.AlterGroupMemberRequest.relation:type_name -> base.RelationEntry
	0,  // 1: evloopio.AlterGroupMemberResponse.code:type_name -> evloopio.AlterGroupMemberResponse.AlterGroupMemberResponseCode
	12, // 2: evloopio.SubscribeGroupRequest.session:type_name -> base.SessionEntry
	1,  // 3: evloopio.SubscribeGroupResponse.code:type_name -> evloopio.SubscribeGroupResponse.SubscribeGroupCode
	2,  // 4: evloopio.SendMessageResponse.code:type_name -> evloopio.SendMessageResponse.SendMessageCode
	3,  // 5: evloopio.UniversalGroupEvloopInput.alter_group_member:type_name -> evloopio.AlterGroupMemberRequest
	5,  // 6: evloopio.UniversalGroupEvloopInput.subscribe_group:type_name -> evloopio.SubscribeGroupRequest
	7,  // 7: evloopio.UniversalGroupEvloopInput.send_message:type_name -> evloopio.SendMessageRequest
	4,  // 8: evloopio.UniversalGroupEvloopOutput.alter_group_member:type_name -> evloopio.AlterGroupMemberResponse
	6,  // 9: evloopio.UniversalGroupEvloopOutput.subscribe_group:type_name -> evloopio.SubscribeGroupResponse
	8,  // 10: evloopio.UniversalGroupEvloopOutput.send_message:type_name -> evloopio.SendMessageResponse
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_base_evloopio_proto_init() }
func file_base_evloopio_proto_init() {
	if File_base_evloopio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_evloopio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlterGroupMemberRequest); i {
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
		file_base_evloopio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlterGroupMemberResponse); i {
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
		file_base_evloopio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeGroupRequest); i {
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
		file_base_evloopio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeGroupResponse); i {
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
		file_base_evloopio_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_base_evloopio_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
		file_base_evloopio_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniversalGroupEvloopInput); i {
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
		file_base_evloopio_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniversalGroupEvloopOutput); i {
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
	file_base_evloopio_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*UniversalGroupEvloopInput_AlterGroupMember)(nil),
		(*UniversalGroupEvloopInput_SubscribeGroup)(nil),
		(*UniversalGroupEvloopInput_SendMessage)(nil),
	}
	file_base_evloopio_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*UniversalGroupEvloopOutput_AlterGroupMember)(nil),
		(*UniversalGroupEvloopOutput_SubscribeGroup)(nil),
		(*UniversalGroupEvloopOutput_SendMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_base_evloopio_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_evloopio_proto_goTypes,
		DependencyIndexes: file_base_evloopio_proto_depIdxs,
		EnumInfos:         file_base_evloopio_proto_enumTypes,
		MessageInfos:      file_base_evloopio_proto_msgTypes,
	}.Build()
	File_base_evloopio_proto = out.File
	file_base_evloopio_proto_rawDesc = nil
	file_base_evloopio_proto_goTypes = nil
	file_base_evloopio_proto_depIdxs = nil
}

var _ context.Context
