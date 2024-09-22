package rpcserver

import (
	"context"
	"pigeon/kitex_gen/service/imgateway"
	"sync"
)

type RPCContext struct {
	// 保存sessionId(string)->eventloop
	evloopRoute *sync.Map
}

type RPCServer struct {
	RPCContext
}

func PushMessage(ctx context.Context, req *imgateway.PushMessageReq) (res *imgateway.
	PushMessageResp, err error) {
	return &imgateway.PushMessageResp{}, nil
}
