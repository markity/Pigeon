// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: service/im-chat-evloop.proto

package imchatevloop

import (
	context "context"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	base "pigeon/kitex_gen/service/base"
	evloopio "pigeon/kitex_gen/service/evloopio"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 创建群组, 由im-relation调用
type CreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version      int64  `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	GroupId      string `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	GroupOwnerId string `protobuf:"bytes,3,opt,name=group_owner_id,json=groupOwnerId,proto3" json:"group_owner_id,omitempty"`
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_im_chat_evloop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_im_chat_evloop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_service_im_chat_evloop_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGroupRequest) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *CreateGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *CreateGroupRequest) GetGroupOwnerId() string {
	if x != nil {
		return x.GroupOwnerId
	}
	return ""
}

type CreateGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success   bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Version   int64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	CreatedAt int64 `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *CreateGroupResponse) Reset() {
	*x = CreateGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_im_chat_evloop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupResponse) ProtoMessage() {}

func (x *CreateGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_im_chat_evloop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateGroupResponse) Descriptor() ([]byte, []int) {
	return file_service_im_chat_evloop_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGroupResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateGroupResponse) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *CreateGroupResponse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type UniversalGroupEvloopRequestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version int64                               `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	GroupId string                              `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Input   *evloopio.UniversalGroupEvloopInput `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *UniversalGroupEvloopRequestReq) Reset() {
	*x = UniversalGroupEvloopRequestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_im_chat_evloop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniversalGroupEvloopRequestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniversalGroupEvloopRequestReq) ProtoMessage() {}

func (x *UniversalGroupEvloopRequestReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_im_chat_evloop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniversalGroupEvloopRequestReq.ProtoReflect.Descriptor instead.
func (*UniversalGroupEvloopRequestReq) Descriptor() ([]byte, []int) {
	return file_service_im_chat_evloop_proto_rawDescGZIP(), []int{2}
}

func (x *UniversalGroupEvloopRequestReq) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *UniversalGroupEvloopRequestReq) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *UniversalGroupEvloopRequestReq) GetInput() *evloopio.UniversalGroupEvloopInput {
	if x != nil {
		return x.Input
	}
	return nil
}

type UniversalGroupEvloopRequestResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool                                 `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Version int64                                `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Output  *evloopio.UniversalGroupEvloopOutput `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *UniversalGroupEvloopRequestResp) Reset() {
	*x = UniversalGroupEvloopRequestResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_im_chat_evloop_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniversalGroupEvloopRequestResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniversalGroupEvloopRequestResp) ProtoMessage() {}

func (x *UniversalGroupEvloopRequestResp) ProtoReflect() protoreflect.Message {
	mi := &file_service_im_chat_evloop_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniversalGroupEvloopRequestResp.ProtoReflect.Descriptor instead.
func (*UniversalGroupEvloopRequestResp) Descriptor() ([]byte, []int) {
	return file_service_im_chat_evloop_proto_rawDescGZIP(), []int{3}
}

func (x *UniversalGroupEvloopRequestResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UniversalGroupEvloopRequestResp) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *UniversalGroupEvloopRequestResp) GetOutput() *evloopio.UniversalGroupEvloopOutput {
	if x != nil {
		return x.Output
	}
	return nil
}

// 迁移相关, 幂等接口可以重复调用
type DoMigrateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *DoMigrateReq) Reset() {
	*x = DoMigrateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_im_chat_evloop_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoMigrateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoMigrateReq) ProtoMessage() {}

func (x *DoMigrateReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_im_chat_evloop_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoMigrateReq.ProtoReflect.Descriptor instead.
func (*DoMigrateReq) Descriptor() ([]byte, []int) {
	return file_service_im_chat_evloop_proto_rawDescGZIP(), []int{4}
}

func (x *DoMigrateReq) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type DoMigrateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok          bool                                         `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	GroupId     string                                       `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	OwnerId     string                                       `protobuf:"bytes,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	CreatedAt   int64                                        `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	SeqId       int64                                        `protobuf:"varint,5,opt,name=seq_id,json=seqId,proto3" json:"seq_id,omitempty"`
	Relations   map[string]*DoMigrateResp_RelationInfo       `protobuf:"bytes,6,rep,name=relations,proto3" json:"relations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Subscribers map[string]*DoMigrateResp_UserSubscribeEntry `protobuf:"bytes,7,rep,name=subscribers,proto3" json:"subscribers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DoMigrateResp) Reset() {
	*x = DoMigrateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_im_chat_evloop_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoMigrateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoMigrateResp) ProtoMessage() {}

