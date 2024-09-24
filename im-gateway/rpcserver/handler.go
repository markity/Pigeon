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

	loop := loop_.(eventloop.EventLoop)
	okChan := make(chan bool, 1)
	data := protocol.PackData(protocol.MustEncodePacket(&protocol.S2CPushMessagePacket{
		Data: v,
	}))
	loop.RunInLoop(func() {
		tcpserver.PushMessage(loop, req.SessionId, data, okChan)
	})

	return &imgateway.PushMessageResp{
		Success: <-okChan,
	}, nil
}

// 定位event loop, 并且打入事件循环就ok
func (server *RPCServer) OtherDeviceKick(ctx context.Context, req *imgateway.OtherDeviceKickReq) (
	res *imgateway.OtherDeviceKickResp, err error) {
	loop_, ok := server.EvloopRoute.Load(req.ToSession)
	if !ok {
		return &imgateway.OtherDeviceKickResp{
			Success: false,
		}, err
	}

	loop := loop_.(eventloop.EventLoop)

	data := protocol.PackData(protocol.PackData(protocol.MustEncodePacket(&protocol.S2COtherDeviceKickNotify{
		FromSessionId:   req.FromSession,
		FromSessionDesc: req.FromSessionDesc,
	})))
	okChan := make(chan bool, 1)
	loop.RunInLoop(func() {
		tcpserver.OtherDeveiceKick(loop, req.ToSession, data, okChan)
	})

	return &imgateway.OtherDeviceKickResp{
		Success: <-okChan,
	}, nil
}

func (server *RPCContext) BroadcastDeviceInfo(ctx context.Context, req *imgateway.BroadcastDeviceInfoReq) (
	res *imgateway.BroadcastDeviceInfoResp, err error) {

	loop_, ok := server.EvloopRoute.Load(req.SessionId)
	if !ok {
		return &imgateway.BroadcastDeviceInfoResp{
			Success: false,
		}, err
	}
	devs := make([]*protocol.DeviceSessionEntry, 0, len(req.Sessions))
	for _, v := range req.Sessions {
		devs = append(devs, &protocol.DeviceSessionEntry{
			SessionId:  v.SessionId,
			LoginAt:    v.LoginAt,
			DeviceDesc: v.DeviceDesc,
		})
	}

	loop := loop_.(eventloop.EventLoop)
	data := protocol.PackData(protocol.MustEncodePacket(&protocol.S2CDeviceInfoBroadcastPacket{
		Version: req.Version,
		Devices: devs,
	}))
	okChan := make(chan bool, 1)
	loop.RunInLoop(func() {
		tcpserver.BroadcastDeviceInfo(loop, req.SessionId, data, okChan)
	})

	return &imgateway.BroadcastDeviceInfoResp{
		Success: <-okChan,
	}, nil
}
