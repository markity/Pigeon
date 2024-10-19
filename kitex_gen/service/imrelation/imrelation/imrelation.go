// Code generated by Kitex v0.9.1. DO NOT EDIT.

package imrelation

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	imrelation "pigeon/kitex_gen/service/imrelation"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateGroup": kitex.NewMethodInfo(
		createGroupHandler,
		newCreateGroupArgs,
		newCreateGroupResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetGroupInfo": kitex.NewMethodInfo(
		getGroupInfoHandler,
		newGetGroupInfoArgs,
		newGetGroupInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"FetchAllRelations": kitex.NewMethodInfo(
		fetchAllRelationsHandler,
		newFetchAllRelationsArgs,
		newFetchAllRelationsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"FetchAllApplications": kitex.NewMethodInfo(
		fetchAllApplicationsHandler,
		newFetchAllApplicationsArgs,
		newFetchAllApplicationsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ApplyGroup": kitex.NewMethodInfo(
		applyGroupHandler,
		newApplyGroupArgs,
		newApplyGroupResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"HandleApply": kitex.NewMethodInfo(
		handleApplyHandler,
		newHandleApplyArgs,
		newHandleApplyResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"QuitGroup": kitex.NewMethodInfo(
		quitGroupHandler,
		newQuitGroupArgs,
		newQuitGroupResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	iMRelationServiceInfo                = NewServiceInfo()
	iMRelationServiceInfoForClient       = NewServiceInfoForClient()
	iMRelationServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return iMRelationServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return iMRelationServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return iMRelationServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "IMRelation"
	handlerType := (*imrelation.IMRelation)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "imrelation",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func createGroupHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(imrelation.CreateGroupReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(imrelation.IMRelation).CreateGroup(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CreateGroupArgs:
		success, err := handler.(imrelation.IMRelation).CreateGroup(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateGroupResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCreateGroupArgs() interface{} {
	return &CreateGroupArgs{}
}

func newCreateGroupResult() interface{} {
	return &CreateGroupResult{}
}

type CreateGroupArgs struct {
	Req *imrelation.CreateGroupReq
}

func (p *CreateGroupArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(imrelation.CreateGroupReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateGroupArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateGroupArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateGroupArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateGroupArgs) Unmarshal(in []byte) error {
	msg := new(imrelation.CreateGroupReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateGroupArgs_Req_DEFAULT *imrelation.CreateGroupReq

func (p *CreateGroupArgs) GetReq() *imrelation.CreateGroupReq {
	if !p.IsSetReq() {
		return CreateGroupArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateGroupArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateGroupArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateGroupResult struct {
	Success *imrelation.CreateGroupResp
}

var CreateGroupResult_Success_DEFAULT *imrelation.CreateGroupResp

func (p *CreateGroupResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(imrelation.CreateGroupResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateGroupResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateGroupResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateGroupResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateGroupResult) Unmarshal(in []byte) error {
	msg := new(imrelation.CreateGroupResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateGroupResult) GetSuccess() *imrelation.CreateGroupResp {
	if !p.IsSetSuccess() {
		return CreateGroupResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateGroupResult) SetSuccess(x interface{}) {
	p.Success = x.(*imrelation.CreateGroupResp)
}

func (p *CreateGroupResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateGroupResult) GetResult() interface{} {
	return p.Success
}

func getGroupInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(imrelation.GetGroupInfoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(imrelation.IMRelation).GetGroupInfo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetGroupInfoArgs:
		success, err := handler.(imrelation.IMRelation).GetGroupInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetGroupInfoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetGroupInfoArgs() interface{} {
	return &GetGroupInfoArgs{}
}

func newGetGroupInfoResult() interface{} {
	return &GetGroupInfoResult{}
}

type GetGroupInfoArgs struct {
	Req *imrelation.GetGroupInfoReq
}

func (p *GetGroupInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(imrelation.GetGroupInfoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetGroupInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetGroupInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetGroupInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetGroupInfoArgs) Unmarshal(in []byte) error {
	msg := new(imrelation.GetGroupInfoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetGroupInfoArgs_Req_DEFAULT *imrelation.GetGroupInfoReq

func (p *GetGroupInfoArgs) GetReq() *imrelation.GetGroupInfoReq {
	if !p.IsSetReq() {
		return GetGroupInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetGroupInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetGroupInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetGroupInfoResult struct {
	Success *imrelation.GetGroupInfoResp
}

var GetGroupInfoResult_Success_DEFAULT *imrelation.GetGroupInfoResp

func (p *GetGroupInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(imrelation.GetGroupInfoResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetGroupInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetGroupInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetGroupInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetGroupInfoResult) Unmarshal(in []byte) error {
	msg := new(imrelation.GetGroupInfoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetGroupInfoResult) GetSuccess() *imrelation.GetGroupInfoResp {
	if !p.IsSetSuccess() {
		return GetGroupInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetGroupInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*imrelation.GetGroupInfoResp)
}

func (p *GetGroupInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetGroupInfoResult) GetResult() interface{} {
	return p.Success
}

func fetchAllRelationsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(imrelation.FetchAllRelationsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(imrelation.IMRelation).FetchAllRelations(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *FetchAllRelationsArgs:
		success, err := handler.(imrelation.IMRelation).FetchAllRelations(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FetchAllRelationsResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newFetchAllRelationsArgs() interface{} {
	return &FetchAllRelationsArgs{}
}

func newFetchAllRelationsResult() interface{} {
	return &FetchAllRelationsResult{}
}

type FetchAllRelationsArgs struct {
	Req *imrelation.FetchAllRelationsReq
}

func (p *FetchAllRelationsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(imrelation.FetchAllRelationsReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FetchAllRelationsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FetchAllRelationsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FetchAllRelationsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *FetchAllRelationsArgs) Unmarshal(in []byte) error {
	msg := new(imrelation.FetchAllRelationsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FetchAllRelationsArgs_Req_DEFAULT *imrelation.FetchAllRelationsReq

func (p *FetchAllRelationsArgs) GetReq() *imrelation.FetchAllRelationsReq {
	if !p.IsSetReq() {
		return FetchAllRelationsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FetchAllRelationsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FetchAllRelationsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FetchAllRelationsResult struct {
	Success *imrelation.FetchAllRelationsResp
}

var FetchAllRelationsResult_Success_DEFAULT *imrelation.FetchAllRelationsResp

func (p *FetchAllRelationsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(imrelation.FetchAllRelationsResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FetchAllRelationsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FetchAllRelationsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FetchAllRelationsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *FetchAllRelationsResult) Unmarshal(in []byte) error {
	msg := new(imrelation.FetchAllRelationsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FetchAllRelationsResult) GetSuccess() *imrelation.FetchAllRelationsResp {
	if !p.IsSetSuccess() {
		return FetchAllRelationsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FetchAllRelationsResult) SetSuccess(x interface{}) {
	p.Success = x.(*imrelation.FetchAllRelationsResp)
}

func (p *FetchAllRelationsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FetchAllRelationsResult) GetResult() interface{} {
	return p.Success
}

func fetchAllApplicationsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(imrelation.FetchAllApplicationsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(imrelation.IMRelation).FetchAllApplications(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *FetchAllApplicationsArgs:
		success, err := handler.(imrelation.IMRelation).FetchAllApplications(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FetchAllApplicationsResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newFetchAllApplicationsArgs() interface{} {
	return &FetchAllApplicationsArgs{}
}

func newFetchAllApplicationsResult() interface{} {
	return &FetchAllApplicationsResult{}
}

type FetchAllApplicationsArgs struct {
	Req *imrelation.FetchAllApplicationsReq
}

func (p *FetchAllApplicationsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(imrelation.FetchAllApplicationsReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FetchAllApplicationsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FetchAllApplicationsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FetchAllApplicationsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *FetchAllApplicationsArgs) Unmarshal(in []byte) error {
	msg := new(imrelation.FetchAllApplicationsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FetchAllApplicationsArgs_Req_DEFAULT *imrelation.FetchAllApplicationsReq

func (p *FetchAllApplicationsArgs) GetReq() *imrelation.FetchAllApplicationsReq {
	if !p.IsSetReq() {
		return FetchAllApplicationsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FetchAllApplicationsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FetchAllApplicationsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FetchAllApplicationsResult struct {
	Success *imrelation.FetchAllApplicationsResp
}

var FetchAllApplicationsResult_Success_DEFAULT *imrelation.FetchAllApplicationsResp

func (p *FetchAllApplicationsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(imrelation.FetchAllApplicationsResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FetchAllApplicationsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FetchAllApplicationsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FetchAllApplicationsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *FetchAllApplicationsResult) Unmarshal(in []byte) error {
	msg := new(imrelation.FetchAllApplicationsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FetchAllApplicationsResult) GetSuccess() *imrelation.FetchAllApplicationsResp {
	if !p.IsSetSuccess() {
		return FetchAllApplicationsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FetchAllApplicationsResult) SetSuccess(x interface{}) {
	p.Success = x.(*imrelation.FetchAllApplicationsResp)
}

func (p *FetchAllApplicationsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FetchAllApplicationsResult) GetResult() interface{} {
	return p.Success
}

func applyGroupHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(imrelation.ApplyGroupReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(imrelation.IMRelation).ApplyGroup(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ApplyGroupArgs:
		success, err := handler.(imrelation.IMRelation).ApplyGroup(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ApplyGroupResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newApplyGroupArgs() interface{} {
	return &ApplyGroupArgs{}
}

func newApplyGroupResult() interface{} {
	return &ApplyGroupResult{}
}

type ApplyGroupArgs struct {
	Req *imrelation.ApplyGroupReq
}

func (p *ApplyGroupArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(imrelation.ApplyGroupReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ApplyGroupArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ApplyGroupArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ApplyGroupArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ApplyGroupArgs) Unmarshal(in []byte) error {
	msg := new(imrelation.ApplyGroupReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ApplyGroupArgs_Req_DEFAULT *imrelation.ApplyGroupReq

func (p *ApplyGroupArgs) GetReq() *imrelation.ApplyGroupReq {
	if !p.IsSetReq() {
		return ApplyGroupArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ApplyGroupArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ApplyGroupArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ApplyGroupResult struct {
	Success *imrelation.ApplyGroupResp
}

var ApplyGroupResult_Success_DEFAULT *imrelation.ApplyGroupResp

func (p *ApplyGroupResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(imrelation.ApplyGroupResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ApplyGroupResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ApplyGroupResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ApplyGroupResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ApplyGroupResult) Unmarshal(in []byte) error {
	msg := new(imrelation.ApplyGroupResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ApplyGroupResult) GetSuccess() *imrelation.ApplyGroupResp {
	if !p.IsSetSuccess() {
		return ApplyGroupResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ApplyGroupResult) SetSuccess(x interface{}) {
	p.Success = x.(*imrelation.ApplyGroupResp)
}

func (p *ApplyGroupResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ApplyGroupResult) GetResult() interface{} {
	return p.Success
}

func handleApplyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(imrelation.HandleApplyReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(imrelation.IMRelation).HandleApply(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *HandleApplyArgs:
		success, err := handler.(imrelation.IMRelation).HandleApply(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*HandleApplyResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newHandleApplyArgs() interface{} {
	return &HandleApplyArgs{}
}

func newHandleApplyResult() interface{} {
	return &HandleApplyResult{}
}

type HandleApplyArgs struct {
	Req *imrelation.HandleApplyReq
}

func (p *HandleApplyArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(imrelation.HandleApplyReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *HandleApplyArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *HandleApplyArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *HandleApplyArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *HandleApplyArgs) Unmarshal(in []byte) error {
	msg := new(imrelation.HandleApplyReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var HandleApplyArgs_Req_DEFAULT *imrelation.HandleApplyReq

func (p *HandleApplyArgs) GetReq() *imrelation.HandleApplyReq {
	if !p.IsSetReq() {
		return HandleApplyArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *HandleApplyArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *HandleApplyArgs) GetFirstArgument() interface{} {
	return p.Req
}

type HandleApplyResult struct {
	Success *imrelation.HandleApplyResp
}

var HandleApplyResult_Success_DEFAULT *imrelation.HandleApplyResp

func (p *HandleApplyResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(imrelation.HandleApplyResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *HandleApplyResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *HandleApplyResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *HandleApplyResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *HandleApplyResult) Unmarshal(in []byte) error {
	msg := new(imrelation.HandleApplyResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *HandleApplyResult) GetSuccess() *imrelation.HandleApplyResp {
	if !p.IsSetSuccess() {
		return HandleApplyResult_Success_DEFAULT
	}
	return p.Success
}

func (p *HandleApplyResult) SetSuccess(x interface{}) {
	p.Success = x.(*imrelation.HandleApplyResp)
}

func (p *HandleApplyResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *HandleApplyResult) GetResult() interface{} {
	return p.Success
}

func quitGroupHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(imrelation.QuitGroupReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(imrelation.IMRelation).QuitGroup(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *QuitGroupArgs:
		success, err := handler.(imrelation.IMRelation).QuitGroup(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QuitGroupResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newQuitGroupArgs() interface{} {
	return &QuitGroupArgs{}
}

func newQuitGroupResult() interface{} {
	return &QuitGroupResult{}
}

type QuitGroupArgs struct {
	Req *imrelation.QuitGroupReq
}

func (p *QuitGroupArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(imrelation.QuitGroupReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *QuitGroupArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *QuitGroupArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *QuitGroupArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *QuitGroupArgs) Unmarshal(in []byte) error {
	msg := new(imrelation.QuitGroupReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QuitGroupArgs_Req_DEFAULT *imrelation.QuitGroupReq

func (p *QuitGroupArgs) GetReq() *imrelation.QuitGroupReq {
	if !p.IsSetReq() {
		return QuitGroupArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QuitGroupArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *QuitGroupArgs) GetFirstArgument() interface{} {
	return p.Req
}

type QuitGroupResult struct {
	Success *imrelation.QuitGroupResp
}

var QuitGroupResult_Success_DEFAULT *imrelation.QuitGroupResp

func (p *QuitGroupResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(imrelation.QuitGroupResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *QuitGroupResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *QuitGroupResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *QuitGroupResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *QuitGroupResult) Unmarshal(in []byte) error {
	msg := new(imrelation.QuitGroupResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QuitGroupResult) GetSuccess() *imrelation.QuitGroupResp {
	if !p.IsSetSuccess() {
		return QuitGroupResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QuitGroupResult) SetSuccess(x interface{}) {
	p.Success = x.(*imrelation.QuitGroupResp)
}

func (p *QuitGroupResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *QuitGroupResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateGroup(ctx context.Context, Req *imrelation.CreateGroupReq) (r *imrelation.CreateGroupResp, err error) {
	var _args CreateGroupArgs
	_args.Req = Req
	var _result CreateGroupResult
	if err = p.c.Call(ctx, "CreateGroup", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetGroupInfo(ctx context.Context, Req *imrelation.GetGroupInfoReq) (r *imrelation.GetGroupInfoResp, err error) {
	var _args GetGroupInfoArgs
	_args.Req = Req
	var _result GetGroupInfoResult
	if err = p.c.Call(ctx, "GetGroupInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FetchAllRelations(ctx context.Context, Req *imrelation.FetchAllRelationsReq) (r *imrelation.FetchAllRelationsResp, err error) {
	var _args FetchAllRelationsArgs
	_args.Req = Req
	var _result FetchAllRelationsResult
	if err = p.c.Call(ctx, "FetchAllRelations", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FetchAllApplications(ctx context.Context, Req *imrelation.FetchAllApplicationsReq) (r *imrelation.FetchAllApplicationsResp, err error) {
	var _args FetchAllApplicationsArgs
	_args.Req = Req
	var _result FetchAllApplicationsResult
	if err = p.c.Call(ctx, "FetchAllApplications", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ApplyGroup(ctx context.Context, Req *imrelation.ApplyGroupReq) (r *imrelation.ApplyGroupResp, err error) {
	var _args ApplyGroupArgs
	_args.Req = Req
	var _result ApplyGroupResult
	if err = p.c.Call(ctx, "ApplyGroup", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) HandleApply(ctx context.Context, Req *imrelation.HandleApplyReq) (r *imrelation.HandleApplyResp, err error) {
	var _args HandleApplyArgs
	_args.Req = Req
	var _result HandleApplyResult
	if err = p.c.Call(ctx, "HandleApply", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QuitGroup(ctx context.Context, Req *imrelation.QuitGroupReq) (r *imrelation.QuitGroupResp, err error) {
	var _args QuitGroupArgs
	_args.Req = Req
	var _result QuitGroupResult
	if err = p.c.Call(ctx, "QuitGroup", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
