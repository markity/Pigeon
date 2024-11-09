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

func handleSendApply(ctx *handle.HandleContext, req *imrelay.BizMessageReq) {
	var sendMsgReq bizprotocol.BizSendApply
	err := json.Unmarshal(req.Data, &sendMsgReq)
	if err != nil {
		log.Printf("failed to unmarshal send apply request, err: %v, data: %v\n", err, string(req.Data))
		return
	}

	err = json.Unmarshal(req.Data, &sendApplyReq)
	if err != nil {
		log.Printf("failed to unmarshal pull relation request, err: %v, data: %v\n", err, string(req.Data))
		return
	}

	_, err = ctx.RelationCli.ApplyGroup(context.Background(), &imrelation.ApplyGroupReq{
		Session:  req.Session,
		EchoCode: req.EchoCode,
		GroupId:  sendApplyReq.GroupId,
		ApplyMsg: sendApplyReq.ApplyMsg,
	})
	if err != nil {
		log.Printf("failed to call apply group rpc: %v\n", err)
		return
	}
}
