package rpcserver

import (
	"context"
	"encoding/json"
	"sync"

	"pigeon/im-gateway/protocol"
	"pigeon/im-gateway/tcpserver"
	"pigeon/kitex_gen/service/imgateway"

	eventloop "github.com/markity/go-reactor/pkg/event_loop"
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
	loop_, ok := server.EvloopRoute.Load(req.SessionId)
	if !ok {
		return &imgateway.PushMessageResp{
			Success: false,
		}, err
	}

	var v map[string]interface{}
	if err := json.Unmarshal(req.Data, &v); err != nil {
		return nil, err
	}

	data := protocol.PackData(protocol.MustEncodePacket(&protocol.S2CPushMessagePacket{
		Data: v,
	}))

	okChan := make(chan bool, 1)
	loop := loop_.(eventloop.EventLoop)
	loop.RunInLoop(func() {
		tcpserver.PushMessage(loop, req.SessionId, data, okChan)
	})
	<-okChan
	return &imgateway.PushMessageResp{
		Success: <-okChan,
	}, nil
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
