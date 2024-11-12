package bizprotocol

// 客户端发给服务端的包体

// chat-create-group
type BizCreateGroup struct {
	// empty
}

func (*BizCreateGroup) String() string {
	return "chat-create-group"
}

type BizPullRelations struct {
	// empty
}

func (*BizPullRelations) String() string {
	return "chat-pull-relations"
}

type BizPullApply struct {
	// empty
}

func (*BizPullApply) String() string {
	return "chat-pull-apply"
}

type BizSendMessage struct {
	GroupId         string `json:"group_id"`
	Msg             string `json:"msg"`
	CheckIdempotent bool   `json:"check_idempotent"`
	IdempotentKey   string `json:"idempotent_key"`
}

func (*BizSendMessage) String() string {
	return "chat-send-message"
}

type BizSendApply struct {
	GroupId  string `json:"group_id"`
	ApplyMsg string `json:"apply_msg"`
}

func (*BizSendApply) String() string {
	return "chat-send-apply"
}

type BizHandleApply struct {
	UserId  string `json:"user_id"`
	GroupId string `json:"group_id"`
	Accept  bool   `json:"accept"`
}

func (*BizHandleApply) String() string {
	return "chat-handle-apply"
}

type BizSub struct {
	GroupId string `json:"group_id"`
}

func (*BizSub) String() string {
	return "chat-sub"
}

type BizPullMessage struct {
	GroupId  string `json:"group_id"`
	MaxSeqId int64  `json:"max_seq_id"`
	Limit    int64  `json:"limit"`
}

func (*BizPullMessage) String() string {
	return "chat-pull-msg"
}
