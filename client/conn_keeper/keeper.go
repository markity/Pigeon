package connkeeper

import (
	"encoding/binary"
	"io"
	"net"
	"pigeon/client/conn_keeper/events"
	bizprotocol "pigeon/common/biz_protocol"
	"pigeon/common/protocol"
	unboundedqueue "pigeon/common/unbounded_queue"
	"sync"
	"time"
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
		d, err := protocol.ParseC2SPacket(data)
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
	waitLoginBufferPackets := make([]interface{}, 0)

	var sessionsFull []*protocol.DeviceSessionEntry
	var relationAndSeqFull map[string]*events.RelationAndSeqEntry

	timerChan := ck.hbTicker.C
	for {
		select {
		case <-ck.timeoutTimer.C:
			ck.conn.Close()
		case <-timerChan:
			ck.sendMsgUq.Push(protocol.PackData(protocol.MustEncodePacket(&protocol.HeartbeatPacket{})))
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
			if kickPkt, ok := packet.(*protocol.S2COtherDeviceKickNotifyPacket); ok {
				ck.eventChan <- &events.EventClose{
					CloseReason:     events.CloseReasonOtherKick,
					FromSession:     kickPkt.FromSessionId,
					FromSessionDesc: kickPkt.FromSessionDesc,
				}
				return
			}
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
				pullRelationsPkt := protocol.C2SBizMessagePacket{
					BizType: (&bizprotocol.BizPullRelations{}).String(),
					Data:    &bizprotocol.BizPullRelations{},
				}
				pullRelationsPkt.SetEchoCode("pull-relation")
				ck.sendMsgUq.Push(protocol.PackData(protocol.MustEncodePacket(&protocol.C2SBizMessagePacket{})))
			}
			if pushPkt, ok := packet.(*protocol.S2CPushMessagePacket); ok {
	
			}

			switch stateNow {
			case "waiting-login-resp":
				switch respPacket := packet.(type) {
				case *protocol.S2CLoginRespPacket:

					// 登录完成, 拉全量
					stateNow = "pulling-full"
					// 多设备管理全量
					sessionsFull = respPacket.Sessions
					// 关系管理全量
					pullRelationsPkt := protocol.C2SBizMessagePacket{
						BizType: (&bizprotocol.BizPullRelations{}).String(),
						Data:    &bizprotocol.BizPullRelations{},
					}
					pullRelationsPkt.SetEchoCode("pull-relation")
					ck.sendMsgUq.Push(protocol.PackData(protocol.MustEncodePacket(&protocol.C2SBizMessagePacket{})))
				default:
					waitLoginBufferPackets = append(waitLoginBufferPackets, respPacket)
				}
			case "pull-relation":
				// 等着push下发
				switch respPacket := packet.(type) {
				case *protocol.S2CPushMessagePacket:
					if respPacket.Code != protocol.LoginRespCodeSuccess {
						ck.eventChan <- &events.EventClose{
							CloseReason: events.CloseReasonLoginFail,
						}
						return
					}

					// 登录完成, 拉全量
					stateNow = "pulling-full"
					// 多设备管理全量
					sessionsFull = respPacket.Sessions
					// 关系管理全量
					pullRelationsPkt := protocol.C2SBizMessagePacket{
						BizType: (&bizprotocol.BizPullRelations{}).String(),
						Data:    &bizprotocol.BizPullRelations{},
					}
					pullRelationsPkt.SetEchoCode("pull-relation")
					ck.sendMsgUq.Push(protocol.PackData(protocol.MustEncodePacket(&protocol.C2SBizMessagePacket{})))
				default:
					waitLoginBufferPackets = append(waitLoginBufferPackets, respPacket)
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
