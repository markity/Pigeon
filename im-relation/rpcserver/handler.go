package rpcserver

import (
	"pigeon/im-relation/db"
	"pigeon/kitex_gen/service/imauthroute/imauthroute"
	"pigeon/kitex_gen/service/imrelay/imrelay"
)

type RPCContext struct {
	DB           *db.DB
	RelayCli     imrelay.Client
	AuthRouteCli imauthroute.Client
}

type RPCServer struct {
	RPCContext
}
