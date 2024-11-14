package connkeeper

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"pigeon/client/conn_keeper/events"
	bizprotocol "pigeon/common/biz_protocol"
	"pigeon/common/protocol"
	pushprotocol "pigeon/common/push_protocol"
	unboundedqueue "pigeon/common/unbounded_queue"
	"pigeon/kitex_gen/service/base"
)

type ConnKeeper struct {
	// 输入参数
	conn   net.Conn
	config Config

	startOnce sync.Once

	hbTicker       *time.Ticker
	sendMsgUq      *unboundedqueue.UnboundedQueen
	timeoutTimer   *time.Timer
	ioErrChan      chan error
	readerRecvChan chan interface{}
	eventChan      chan interface{}
}

type Config struct {
	// TimerConfig
	HbInterval time.Duration
	HbTimeout  time.Duration

	// UserConfig
	Username   string
	Password   string
	DeviceDesc string
}

func (ck *ConnKeeper) goWriter() {
	for {
		data := ck.sendMsgUq.PopBlock().([]byte)
		_, err := ck.conn.Write(data)
		if err != nil {
			ck.ioErrChan <- err
			return
		}
	}
}

func (ck *ConnKeeper) goReader() {
	lengthPrefix := make([]byte, 4)
	for {
		_, err := io.ReadFull(ck.conn, lengthPrefix)
		if err != nil {
			ck.conn.Close()
			ck.ioErrChan <- err
			return
		}
		l := binary.LittleEndian.Uint32(lengthPrefix)
		data := make([]byte, l)
		_, err = io.ReadFull(ck.conn, data)
		if err != nil {
			ck.conn.Close()
			ck.ioErrChan <- err
			return
		}
		d, err := protocol.ParseS2CPacket(data)
		if err != nil {
			ck.conn.Close()
			ck.ioErrChan <- err
			return
		}
		ck.readerRecvChan <- d
	}
}

func (ck *ConnKeeper) resetTimer() {
	if ck.timeoutTimer != nil {
		ck.timeoutTimer.Stop()
	}
	ck.timeoutTimer = time.NewTimer(ck.config.HbTimeout)
}

func (ck *ConnKeeper) sendHb() {
	ck.sendMsgUq.Push(protocol.PackData(protocol.MustEncodePacket(&protocol.HeartbeatPacket{})))
}

