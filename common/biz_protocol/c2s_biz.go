package bizprotocol

// chat-create-group
type BizCreateGroup struct {
	// empty
}

type BizPullRelations struct {
}

type BizSendMessage struct {
	GroupId string `json:"group_id"`
	Msg     string `json:"msg"`
}

type BizSendApply struct {
	GroupId  string `json:"group_id"`
	ApplyMsg string `json:"apply_msg"`
}

type BizHandleApply struct {
	UserId  string `json:"user_id"`
	GroupId string `json:"group_id"`
	Accept  bool   `json:"accept"`
}

type BizSub struct {
	GroupId string `json:"group_id"`
}
