package bizpush

import (
	"pigeon/common/push"
)

type BizPusher struct {
	pushMan *push.PushManager
}

func NewBizPusher(pushMan *push.PushManager) *BizPusher {
	return &BizPusher{
		pushMan: pushMan,
	}
}
