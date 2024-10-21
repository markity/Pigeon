package evloop

import (
	"fmt"
	"log"
	relationmanager "pigeon/im-chat-evloop/evloop/relation_manager"
	subscribemanager "pigeon/im-chat-evloop/evloop/subscribe_manager"
	"pigeon/im-chat-evloop/push"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/evloopio"
	"pigeon/kitex_gen/service/imchatevloop"
	"sync"
	"time"
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
	chatId  string
	ownerId string

	createdAt int64

	seqId int64

	// 全量群员信息, key是userId
	relationManager *relationmanager.RelationManager

	// 订阅者信息, key是userId
	subscribeManager *subscribemanager.SubscribeManager

	// 作用1: 保护queue
	// 作用2: 变更status
	queueMu sync.Mutex
	cond    *sync.Cond
	status  chatEvloopStatus
	queue   []*evInput
}

type evInput struct {
	input  interface{}
	output chan interface{}
}

type evInputMove struct{}

type evInputStop struct{}

type evInputUniversal struct {
	Input *evloopio.UniversalGroupEvloopInput
}

// 拿到evloop的全部状态, 用于迁移
type EvOutputMove struct {
	err       error // queueMessageError, 可能是stop
	ChatId    string
	OwnerId   string
	CreatedAt int64
	SeqId     int64
	Relations map[string]*base.RelationEntry
	Subs      map[string](map[string]*subscribemanager.Sub)
}

type EvOutputStop struct {
	err error // queueMessageError, 可能是running
}

type EvOutputUniversal struct {
	err    error // queueMessageError, 可能是stop或moving
	Output *evloopio.UniversalGroupEvloopOutput
}

type NewChatEvLoopInput struct {
	ChatId string
	// 群主id
	OwnerId string
}

// 群聊创建时间由chat-evloop server生成
func NewChatEvLoopAndStart(in *NewChatEvLoopInput) (loop *ChatEvLoop, createdAt int64) {
	now := time.Now()
	lp := &ChatEvLoop{
		chatId:  in.ChatId,
		ownerId: in.OwnerId,
		relationManager: relationmanager.NewRelationManager(&base.RelationEntry{
			GroupId:         in.ChatId,
			UserId:          in.OwnerId,
			Status:          base.RelationStatus_RELATION_STATUS_OWNER,
			ChangeType:      base.RelationChangeType_RELATION_CHANGE_TYPE_CREATE_GROUP,
			RelationVersion: 1,
			ChangeAt:        now.UnixMilli(),
		}),
		subscribeManager: subscribemanager.NewSubscribeManager(),
		status:           statusRunning,
		queue:            make([]*evInput, 0, 1024),
	}
	lp.cond = sync.NewCond(&lp.queueMu)
	lp.start()
	return lp, now.UnixMilli()
}

