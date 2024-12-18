package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"time"

	"pigeon/common/protocol"

	"github.com/google/uuid"
	interactive "github.com/markity/Interactive-Console"
)

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
	hideHeartbeat := false

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

					pack, err := protocol.ParseS2CPacket(bs)
					if err != nil {
						connErrChan <- errors.New("parse packet error: " + err.Error())
						conn.Close()
						return
					}
					packetRecvChan <- pack
				}
			}()
			win.SetBlockInput(false)
		case cmd := <-cmdChan:
			win.SendLineBack(cmd)
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
						if !hideHeartbeat {
							win.SendLineBack("auto send heartbeat")
						}
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
				deviceDesc := c.DeviceDesc
				echoCode := ""
				if c.EchoCode == nil {
					echoCode = uuid.NewString()
				} else {
					echoCode = *c.EchoCode
				}
				win.SendLineBack("send login packet, username: " + username + ", password: " + password + ", deviceDesc" + c.DeviceDesc + ", echoCode: " + echoCode)
				var p = &protocol.C2SLoginPacket{
					Username:   username,
					Password:   password,
					DeviceDesc: deviceDesc,
				}

				p.SetEchoCode(echoCode)
				uq.Push(protocol.PackData(protocol.MustEncodePacket(p)))
			case *cmdLogout:
				echoCode := ""
				if c.EchoCode == nil {
					echoCode = uuid.NewString()
				} else {
					echoCode = *c.EchoCode
				}
				win.SendLineBack("send logout packet, echoCode: " + echoCode)
				var p = &protocol.C2SLogoutPacket{}
				p.SetEchoCode(echoCode)
				uq.Push(protocol.PackData(protocol.MustEncodePacket(p)))
			case *helpCmd:
				win.SendLineBack("help:")
				win.SendLineBack("    ah <on|off>: auto send heartbeat per second")
				win.SendLineBack("    hb: send heartbeat immediately")
				win.SendLineBack("    login <username> <password> <device-desc> [echoCode]: send login packet")
				win.SendLineBack("    logout [echoCode]: send logout packet")
				win.SendLineBack("    kick <sessionId> [echoCode]: send kick other device packet")
				win.SendLineBack("    status: check login status")
				win.SendLineBack("    sendbiz <bizType> <data> [echoCode]")
				win.SendLineBack("    hidehb: hide or unhide heartbeat info")
				win.SendLineBack("    exit: exit")
				win.SendLineBack("    help: show this message")
			case *clearCmd:
				win.Clear()
			case *exitCmd:
				win.Stop()
				return
			case *cmdKickOtherSession:
				sessionId := c.SessionId
				echoCode := ""
				if c.EchoCode == nil {
					echoCode = uuid.NewString()
				} else {
					echoCode = *c.EchoCode
				}
				win.SendLineBack("send kick other session packet, sessionId: " + sessionId + ", echoCode: " + echoCode)
				var p = &protocol.C2SKickOtherDevicePacket{
					SessionId: sessionId,
				}
				p.SetEchoCode(echoCode)
				uq.Push(protocol.PackData(protocol.MustEncodePacket(p)))
			case *statusCmd:
				echoCode := ""
				if c.EchoCode == nil {
					echoCode = uuid.NewString()
				} else {
					echoCode = *c.EchoCode
				}
				var p = &protocol.C2SQueryStatusPacket{}
				win.SendLineBack("send status command packet, echoCode: " + echoCode)
				p.SetEchoCode(echoCode)
				uq.Push(protocol.PackData(protocol.MustEncodePacket(p)))
			case *hideHeartbeatInfo:
				if hideHeartbeat {
					win.SendLineBack("show heartbeat")
				} else {
					win.SendLineBack("hide heartbeat")
				}
				hideHeartbeat = !hideHeartbeat
			case *cmdSendBiz:
				echoCode := ""
				if c.EchoCode == nil {
					echoCode = uuid.NewString()
				} else {
					echoCode = *c.EchoCode
				}
				var m map[string]interface{}
				err := json.Unmarshal([]byte(c.Data), &m)
				if err != nil {
					win.SendLineBack("invalid json data: " + err.Error())
					break
				}
				win.SendLineBack("send send biz packet, echoCode: " + echoCode)
				var p = &protocol.C2SBizMessagePacket{
					BizType: c.Biz,
					Data:    m,
				}
				p.SetEchoCode(echoCode)
				uq.Push(protocol.PackData(protocol.MustEncodePacket(p)))
			case *emptyCmd:
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
			if !hideHeartbeat {
				win.SendLineBack("auto send heartbeat")
			}
		case err := <-connErrChan:
			win.SendLineBack("conn lost: " + err.Error())
			win.SetBlockInput(true)
			win.Stop()
			fmt.Println("conn lost: " + err.Error())
			return
		case packet := <-packetRecvChan:
			switch pkt := packet.(type) {
			case *protocol.HeartbeatPacket:
				if !hideHeartbeat {
					win.SendLineBack("recv: packet heartbeat")
				}
			case *protocol.S2CDeviceInfoBroadcastPacket:
				win.SendLineBack("recv: packet device info broadcast")
				version := pkt.Version
				devices := pkt.Sessions
				win.SendLineBack("    version: " + fmt.Sprint(version))
				win.SendLineBack("    devices: " + fmt.Sprint(len(devices)))
				for i, device := range devices {
					win.SendLineBack(fmt.Sprintf("        device[%v]", i))
					win.SendLineBack(fmt.Sprintf("            sessionId: %v", device.SessionId))
					win.SendLineBack(fmt.Sprintf("            loginAt: %v", device.LoginAt))
					win.SendLineBack(fmt.Sprintf("            deviceDesc: %v", device.DeviceDesc))
				}
			case *protocol.S2CKickOhterDeviceRespPacket:
				win.SendLineBack("recv: packet kick resp")
				success := pkt.Success
				ok := pkt.KickOK
				version := pkt.Version
				devices := pkt.Sessions
				ec := pkt.EchoCode()
				win.SendLineBack("    success: " + fmt.Sprint(success))
				win.SendLineBack("    kickOk: " + fmt.Sprint(ok))
				win.SendLineBack("    version: " + fmt.Sprint(version))
				win.SendLineBack("    echoCode: " + ec)
				win.SendLineBack("    devices: " + fmt.Sprint(len(devices)))
				for i, device := range devices {
					win.SendLineBack(fmt.Sprintf("        device[%v]", i))
					win.SendLineBack(fmt.Sprintf("            sessionId: %v", device.SessionId))
					win.SendLineBack(fmt.Sprintf("            loginAt: %v", device.LoginAt))
					win.SendLineBack(fmt.Sprintf("            deviceDesc: %v", device.DeviceDesc))
				}
			case *protocol.S2CLoginRespPacket:
				win.SendLineBack("recv: packet login resp")
				code := pkt.Code
				sessionId := pkt.SessionId
				version := pkt.Version
				devices := pkt.Sessions
				ec := pkt.EchoCode()
				win.SendLineBack("    code: " + fmt.Sprint(code))
				win.SendLineBack("    sessionId: " + sessionId)
				win.SendLineBack("    version: " + fmt.Sprint(version))
				win.SendLineBack("    echoCode: " + ec)
				win.SendLineBack("    devices: " + fmt.Sprint(len(devices)))
				for i, device := range devices {
					win.SendLineBack(fmt.Sprintf("        device[%v]", i))
					win.SendLineBack(fmt.Sprintf("            sessionId: %v", device.SessionId))
					win.SendLineBack(fmt.Sprintf("            loginAt: %v", device.LoginAt))
					win.SendLineBack(fmt.Sprintf("            deviceDesc: %v", device.DeviceDesc))
				}
			case *protocol.S2CLogoutRespPacket:
				win.SendLineBack("recv: packet logout resp")
				success := pkt.Success
				version := pkt.Version
				devices := pkt.Sessions
				ec := pkt.EchoCode()
				win.SendLineBack("    success: " + fmt.Sprint(success))
				win.SendLineBack("    version: " + fmt.Sprint(version))
				win.SendLineBack("    echoCode: " + ec)
				win.SendLineBack("    devices: " + fmt.Sprint(len(devices)))
				for i, device := range devices {
					win.SendLineBack(fmt.Sprintf("        device[%v]", i))
					win.SendLineBack(fmt.Sprintf("            sessionId: %v", device.SessionId))
					win.SendLineBack(fmt.Sprintf("            loginAt: %v", device.LoginAt))
					win.SendLineBack(fmt.Sprintf("            deviceDesc: %v", device.DeviceDesc))
				}
			case *protocol.S2COtherDeviceKickNotifyPacket:
				win.SendLineBack("recv: packet other device kick notify")
				fromDesc := pkt.FromSessionDesc
				fromId := pkt.FromSessionId
				win.SendLineBack("    fromSessionId: " + fromId)
				win.SendLineBack("    fromSessionDesc: " + fromDesc)
			case *protocol.S2CQueryStatusRespPacket:
				win.SendLineBack("recv: packet status")
				status := pkt.Status
				username := pkt.Username
				sessionId := pkt.SessionId
				devices := pkt.Sessions
				ec := pkt.EchoCode()
				win.SendLineBack("    status: " + status)
				win.SendLineBack("    username: " + username)
				win.SendLineBack("    sessionId: " + sessionId)
				win.SendLineBack("    echoCode: " + ec)
				win.SendLineBack("    version: " + fmt.Sprint(pkt.Version))
				win.SendLineBack("    devices: " + fmt.Sprint(len(devices)))
				for i, device := range devices {
					win.SendLineBack(fmt.Sprintf("        device[%v]", i))
					win.SendLineBack(fmt.Sprintf("            sessionId: %v", device.SessionId))
					win.SendLineBack(fmt.Sprintf("            loginAt: %v", device.LoginAt))
					win.SendLineBack(fmt.Sprintf("            deviceDesc: %v", device.DeviceDesc))
				}
			case *protocol.S2CPushMessagePacket:
				win.SendLineBack("recv: packet push meesage")
				pushType := pkt.PushType
				data := pkt.Data.([]byte)
				win.SendLineBack("    pushType: " + pushType)
				win.SendLineBack("    pushData: " + string(data))
				win.SendLineBack("    echoCode: " + pkt.EchoCode())
			}
		}
	}
}
