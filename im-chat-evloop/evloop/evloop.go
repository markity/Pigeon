package evloop

import (
	"fmt"
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

type SubscriberInfo struct {
	GwAddrPort               string
	SessionId                string
	SubscribeRelationVersion int64
}

type MemberInfo struct {
	UserId          string
	RelationVersion int64
	Status          base.RelationStatus
	ChangeType      base.RelationChangeType
}

type ChatEvLoop struct {
	chatId  string
	ownerId string

	seqId int64

	// 全量群员信息, key是userId
	relations map[string]*MemberInfo

	// 订阅者信息, key是userId
	subscribers map[string][]*SubscriberInfo

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

type evOutputMove struct {
	Err       error // queueMessageError, 可能是stop
	OwnerId   string
	SeqId     int64
	Subs      map[string][]*SubscriberInfo
	Relations map[string]*MemberInfo
}

type evOutputStop struct {
	Err error // queueMessageError, 可能是running
}

type evOutputUniversal struct {
	Err    error // queueMessageError, 可能是stop或moving
	Output *evloopio.UniversalGroupEvloopOutput
}

type NewChatEvLoopInput struct {
	ChatId string
	// 群主id
	OwnerId string
}

func NewChatEvLoopAndStart(in *NewChatEvLoopInput) *ChatEvLoop {
	mems := make(map[string]*MemberInfo)
	mems[in.OwnerId] = &MemberInfo{
		UserId:          in.OwnerId,
		RelationVersion: 1,
		Status:          base.RelationStatus_RELATION_STATUS_OWNER,
	}
	lp := &ChatEvLoop{
		chatId:      in.ChatId,
		ownerId:     in.OwnerId,
		relations:   mems,
		subscribers: make(map[string][]*SubscriberInfo),
		status:      statusRunning,
		queue:       make([]*evInput, 0, 1024),
	}
	lp.cond = sync.NewCond(&lp.queueMu)
	lp.start()
	return lp
}

func NewMigrateEvLoop(resp *imchatevloop.DoMigrateResp) *ChatEvLoop {
	members := make(map[string]*MemberInfo)
	for _, v := range resp.Relations {
		members[v.MemberId] = &MemberInfo{
			UserId:          v.MemberId,
			RelationVersion: v.RelationVersion,
			Status:          v.Status,
			ChangeType:      v.ChangeType,
		}
	}

	subscriberInfos := make(map[string][]*SubscriberInfo)
	for k, v := range resp.Subscribers {
		subscriberInfos[k] = make([]*SubscriberInfo, 0)
		for _, ent := range v.SessionEntries {
			subscriberInfos[k] = append(subscriberInfos[k], &SubscriberInfo{
				GwAddrPort:               ent.GwAddrport,
				SessionId:                ent.SessionId,
				SubscribeRelationVersion: ent.OnSubRelationVersion,
			})
		}
	}

	lp := &ChatEvLoop{
		chatId:      resp.GroupId,
		ownerId:     resp.OwnerId,
		relations:   members,
		subscribers: subscriberInfos,
		status:      statusRunning,
	}
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
	*evloopio.UniversalGroupEvloopOutput, error) {
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
	out := output.(*evOutputUniversal)
	fmt.Println("queue message out")
	return out.Output, out.Err
}

type MoveOutput struct {
	OwnerId   string
	SeqId     int64
	Subs      map[string][]*SubscriberInfo
	Relations map[string]*MemberInfo
}

func (c *ChatEvLoop) Move() (*MoveOutput, error) {
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
	out := output.(*evOutputMove)
	mo := &MoveOutput{
		OwnerId:   out.OwnerId,
		SeqId:     out.SeqId,
		Subs:      out.Subs,
		Relations: out.Relations,
	}
	if out.Err != nil {
		return nil, out.Err
	}
	return mo, nil
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
	out := output.(*evOutputStop)
	return out.Err
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
						request.output <- &evOutputUniversal{
							Err:    err,
							Output: nil,
						}
						continue
					}

					switch inputSpec := spec.Input.Input.(type) {
					case *evloopio.UniversalGroupEvloopInput_AlterGroupMember:
						c.relations[inputSpec.AlterGroupMember.GroupId] = &MemberInfo{
							UserId:          inputSpec.AlterGroupMember.MemberId,
							RelationVersion: inputSpec.AlterGroupMember.RelationVersion,
							Status:          inputSpec.AlterGroupMember.Status,
							ChangeType:      inputSpec.AlterGroupMember.ChangeType,
						}
						// 然后需要将以前的订阅者全部删除
						subs := c.subscribers[inputSpec.AlterGroupMember.MemberId]
						if len(subs) != 0 && subs[0].SubscribeRelationVersion < inputSpec.AlterGroupMember.RelationVersion {
							c.subscribers[inputSpec.AlterGroupMember.MemberId] = make([]*SubscriberInfo, 0)
						}
						request.output <- &evOutputUniversal{
							Err: nil,
							Output: &evloopio.UniversalGroupEvloopOutput{
								Output: &evloopio.UniversalGroupEvloopOutput_AlterGroupMember{
									AlterGroupMember: &evloopio.AlterGroupMemberResponse{
										Code:            evloopio.AlterGroupMemberResponse_OK,
										RelationVersion: spec.Input.GetAlterGroupMember().RelationVersion,
										CurrentSeqId:    c.seqId,
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
						startBroadcastSeqId(c.subscribers, inputSpec.SendMessage.GroupId, msgSeq, now)
					case *evloopio.UniversalGroupEvloopInput_SubscribeGroup:
						// 判断relation是否ok
						relation := c.relations[inputSpec.SubscribeGroup.UserId]
						if relation == nil || relation.Status ==
							base.RelationStatus_RELATION_STATUS_NOT_IN_GROUP {
							request.output <- &evOutputUniversal{
								Err: nil,
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
									GwAddrPort:      inputSpec.SubscribeGroup.GwAdvertiseAddrPort,
									SessionId:       inputSpec.SubscribeGroup.SessionId,
									GroupId:         inputSpec.SubscribeGroup.GroupId,
									SubOk:           false,
									RelationVersion: 0,
									SeqId:           -1,
								})
							}()
							break
						}

						// relation ok, 注册进去
						c.subscribers[inputSpec.SubscribeGroup.UserId] = append(
							c.subscribers[inputSpec.SubscribeGroup.UserId],
							&SubscriberInfo{
								GwAddrPort:               inputSpec.SubscribeGroup.GwAdvertiseAddrPort,
								SessionId:                inputSpec.SubscribeGroup.SessionId,
								SubscribeRelationVersion: relation.RelationVersion,
							})
						request.output <- &evOutputUniversal{
							Err: nil,
							Output: &evloopio.UniversalGroupEvloopOutput{
								Output: &evloopio.UniversalGroupEvloopOutput_SubscribeGroup{
									SubscribeGroup: &evloopio.SubscribeGroupResponse{
										Code:       evloopio.SubscribeGroupResponse_OK,
										RelationId: relation.RelationVersion,
										MaxSeqId:   c.seqId,
									},
								},
							},
						}
						go func() {
							push.SeqResp(&push.SubRespInput{
								GwAddrPort:      inputSpec.SubscribeGroup.GwAdvertiseAddrPort,
								SessionId:       inputSpec.SubscribeGroup.SessionId,
								GroupId:         inputSpec.SubscribeGroup.GroupId,
								SubOk:           true,
								RelationVersion: relation.RelationVersion,
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
						request.output <- &evOutputMove{
							Err:       nil,
							OwnerId:   c.ownerId,
							SeqId:     c.seqId,
							Subs:      c.subscribers,
							Relations: c.relations,
						}
						fmt.Println("move")
					}
				case *evInputStop:
					if status != statusBeMoving {
						request.output <- &evOutputStop{
							Err: &errRunning,
						}
					}
					status = statusStop
					c.queueMu.Lock()
					c.status = statusStop
					requests2 = c.queue
					c.queue = nil
					c.queueMu.Unlock()
					request.output <- &evOutputStop{
						Err: nil,
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
						request.output <- &evOutputUniversal{
							Err:    &errStop,
							Output: nil,
						}
						fmt.Println("universal")
					case *evInputMove:
						c.queueMu.Lock()
						c.status = statusBeMoving
						c.queueMu.Unlock()
						status = statusBeMoving
						request.output <- &evOutputMove{
							Err: &errStop,
						}
						fmt.Println("move")
					case *evInputStop:
						request.output <- &evOutputStop{
							Err: nil,
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

func startBroadcastSeqId(subs map[string][]*SubscriberInfo, groupId string, seqId int64, sendAt time.Time) {
	// 先拷贝subs, 再异步下行seqid
	newSubs := make([]*SubscriberInfo, 0)
	for _, v := range subs {
		for _, ent := range v {
			s := *ent
			newSubs = append(newSubs, &s)
		}
	}

	go func() {
		for _, sub := range newSubs {
			push.SeqNotify(&push.SeqNotifyInput{
				GwAddrPort: sub.GwAddrPort,
				SessionId:  sub.SessionId,
				SeqId:      seqId,
				GroupId:    groupId,
				SendAt:     sendAt.UnixMilli(),
			})
		}
	}()
}
