package evloop

import (
	"fmt"
	"log"
	"time"

	"pigeon/im-chat-evloop/bizpush"
	"pigeon/im-chat-evloop/db"
	"pigeon/im-chat-evloop/db/model"
	"pigeon/kitex_gen/service/evloopio"
)

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
					// TODO: 暂时不支持解散群聊
					// case *evloopio.UniversalGroupEvloopInput_DisbandGroup:
					case *evloopio.UniversalGroupEvloopInput_SendMessage:
						// 1.做幂等检查
						key := fmt.Sprintf("%v-%v", inputSpec.SendMessage.Session.Username, inputSpec.SendMessage.IdempotentKey)
						if inputSpec.SendMessage.CheckIdempotent {
							m, err := db.GetMessageByIdempotentKey(c.db, key)
							if err != nil {
								fmt.Printf("failed to GetMessageByIdempotentKey: %v\n", err)
								break
							}
							if m != nil {
								go c.bPush.SendMessageResp(&bizpush.SendMessageRespInput{
									Session:  inputSpec.SendMessage.Session,
									EchoCode: inputSpec.SendMessage.EchoCode,
									Code:     evloopio.SendMessageResponse_OK,
									SeqId:    m.SeqId,
								})
								break
							}
						}

						// 2.权限检查
						version, canSub := c.relationManager.CanSubscribe(inputSpec.SendMessage.Session.Username)
						if !canSub {
							go c.bPush.SendMessageResp(&bizpush.SendMessageRespInput{
								Session:         inputSpec.SendMessage.Session,
								EchoCode:        inputSpec.SendMessage.EchoCode,
								RelationVersion: version,
								Code:            evloopio.SendMessageResponse_NO_PERMISSION,
								SeqId:           0,
							})
						} else {
							// 先定序
							c.seqId++
							msgSeq := c.seqId
							now := time.Now()
							db.InsertMessage(c.db, &model.MessageModel{
								OwnerId:       inputSpec.SendMessage.Session.Username,
								GroupId:       c.chatId,
								SeqId:         msgSeq,
								MsgId:         c.idGen.Generate().Base64(),
								Data:          string(inputSpec.SendMessage.MessageData),
								IdempotentKey: key,
								CreatedAt:     now.UnixMilli(),
							})
							// 广播seq id
							go c.bPush.SendMessageResp(&bizpush.SendMessageRespInput{
								Session:  inputSpec.SendMessage.Session,
								EchoCode: inputSpec.SendMessage.EchoCode,
								Code:     evloopio.SendMessageResponse_OK,
								SeqId:    msgSeq,
							})
							c.bPush.StartBroadcastSeqId(c.subscribeManager.SnapshotAllSubs(), c.chatId, msgSeq, now)
						}
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
								c.bPush.SubResp(&bizpush.SubRespInput{
									Session:         inputSpec.SubscribeGroup.Session,
									EchoCode:        inputSpec.SubscribeGroup.EchoCode,
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
							c.bPush.SubResp(&bizpush.SubRespInput{
								Session:         inputSpec.SubscribeGroup.Session,
								EchoCode:        inputSpec.SubscribeGroup.EchoCode,
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