func (x *DoMigrateResp) ProtoReflect() protoreflect.Message {
	mi := &file_service_im_chat_evloop_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoMigrateResp.ProtoReflect.Descriptor instead.
func (*DoMigrateResp) Descriptor() ([]byte, []int) {
	return file_service_im_chat_evloop_proto_rawDescGZIP(), []int{5}
}

func (x *DoMigrateResp) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *DoMigrateResp) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *DoMigrateResp) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *DoMigrateResp) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *DoMigrateResp) GetSeqId() int64 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

func (x *DoMigrateResp) GetRelations() map[string]*DoMigrateResp_RelationInfo {
	if x != nil {
		return x.Relations
	}
	return nil
}

func (x *DoMigrateResp) GetSubscribers() map[string]*DoMigrateResp_UserSubscribeEntry {
	if x != nil {
		return x.Subscribers
	}
	return nil
}

type MigrateDoneReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *MigrateDoneReq) Reset() {
	*x = MigrateDoneReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_im_chat_evloop_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrateDoneReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrateDoneReq) ProtoMessage() {}

func (x *MigrateDoneReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_im_chat_evloop_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrateDoneReq.ProtoReflect.Descriptor instead.
func (*MigrateDoneReq) Descriptor() ([]byte, []int) {
	return file_service_im_chat_evloop_proto_rawDescGZIP(), []int{6}
}

func (x *MigrateDoneReq) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type MigrateDoneResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *MigrateDoneResp) Reset() {
	*x = MigrateDoneResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_im_chat_evloop_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrateDoneResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrateDoneResp) ProtoMessage() {}

func (x *MigrateDoneResp) ProtoReflect() protoreflect.Message {
	mi := &file_service_im_chat_evloop_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrateDoneResp.ProtoReflect.Descriptor instead.
func (*MigrateDoneResp) Descriptor() ([]byte, []int) {
	return file_service_im_chat_evloop_proto_rawDescGZIP(), []int{7}
}

func (x *MigrateDoneResp) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type DoMigrateResp_RelationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relation *base.RelationEntry `protobuf:"bytes,1,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *DoMigrateResp_RelationInfo) Reset() {
	*x = DoMigrateResp_RelationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_im_chat_evloop_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoMigrateResp_RelationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoMigrateResp_RelationInfo) ProtoMessage() {}

func (x *DoMigrateResp_RelationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_service_im_chat_evloop_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoMigrateResp_RelationInfo.ProtoReflect.Descriptor instead.
func (*DoMigrateResp_RelationInfo) Descriptor() ([]byte, []int) {
	return file_service_im_chat_evloop_proto_rawDescGZIP(), []int{5, 0}
}

func (x *DoMigrateResp_RelationInfo) GetRelation() *base.RelationEntry {
	if x != nil {
		return x.Relation
	}
	return nil
}

type DoMigrateResp_UserSubscribeEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*DoMigrateResp_UserSubscribeEntry_SubscribeEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *DoMigrateResp_UserSubscribeEntry) Reset() {
	*x = DoMigrateResp_UserSubscribeEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_im_chat_evloop_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoMigrateResp_UserSubscribeEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoMigrateResp_UserSubscribeEntry) ProtoMessage() {}

func (x *DoMigrateResp_UserSubscribeEntry) ProtoReflect() protoreflect.Message {
	mi := &file_service_im_chat_evloop_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoMigrateResp_UserSubscribeEntry.ProtoReflect.Descriptor instead.
func (*DoMigrateResp_UserSubscribeEntry) Descriptor() ([]byte, []int) {
	return file_service_im_chat_evloop_proto_rawDescGZIP(), []int{5, 1}
}

func (x *DoMigrateResp_UserSubscribeEntry) GetEntries() []*DoMigrateResp_UserSubscribeEntry_SubscribeEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type DoMigrateResp_UserSubscribeEntry_SubscribeEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnSubRelationVersion int64              `protobuf:"varint,1,opt,name=on_sub_relation_version,json=onSubRelationVersion,proto3" json:"on_sub_relation_version,omitempty"`
	Session              *base.SessionEntry `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *DoMigrateResp_UserSubscribeEntry_SubscribeEntry) Reset() {
	*x = DoMigrateResp_UserSubscribeEntry_SubscribeEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_im_chat_evloop_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoMigrateResp_UserSubscribeEntry_SubscribeEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoMigrateResp_UserSubscribeEntry_SubscribeEntry) ProtoMessage() {}

