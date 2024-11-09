package rpcserver

import (
	"pigeon/im-relation/db"
	"pigeon/im-relation/rds"
	"pigeon/kitex_gen/service/imauthroute/imauthroute"
	"pigeon/kitex_gen/service/imrelay/imrelay"
)

type RPCContext struct {
	DB           *db.DB
	RelayCli     imrelay.Client
	AuthRouteCli imauthroute.Client
	RdsAct       *rds.RdsAction
}

type RPCServer struct {
	RPCContext
}
