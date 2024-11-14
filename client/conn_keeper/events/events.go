package events

import (
	"pigeon/common/protocol"
	"pigeon/kitex_gen/service/base"
)

type CloseReason int

const (
	closeReasonUnUsed                = 0
	CloseReasonLoginFail CloseReason = iota
	CloseReasonUserNetwork
	CloseReasonOtherKick
)

type RelationAndSeqEntry struct {
	base.RelationEntry
	SeqId int64
}

type EventLoginResult struct {
	Code protocol.LoginRespCode
	// 自己的session
	Session string
	// 设备管理全量
	Devices []*protocol.DeviceSessionEntry
	// 群聊关系全量, 某个人关于某个群的关系信息, 以及能拉seq的大小
	RelationsAndSeq map[string]*RelationAndSeqEntry
}

type EventClose struct {
	CloseReason CloseReason

	FromSession     string
	FromSessionDesc string
}