func (x *DoMigrateResp_UserSubscribeEntry_SubscribeEntry) ProtoReflect() protoreflect.Message {
	mi := &file_service_im_chat_evloop_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoMigrateResp_UserSubscribeEntry_SubscribeEntry.ProtoReflect.Descriptor instead.
func (*DoMigrateResp_UserSubscribeEntry_SubscribeEntry) Descriptor() ([]byte, []int) {
	return file_service_im_chat_evloop_proto_rawDescGZIP(), []int{5, 1, 0}
}

func (x *DoMigrateResp_UserSubscribeEntry_SubscribeEntry) GetOnSubRelationVersion() int64 {
	if x != nil {
		return x.OnSubRelationVersion
	}
	return 0
}

func (x *DoMigrateResp_UserSubscribeEntry_SubscribeEntry) GetSession() *base.SessionEntry {
	if x != nil {
		return x.Session
	}
	return nil
}

var File_service_im_chat_evloop_proto protoreflect.FileDescriptor

var file_service_im_chat_evloop_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6d, 0x2d, 0x63, 0x68, 0x61,
	0x74, 0x2d, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x69, 0x6d, 0x63, 0x68, 0x61, 0x74, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x1a, 0x0f, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x68, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x90, 0x01, 0x0a,
	0x1e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45,
	0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x69, 0x6f, 0x2e, 0x55,
	0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x76, 0x6c,
	0x6f, 0x6f, 0x70, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22,
	0x93, 0x01, 0x0a, 0x1f, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x45, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70,
	0x69, 0x6f, 0x2e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x45, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x29, 0x0a, 0x0c, 0x44, 0x6f, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x22, 0xa5, 0x06, 0x0a, 0x0d, 0x44, 0x6f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02,
	0x6f, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64, 0x12, 0x48,
	0x0a, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6d, 0x63, 0x68, 0x61, 0x74, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70,
	0x2e, 0x44, 0x6f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4e, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x69, 0x6d, 0x63, 0x68, 0x61, 0x74, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x44, 0x6f, 0x4d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x1a, 0x3f, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0xe4, 0x01, 0x0a, 0x12, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x57, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3d, 0x2e, 0x69, 0x6d, 0x63, 0x68, 0x61, 0x74, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70,
	0x2e, 0x44, 0x6f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x75, 0x0a, 0x0e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x35, 0x0a, 0x17, 0x6f,
	0x6e, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6f, 0x6e,
	0x53, 0x75, 0x62, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x1a, 0x66, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x3e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6d, 0x63, 0x68, 0x61, 0x74, 0x65, 0x76, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x44, 0x6f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x6e, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x44,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x69, 0x6d, 0x63, 0x68, 0x61, 0x74, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x44, 0x6f, 0x4d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2b, 0x0a, 0x0e, 0x4d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x65, 0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65,
	0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0xf0, 0x02, 0x0a, 0x0c, 0x49, 0x4d, 0x43,
	0x68, 0x61, 0x74, 0x45, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x12, 0x52, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x2e, 0x69, 0x6d, 0x63, 0x68, 0x61,
	0x74, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6d, 0x63,
	0x68, 0x61, 0x74, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7a, 0x0a,
	0x1b, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45,
	0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x2e, 0x69,
	0x6d, 0x63, 0x68, 0x61, 0x74, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x55, 0x6e, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x76, 0x6c, 0x6f, 0x6f, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x2d, 0x2e, 0x69, 0x6d, 0x63,
	0x68, 0x61, 0x74, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x44, 0x0a, 0x09, 0x44, 0x6f, 0x4d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x69, 0x6d, 0x63, 0x68, 0x61, 0x74, 0x65,
	0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x44, 0x6f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x69, 0x6d, 0x63, 0x68, 0x61, 0x74, 0x65, 0x76, 0x6c, 0x6f, 0x6f,
	0x70, 0x2e, 0x44, 0x6f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x4a, 0x0a, 0x0b, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x1c,
	0x2e, 0x69, 0x6d, 0x63, 0x68, 0x61, 0x74, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x4d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x69,
	0x6d, 0x63, 0x68, 0x61, 0x74, 0x65, 0x76, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x4d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x65, 0x44, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x27, 0x5a, 0x25, 0x70,
	0x69, 0x67, 0x65, 0x6f, 0x6e, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6d, 0x63, 0x68, 0x61, 0x74, 0x65, 0x76,
	0x6c, 0x6f, 0x6f, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_im_chat_evloop_proto_rawDescOnce sync.Once
	file_service_im_chat_evloop_proto_rawDescData = file_service_im_chat_evloop_proto_rawDesc
)

