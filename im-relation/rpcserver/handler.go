package rpcserver

import (
	"encoding/json"

	"pigeon/im-auth-route/rds"
	"pigeon/im-relation/db"
	"pigeon/kitex_gen/service/imauthroute/imauthroute"
	"pigeon/kitex_gen/service/imrelay/imrelay"
)

type RPCContext struct {
	DB           *db.DB
	Rds          *rds.RdsAction
	RelayCli     imrelay.Client
	AuthRouteCli imauthroute.Client
}

type RPCServer struct {
	RPCContext
}

func mustMarshalToBytes(v interface{}) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return data
}
