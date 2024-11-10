package chat

import (
	"context"
	"encoding/json"
	"log"
	bizprotocol "pigeon/common/biz_protocol"
	"pigeon/im-relay/api"
	"pigeon/im-relay/handle"
	"pigeon/kitex_gen/service/evloopio"
	"pigeon/kitex_gen/service/imchatevloop"
	"pigeon/kitex_gen/service/imrelation"
	"pigeon/kitex_gen/service/imrelay"
)

func handleSendMsg(ctx *handle.HandleContext, req *imrelay.BizMessageReq) {
	var sendMsgReq bizprotocol.BizSendMessage
	err := json.Unmarshal(req.Data, &sendMsgReq)
	if err != nil {
		log.Printf("failed to unmarshal send msg request, err: %v, data: %v\n", err, string(req.Data))
		return
	}
	// TODO: 此处可以重试的, 操作是幂等的
	resp, err := ctx.RelationCli.GetGroupInfo(context.Background(), &imrelation.GetGroupInfoReq{
		GroupId: sendMsgReq.GroupId,
	})
	if err != nil {
		log.Printf("failed to call get group info rpc: %v\n", err)
		return
	}

	if !resp.Exists {
		return
	}

	for {
		evloopSpec, version := ctx.EvCfgWatcher.GetNode(sendMsgReq.GroupId)
		cli := api.MustNewChatEvLoopCliFromAdAddr(evloopSpec.IPPort)
		resp, err := cli.UniversalGroupEvloopRequest(context.Background(), &imchatevloop.UniversalGroupEvloopRequestReq{
			Version: version,
			GroupId: sendMsgReq.GroupId,
			Input: &evloopio.UniversalGroupEvloopInput{
				Input: &evloopio.UniversalGroupEvloopInput_SendMessage{
					SendMessage: &evloopio.SendMessageRequest{
						Session:         req.Session,
						EchoCode:        req.EchoCode,
						MessageData:     []byte(sendMsgReq.Msg),
						CheckIdempotent: sendMsgReq.CheckIdempotent,
						IdempotentKey:   sendMsgReq.IdempotentKey,
					},
				},
			},
		})
		if err != nil {
			log.Printf("failed to call UniversalGroupEvloopRequest rpc, err: %v\n", err)
			break
		}
		if !resp.Success {
			ctx.EvCfgWatcher.ForceUpdate(resp.Version)
			continue
		}
		break
	}
}
