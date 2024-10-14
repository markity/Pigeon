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

func HandleChat(ctx *handle.HandleContext, req *imrelay.BizMessageReq) {
	switch req.Biz {
	// 创建群聊
	case "chat-create-group":
		var createGroupReq bizprotocol.BizCreateGroup
		err := json.Unmarshal(req.Data, &createGroupReq)
		if err != nil {
			log.Printf("failed to unmarshal create group request, err: %v, data: %v\n", err, string(req.Data))
		}

		_, err = ctx.RelationCli.CreateGroup(context.Background(), &imrelation.CreateGroupReq{
			Session:  req.Session,
			EchoCode: req.EchoCode,
		})
		if err != nil {
			log.Printf("failed to call create group rpc: %v\n", err)
		}
	// 群聊发消息
	case "chat-group-send-msg":
	// 发送加入群聊命令
	case "chat-group-send-apply":
	// 处理加群请求
	case "chat-group-handle-apply":
	// 退出群聊
	case "chat-group-quit":
	// 解散群聊
	case "chat-group-disband":
	// 拉全量关系
	case "chat-pull-relation":
	// 拉全量申请
	case "chat-pull-join":
	}
}
