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

func handleSub(ctx *handle.HandleContext, req *imrelay.BizMessageReq) {
	var subReq bizprotocol.BizSub
	err := json.Unmarshal(req.Data, &subReq)
	if err != nil {
		log.Printf("failed to unmarshal sub group request, err: %v, data: %v\n", err, string(req.Data))
		return
	}
	// TODO 这里可以无限重试
	resp, err := ctx.RelationCli.GetGroupInfo(context.Background(), &imrelation.GetGroupInfoReq{
		GroupId: subReq.GroupId,
	})
	if err != nil {
		log.Printf("failed to call GetGroupInfo rpc, err: %v\n", err)

		return
	}
	if !resp.Exists {
		return
	}

	for {
		evloopSpec, version := ctx.EvCfgWatcher.GetNode(subReq.GroupId)
		cli := api.MustNewChatEvLoopCliFromAdAddr(evloopSpec.IPPort)
		resp, err := cli.UniversalGroupEvloopRequest(context.Background(), &imchatevloop.UniversalGroupEvloopRequestReq{
			Version: version,
			GroupId: subReq.GroupId,
			Input: &evloopio.UniversalGroupEvloopInput{
				Input: &evloopio.UniversalGroupEvloopInput_SubscribeGroup{
					SubscribeGroup: &evloopio.SubscribeGroupRequest{
						Session:  req.Session,
						EchoCode: req.EchoCode,
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
