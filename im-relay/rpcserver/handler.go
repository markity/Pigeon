package rpcserver

import (
	"context"
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
		// todo
	}()
	return &imrelay.BizMessageResp{}, nil
}
