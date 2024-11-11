package bizpush

import (
	"pigeon/common/push"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/evloopio"
	"time"
)

type BizPusher struct {
	pushMan *push.PushManager
}

func NewBizPusher(pushMan *push.PushManager) *BizPusher {
	return &BizPusher{
		pushMan: pushMan,
	}
}

type SubRespInput struct {
	Session         *base.SessionEntry
	EchoCode        string
	GroupId         string
	SubOk           bool
	RelationVersion int64
	SeqId           int64
}

func (bp *BizPusher) SubResp(input *SubRespInput) {
	m := map[string]interface{}{
		"group_id":         input.GroupId,
		"sub_ok":           input.SubOk,
		"relation_version": input.RelationVersion,
		"seq_id":           input.SeqId,
	}
	bp.pushMan.PushToSessionByMap(input.Session, "push-sub-resp", input.EchoCode, m)
}

type SeqNotifyInput struct {
	Session *base.SessionEntry
	SeqId   int64
	GroupId string
	SendAt  int64
}

func (bp *BizPusher) SeqNotify(input *SeqNotifyInput) {
	m := map[string]interface{}{
		"seq_id":   input.SeqId,
		"group_id": input.GroupId,
		"send_at":  input.SendAt,
	}
	bp.pushMan.PushToSessionByMap(input.Session, "push-seq-notify", "", m)
}

type SendMessageRespInput struct {
	Session         *base.SessionEntry
	EchoCode        string
	RelationVersion int64
	Code            evloopio.SendMessageResponse_SendMessageCode // 0发送成功, 1幂等检查已发送, 2无权限
	// code为 2 时, SeqId无意义
	SeqId int64
}

func (bp *BizPusher) SendMessageResp(input *SendMessageRespInput) {
	m := map[string]interface{}{
		"code":   input.Code,
		"seq_id": input.SeqId,
	}
	bp.pushMan.PushToSessionByMap(input.Session, "push-send-msg-resp", input.EchoCode, m)
}

func (bp *BizPusher) StartBroadcastSeqId(subs []*base.SessionEntry, groupId string, seqId int64, sendAt time.Time) {
	// 先拷贝subs, 再异步下行seqid
	go func() {
		for _, sub := range subs {
			bp.SeqNotify(&SeqNotifyInput{
				Session: sub,
				SeqId:   seqId,
				GroupId: groupId,
				SendAt:  sendAt.UnixMilli(),
			})
		}
	}()
}
