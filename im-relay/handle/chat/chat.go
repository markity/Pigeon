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
			return
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
		// var sendMsgReq bizprotocol.BizSendMessage
		// err := json.Unmarshal(req.Data, &sendMsgReq)
		// if err != nil {
		// 	log.Printf("failed to unmarshal send msg request, err: %v, data: %v\n", err, string(req.Data))
		// 	return
		// }
		// resp, err := ctx.RelationCli.GetGroupInfo(context.Background(), &imrelation.GetGroupInfoReq{
		// 	GroupId: sendMsgReq.GroupId,
		// })
		// if err != nil {
		// 	log.Printf("failed to call get group info rpc: %v\n", err)
		// 	return
		// }

	// 发送加入群聊命令
	case "chat-group-send-apply":
		var sendApplyReq bizprotocol.BizSendApply
		err := json.Unmarshal(req.Data, &sendApplyReq)
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
	// 处理加群请求
	case "chat-group-handle-apply":
	// 退出群聊
	case "chat-group-quit":
	// 解散群聊
	case "chat-group-disband":
	// 拉全量关系
	case "chat-pull-relation":
		var pullRelationsReq bizprotocol.BizPullRelations
		err := json.Unmarshal(req.Data, &pullRelationsReq)
		if err != nil {
			log.Printf("failed to unmarshal pull relation request, err: %v, data: %v\n", err, string(req.Data))
			return
		}

		_, err = ctx.RelationCli.FetchAllRelations(context.Background(), &imrelation.FetchAllRelationsReq{
			Session:  req.Session,
			EchoCode: req.EchoCode,
		})
		if err != nil {
			log.Printf("failed to call fetch all relations group rpc: %v\n", err)
		}

	// 拉全量申请
	case "chat-pull-join":
		var pullJoinReq bizprotocol.BizPullRelations
		err := json.Unmarshal(req.Data, &pullJoinReq)
		if err != nil {
			log.Printf("failed to unmarshal pull join request, err: %v, data: %v\n", err, string(req.Data))
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
}
