package chat

import (
	"context"
	"encoding/json"
	"log"

	bizprotocol "pigeon/common/biz_protocol"
	"pigeon/im-relay/handle"
	"pigeon/kitex_gen/service/imrelation"
	"pigeon/kitex_gen/service/imrelay"
)

func handlePullRelations(ctx *handle.HandleContext, req *imrelay.BizMessageReq) {
	var pullRelationsReq bizprotocol.BizPullRelations
	err := json.Unmarshal(req.Data, &pullRelationsReq)
	if err != nil {
		log.Printf("failed to unmarshal pull relation request, err: %v, data: %v\n", err, string(req.Data))
		return
	}

	_, err = ctx.RelationCli.FetchAllRelations(context.Background(), &imrelation.FetchAllRelationsReq{
		Session:  req.Session,
		EchoCode: req.EchoCode,
	})
	if err != nil {
		log.Printf("failed to call fetch all relations group rpc: %v\n", err)
	}
}
