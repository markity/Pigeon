package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"pigeon/im-gateway/protocol"
	"strings"
	"sync"
	"time"

	interactive "github.com/markity/Interactive-Console"
)

// ah on/off
type cmdAutoHeartbeat struct {
	On *bool
}

// hb
type cmdHeartbeat struct{}

// login username password echocode
type cmdLogin struct {
	Username string
	Password string
	EchoCode string
}

type cmdLogout struct {
	EchoCode string
}

type cmdKickOtherSession struct {
	SessionId string
	EchoCode  string
}

type exitCmd struct{}

type helpCmd struct {
}

func parseCommand(line string) interface{} {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil
	}

	cmds := strings.Split(line, " ")
	switch cmds[0] {
	case "ah":
		if len(cmds) == 1 {
			return &cmdAutoHeartbeat{On: nil}
		}
		if len(cmds) == 2 && (cmds[1] == "on" || cmds[1] == "off") {
			b := cmds[1] == "on"
			return &cmdAutoHeartbeat{On: &b}
		}
		return nil
	case "hb":
		return &cmdHeartbeat{}
	case "login":
		if len(cmds) == 4 {
			return &cmdLogin{Username: cmds[1], Password: cmds[2], EchoCode: cmds[3]}
		}
	case "logout":
		if len(cmds) == 2 {
			return &cmdLogout{EchoCode: cmds[1]}
		}
	case "kick":
		if len(cmds) == 3 {
			return &cmdKickOtherSession{SessionId: cmds[1], EchoCode: cmds[2]}
		}
	case "exit":
		if len(cmds) == 1 {
			return &exitCmd{}
		}
	case "help":
		if len(cmds) == 1 {
			return &helpCmd{}
		}
	}
	return nil

}
func main() {
	promptStyle := interactive.GetDefaultSytleAttr()
	promptStyle.Bold = true
	promptStyle.Foreground = interactive.ColorDarkGreen
	win := interactive.Run(interactive.Config{
		Prompt:               '>',
		PromptStyle:          promptStyle,
		BlockInputAfterRun:   false,
		BlockInputAfterEnter: true,
		TraceAfterRun:        true,
		EventHandleMask:      interactive.EventMaskKeyCtrlC,
	})

	win.SendLineBack("connecting to server...")
	win.SetBlockInput(true)

	var connOkChan = make(chan struct{}, 1)

	var conn net.Conn
	go func() {
		for {
			var err error
			conn, err = net.Dial("tcp", "127.0.0.1:3501")
			if err != nil {
				win.SendLineBack("fail, retrying: " + err.Error())
				time.Sleep(time.Second)
				continue
			}
			connOkChan <- struct{}{}
			break
		}
	}()

	autoHeartbeat := false

	uq := NewUnboundedQueen()
	cmdChan := win.GetCmdChan()
	eventChan := win.GetEventChan()
	var heartbeatTikcer *time.Ticker
	var heartbeatChan <-chan time.Time
	connErrChan := make(chan error, 2)
	packetRecvChan := make(chan interface{})
	for {
		select {
		case <-connOkChan:
			win.SendLineBack("connected!")
			// 负责发消息
			go func() {
				for {
					_, err := conn.Write(uq.PopBlock().([]byte))
					if err != nil {
						connErrChan <- err
						return
					}
				}
			}()
			// 负责收消息
			go func() {
				lengthData := [4]byte{}
				for {
					_, err := io.ReadFull(conn, lengthData[:])
					if err != nil {
						connErrChan <- err
						return
					}
					length := binary.LittleEndian.Uint32(lengthData[:])
					bs := make([]byte, length)
					_, err = io.ReadFull(conn, bs)
					if err != nil {
						connErrChan <- err
						return
					}

					pack, ok := protocol.ParseS2CPacket(bs)
					if !ok {
						connErrChan <- errors.New("parse packet error")
						conn.Close()
						return
					}
					packetRecvChan <- pack
				}
			}()
			win.SetBlockInput(false)
		case cmd := <-cmdChan:
			cmdIface := parseCommand(cmd)
			switch c := cmdIface.(type) {
			case *cmdAutoHeartbeat:
				originHeartbeat := autoHeartbeat
				if c.On != nil {
					autoHeartbeat = *cmdIface.(*cmdAutoHeartbeat).On
				} else {
					autoHeartbeat = !autoHeartbeat
				}
				if originHeartbeat != autoHeartbeat {
					if !autoHeartbeat {
						win.SendLineBack("auto heartbeat off")
						heartbeatTikcer.Stop()
						heartbeatTikcer = nil
					} else {
						win.SendLineBack("auto heartbeat on")
						uq.Push(protocol.PackData(protocol.MustEncodePacket(&protocol.HeartbeatPacket{})))
						win.SendLineBack("auto send heartbeat")
						heartbeatTikcer = time.NewTicker(time.Second * 1)
						heartbeatChan = heartbeatTikcer.C
					}
				} else {
					win.SendLineBack("nothing happened")
				}
			case *cmdHeartbeat:
				win.SendLineBack("send heartbeat")
				uq.Push(protocol.PackData(protocol.MustEncodePacket(&protocol.HeartbeatPacket{})))
			case *cmdLogin:
				username := c.Username
				password := c.Password
				echoCode := c.EchoCode
				win.SendLineBack("send login packet, username: " + username + ", password: " + password + ", echoCode: " + echoCode)
				var p = &protocol.C2SLoginPacket{
					Username: username,
					Password: password,
				}
				p.SetEchoCode(echoCode)
				uq.Push(protocol.PackData(protocol.MustEncodePacket(p)))
			case *cmdLogout:
				echoCode := c.EchoCode
				win.SendLineBack("send logout packet, echoCode: " + echoCode)
				var p = &protocol.C2SLogoutPacket{}
				p.SetEchoCode(echoCode)
				uq.Push(protocol.PackData(protocol.MustEncodePacket(p)))
			case *helpCmd:
				win.SendLineBack("help:")
				win.SendLineBack("    ah on|off: auto send heartbeat per second")
				win.SendLineBack("    hb: send heartbeat immediately")
				win.SendLineBack("    login username password echoCode: send login packet")
				win.SendLineBack("    logout echoCode: send logout packet")
				win.SendLineBack("    kick sessionId echoCode: send kick other device packet")
				win.SendLineBack("    exit: exit")
				win.SendLineBack("    help: show this message")
			case *exitCmd:
				win.Stop()
				return
			case *cmdKickOtherSession:
				sessionId := c.SessionId
				echoCode := c.EchoCode
				win.SendLineBack("send kick other session packet, sessionId: " + sessionId + ", echoCode: " + echoCode)
				var p = &protocol.C2SKickOhterDevicePacket{
					SessionId: sessionId,
				}
				p.SetEchoCode(echoCode)
				uq.Push(protocol.PackData(protocol.MustEncodePacket(p)))
			default:
				win.SendLineBack("unknown command")
			}
			win.SetBlockInput(false)
		case ev := <-eventChan:
			switch ev.(type) {
			case *interactive.EventKeyCtrlC:
				win.Stop()
				return
			}
		case <-heartbeatChan:
			uq.Push(protocol.PackData(protocol.MustEncodePacket(&protocol.HeartbeatPacket{})))
			win.SendLineBack("auto send heartbeat")
		case err := <-connErrChan:
			win.SendLineBack("conn lost: " + err.Error())
			win.SetBlockInput(true)
			win.Stop()
			fmt.Println("conn lost: " + err.Error())
			return
		case packet := <-packetRecvChan:
			win.SendLineBack(fmt.Sprint(packet))
		}
	}
}

type unboundedQueen struct {
	queen []interface{}
	l     sync.Mutex
	cond  *sync.Cond
}

func NewUnboundedQueen() *unboundedQueen {
	uq := unboundedQueen{}
	uq.queen = make([]interface{}, 0)
	uq.cond = sync.NewCond(&uq.l)
	return &uq
}

func (uq *unboundedQueen) Push(i interface{}) {
	uq.l.Lock()
	uq.queen = append(uq.queen, i)
	uq.l.Unlock()
	uq.cond.Broadcast()
}

func (uq *unboundedQueen) PopBlock() interface{} {
	uq.l.Lock()
	for {
		if len(uq.queen) == 0 {
			uq.cond.Wait()
		} else {
			ret := uq.queen[0]
			uq.queen = uq.queen[1:]
			uq.l.Unlock()
			return ret
		}
	}
}
