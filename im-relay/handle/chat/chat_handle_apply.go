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

func handleHandleApply(ctx *handle.HandleContext, req *imrelay.BizMessageReq) {
	var handleApplyReq bizprotocol.BizHandleApply
	err := json.Unmarshal(req.Data, &handleApplyReq)
	if err != nil {
		log.Printf("failed to unmarshal handle apply request, err: %v, data: %v\n", err, string(req.Data))
		return
	}
	_, err = ctx.RelationCli.HandleApply(context.Background(), &imrelation.HandleApplyReq{
		Session:  req.Session,
		EchoCode: req.EchoCode,
		UserId:   handleApplyReq.UserId,
		GroupId:  handleApplyReq.GroupId,
		Accept:   handleApplyReq.Accept,
	})
	if err != nil {
		log.Printf("failed to call handle apply group rpc: %v\n", err)
		return
	}
}
