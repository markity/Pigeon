package rpcserver

import (
	"context"
	"strings"

	"pigeon/im-relay/handle"
	"pigeon/im-relay/handle/chat"
	"pigeon/im-relay/handle/echo"
	"pigeon/kitex_gen/service/imrelation/imrelation"
	"pigeon/kitex_gen/service/imrelay"
)

type RPCContext struct {
	RelationCli imrelation.Client
}

type RPCServer struct {
	RPCContext
}

func (s *RPCServer) BizMessage(ctx context.Context,
	req *imrelay.BizMessageReq) (res *imrelay.BizMessageResp, err error) {
	// 请求直接异步出去
	go func() {
		s.handleBizMessage(req)
	}()
	return &imrelay.BizMessageResp{}, nil
}

func (s *RPCServer) handleBizMessage(req *imrelay.BizMessageReq) {
	splits := strings.Split(req.Biz, "-")
	if len(splits) < 1 {
		return
	}
	switch splits[0] {
	case "echo":
		echo.HandleEcho(&handle.HandleContext{RelationCli: s.RelationCli}, req)
	case "chat":
		chat.HandleChat(&handle.HandleContext{RelationCli: s.RelationCli}, req)
	default:
	}
}
