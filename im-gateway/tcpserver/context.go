package tcpserver

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"pigeon/im-gateway/protocol"
	"pigeon/kitex_gen/service/imauthroute"
	authroute "pigeon/kitex_gen/service/imauthroute/imauthroute"

	goreactor "github.com/markity/go-reactor"
	eventloop "github.com/markity/go-reactor/pkg/event_loop"
)

const evloopContextKey = "_ev_loop_ctx"
const connStateKey = "_ev_conn_state"

type EvloopContext struct {
	RelayCli          interface{}
	AuthRouteCli      authroute.Client
	ConnMetrics       *atomic.Int64
	HeartbeatInterval time.Duration
	HeartbeatTimeout  time.Duration
	EvloopRoute       *sync.Map
	RPCAdAddr         string

	// 私有的变量, setup的时候重新创建
	// 这个变量代表已经登录的session的状态
	LoginedConnInfo map[string]*ConnState
}

type StateCode int

const (
	StateCodeUnLogin StateCode = iota
	StateCodeLogin
)

type ConnState struct {
	// 状态码
	StateCode StateCode
	// 无论什么状态, heartbeat都是会发送的, 启动一个定时器不断发送
	HeartbeatTimerId int
	// 心跳超时定时器, 当收到客户端的心跳包时重试
	HeartbeatTimeoutTimerId int
	// Conn句柄
	Conn goreactor.TCPConnection

	// 已登录状态下有效的变量
	Username  *string
	SessionId *string
}

func SetUpEvLoopContext(loop eventloop.EventLoop, ctx EvloopContext) {
	ctx.LoginedConnInfo = make(map[string]*ConnState)
	loop.SetContext(evloopContextKey, &ctx)
}

func MustLoadEvLoopContext(loop eventloop.EventLoop) *EvloopContext {
	c, ok := loop.GetContext(evloopContextKey)
	if !ok {
		panic("get context")
	}
	return c.(*EvloopContext)
}

func SetUpConn(conn goreactor.TCPConnection) {
	bs := protocol.PackData(protocol.MustEncodePacket(&protocol.HeartbeatPacket{}))
	heartbeatTimerId := conn.GetEventLoop().RunAt(time.Now(), MustLoadEvLoopContext(conn.GetEventLoop()).HeartbeatInterval, func(timerID int) {
		conn.Send(bs)
	})
	heartbeatTimeoutTimerId := conn.GetEventLoop().RunAt(time.Now().
		Add(MustLoadEvLoopContext(conn.GetEventLoop()).HeartbeatTimeout), 0, func(timerID int) {
		conn.ForceClose()
		log.Printf("timeout: force close\n")
	})
	conn.SetContext(connStateKey, &ConnState{
		StateCode:               StateCodeUnLogin,
		HeartbeatTimerId:        heartbeatTimerId,
		HeartbeatTimeoutTimerId: heartbeatTimeoutTimerId,
		SessionId:               nil,
		Conn:                    conn,
	})
}

// 删除定时器, 释放eventloop的conn路由
func ReleaseConn(conn goreactor.TCPConnection) {
	ctx := MustLoadEvLoopContext(conn.GetEventLoop())

	state := MustGetConnStateFromConn(conn)
	ok1 := conn.GetEventLoop().CancelTimer(state.HeartbeatTimerId)
	ok2 := conn.GetEventLoop().CancelTimer(state.HeartbeatTimeoutTimerId)
	if !ok1 || !ok2 {
		panic("check me")
	}
	if state.StateCode == StateCodeLogin {
		_, ok := MustLoadEvLoopContext(conn.GetEventLoop()).LoginedConnInfo[*state.SessionId]
		if !ok {
			panic("check me")
		}

		// 调用rpc 删除evloop-route中的路由表
		resp, err := ctx.AuthRouteCli.Logout(context.Background(), &imauthroute.LogoutReq{
			SessionId: *state.SessionId,
			Username:  *state.Username,
		})
		if err != nil || !resp.Success {
			fmt.Printf("failed to logout: resp: %v, err: %v\n", resp, err)
		}

		delete(MustLoadEvLoopContext(conn.GetEventLoop()).LoginedConnInfo, *state.SessionId)
		ctx.EvloopRoute.Delete(*state.SessionId)
	}
}

func MustGetConnStateFromConn(conn goreactor.TCPConnection) *ConnState {
	v, ok := conn.GetContext(connStateKey)
	if !ok {
		panic("get context")
	}
	return v.(*ConnState)
}
