package rpcserver

import (
	"context"
	"encoding/json"
	"sync"

	"pigeon/common/protocol"
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
		return &imgateway.PushMessageResp{}, nil
	}

	var v map[string]interface{}
	if err := json.Unmarshal(req.Data, &v); err != nil {
		return nil, err
	}

	loop := loop_.(eventloop.EventLoop)
	data := protocol.PackData(protocol.MustEncodePacket(&protocol.S2CPushMessagePacket{
		Data:     v,
		PushType: req.PushType,
	}, req.EchoCode))
	loop.RunInLoop(func() {
		tcpserver.PushMessage(loop, req.SessionId, data)
	})

	return &imgateway.PushMessageResp{}, nil
}

// 定位event loop, 并且打入事件循环就ok
func (server *RPCServer) OtherDeviceKick(ctx context.Context, req *imgateway.OtherDeviceKickReq) (
	res *imgateway.OtherDeviceKickResp, err error) {
	loop_, ok := server.EvloopRoute.Load(req.ToSession)
	if !ok {
		return &imgateway.OtherDeviceKickResp{}, nil
	}

	loop := loop_.(eventloop.EventLoop)

	data := protocol.PackData(protocol.MustEncodePacket(&protocol.S2COtherDeviceKickNotifyPacket{
		FromSessionId:   req.FromSession,
		FromSessionDesc: req.FromSessionDesc,
	}))
	loop.RunInLoop(func() {
		tcpserver.OtherDeveiceKick(loop, req.ToSession, data)
	})

	return &imgateway.OtherDeviceKickResp{}, nil
}

func (server *RPCContext) BroadcastDeviceInfo(ctx context.Context, req *imgateway.BroadcastDeviceInfoReq) (
	res *imgateway.BroadcastDeviceInfoResp, err error) {

	loop_, ok := server.EvloopRoute.Load(req.SessionId)
	if !ok {
		return &imgateway.BroadcastDeviceInfoResp{}, nil
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
		Version:  req.Version,
		Sessions: devs,
	}))
	loop.RunInLoop(func() {
		tcpserver.BroadcastDeviceInfo(loop, req.SessionId, data)
	})
	return &imgateway.BroadcastDeviceInfoResp{}, nil
}
