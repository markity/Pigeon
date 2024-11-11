package bizpush

import (
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/imrelation"
)

type CreateGroupRespInput struct {
	Session   *base.SessionEntry
	EchoCode  string
	OwnerId   string
	GroupId   string
	CreatedAt int64
}

// 创建群聊请求的resp
func (bp *BizPusher) CreateGroupResp(input *CreateGroupRespInput) {
	m := map[string]interface{}{
		"owner_id":   input.OwnerId,
		"group_id":   input.GroupId,
		"created_at": input.CreatedAt,
	}
	bp.pushMan.PushToSessionByMap(input.Session, "push-create-group-resp", input.EchoCode, m)
}

type FetchAllRelationsRespInput struct {
	Session   *base.SessionEntry
	EchoCode  string
	Relations []*imrelation.RelationEntry
}

type relationEntry struct {
	Username           string                  `json:"username"`
	GroupId            string                  `json:"group_id"`
	RelationVersion    int64                   `json:"relation_version"`
	RelationStatus     base.RelationStatus     `json:"relation_status"`
	RelationChangeType base.RelationChangeType `json:"relation_change_type"`
	UpdatedAt          int64                   `json:"updated_at"`
}

func (bp *BizPusher) FetchAllRelationsResp(input *FetchAllRelationsRespInput) {
	myRelations := make([]*relationEntry, 0, len(input.Relations))
	for _, v := range input.Relations {
		myRelations = append(myRelations, &relationEntry{
			Username:           v.UserId,
			GroupId:            v.GroupId,
			RelationVersion:    v.RelationVersion,
			RelationStatus:     v.Status,
			RelationChangeType: v.ChangeType,
			UpdatedAt:          v.UpdatedAt,
		})
	}
	m := map[string]interface{}{
		"relations": myRelations,
	}
	bp.pushMan.PushToSessionByMap(input.Session, "push-fetch-relations-resp", input.EchoCode, m)
}

type FetchAllAppliesRespInput struct {
	Session  *base.SessionEntry
	EchoCode string
	Applies  []*imrelation.ApplyEntry
}

type applyEntry struct {
	Username     string           `json:"username"`
	GroupId      string           `json:"group_id"`
	ApplyVersion int64            `json:"version"`
	ApplyStatus  base.ApplyStatus `json:"apply_status"`
	UpdatedAt    int64            `json:"updated_at"`
	Msg          string           `json:"msg"`
}

func (bp *BizPusher) FetchAllAppliesResp(input *FetchAllAppliesRespInput) {
	myRelations := make([]*applyEntry, 0, len(input.Applies))
	for _, v := range input.Applies {
		myRelations = append(myRelations, &applyEntry{
			Username:     v.UserId,
			GroupId:      v.GroupId,
			ApplyVersion: v.ApplyVersion,
			ApplyStatus:  v.Status,
			UpdatedAt:    v.ApplyAt,
			Msg:          v.ApplyMsg,
		})
	}
	m := map[string]interface{}{
		"applies": myRelations,
	}
	bp.pushMan.PushToSessionByMap(input.Session, "push-fetch-applies-resp", input.EchoCode, m)
}

type ApplyGroupRespInput struct {
	Session  *base.SessionEntry
	EchoCode string
	Code     imrelation.ApplyGroupResp_ApplyGroupRespCode
}

func (bp *BizPusher) ApplyGroupResp(input *ApplyGroupRespInput) {
	m := map[string]interface{}{
		"code": input.Code,
	}
	bp.pushMan.PushToSessionByMap(input.Session, "push-apply-group-resp", input.EchoCode, m)
}

type HandleApplyRespInput struct {
	Session      *base.SessionEntry
	EchoCode     string
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
	m := map[string]interface{}{
		"code":             input.Code,
		"relation_version": input.RelationVersion,
		"apply_msg":        input.ApplyMsg,
		"apply_at":         input.ApplyAt,
		"apply_version":    input.ApplyVersion,
		"handle_at":        input.HandleAt,
	}
	bp.pushMan.PushToSessionByMap(input.Session, "push-handle-apply-resp", input.EchoCode, m)
}
