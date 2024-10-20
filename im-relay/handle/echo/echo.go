package echo

import (
	"context"
	"log"
	"pigeon/im-relay/api"
	"pigeon/im-relay/handle"
	"pigeon/kitex_gen/service/imgateway"
	"pigeon/kitex_gen/service/imrelay"
)

func HandleEcho(ctx *handle.HandleContext, req *imrelay.BizMessageReq) {
	if req.Biz == "echo" {
		cli := api.MustNewIMGatewayClientFromAdAddr(req.Session.GwAdvertiseAddrport)
		_, err := cli.PushMessage(context.Background(), &imgateway.PushMessageReq{
			SessionId: req.Session.SessionId,
			PushType:  "push-echo",
			EchoCode:  req.EchoCode,
			Data:      req.Data,
		})
		if err != nil {
			log.Printf("failed to push echo message: %v\n", err)
		}
	}
}