func (ck *ConnKeeper) goStart() {
	// 第一件事是, 先发登录包
	loginPkt := protocol.C2SLoginPacket{
		Username:   ck.config.Username,
		Password:   ck.config.Password,
		DeviceDesc: ck.config.DeviceDesc,
	}
	loginPkt.SetEchoCode("login-pkt")
	loginPktBs := protocol.PackData(protocol.MustEncodePacket(&loginPkt))
	ck.sendMsgUq.Push(loginPktBs)

	stateNow := "waiting-login-resp"
	// waitLoginBufferPackets := make([]interface{}, 0)

	// pulling-full状态下使用的变量
	var sessionsFull []*protocol.DeviceSessionEntry
	fmt.Println(sessionsFull)

	// string是群聊id
	var subRecv int = 0
	var relationAndSeqFull map[string]*events.RelationAndSeqEntry
	var applyFull []*base.ApplyEntry

	timerChan := ck.hbTicker.C
	ck.sendHb()
	for {
		select {
		case <-ck.timeoutTimer.C:
			ck.conn.Close()
		case <-timerChan:
			ck.sendHb()
		case <-ck.ioErrChan:
			ck.eventChan <- &events.EventClose{
				CloseReason: events.CloseReasonUserNetwork,
			}
			return
		case packet := <-ck.readerRecvChan:
			// timer重启定时器
			if _, ok := packet.(*protocol.HeartbeatPacket); ok {
				ck.resetTimer()
				continue
			}

			// 被踢
			if kickPkt, ok := packet.(*protocol.S2COtherDeviceKickNotifyPacket); ok {
				ck.eventChan <- &events.EventClose{
					CloseReason:     events.CloseReasonOtherKick,
					FromSession:     kickPkt.FromSessionId,
					FromSessionDesc: kickPkt.FromSessionDesc,
				}
				return
			}

			// login resp packet
			if loginRespPacket, ok := packet.(*protocol.S2CLoginRespPacket); ok {
				if stateNow != "waiting-login-resp" {
					panic(stateNow)
				}

				if loginRespPacket.Code != protocol.LoginRespCodeSuccess {
					ck.eventChan <- &events.EventClose{
						CloseReason: events.CloseReasonLoginFail,
					}
					return
				}

				// 登录完成, 拉全量
				stateNow = "pulling-full"

				// 多设备管理全量
				sessionsFull = loginRespPacket.Sessions

				// relations全量请求
				pullRelationsPkt := protocol.C2SBizMessagePacket{
					BizType: (&bizprotocol.BizPullRelations{}).String(),
					Data:    &bizprotocol.BizPullRelations{},
				}
				pullRelationsPkt.SetEchoCode("pull-relation")
				ck.sendMsgUq.Push(protocol.PackData(protocol.MustEncodePacket(&pullRelationsPkt)))

				// apply全量请求
				applyPkt := protocol.C2SBizMessagePacket{
					BizType: (&bizprotocol.BizPullApply{}).String(),
					Data:    &bizprotocol.BizPullApply{},
				}
				applyPkt.SetEchoCode("pull-apply")
				fmt.Println(string(protocol.MustEncodePacket(&applyPkt)))
				ck.sendMsgUq.Push(protocol.PackData(protocol.MustEncodePacket(&applyPkt)))
			}

			// biz packet
			if pushPacket, ok := packet.(*protocol.S2CPushMessagePacket); ok {
				pushPkt, ec, err := pushprotocol.ParsePush(pushPacket)
				fmt.Printf("%T %v\n", pushPkt, ec)
				if err != nil {
					ck.conn.Close()
					ck.eventChan <- &events.EventClose{
						CloseReason: events.CloseReasonUserNetwork,
					}
					return
				}
				switch pkt := pushPkt.(type) {
				case (*pushprotocol.FetchAllRelationsResp):
					if stateNow != "pulling-full" {
						ck.conn.Close()
						ck.eventChan <- &events.EventClose{
							CloseReason: events.CloseReasonUserNetwork,
						}
						return
					}

					relationAndSeqFull = make(map[string]*events.RelationAndSeqEntry)
					for _, v := range pkt.Relations {
						relationAndSeqFull[v.GroupId] = &events.RelationAndSeqEntry{
							RelationEntry: v,
							SeqId:         0,
						}
						sendPkt := &protocol.C2SBizMessagePacket{
							BizType: (&bizprotocol.BizSub{}).String(),
							Data: &bizprotocol.BizSub{
								GroupId: v.GroupId,
							},
						}
						sendPkt.SetEchoCode("sub-" + v.GroupId)
						ck.sendMsgUq.Push(protocol.PackData(protocol.MustEncodePacket(sendPkt)))
					}
					if relationAndSeqFull != nil && len(relationAndSeqFull) == subRecv && applyFull != nil {
						stateNow = "ruinng"
						fmt.Println("running")
					}
				case (*pushprotocol.FetchAllApplicationsResp):
					if stateNow != "pulling-full" {
						ck.conn.Close()
						ck.eventChan <- &events.EventClose{
							CloseReason: events.CloseReasonUserNetwork,
						}
						return
					}

					applyFull = pkt.Applications
					if relationAndSeqFull != nil && len(relationAndSeqFull) == subRecv && applyFull != nil {
						stateNow = "running"
						fmt.Println("running")
					}
				case (*pushprotocol.SubResp):
					if stateNow == "pulling-full" {
						subRecv++
						if pkt.RelationVersion >= relationAndSeqFull[pkt.GroupId].RelationVersion {
							relationAndSeqFull[pkt.GroupId].SeqId = pkt.SeqId
							relationAndSeqFull[pkt.GroupId].RelationVersion = pkt.RelationVersion
						}
						if relationAndSeqFull != nil && len(relationAndSeqFull) == subRecv && applyFull != nil {
							stateNow = "running"
							fmt.Println("running")
						}
					} else if stateNow == "running" {

					} else {
						ck.conn.Close()
						ck.eventChan <- &events.EventClose{
							CloseReason: events.CloseReasonUserNetwork,
						}
						return
					}
				}
			}
		}
	}
}

func (ck *ConnKeeper) Start() chan interface{} {
	ck.startOnce.Do(func() {
		ck.hbTicker = time.NewTicker(ck.config.HbInterval)
		ck.resetTimer()
		go ck.goWriter()
		go ck.goReader()
		go ck.goStart()
	})
	return ck.eventChan
}

// 简单的关闭conn就能实现整体关闭
func (ck *ConnKeeper) Close() {
	ck.conn.Close()
}

func NewConnKeeper(conn net.Conn, cfg Config) *ConnKeeper {
	return &ConnKeeper{
		conn: conn,

		config: cfg,

		sendMsgUq:      unboundedqueue.NewUnboundedQueen(),
		readerRecvChan: make(chan interface{}, 1024),
		ioErrChan:      make(chan error, 2),
		eventChan:      make(chan interface{}, 1024),
	}
}
