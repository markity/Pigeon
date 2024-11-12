package chat

import (
	"context"
	"encoding/json"
	"log"
	bizprotocol "pigeon/common/biz_protocol"
	"pigeon/im-relay/bizpush"
	"pigeon/im-relay/db"
	"pigeon/im-relay/handle"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/imrelation"
	"pigeon/kitex_gen/service/imrelay"
)

func handlePullMessages(ctx *handle.HandleContext, req *imrelay.BizMessageReq) {
	var pullMsgReq bizprotocol.BizPullMessage
	err := json.Unmarshal(req.Data, &pullMsgReq)
	if err != nil {
		log.Printf("failed to unmarshal pull messages request, err: %v, data: %v\n", err, string(req.Data))
		return
	}

	resp, err := ctx.RelationCli.GetGroupInfo(context.Background(), &imrelation.GetGroupInfoReq{
		GroupId: pullMsgReq.GroupId,
	})
	if err != nil {
		log.Printf("failed to GetGroupInfo: %v\n", err)
		return
	}
	if !resp.Exists {
		ctx.BPush.PullMsgResp(&bizpush.PullMsgRespInput{
			Session:  req.Session,
			EchoCode: req.EchoCode,
			Code:     bizpush.NoPermission,
			GroupId:  pullMsgReq.GroupId,
			Data:     make([]*bizpush.MsgEntry, 0),
		})
		return
	}

	resp2, err := ctx.RelationCli.GetGroupMemberInfo(context.Background(), &imrelation.GetGroupMemberInfoReq{
		GroupId:  pullMsgReq.GroupId,
		Username: req.Session.Username,
	})
	if err != nil {
		log.Printf("failed to GetGroupMemberInfo: %v\n", err)
		return
	}
	if resp2.Stauts == base.RelationStatus_RELATION_STATUS_NOT_IN_GROUP {
		ctx.BPush.PullMsgResp(&bizpush.PullMsgRespInput{
			Session:  req.Session,
			EchoCode: req.EchoCode,
			Code:     bizpush.NoPermission,
			GroupId:  pullMsgReq.GroupId,
			Data:     make([]*bizpush.MsgEntry, 0),
		})
		return
	}

	msgs, err := db.GetMessages(ctx.DB, pullMsgReq.GroupId, pullMsgReq.MaxSeqId, pullMsgReq.Limit)
	if err != nil {
		log.Printf("failed to GetMessages: %v\n", err)
		return
	}

	var start int64
	var end int64
	if len(msgs) != 0 {
		start = msgs[0].SeqId
		end = msgs[len(msgs)-1].SeqId
	}
	toMsgs := make([]*bizpush.MsgEntry, 0, len(msgs))
	for _, v := range msgs {
		toMsgs = append(toMsgs, &bizpush.MsgEntry{
			MsgData:   v.Data,
			Sender:    v.OwnerId,
			MsgSeq:    v.SeqId,
			CreatedAt: v.CreatedAt,
		})
	}

	ctx.BPush.PullMsgResp(&bizpush.PullMsgRespInput{
		Session:    req.Session,
		EchoCode:   req.EchoCode,
		Code:       bizpush.OK,
		GroupId:    pullMsgReq.GroupId,
		StartSeqId: start,
		EndSeqId:   end,
		Data:       toMsgs,
	})
}
