package rpcserver

import (
	"context"
	distributelock "pigeon/common/distribute_lock"
	"pigeon/kitex_gen/service/imauthroute"

	"gorm.io/gorm"
)

type RPCContext struct {
	DB   *gorm.DB
	Lock *distributelock.DisLockClient
}

type RPCServer struct {
	RPCContext
}

func (server *RPCServer) Login(ctx context.Context, req *imauthroute.LoginReq) (res *imauthroute.LoginResp, err error) {
	return nil, nil
}
func (server *RPCServer) Logout(ctx context.Context, req *imauthroute.LogoutReq) (res *imauthroute.LogoutResp, err error) {
	return nil, nil
}
func (server *RPCServer) ForceOffline(ctx context.Context, req *imauthroute.ForceOfflineReq) (res *imauthroute.ForceOfflineResp, err error) {
	return nil, nil
}
func (server *RPCServer) QuerySessionRoute(ctx context.Context, req *imauthroute.QuerySessionRouteReq) (res *imauthroute.QuerySessionRouteResp, err error) {
	return nil, nil
}
func (server *RPCServer) QueryUserRoute(ctx context.Context, req *imauthroute.QueryUserRouteReq) (res *imauthroute.QueryUserRouteResp, err error) {
	return nil, nil
}
