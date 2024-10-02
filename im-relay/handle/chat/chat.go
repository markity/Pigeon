package chat

import (
	"context"
	"log"

	"pigeon/im-relay/handle"
	"pigeon/kitex_gen/service/imrelation"
	"pigeon/kitex_gen/service/imrelay"
)

func HandleChat(ctx *handle.HandleContext, req *imrelay.BizMessageReq) {
	switch req.Biz {
	// 创建群聊
	case "chat-group-create":
		_, err := ctx.RelationCli.CreateGroup(context.Background(), &imrelation.CreateGroupReq{
			Session: req.Session,
		})
		if err != nil {
			log.Printf("failed to call create group rpc, will retry: %v\n", err)
		}
	// 群聊发消息
	case "chat-group-send-msg":
	// 发送加入群聊命令
	case "chat-group-send-join":
	// 处理加群请求
	case "chat-group-handle-join":
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
