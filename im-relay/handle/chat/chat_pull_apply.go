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

func handlePullApply(ctx *handle.HandleContext, req *imrelay.BizMessageReq) {
	var pullApplyReq bizprotocol.BizPullApply
	err := json.Unmarshal(req.Data, &pullApplyReq)
	if err != nil {
		log.Printf("failed to unmarshal pull apply request, err: %v, data: %v\n", err, string(req.Data))
		return
	}

	_, err = ctx.RelationCli.FetchAllApplications(context.Background(), &imrelation.FetchAllApplicationsReq{
		Session:  req.Session,
		EchoCode: req.EchoCode,
	})
	if err != nil {
		log.Printf("failed to call fetch all application group rpc: %v\n", err)
	}
}
