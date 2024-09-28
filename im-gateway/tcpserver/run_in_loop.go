package tcpserver

import (
	"fmt"

	eventloop "github.com/markity/go-reactor/pkg/event_loop"
)

// 此处data为[]byte类型, 不要在event loop做太多操作, 否则会阻塞event loop
func PushMessage(evloop eventloop.EventLoop, sessionId string, data []byte) {
	evLoopCtx := MustLoadEvLoopContext(evloop)
	connInfo, ok := evLoopCtx.LoginedConnInfo[sessionId]
	if !ok {
		return
	}
	connInfo.Conn.Send(data)
}

func OtherDeveiceKick(evloop eventloop.EventLoop, toSessionId string, kickMessage []byte) {
	fmt.Println("other device kick")
	evLoopCtx := MustLoadEvLoopContext(evloop)
	connInfo, ok := evLoopCtx.LoginedConnInfo[toSessionId]
	if !ok {
		return
	}

	connInfo.StateCode = StateCodeUnLogin
	connInfo.SessionId = nil
	connInfo.Username = nil
	delete(evLoopCtx.LoginedConnInfo, toSessionId)
	evLoopCtx.EvloopRoute.Delete(toSessionId)
	connInfo.Conn.Send(kickMessage)
}

func BroadcastDeviceInfo(evloop eventloop.EventLoop, toSessionId string, data []byte) {
	evLoopCtx := MustLoadEvLoopContext(evloop)
	connInfo, ok := evLoopCtx.LoginedConnInfo[toSessionId]
	if !ok {
		return
	}

	connInfo.Conn.Send(data)
}
