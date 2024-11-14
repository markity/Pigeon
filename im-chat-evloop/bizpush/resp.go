package bizpush

import (
	pushprotocol "pigeon/common/push_protocol"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/evloopio"
)

type SubRespInput struct {
	Session  *base.SessionEntry
	EchoCode string

	GroupId         string
	SubOk           bool
	RelationVersion int64
	SeqId           int64
}

func (bp *BizPusher) SubResp(input *SubRespInput) {
	m := &pushprotocol.SubResp{
		GroupId:         input.GroupId,
		SubOk:           input.SubOk,
		RelationVersion: input.RelationVersion,
		SeqId:           input.SeqId,
	}
	bp.pushMan.PushToSessionByAny(input.Session, m.String(), input.EchoCode, m)
}

type SendMessageRespInput struct {
	Session  *base.SessionEntry
	EchoCode string

	RelationVersion int64
	Code            evloopio.SendMessageResponse_SendMessageCode // 0发送成功, 1幂等检查已发送, 2无权限
	// code为 2 时, SeqId无意义
	SeqId int64
}

func (bp *BizPusher) SendMessageResp(input *SendMessageRespInput) {
	m := &pushprotocol.SendMessageResp{
		RelationVersion: input.RelationVersion,
		Code:            input.Code,
		SeqId:           input.SeqId,
	}
	bp.pushMan.PushToSessionByAny(input.Session, m.String(), input.EchoCode, m)
}
