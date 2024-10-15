package rpcserver

import (
	"context"
	"pigeon/kitex_gen/service/imchatevloop"
)

type RPCServer struct {
}

func (*RPCServer) CreateGroup(ctx context.Context, req *imchatevloop.CreateGroupRequest) (
	res *imchatevloop.CreateGroupResponse, err error) {
	return
}
func (*RPCServer) UniversalGroupEvloopRequest(ctx context.Context,
	req *imchatevloop.UniversalGroupEvloopRequestReq) (res *imchatevloop.UniversalGroupEvloopRequestResp, err error) {
	return
}
