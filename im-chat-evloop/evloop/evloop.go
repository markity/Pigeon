package evloop

import (
	"pigeon/kitex_gen/service/evloopio"
	"pigeon/kitex_gen/service/imrelation"
)

type subscriberInfo struct {
	GwAddrPort               string
	SessionId                string
	SubscribeRelationVersion int64
}

type memberInfo struct {
	UserId          string
	RelationVersion int64
	Status          imrelation.RelationStatus
}

type chatEvLoop struct {
	chatId  string
	ownerId string

	// 全量群员信息, key是userId
	members map[string]*memberInfo

	// 订阅者信息, key是userId
	subscribers map[string][]*subscriberInfo
}

type NewChatEvLoopInput struct {
	ChatId string
	// 群主id
	OwnerId string
}

func NewChatEvLoop(in *NewChatEvLoopInput) *chatEvLoop {
	return &chatEvLoop{
		chatId:      in.ChatId,
		ownerId:     in.OwnerId,
		members:     make(map[string]*memberInfo),
		subscribers: make(map[string][]*subscriberInfo),
	}
}

func (c *chatEvLoop) QueueMessage(msg *evloopio.UniversalGroupEvloopInput)

func (c *chatEvLoop) Start() {
	go func() {
	}()
}
