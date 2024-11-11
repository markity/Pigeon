package rpcserver

import (
	"pigeon/im-relation/bizpush"
	"pigeon/im-relation/db"
	"pigeon/im-relation/rds"
	"pigeon/kitex_gen/service/imrelay/imrelay"
)

type RPCContext struct {
	DB       *db.DB
	RelayCli imrelay.Client
	RdsAct   *rds.RdsAction
	BPush    *bizpush.BizPusher
}

type RPCServer struct {
	RPCContext
}
