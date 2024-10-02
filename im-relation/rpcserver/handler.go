package rpcserver

import (
	"pigeon/im-auth-route/db"
	"pigeon/im-auth-route/rds"
)

type RPCContext struct {
	DB  *db.DB
	Rds *rds.RdsAction
}

type RPCServer struct {
	RPCContext
}
