package pushprotocol

import (
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/evloopio"
	"pigeon/kitex_gen/service/imrelation"
)

type SubResp struct {
	GroupId         string `json:"group_id"`
	SubOk           bool   `json:"sub_ok"`
	RelationVersion int64  `json:"relation_version"`
	SeqId           int64  `json:"seq_id"`
}

func (*SubResp) String() string {
	return "push-sub-resp"
}

type SendMessageResp struct {
	RelationVersion int64                                        `json:"relation_version"`
	Code            evloopio.SendMessageResponse_SendMessageCode `json:"code"`
	SeqId           int64                                        `json:"seq_id"`
}

func (*SendMessageResp) String() string {
	return "push-send-msg-resp"
}

type CreateGroupResp struct {
	OwnerId   string `json:"owner_id"`
	GroupId   string `json:"group_id"`
	CreatedAt int64  `json:"created_at"`
}

func (*CreateGroupResp) String() string {
	return "push-create-group-resp"
}

type FetchAllRelationsResp struct {
	Relations []*base.RelationEntry `json:"relations"`
}

func (*FetchAllRelationsResp) String() string {
	return "push-fetch-relations-resp"
}

type FetchAllApplicationsResp struct {
	Applications []*base.ApplyEntry `json:"applications"`
}

func (*FetchAllApplicationsResp) String() string {
	return "push-fetch-applications-resp"
}

type ApplyGroupResp struct {
	Code imrelation.ApplyGroupResp_ApplyGroupRespCode `json:"code"`
}

func (*ApplyGroupResp) String() string {
	return "push-apply-group-resp"
}

type HandleApplyResp struct {
	Code            imrelation.HandleApplyResp_HandleApplyRespCode `json:"code"`
	RelationVersion int64                                          `json:"relation_version"`
	ApplyMsg        string                                         `json:"apply_msg"`
	ApplyAt         int64                                          `json:"apply_at"`
	ApplyVersion    int64                                          `json:"apply_version"`
	HandleAt        int64                                          `json:"handle_at"`
}

func (*HandleApplyResp) String() string {
	return "push-handle-apply-resp"
}
