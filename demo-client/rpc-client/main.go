package main

import (
	"context"
	"fmt"
	regetcd "pigeon/common/kitex_registry/etcd"
	"pigeon/im-auth-route/api"
	authroute "pigeon/kitex_gen/service/imauthroute"
	"pigeon/kitex_gen/service/imauthroute/imauthroute"
	"pigeon/kitex_gen/service/imgateway"

	"strings"

	"github.com/cloudwego/kitex/client"
	interactive "github.com/markity/Interactive-Console"
)

type helpCmd struct{}
type exitCmd struct{}
type emptyCmd struct{}
type unknownCmd struct{}
type clearCmd struct{}

type queryUserRouteCmd struct {
	Username string
}

type querySessionRouteCmd struct {
	SessionId string
}

type pushCmd struct {
	GwAddrPort string
	SessionId  string
	PushType   string
	EchoCode   string
	Msg        string
}

func parseCommand(line string) interface{} {
	line = strings.TrimSpace(line)
	if line == "" {
		return &emptyCmd{}
	}

	cmds := strings.Split(line, " ")
	switch cmds[0] {
	case "help":
		if len(cmds) == 1 {
			return &helpCmd{}
		}
	case "exit":
		if len(cmds) == 1 {
			return &exitCmd{}
		}
	case "qu":
		if len(cmds) == 2 {
			return &queryUserRouteCmd{
				Username: cmds[1],
			}
		}
	case "qs":
		if len(cmds) == 2 {
			return &querySessionRouteCmd{
				SessionId: cmds[1],
			}
		}
	case "push":
		if len(cmds) == 6 {
			return &pushCmd{
				GwAddrPort: cmds[1],
				SessionId:  cmds[2],
				PushType:   cmds[3],
				EchoCode:   cmds[4],
				Msg:        cmds[5],
			}
		}
	case "clear":
		if len(cmds) == 1 {
			return &clearCmd{}
		}
	}
	return &unknownCmd{}
}

func main() {
	res, err := regetcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	cli, err := imauthroute.NewClient("im-authroute", client.WithResolver(res))
	if err != nil {
		panic(err)
	}

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

	win.SetBlockInput(false)

	cmdChan := win.GetCmdChan()
	eventChan := win.GetEventChan()

	for {
		select {
		case cmd := <-cmdChan:
			c := parseCommand(cmd)
			switch pkt := c.(type) {
			case *helpCmd:
				win.SendLineBack("help:")
				win.SendLineBack("    clear: clear screen")
				win.SendLineBack("    qs username: query session route")
				win.SendLineBack("    qu username: query user route")
				win.SendLineBack("    push gwAddr session pushType echoCode msg: push msg")
				win.SendLineBack("    exit: exit")
				win.SendLineBack("    help: print help")
			case *exitCmd:
				win.Stop()
				return
			case *emptyCmd:
			case *querySessionRouteCmd:
				win.SendLineBack("query session route, session=" + pkt.SessionId)
				resp, err := cli.QuerySessionRoute(context.Background(), &authroute.QuerySessionRouteReq{
					SessionId: pkt.SessionId,
				})
				if err != nil {
					win.SendLineBack("error: " + err.Error())
				} else {
					win.SendLineBack("success:")
					if resp.Route != nil {
						win.SendLineBack("    username: " + resp.Route.Username)
						win.SendLineBack("    sessionId: " + resp.Route.SessionId)
						win.SendLineBack("    deviceDesc: " + resp.Route.DeviceDesc)
						win.SendLineBack("    gwAdAddrPort: " + resp.Route.GwAdvertiseAddrport)
					} else {
						win.SendLineBack("    no route found")
					}
				}
			case *queryUserRouteCmd:
				win.SendLineBack("query user route, username=" + pkt.Username)
				resp, err := cli.QueryUserRoute(context.Background(), &authroute.QueryUserRouteReq{
					Username: pkt.Username,
				})
				if err != nil {
					win.SendLineBack("error: " + err.Error())
				} else {
					win.SendLineBack("success: " + fmt.Sprint(len(resp.Routes)) + " routes")
					for i, v := range resp.Routes {
						win.SendLineBack("    route[" + fmt.Sprint(i) + "]")
						win.SendLineBack("        username: " + v.Username)
						win.SendLineBack("        sessionId: " + v.SessionId)
						win.SendLineBack("        deviceDesc: " + v.DeviceDesc)
						win.SendLineBack("        gwAdAddrPort: " + v.GwAdvertiseAddrport)
					}
				}
			case *pushCmd:
				gwCli := api.NewGatewayClientFromAdAddr(pkt.GwAddrPort)
				_, err := gwCli.PushMessage(context.Background(), &imgateway.PushMessageReq{
					SessionId: pkt.SessionId,
					PushType:  pkt.PushType,
					Data:      []byte(pkt.Msg),
					EchoCode:  pkt.EchoCode,
				})
				if err != nil {
					win.SendLineBack("error: " + err.Error())
				} else {
					win.SendLineBack("success")
				}
			case *clearCmd:
				win.Clear()
			case *unknownCmd:
				win.SendLineBack("unknown cmd")
			}
		case <-eventChan:
			win.Stop()
			return
		}
		win.SetBlockInput(false)
	}
}
