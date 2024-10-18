package evloop

import (
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/evloopio"
	"pigeon/kitex_gen/service/imchatevloop"
)

type chatEvloopStatus int

const (
	// 可以正常接受请求
	statusRunning chatEvloopStatus = iota
	// 收到migrate指令, 开始迁移, 此时不接受请求
	statusBeMoving
	// 迁移完成, eventloop停止且eventloop停止工作
	statusStop
)

type subscriberInfo struct {
	GwAddrPort               string
	SessionId                string
	SubscribeRelationVersion int64
}

type memberInfo struct {
	UserId          string
	RelationVersion int64
	Status          base.RelationStatus
	ChangeType      base.RelationChangeType
}

type ChatEvLoop struct {
	chatId  string
	ownerId string

	// 全量群员信息, key是userId
	members map[string]*memberInfo

	// 订阅者信息, key是userId
	subscribers map[string][]*subscriberInfo

	status chatEvloopStatus
}

type NewChatEvLoopInput struct {
	ChatId string
	// 群主id
	OwnerId string
}

func NewChatEvLoopAndStart(in *NewChatEvLoopInput) *ChatEvLoop {
	mems := make(map[string]*memberInfo)
	mems[in.OwnerId] = &memberInfo{
		UserId:          in.OwnerId,
		RelationVersion: 1,
		Status:          base.RelationStatus_RELATION_STATUS_OWNER,
	}
	lp := &ChatEvLoop{
		chatId:      in.ChatId,
		ownerId:     in.OwnerId,
		members:     make(map[string]*memberInfo),
		subscribers: make(map[string][]*subscriberInfo),
		status:      statusRunning,
	}
	lp.Start()
	return lp
}

func NewMigrateEvLoop(resp *imchatevloop.DoMigrateResp) *ChatEvLoop {

	members := make(map[string]*memberInfo)
	for _, v := range resp.Relations {
		members[v.MemberId] = &memberInfo{
			UserId:          v.MemberId,
			RelationVersion: v.RelationVersion,
			Status:          v.Status,
			ChangeType:      v.ChangeType,
		}
	}

	subscriberInfos := make(map[string][]*subscriberInfo)
	for _, v := range resp.Subscribers {
		subscriberInfos[v.MemberId] = append(subscriberInfos[v.MemberId], &subscriberInfo{
			GwAddrPort:               v.GwAddrport,
			SessionId:                v.SessionId,
			SubscribeRelationVersion: v.OnSubRelationVersion,
		})
	}

	lp := &ChatEvLoop{
		chatId:      resp.GroupId,
		ownerId:     resp.OwnerId,
		members:     members,
		subscribers: subscriberInfos,
		status:      statusRunning,
	}
	lp.Start()
	return lp
}

type queueMessageError struct {
	typ int
}

func (c *queueMessageError) Error() string {
	if c.typ == 1 {
		return "migrating"
	} else if c.typ == 2 {
		return "stop"
	}
	panic("unexpected")
}

func IsErrMigrating(e error) bool {
	if ee, ok := e.(*queueMessageError); ok {
		return ee.typ == 1
	}
	return false
}

func IsErrStop(e error) bool {
	if ee, ok := e.(*queueMessageError); ok {
		return ee.typ == 2
	}
	return false
}

func (c *ChatEvLoop) QueueMessage(msg *evloopio.UniversalGroupEvloopInput) (output *evloopio.UniversalGroupEvloopOutput, err error) {

}

func (c *ChatEvLoop) Start() {
	go func() {
	}()
}
