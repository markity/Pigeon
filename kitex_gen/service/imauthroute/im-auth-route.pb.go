// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: idl/service/im-auth-route.proto

package imauthroute

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

type SessionEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginAt           int64  `protobuf:"varint,1,opt,name=LoginAt,proto3" json:"LoginAt,omitempty"`
	Username          string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	SessionId         string `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	DeviceDesc        string `protobuf:"bytes,4,opt,name=device_desc,json=deviceDesc,proto3" json:"device_desc,omitempty"`
	AdvertiseAddrPort string `protobuf:"bytes,5,opt,name=advertiseAddrPort,proto3" json:"advertiseAddrPort,omitempty"`
}

func (x *SessionEntry) Reset() {
	*x = SessionEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_service_im_auth_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionEntry) ProtoMessage() {}

func (x *SessionEntry) ProtoReflect() protoreflect.Message {
	mi := &file_idl_service_im_auth_route_proto_msgTypes[0]
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
	return file_idl_service_im_auth_route_proto_rawDescGZIP(), []int{0}
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

func (x *SessionEntry) GetAdvertiseAddrPort() string {
	if x != nil {
		return x.AdvertiseAddrPort
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// 多设备管理, 用来描述设备信息的字符串, 业务方自行上传
	DeviceDesc string `protobuf:"bytes,3,opt,name=device_desc,json=deviceDesc,proto3" json:"device_desc,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_service_im_auth_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_service_im_auth_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_idl_service_im_auth_route_proto_rawDescGZIP(), []int{1}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginReq) GetDeviceDesc() string {
	if x != nil {
		return x.DeviceDesc
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// 仅当success==true时才有意义
	SessionId string `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// 用来做多设备管理的, 通过version避免乱序
	Version int64 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	// 会包括自身
	Sessions []*SessionEntry `protobuf:"bytes,4,rep,name=sessions,proto3" json:"sessions,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_service_im_auth_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_service_im_auth_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_idl_service_im_auth_route_proto_rawDescGZIP(), []int{2}
}

func (x *LoginResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LoginResp) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *LoginResp) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *LoginResp) GetSessions() []*SessionEntry {
	if x != nil {
		return x.Sessions
	}
	return nil
}

type LogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *LogoutReq) Reset() {
	*x = LogoutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_service_im_auth_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReq) ProtoMessage() {}

func (x *LogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_service_im_auth_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReq.ProtoReflect.Descriptor instead.
func (*LogoutReq) Descriptor() ([]byte, []int) {
	return file_idl_service_im_auth_route_proto_rawDescGZIP(), []int{3}
}

func (x *LogoutReq) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type LogoutResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 如果不存在这个登录项, 那么就false
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *LogoutResp) Reset() {
	*x = LogoutResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_service_im_auth_route_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResp) ProtoMessage() {}

func (x *LogoutResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_service_im_auth_route_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResp.ProtoReflect.Descriptor instead.
func (*LogoutResp) Descriptor() ([]byte, []int) {
	return file_idl_service_im_auth_route_proto_rawDescGZIP(), []int{4}
}

func (x *LogoutResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 踢人请求
type ForceOfflineReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SelfSessionId   string `protobuf:"bytes,1,opt,name=self_session_id,json=selfSessionId,proto3" json:"self_session_id,omitempty"`
	RemoteSessionId string `protobuf:"bytes,2,opt,name=remote_session_id,json=remoteSessionId,proto3" json:"remote_session_id,omitempty"`
}

func (x *ForceOfflineReq) Reset() {
	*x = ForceOfflineReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_service_im_auth_route_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceOfflineReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceOfflineReq) ProtoMessage() {}

func (x *ForceOfflineReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_service_im_auth_route_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceOfflineReq.ProtoReflect.Descriptor instead.
func (*ForceOfflineReq) Descriptor() ([]byte, []int) {
	return file_idl_service_im_auth_route_proto_rawDescGZIP(), []int{5}
}

func (x *ForceOfflineReq) GetSelfSessionId() string {
	if x != nil {
		return x.SelfSessionId
	}
	return ""
}

func (x *ForceOfflineReq) GetRemoteSessionId() string {
	if x != nil {
		return x.RemoteSessionId
	}
	return ""
}

type ForceOfflineResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 是否成功踢出
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// 踢出ok后会返回当前此用户的在线全量
	Version int64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	// 会包括自身
	Sessions []*SessionEntry `protobuf:"bytes,3,rep,name=sessions,proto3" json:"sessions,omitempty"`
}

func (x *ForceOfflineResp) Reset() {
	*x = ForceOfflineResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_service_im_auth_route_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForceOfflineResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForceOfflineResp) ProtoMessage() {}

func (x *ForceOfflineResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_service_im_auth_route_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForceOfflineResp.ProtoReflect.Descriptor instead.
func (*ForceOfflineResp) Descriptor() ([]byte, []int) {
	return file_idl_service_im_auth_route_proto_rawDescGZIP(), []int{6}
}

func (x *ForceOfflineResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ForceOfflineResp) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ForceOfflineResp) GetSessions() []*SessionEntry {
	if x != nil {
		return x.Sessions
	}
	return nil
}

type QuerySessionRouteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *QuerySessionRouteReq) Reset() {
	*x = QuerySessionRouteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_service_im_auth_route_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySessionRouteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySessionRouteReq) ProtoMessage() {}

