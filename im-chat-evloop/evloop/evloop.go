package evloop

import (
	"fmt"
	"sync"
	"time"

	"pigeon/im-chat-evloop/bizpush"
	relationmanager "pigeon/im-chat-evloop/evloop/relation_manager"
	subscribemanager "pigeon/im-chat-evloop/evloop/subscribe_manager"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/evloopio"
	"pigeon/kitex_gen/service/imchatevloop"

	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
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

type ChatEvLoop struct {
	// 群聊id
	chatId string

	// 群主username
	ownerId string

	// 创建时间, 群聊创建时间, 消息发送时间都由evloop server定
	createdAt int64

	// 当前消息序列号
	seqId int64

	// 全量群员信息, key是userId
	relationManager *relationmanager.RelationManager

	// 订阅者信息, key是userId
	subscribeManager *subscribemanager.SubscribeManager

	// 用来生成消息的幂等唯一id, 雪花算法
	idGen *snowflake.Node

	// push工具类
	bPush *bizpush.BizPusher

	// 数据库
	db *gorm.DB

	// 作用1: 保护queue
	// 作用2: 变更status
	queueMu sync.Mutex
	cond    *sync.Cond
	status  chatEvloopStatus
	queue   []*evInput
}

func (c *ChatEvLoop) GetCreatedAt() int64 {
	return c.createdAt
}

// 群聊创建时间由chat-evloop server生成
func NewChatEvLoopAndStart(in *NewChatEvLoopInput) (loop *ChatEvLoop) {
	now := time.Now()
	lp := &ChatEvLoop{
		chatId:    in.ChatId,
		ownerId:   in.OwnerId,
		createdAt: now.UnixMilli(),
		relationManager: relationmanager.NewRelationManager(&base.RelationEntry{
			GroupId:         in.ChatId,
			UserId:          in.OwnerId,
			Status:          base.RelationStatus_RELATION_STATUS_OWNER,
			ChangeType:      base.RelationChangeType_RELATION_CHANGE_TYPE_CREATE_GROUP,
			RelationVersion: 1,
			ChangeAt:        now.UnixMilli(),
		}),
		subscribeManager: subscribemanager.NewSubscribeManager(),
		bPush:            in.PushMan,
		db:               in.DB,
		idGen:            in.Snowflake,
		status:           statusRunning,
		queue:            make([]*evInput, 0, 1024),
	}
	lp.cond = sync.NewCond(&lp.queueMu)
	lp.start()
	return lp
}

func NewMigrateEvLoopAndStart(resp *imchatevloop.DoMigrateResp, pushMan *bizpush.BizPusher, db *gorm.DB, sn *snowflake.Node) *ChatEvLoop {
	relationMan := relationmanager.NewRelationManagerFromMigrage(resp)
	subscrberMan := subscribemanager.NewSubscrbieManagerFromMigrage(resp)

	lp := &ChatEvLoop{
		chatId:           resp.GroupId,
		ownerId:          resp.OwnerId,
		createdAt:        resp.CreatedAt,
		relationManager:  relationMan,
		subscribeManager: subscrberMan,
		bPush:            pushMan,
		db:               db,
		idGen:            sn,
		queueMu:          sync.Mutex{},
		cond:             nil,
		status:           statusRunning,
		queue:            make([]*evInput, 0, 1024),
	}
	lp.cond = sync.NewCond(&lp.queueMu)
	lp.start()
	return lp
}

type queueMessageError struct {
	typ int
}

var errMigrating queueMessageError = queueMessageError{
	typ: 1,
}

var errStop queueMessageError = queueMessageError{
	typ: 2,
}

var errRunning queueMessageError = queueMessageError{
	typ: 3,
}

func (c *queueMessageError) Error() string {
	if c.typ == 1 {
		return "migrating"
	} else if c.typ == 2 {
		return "stop"
	} else if c.typ == 3 {
		return "running"
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

func IsErrRunning(e error) bool {
	if ee, ok := e.(*queueMessageError); ok {
		return ee.typ == 3
	}
	return false
}

func (c *ChatEvLoop) QueueMessage(msg *evloopio.UniversalGroupEvloopInput) (
	*EvOutputUniversal, error) {
	fmt.Println("queue message in")
	outChan := make(chan interface{}, 1)
	c.queueMu.Lock()
	if c.status != statusRunning {
		status := c.status
		c.queueMu.Unlock()
		switch status {
		case statusBeMoving:
			return nil, &errMigrating
		case statusStop:
			return nil, &errStop
		}
	}
	c.queue = append(c.queue,
		&evInput{
			input: &evInputUniversal{
				Input: msg,
			},
			output: outChan,
		})
	c.cond.Broadcast()
	c.queueMu.Unlock()

	fmt.Println("waiting")

	output := <-outChan
	out := output.(*EvOutputUniversal)
	fmt.Println("queue message out")
	return out, nil
}

func (c *ChatEvLoop) Move() (*EvOutputMove, error) {
	fmt.Println("lock")
	c.queueMu.Lock()
	status := c.status
	switch status {
	case statusStop:
		c.queueMu.Unlock()
		return nil, &errStop
	}
	outChan := make(chan interface{}, 1)
	c.queue = append(c.queue, &evInput{input: &evInputMove{}, output: outChan})
	fmt.Println("unlock")
	c.cond.Broadcast()
	c.queueMu.Unlock()

	output := <-outChan
	out := output.(*EvOutputMove)
	if out.err != nil {
		return nil, out.err
	}
	return out, nil
}

func (c *ChatEvLoop) Stop() error {
	c.queueMu.Lock()
	status := c.status
	if status == statusRunning {
		c.queueMu.Unlock()
		return &errRunning
	}
	if status == statusStop {
		c.queueMu.Unlock()
		return nil
	}

	outChan := make(chan interface{}, 1)
	c.queue = append(c.queue, &evInput{input: &evInputStop{}, output: outChan})
	c.cond.Broadcast()
	c.queueMu.Unlock()

	output := <-outChan
	out := output.(*EvOutputStop)
	return out.err
}
