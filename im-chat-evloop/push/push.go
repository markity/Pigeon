package push

import (
	"context"
	"encoding/json"
	"log"
	"pigeon/im-relation/api"
	"pigeon/kitex_gen/service/imgateway"
	"time"
)

type SubRespInput struct {
	GwAddrPort      string
	SessionId       string
	GroupId         string
	SubOk           bool
	RelationVersion int64
	SeqId           int64
}

func SeqResp(input *SubRespInput) {
	cli := api.NewGatewayClientFromAdAddr(input.GwAddrPort)
	for {
		_, err := cli.PushMessage(context.Background(), &imgateway.PushMessageReq{
			SessionId: input.SessionId,
			PushType:  "push-sub-resp",
			EchoCode:  "",
			Data: mustMarshal(map[string]interface{}{
				"group_id":         input.GroupId,
				"sub_ok":           input.SubOk,
				"relation_version": input.RelationVersion,
				"seq_id":           input.SeqId,
			}),
		})
		if err != nil {
			log.Printf("push SeqResp: %v\n", err)
			time.Sleep(time.Millisecond * 50)
			continue
		}
		break
	}
}

type SeqNotifyInput struct {
	GwAddrPort string
	SessionId  string
	SeqId      int64
	GroupId    string
	SendAt     int64
}

func SeqNotify(input *SeqNotifyInput) {
	cli := api.NewGatewayClientFromAdAddr(input.GwAddrPort)
	for {
		_, err := cli.PushMessage(context.Background(), &imgateway.PushMessageReq{
			SessionId: input.SessionId,
			PushType:  "push-seq-notify",
			EchoCode:  "",
			Data: mustMarshal(map[string]interface{}{
				"seq_id":   input.SeqId,
				"group_id": input.GroupId,
				"send_at":  input.SendAt,
			}),
		})
		if err != nil {
			log.Printf("push SeqNotify: %v\n", err)
			time.Sleep(time.Millisecond * 50)
			continue
		}
		break
	}
}

func mustMarshal(obj interface{}) []byte {
	bs, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	return bs
}
