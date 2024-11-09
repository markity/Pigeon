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

func handleCreateGroup(ctx *handle.HandleContext, req *imrelay.BizMessageReq) {
	var createGroupReq bizprotocol.BizCreateGroup
	err := json.Unmarshal(req.Data, &createGroupReq)
	if err != nil {
		log.Printf("failed to unmarshal create group request, err: %v, data: %v\n", err, string(req.Data))
		return
	}
	_, err = ctx.RelationCli.CreateGroup(context.Background(), &imrelation.CreateGroupReq{
		Session:  req.Session,
		EchoCode: req.EchoCode,
	})
	if err != nil {
		log.Printf("failed to call create group rpc: %v\n", err)
	}
}