func NewMigrateEvLoop(resp *imchatevloop.DoMigrateResp) *ChatEvLoop {
	relationMan := relationmanager.NewRelationManagerFromMigrage(resp)
	subscrberMan := subscribemanager.NewSubscrbieManagerFromMigrage(resp)

	lp := &ChatEvLoop{
		chatId:           resp.GroupId,
		ownerId:          resp.OwnerId,
		createdAt:        resp.CreatedAt,
		relationManager:  relationMan,
		subscribeManager: subscrberMan,
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

func (c *ChatEvLoop) start() {
	go func() {
		status := statusRunning
		for {
			fmt.Println("lock")
			c.queueMu.Lock()
			if len(c.queue) == 0 {
				c.cond.Wait()
				c.queueMu.Unlock()
				continue
			}
			requests := c.queue
			c.queue = make([]*evInput, 0, 1024)
			fmt.Println("unlock")
			c.queueMu.Unlock()

			var requests2 []*evInput

			// 处理前, 知道是running的状态
			for _, request := range requests {
				input := request.input
				switch spec := input.(type) {
				case *evInputUniversal:
					var err error
					if status == statusBeMoving {
						err = &errMigrating
					} else if status == statusStop {
						err = &errStop
					}
					if err != nil {
						request.output <- &EvOutputUniversal{
							err:    err,
							Output: nil,
						}
						continue
					}

					switch inputSpec := spec.Input.Input.(type) {
					case *evloopio.UniversalGroupEvloopInput_AlterGroupMember:
						err := c.relationManager.UpdateRelation(inputSpec.AlterGroupMember.Relation)
						if err != nil {
							log.Printf("update relation err: %v\n", err)
						}
						// 然后需要将以前的订阅者全部删除
						c.subscribeManager.RemoveOldSubs(inputSpec.AlterGroupMember.Relation.UserId, inputSpec.AlterGroupMember.Relation.RelationVersion)
						request.output <- &EvOutputUniversal{
							err: nil,
							Output: &evloopio.UniversalGroupEvloopOutput{
								Output: &evloopio.UniversalGroupEvloopOutput_AlterGroupMember{
									AlterGroupMember: &evloopio.AlterGroupMemberResponse{
										Code:            evloopio.AlterGroupMemberResponse_OK,
										RelationVersion: spec.Input.GetAlterGroupMember().Relation.RelationVersion,
										CurrentSeqId:    c.seqId,
										ChangeAt:        inputSpec.AlterGroupMember.Relation.ChangeAt,
									},
								},
							},
						}
					// 暂时不支持解散群聊
					// case *evloopio.UniversalGroupEvloopInput_DisbandGroup:
					case *evloopio.UniversalGroupEvloopInput_SendMessage:
						// 先定序
						c.seqId++
						msgSeq := c.seqId
						// TODO 需要存消息
						// store(inputSpec, msgSeq)
						// 广播seq id
						now := time.Now()
						startBroadcastSeqId(c.subscribeManager.SnapshotAllSubs(), c.chatId, msgSeq, now)
					case *evloopio.UniversalGroupEvloopInput_SubscribeGroup:
						// 判断relation是否ok
						relationVersion, ok := c.relationManager.CanSubscribe(inputSpec.SubscribeGroup.Session.Username)
						if !ok {
							request.output <- &EvOutputUniversal{
								err: nil,
								Output: &evloopio.UniversalGroupEvloopOutput{
									Output: &evloopio.UniversalGroupEvloopOutput_SubscribeGroup{
										SubscribeGroup: &evloopio.SubscribeGroupResponse{
											Code: evloopio.SubscribeGroupResponse_NO_PERMISSION,
										},
									},
								},
							}
							go func() {
								push.SeqResp(&push.SubRespInput{
									GwAddrPort:      inputSpec.SubscribeGroup.Session.GwAdvertiseAddrport,
									SessionId:       inputSpec.SubscribeGroup.Session.SessionId,
									GroupId:         c.chatId,
									SubOk:           false,
									RelationVersion: 0,
									SeqId:           -1,
								})
							}()
							break
						}

						// relation ok, 注册进去
						c.subscribeManager.SessionSub(inputSpec.SubscribeGroup.Session.Username, inputSpec.SubscribeGroup.Session.SessionId,
							relationVersion, inputSpec.SubscribeGroup.GetSession())
						request.output <- &EvOutputUniversal{
							err: nil,
							Output: &evloopio.UniversalGroupEvloopOutput{
								Output: &evloopio.UniversalGroupEvloopOutput_SubscribeGroup{
									SubscribeGroup: &evloopio.SubscribeGroupResponse{
										Code:       evloopio.SubscribeGroupResponse_OK,
										RelationId: relationVersion,
										MaxSeqId:   c.seqId,
									},
								},
							},
						}
						go func() {
							push.SeqResp(&push.SubRespInput{
								GwAddrPort:      inputSpec.SubscribeGroup.Session.GwAdvertiseAddrport,
								SessionId:       inputSpec.SubscribeGroup.Session.SessionId,
								GroupId:         c.chatId,
								SubOk:           true,
								RelationVersion: relationVersion,
								SeqId:           c.seqId,
							})
						}()
					}
					fmt.Println("universal")
				case *evInputMove:
					if status == statusRunning || status == statusBeMoving {
						c.queueMu.Lock()
						c.status = statusBeMoving
						c.queueMu.Unlock()
						status = statusBeMoving
						request.output <- &EvOutputMove{
							err:       nil,
							OwnerId:   c.ownerId,
							SeqId:     c.seqId,
							Subs:      c.subscribeManager.GetSubscibers(),
							Relations: c.relationManager.GetRelations(),
						}
						fmt.Println("move")
					}
				case *evInputStop:
					if status != statusBeMoving {
						request.output <- &EvOutputStop{
							err: &errRunning,
						}
					}
					status = statusStop
					c.queueMu.Lock()
					c.status = statusStop
					requests2 = c.queue
					c.queue = nil
					c.queueMu.Unlock()
					request.output <- &EvOutputStop{
						err: nil,
					}
					fmt.Println("stop")
				}
			}

			if status == statusStop {
				fmt.Println(requests2)
				for _, request := range requests2 {
					input := request.input
					switch input.(type) {
					case *evInputUniversal:
						request.output <- &EvOutputUniversal{
							err:    &errStop,
							Output: nil,
						}
						fmt.Println("universal")
					case *evInputMove:
						c.queueMu.Lock()
						c.status = statusBeMoving
						c.queueMu.Unlock()
						status = statusBeMoving
						request.output <- &EvOutputMove{
							err: &errStop,
						}
						fmt.Println("move")
					case *evInputStop:
						request.output <- &EvOutputStop{
							err: nil,
						}
						fmt.Println("stop")
					}
				}
				fmt.Println("evloop exit")
				return
			}
		}
	}()
}

func startBroadcastSeqId(subs []*base.SessionEntry, groupId string, seqId int64, sendAt time.Time) {
	// 先拷贝subs, 再异步下行seqid

	go func() {
		for _, sub := range subs {
			push.SeqNotify(&push.SeqNotifyInput{
				GwAddrPort: sub.GwAdvertiseAddrport,
				SessionId:  sub.SessionId,
				SeqId:      seqId,
				GroupId:    groupId,
				SendAt:     sendAt.UnixMilli(),
			})
		}
	}()
}
