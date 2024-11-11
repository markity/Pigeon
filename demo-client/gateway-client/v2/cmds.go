package main

import "strings"

// ah on/off
type cmdAutoHeartbeat struct {
	On *bool
}

// hb
type cmdHeartbeat struct{}

// login username password echocode
type cmdLogin struct {
	Username   string
	Password   string
	DeviceDesc string
	EchoCode   *string
}

type cmdLogout struct {
	EchoCode *string
}

type cmdKickOtherSession struct {
	SessionId string
	EchoCode  *string
}

type cmdSendBiz struct {
	Biz      string
	Data     string
	EchoCode *string
}

type exitCmd struct{}

type helpCmd struct{}

type clearCmd struct{}

type statusCmd struct {
	EchoCode *string
}

type emptyCmd struct{}

type hideHeartbeatInfo struct{}

func parseCommand(line string) interface{} {
	line = strings.TrimSpace(line)
	if line == "" {
		return &emptyCmd{}
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
		if len(cmds) == 5 {
			var ec = cmds[4]
			return &cmdLogin{Username: cmds[1], Password: cmds[2], DeviceDesc: cmds[3], EchoCode: &ec}
		}

		if len(cmds) == 4 {
			return &cmdLogin{Username: cmds[1], Password: cmds[2], DeviceDesc: cmds[3], EchoCode: nil}
		}
	case "logout":
		if len(cmds) == 2 {
			var ec = cmds[1]
			return &cmdLogout{EchoCode: &ec}
		}
		if len(cmds) == 1 {
			return &cmdLogout{EchoCode: nil}
		}
	case "kick":
		if len(cmds) == 3 {
			var ec = cmds[2]
			return &cmdKickOtherSession{SessionId: cmds[1], EchoCode: &ec}
		}
		if len(cmds) == 2 {
			return &cmdKickOtherSession{SessionId: cmds[1], EchoCode: nil}
		}
	case "exit":
	case "quit":
	case "q":
		if len(cmds) == 1 {
			return &exitCmd{}
		}
	case "clear":
		if len(cmds) == 1 {
			return &clearCmd{}
		}
	case "help":
		if len(cmds) == 1 {
			return &helpCmd{}
		}
	case "status":
		if len(cmds) == 2 {
			var ec = cmds[1]
			return &statusCmd{
				EchoCode: &ec,
			}
		}
		if len(cmds) == 1 {
			return &statusCmd{
				EchoCode: nil,
			}
		}
	case "sendbiz":
		if len(cmds) == 4 {
			var ec = cmds[3]
			return &cmdSendBiz{
				Biz:      cmds[1],
				Data:     cmds[2],
				EchoCode: &ec,
			}
		}
		if len(cmds) == 3 {
			return &cmdSendBiz{
				Biz:      cmds[1],
				Data:     cmds[2],
				EchoCode: nil,
			}
		}
	case "hidehb":
		if len(cmds) == 1 {
			return &hideHeartbeatInfo{}
		}
	}
	return nil

}
