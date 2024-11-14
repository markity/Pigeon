package pushprotocol

import (
	"pigeon/kitex_gen/service/base"
)

type ApplyGroupNotify struct {
	Username     string `json:"username"`
	GroupId      string `json:"group_id"`
	ApplyVersion int64  `json:"apply_version"`
	ApplyMsg     string `json:"apply_msg"`
	ApplyAt      int64  `json:"apply_at"`
}

func (*ApplyGroupNotify) String() string {
	return "push-apply-notify"
}

type HandleApplyNotify struct {
	Username     string           `json:"username"`
	GroupId      string           `json:"group_id"`
	ApplyVersion int64            `json:"apply_version"`
	ApplyMsg     string           `json:"apply_msg"`
	ApplyStatus  base.ApplyStatus `json:"apply_status"`
	ApplyAt      int64            `json:"apply_at"`
	HandleAt     int64            `json:"handle_at"`
}

func (*HandleApplyNotify) String() string {
	return "push-handle-apply-notify"
}

type RelationChangeNotify struct {
	Username           string                  `json:"username"`
	GroupId            string                  `json:"group_id"`
	RelationVersion    int64                   `json:"relation_version"`
	RelationStatus     base.RelationStatus     `json:"relation_status"`
	RelationChangeType base.RelationChangeType `json:"relation_change_type"`
	UpdatedAt          int64                   `json:"updated_at"`
}

func (*RelationChangeNotify) String() string {
	return "push-relation-change-notify"
}

type SeqNotify struct {
	SeqId   int64  `json:"seq_id"`
	GroupId string `json:"group_id"`
	SendAt  int64  `json:"send_at"`
}

func (*SeqNotify) String() string {
	return "push-seq-notify"
}
