package bizpush

import (
	pushprotocol "pigeon/common/push_protocol"
	"pigeon/kitex_gen/service/base"
	"time"
)

type SeqNotifyInput struct {
	Session *base.SessionEntry

	SeqId   int64
	GroupId string
	SendAt  int64
}

func (bp *BizPusher) SeqNotify(input *SeqNotifyInput) {
	m := &pushprotocol.SeqNotify{
		SeqId:   input.SeqId,
		GroupId: input.GroupId,
		SendAt:  input.SendAt,
	}
	bp.pushMan.PushToSessionByAny(input.Session, m.String(), "", m)
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
