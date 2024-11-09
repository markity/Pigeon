package chat

import (
	bizprotocol "pigeon/common/biz_protocol"
	"pigeon/im-relay/handle"
	"pigeon/kitex_gen/service/imrelay"
)

var createGroupReq bizprotocol.BizCreateGroup
var subReq bizprotocol.BizSub
var sendMsgReq bizprotocol.BizSendMessage
var sendApplyReq bizprotocol.BizSendApply
var handleApplyReq bizprotocol.BizHandleApply
var pullRelationsReq bizprotocol.BizPullRelations
var pullApplyReq bizprotocol.BizPullApply

func HandleChat(ctx *handle.HandleContext, req *imrelay.BizMessageReq) {
	switch req.Biz {
	// 创建群聊
	case createGroupReq.String():
		handleCreateGroup(ctx, req)
	case subReq.String():
		handleSub(ctx, req)
	// 群聊发消息
	case sendMsgReq.String():
		handleSendMsg(ctx, req)
	// 发送加入群聊命令
	case sendApplyReq.String():
		handleSendApply(ctx, req)
	// 处理加群请求
	case handleApplyReq.String():
		handleHandleApply(ctx, req)
	// 退出群聊
	// case "chat-group-quit":
	// todo支持解散群聊
	// // 解散群聊
	// case "chat-group-disband":
	// 拉全量关系
	case pullRelationsReq.String():
		handlePullRelations(ctx, req)
	// 拉全量申请
	case pullApplyReq.String():
		handlePullApply(ctx, req)
	}
}
