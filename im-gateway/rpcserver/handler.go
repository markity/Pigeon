package rpcserver

import (
	"context"
	"pigeon/kitex_gen/service/imgateway"
)

type RPCContext struct {
}

type RPCServer struct {
	RPCContext
}

func PushMessage(ctx context.Context, req *imgateway.PushMessageReq) (res *imgateway.
	PushMessageResp, err error) {
	return &imgateway.PushMessageResp{}, nil
}
