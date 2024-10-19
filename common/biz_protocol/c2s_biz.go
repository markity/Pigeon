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
