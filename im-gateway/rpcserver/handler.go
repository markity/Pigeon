package rpcserver

import (
	"context"
	"pigeon/kitex_gen/service/imgateway"
	"sync"
)

type RPCContext struct {
	// 保存sessionId(string)->eventloop
	EvloopRoute *sync.Map
}

type RPCServer struct {
	RPCContext
}

// 定位event loop, 并且打入事件循环就ok
func (server *RPCServer) PushMessage(ctx context.Context, req *imgateway.PushMessageReq) (res *imgateway.
	PushMessageResp, err error) {
	return &imgateway.PushMessageResp{}, nil
}

// 定位event loop, 并且打入事件循环就ok
func (server *RPCServer) OtherDeviceKick(ctx context.Context, req *imgateway.OtherDeviceKickReq) (
	res *imgateway.OtherDeviceKickResp, err error) {
	return &imgateway.OtherDeviceKickResp{}, nil
}

func (server *RPCContext) BroadcastDeviceInfo(ctx context.Context, req *imgateway.BroadcastDeviceInfoReq) (
	res *imgateway.BroadcastDeviceInfoResp, err error) {
	return &imgateway.BroadcastDeviceInfoResp{}, nil
}