func (x *QuerySessionRouteReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_service_im_auth_route_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySessionRouteReq.ProtoReflect.Descriptor instead.
func (*QuerySessionRouteReq) Descriptor() ([]byte, []int) {
	return file_idl_service_im_auth_route_proto_rawDescGZIP(), []int{7}
}

func (x *QuerySessionRouteReq) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type QuerySessionRouteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// success == true时后面的数据才有效
	Route *SessionEntry `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *QuerySessionRouteResp) Reset() {
	*x = QuerySessionRouteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_service_im_auth_route_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySessionRouteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySessionRouteResp) ProtoMessage() {}

func (x *QuerySessionRouteResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_service_im_auth_route_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySessionRouteResp.ProtoReflect.Descriptor instead.
func (*QuerySessionRouteResp) Descriptor() ([]byte, []int) {
	return file_idl_service_im_auth_route_proto_rawDescGZIP(), []int{8}
}

func (x *QuerySessionRouteResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *QuerySessionRouteResp) GetRoute() *SessionEntry {
	if x != nil {
		return x.Route
	}
	return nil
}

type QueryUserRouteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *QueryUserRouteReq) Reset() {
	*x = QueryUserRouteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_service_im_auth_route_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserRouteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserRouteReq) ProtoMessage() {}

func (x *QueryUserRouteReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_service_im_auth_route_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserRouteReq.ProtoReflect.Descriptor instead.
func (*QueryUserRouteReq) Descriptor() ([]byte, []int) {
	return file_idl_service_im_auth_route_proto_rawDescGZIP(), []int{9}
}

func (x *QueryUserRouteReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type QueryUserRouteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool            `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Routes  []*SessionEntry `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *QueryUserRouteResp) Reset() {
	*x = QueryUserRouteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_service_im_auth_route_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserRouteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserRouteResp) ProtoMessage() {}

func (x *QueryUserRouteResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_service_im_auth_route_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserRouteResp.ProtoReflect.Descriptor instead.
func (*QueryUserRouteResp) Descriptor() ([]byte, []int) {
	return file_idl_service_im_auth_route_proto_rawDescGZIP(), []int{10}
}

func (x *QueryUserRouteResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *QueryUserRouteResp) GetRoutes() []*SessionEntry {
	if x != nil {
		return x.Routes
	}
	return nil
}

var File_idl_service_im_auth_route_proto protoreflect.FileDescriptor

var file_idl_service_im_auth_route_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x64, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6d,
	0x2d, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x69, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0xb2,
	0x01, 0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x44, 0x65, 0x73, 0x63, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69,
	0x73, 0x65, 0x41, 0x64, 0x64, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x41, 0x64, 0x64, 0x72, 0x50,
	0x6f, 0x72, 0x74, 0x22, 0x63, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x22, 0x95, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6d,
	0x61, 0x75, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x2a, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x0a,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x65, 0x0a, 0x0f, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x4f, 0x66, 0x66,
	0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x65, 0x6c, 0x66, 0x5f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x65, 0x6c, 0x66, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x2a, 0x0a, 0x11, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x7d, 0x0a, 0x10, 0x46,
	0x6f, 0x72, 0x63, 0x65, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x35, 0x0a, 0x14, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x62, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x2f, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x2f, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x61, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x32, 0xfc, 0x02, 0x0a, 0x0b, 0x49, 0x4d,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x15, 0x2e, 0x69, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x69, 0x6d, 0x61, 0x75,
	0x74, 0x68, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x39, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x2e, 0x69, 0x6d,
	0x61, 0x75, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x69, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4b, 0x0a, 0x0c,
	0x46, 0x6f, 0x72, 0x63, 0x65, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1c, 0x2e, 0x69,
	0x6d, 0x61, 0x75, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x63, 0x65,
	0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x69, 0x6d, 0x61,
	0x75, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x4f, 0x66,
	0x66, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x5a, 0x0a, 0x11, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x21,
	0x2e, 0x69, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x22, 0x2e, 0x69, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x51, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x69, 0x6d, 0x61, 0x75, 0x74, 0x68,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x69, 0x6d, 0x61, 0x75, 0x74, 0x68,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x26, 0x5a, 0x24, 0x70, 0x69, 0x67, 0x65,
	0x6f, 0x6e, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_service_im_auth_route_proto_rawDescOnce sync.Once
	file_idl_service_im_auth_route_proto_rawDescData = file_idl_service_im_auth_route_proto_rawDesc
)

func file_idl_service_im_auth_route_proto_rawDescGZIP() []byte {
	file_idl_service_im_auth_route_proto_rawDescOnce.Do(func() {
		file_idl_service_im_auth_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_service_im_auth_route_proto_rawDescData)
	})
	return file_idl_service_im_auth_route_proto_rawDescData
}

var file_idl_service_im_auth_route_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_idl_service_im_auth_route_proto_goTypes = []interface{}{
	(*SessionEntry)(nil),          // 0: imauthroute.SessionEntry
	(*LoginReq)(nil),              // 1: imauthroute.LoginReq
	(*LoginResp)(nil),             // 2: imauthroute.LoginResp
	(*LogoutReq)(nil),             // 3: imauthroute.LogoutReq
	(*LogoutResp)(nil),            // 4: imauthroute.LogoutResp
	(*ForceOfflineReq)(nil),       // 5: imauthroute.ForceOfflineReq
	(*ForceOfflineResp)(nil),      // 6: imauthroute.ForceOfflineResp
	(*QuerySessionRouteReq)(nil),  // 7: imauthroute.QuerySessionRouteReq
	(*QuerySessionRouteResp)(nil), // 8: imauthroute.QuerySessionRouteResp
	(*QueryUserRouteReq)(nil),     // 9: imauthroute.QueryUserRouteReq
	(*QueryUserRouteResp)(nil),    // 10: imauthroute.QueryUserRouteResp
}
var file_idl_service_im_auth_route_proto_depIdxs = []int32{
	0,  // 0: imauthroute.LoginResp.sessions:type_name -> imauthroute.SessionEntry
	0,  // 1: imauthroute.ForceOfflineResp.sessions:type_name -> imauthroute.SessionEntry
	0,  // 2: imauthroute.QuerySessionRouteResp.route:type_name -> imauthroute.SessionEntry
	0,  // 3: imauthroute.QueryUserRouteResp.routes:type_name -> imauthroute.SessionEntry
	1,  // 4: imauthroute.IMAuthRoute.Login:input_type -> imauthroute.LoginReq
	3,  // 5: imauthroute.IMAuthRoute.Logout:input_type -> imauthroute.LogoutReq
	5,  // 6: imauthroute.IMAuthRoute.ForceOffline:input_type -> imauthroute.ForceOfflineReq
	7,  // 7: imauthroute.IMAuthRoute.QuerySessionRoute:input_type -> imauthroute.QuerySessionRouteReq
	9,  // 8: imauthroute.IMAuthRoute.QueryUserRoute:input_type -> imauthroute.QueryUserRouteReq
	2,  // 9: imauthroute.IMAuthRoute.Login:output_type -> imauthroute.LoginResp
	4,  // 10: imauthroute.IMAuthRoute.Logout:output_type -> imauthroute.LogoutResp
	6,  // 11: imauthroute.IMAuthRoute.ForceOffline:output_type -> imauthroute.ForceOfflineResp
	8,  // 12: imauthroute.IMAuthRoute.QuerySessionRoute:output_type -> imauthroute.QuerySessionRouteResp
	10, // 13: imauthroute.IMAuthRoute.QueryUserRoute:output_type -> imauthroute.QueryUserRouteResp
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_idl_service_im_auth_route_proto_init() }
func file_idl_service_im_auth_route_proto_init() {
	if File_idl_service_im_auth_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_service_im_auth_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_idl_service_im_auth_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_idl_service_im_auth_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_idl_service_im_auth_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutReq); i {
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
		file_idl_service_im_auth_route_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutResp); i {
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
		file_idl_service_im_auth_route_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForceOfflineReq); i {
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
		file_idl_service_im_auth_route_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForceOfflineResp); i {
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
		file_idl_service_im_auth_route_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySessionRouteReq); i {
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
		file_idl_service_im_auth_route_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySessionRouteResp); i {
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
		file_idl_service_im_auth_route_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserRouteReq); i {
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
		file_idl_service_im_auth_route_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserRouteResp); i {
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
			RawDescriptor: file_idl_service_im_auth_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_service_im_auth_route_proto_goTypes,
		DependencyIndexes: file_idl_service_im_auth_route_proto_depIdxs,
		MessageInfos:      file_idl_service_im_auth_route_proto_msgTypes,
	}.Build()
	File_idl_service_im_auth_route_proto = out.File
	file_idl_service_im_auth_route_proto_rawDesc = nil
	file_idl_service_im_auth_route_proto_goTypes = nil
	file_idl_service_im_auth_route_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.9.1. DO NOT EDIT.

type IMAuthRoute interface {
	Login(ctx context.Context, req *LoginReq) (res *LoginResp, err error)
	Logout(ctx context.Context, req *LogoutReq) (res *LogoutResp, err error)
	ForceOffline(ctx context.Context, req *ForceOfflineReq) (res *ForceOfflineResp, err error)
	QuerySessionRoute(ctx context.Context, req *QuerySessionRouteReq) (res *QuerySessionRouteResp, err error)
	QueryUserRoute(ctx context.Context, req *QueryUserRouteReq) (res *QueryUserRouteResp, err error)
}
