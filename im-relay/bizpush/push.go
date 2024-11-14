package bizpush

import (
	"pigeon/common/push"
	"pigeon/kitex_gen/service/base"
)

type BizPusher struct {
	pushMan *push.PushManager
}

func NewBisPusher(pushMan *push.PushManager) *BizPusher {
	return &BizPusher{
		pushMan: pushMan,
	}
}

type PullMsgRespCode int

const (
	UnUsed PullMsgRespCode = iota
	OK
	NoPermission
)

type PullMsgRespInput struct {
	Session  *base.SessionEntry
	EchoCode string

	Code       PullMsgRespCode
	GroupId    string
	StartSeqId int64
	EndSeqId   int64
	Data       []*MsgEntry
}

type MsgEntry struct {
	MsgData   string
	Sender    string
	MsgSeq    int64
	CreatedAt int64
}

func (bp *BizPusher) PullMsgResp(input *PullMsgRespInput) {
	m := map[string]interface{}{
		"code":         input.Code,
		"group_id":     input.GroupId,
		"start_seq_id": input.StartSeqId,
		"end_seq_id":   input.EndSeqId,
		"data":         input.Data,
	}
	bp.pushMan.PushToSessionByAny(input.Session, "push-pull-resp", input.EchoCode, m)
}