func file_service_im_chat_evloop_proto_rawDescGZIP() []byte {
	file_service_im_chat_evloop_proto_rawDescOnce.Do(func() {
		file_service_im_chat_evloop_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_im_chat_evloop_proto_rawDescData)
	})
	return file_service_im_chat_evloop_proto_rawDescData
}

var file_service_im_chat_evloop_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_service_im_chat_evloop_proto_goTypes = []interface{}{
	(*CreateGroupRequest)(nil),               // 0: imchatevloop.CreateGroupRequest
	(*CreateGroupResponse)(nil),              // 1: imchatevloop.CreateGroupResponse
	(*UniversalGroupEvloopRequestReq)(nil),   // 2: imchatevloop.UniversalGroupEvloopRequestReq
	(*UniversalGroupEvloopRequestResp)(nil),  // 3: imchatevloop.UniversalGroupEvloopRequestResp
	(*DoMigrateReq)(nil),                     // 4: imchatevloop.DoMigrateReq
	(*DoMigrateResp)(nil),                    // 5: imchatevloop.DoMigrateResp
	(*MigrateDoneReq)(nil),                   // 6: imchatevloop.MigrateDoneReq
	(*MigrateDoneResp)(nil),                  // 7: imchatevloop.MigrateDoneResp
	(*DoMigrateResp_RelationInfo)(nil),       // 8: imchatevloop.DoMigrateResp.RelationInfo
	(*DoMigrateResp_UserSubscribeEntry)(nil), // 9: imchatevloop.DoMigrateResp.UserSubscribeEntry
	nil,                                      // 10: imchatevloop.DoMigrateResp.RelationsEntry
	nil,                                      // 11: imchatevloop.DoMigrateResp.SubscribersEntry
	(*DoMigrateResp_UserSubscribeEntry_SubscribeEntry)(nil), // 12: imchatevloop.DoMigrateResp.UserSubscribeEntry.SubscribeEntry
	(*evloopio.UniversalGroupEvloopInput)(nil),              // 13: evloopio.UniversalGroupEvloopInput
	(*evloopio.UniversalGroupEvloopOutput)(nil),             // 14: evloopio.UniversalGroupEvloopOutput
	(*base.RelationEntry)(nil),                              // 15: base.RelationEntry
	(*base.SessionEntry)(nil),                               // 16: base.SessionEntry
}
var file_service_im_chat_evloop_proto_depIdxs = []int32{
	13, // 0: imchatevloop.UniversalGroupEvloopRequestReq.input:type_name -> evloopio.UniversalGroupEvloopInput
	14, // 1: imchatevloop.UniversalGroupEvloopRequestResp.output:type_name -> evloopio.UniversalGroupEvloopOutput
	10, // 2: imchatevloop.DoMigrateResp.relations:type_name -> imchatevloop.DoMigrateResp.RelationsEntry
	11, // 3: imchatevloop.DoMigrateResp.subscribers:type_name -> imchatevloop.DoMigrateResp.SubscribersEntry
	15, // 4: imchatevloop.DoMigrateResp.RelationInfo.relation:type_name -> base.RelationEntry
	12, // 5: imchatevloop.DoMigrateResp.UserSubscribeEntry.entries:type_name -> imchatevloop.DoMigrateResp.UserSubscribeEntry.SubscribeEntry
	8,  // 6: imchatevloop.DoMigrateResp.RelationsEntry.value:type_name -> imchatevloop.DoMigrateResp.RelationInfo
	9,  // 7: imchatevloop.DoMigrateResp.SubscribersEntry.value:type_name -> imchatevloop.DoMigrateResp.UserSubscribeEntry
	16, // 8: imchatevloop.DoMigrateResp.UserSubscribeEntry.SubscribeEntry.session:type_name -> base.SessionEntry
	0,  // 9: imchatevloop.IMChatEvloop.CreateGroup:input_type -> imchatevloop.CreateGroupRequest
	2,  // 10: imchatevloop.IMChatEvloop.UniversalGroupEvloopRequest:input_type -> imchatevloop.UniversalGroupEvloopRequestReq
	4,  // 11: imchatevloop.IMChatEvloop.DoMigrate:input_type -> imchatevloop.DoMigrateReq
	6,  // 12: imchatevloop.IMChatEvloop.MigrateDone:input_type -> imchatevloop.MigrateDoneReq
	1,  // 13: imchatevloop.IMChatEvloop.CreateGroup:output_type -> imchatevloop.CreateGroupResponse
	3,  // 14: imchatevloop.IMChatEvloop.UniversalGroupEvloopRequest:output_type -> imchatevloop.UniversalGroupEvloopRequestResp
	5,  // 15: imchatevloop.IMChatEvloop.DoMigrate:output_type -> imchatevloop.DoMigrateResp
	7,  // 16: imchatevloop.IMChatEvloop.MigrateDone:output_type -> imchatevloop.MigrateDoneResp
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_service_im_chat_evloop_proto_init() }
func file_service_im_chat_evloop_proto_init() {
	if File_service_im_chat_evloop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_im_chat_evloop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupRequest); i {
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
		file_service_im_chat_evloop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupResponse); i {
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
		file_service_im_chat_evloop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniversalGroupEvloopRequestReq); i {
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
		file_service_im_chat_evloop_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniversalGroupEvloopRequestResp); i {
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
		file_service_im_chat_evloop_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoMigrateReq); i {
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
		file_service_im_chat_evloop_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoMigrateResp); i {
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
		file_service_im_chat_evloop_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrateDoneReq); i {
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
		file_service_im_chat_evloop_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrateDoneResp); i {
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
		file_service_im_chat_evloop_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoMigrateResp_RelationInfo); i {
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
		file_service_im_chat_evloop_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoMigrateResp_UserSubscribeEntry); i {
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
		file_service_im_chat_evloop_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoMigrateResp_UserSubscribeEntry_SubscribeEntry); i {
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
			RawDescriptor: file_service_im_chat_evloop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_im_chat_evloop_proto_goTypes,
		DependencyIndexes: file_service_im_chat_evloop_proto_depIdxs,
		MessageInfos:      file_service_im_chat_evloop_proto_msgTypes,
	}.Build()
	File_service_im_chat_evloop_proto = out.File
	file_service_im_chat_evloop_proto_rawDesc = nil
	file_service_im_chat_evloop_proto_goTypes = nil
	file_service_im_chat_evloop_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.1. DO NOT EDIT.

type IMChatEvloop interface {
	CreateGroup(ctx context.Context, req *CreateGroupRequest) (res *CreateGroupResponse, err error)
	UniversalGroupEvloopRequest(ctx context.Context, req *UniversalGroupEvloopRequestReq) (res *UniversalGroupEvloopRequestResp, err error)
	DoMigrate(ctx context.Context, req *DoMigrateReq) (res *DoMigrateResp, err error)
	MigrateDone(ctx context.Context, req *MigrateDoneReq) (res *MigrateDoneResp, err error)
}
