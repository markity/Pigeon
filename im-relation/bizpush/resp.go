package bizpush

import (
	pushprotocol "pigeon/common/push_protocol"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/imrelation"
)

type CreateGroupRespInput struct {
	Session  *base.SessionEntry
	EchoCode string

	OwnerId   string
	GroupId   string
	CreatedAt int64
}

// 创建群聊请求的resp
func (bp *BizPusher) CreateGroupResp(input *CreateGroupRespInput) {
	m := *&pushprotocol.CreateGroupResp{
		OwnerId:   input.OwnerId,
		GroupId:   input.OwnerId,
		CreatedAt: input.CreatedAt,
	}
	bp.pushMan.PushToSessionByAny(input.Session, m.String(), input.EchoCode, m)
}

type FetchAllRelationsRespInput struct {
	Session  *base.SessionEntry
	EchoCode string

	Relations []*base.RelationEntry
}

func (bp *BizPusher) FetchAllRelationsResp(input *FetchAllRelationsRespInput) {
	m := &pushprotocol.FetchAllRelationsResp{
		Relations: input.Relations,
	}
	bp.pushMan.PushToSessionByAny(input.Session, m.String(), input.EchoCode, m)
}

type FetchAllAppliesRespInput struct {
	Session      *base.SessionEntry
	EchoCode     string
	Applications []*base.ApplyEntry
}

func (bp *BizPusher) FetchAllAppliesResp(input *FetchAllAppliesRespInput) {
	m := &pushprotocol.FetchAllApplicationsResp{
		Applications: input.Applications,
	}
	bp.pushMan.PushToSessionByAny(input.Session, m.String(), input.EchoCode, m)
}

type ApplyGroupRespInput struct {
	Session  *base.SessionEntry
	EchoCode string

	Code imrelation.ApplyGroupResp_ApplyGroupRespCode
}

func (bp *BizPusher) ApplyGroupResp(input *ApplyGroupRespInput) {
	m := &pushprotocol.ApplyGroupResp{
		Code: input.Code,
	}
	bp.pushMan.PushToSessionByAny(input.Session, m.String(), input.EchoCode, m)
}

type HandleApplyRespInput struct {
	Session  *base.SessionEntry
	EchoCode string

	Code         imrelation.HandleApplyResp_HandleApplyRespCode
	ApplyMsg     string
	ApplyAt      int64
	ApplyVersion int64
	ApplyStatus  base.ApplyStatus
	HandleAt     int64

	// 如果是接受请求, 且ApplyStatus变为已接受, relation_version不为0
	RelationVersion int64
}

func (bp *BizPusher) HandleApplyResp(input *HandleApplyRespInput) {
	m := &pushprotocol.HandleApplyResp{
		Code:            input.Code,
		RelationVersion: input.RelationVersion,
		ApplyMsg:        input.ApplyMsg,
		ApplyAt:         input.ApplyAt,
		ApplyVersion:    input.ApplyVersion,
		HandleAt:        input.HandleAt,
	}
	bp.pushMan.PushToSessionByAny(input.Session, m.String(), input.EchoCode, m)
}
