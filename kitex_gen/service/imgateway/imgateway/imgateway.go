// Code generated by Kitex v0.9.1. DO NOT EDIT.

package imgateway

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	imgateway "pigeon/kitex_gen/service/imgateway"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"PushMessage": kitex.NewMethodInfo(
		pushMessageHandler,
		newPushMessageArgs,
		newPushMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	iMGatewayServiceInfo                = NewServiceInfo()
	iMGatewayServiceInfoForClient       = NewServiceInfoForClient()
	iMGatewayServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return iMGatewayServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return iMGatewayServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return iMGatewayServiceInfoForClient
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
	serviceName := "IMGateway"
	handlerType := (*imgateway.IMGateway)(nil)
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
		"PackageName": "imgateway",
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

func pushMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(imgateway.PushMessageReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(imgateway.IMGateway).PushMessage(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *PushMessageArgs:
		success, err := handler.(imgateway.IMGateway).PushMessage(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PushMessageResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newPushMessageArgs() interface{} {
	return &PushMessageArgs{}
}

func newPushMessageResult() interface{} {
	return &PushMessageResult{}
}

type PushMessageArgs struct {
	Req *imgateway.PushMessageReq
}

func (p *PushMessageArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(imgateway.PushMessageReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PushMessageArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PushMessageArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PushMessageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *PushMessageArgs) Unmarshal(in []byte) error {
	msg := new(imgateway.PushMessageReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PushMessageArgs_Req_DEFAULT *imgateway.PushMessageReq

func (p *PushMessageArgs) GetReq() *imgateway.PushMessageReq {
	if !p.IsSetReq() {
		return PushMessageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PushMessageArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PushMessageArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PushMessageResult struct {
	Success *imgateway.PushMessageResp
}

var PushMessageResult_Success_DEFAULT *imgateway.PushMessageResp

func (p *PushMessageResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(imgateway.PushMessageResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PushMessageResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PushMessageResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PushMessageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *PushMessageResult) Unmarshal(in []byte) error {
	msg := new(imgateway.PushMessageResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PushMessageResult) GetSuccess() *imgateway.PushMessageResp {
	if !p.IsSetSuccess() {
		return PushMessageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PushMessageResult) SetSuccess(x interface{}) {
	p.Success = x.(*imgateway.PushMessageResp)
}

func (p *PushMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PushMessageResult) GetResult() interface{} {
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

func (p *kClient) PushMessage(ctx context.Context, Req *imgateway.PushMessageReq) (r *imgateway.PushMessageResp, err error) {
	var _args PushMessageArgs
	_args.Req = Req
	var _result PushMessageResult
	if err = p.c.Call(ctx, "PushMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
