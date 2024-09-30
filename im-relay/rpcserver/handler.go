package rpcserver

import (
	"context"
	"log"

	"pigeon/im-relay/api"
	"pigeon/kitex_gen/service/imgateway"
	"pigeon/kitex_gen/service/imrelay"
)

type RPCContext struct {
}

type RPCServer struct {
	RPCContext
}

func (*RPCServer) BizMessage(ctx context.Context,
	req *imrelay.BizMessageReq) (res *imrelay.BizMessageResp, err error) {
	// 请求直接异步出去
	go func() {
		handleBizMessage(req)
	}()
	return &imrelay.BizMessageResp{}, nil
}

func handleBizMessage(req *imrelay.BizMessageReq) {
	if req.Biz == "echo" {
		cli := api.NewGatewayClientFromAdAddr(req.Session.GwAdvertiseAddrPort)
		_, err := cli.PushMessage(context.Background(), &imgateway.PushMessageReq{
			SessionId: req.Session.SessionId,
			PushType:  "push-echo",
			EchoCode:  req.EchoCode,
			Data:      req.Data,
		})
		if err != nil {
			log.Printf("failed to push echo message: %v\n", err)
		}
	}
}
