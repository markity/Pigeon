package tcpserver

import eventloop "github.com/markity/go-reactor/pkg/event_loop"

// 此处data为[]byte类型, 不要在event loop做太多操作, 否则会阻塞event loop
func PushMessage(evloop eventloop.EventLoop, sessionId string, data []byte, okChan chan<- bool) {
	evLoopCtx := MustLoadEvLoopContext(evloop)
	connInfo, ok := evLoopCtx.LoginedConnInfo[sessionId]
	if !ok {
		okChan <- false
		return
	}
	connInfo.Conn.Send(data)
	okChan <- true
}

func OtherDeveiceKick(evloop eventloop.EventLoop, toSessionId string, kickMessage []byte, okChan chan<- bool) {
	evLoopCtx := MustLoadEvLoopContext(evloop)
	connInfo, ok := evLoopCtx.LoginedConnInfo[toSessionId]
	if !ok {
		okChan <- false
		return
	}

	delete(evLoopCtx.LoginedConnInfo, toSessionId)
	connInfo.StateCode = StateCodeUnLogin
	evLoopCtx.EvloopRoute.Delete(toSessionId)
	connInfo.Conn.Send(kickMessage)
	okChan <- true
}

func BroadcastDeviceInfo(evloop eventloop.EventLoop, toSessionId string, data []byte, okChan chan<- bool) {
	evLoopCtx := MustLoadEvLoopContext(evloop)
	connInfo, ok := evLoopCtx.LoginedConnInfo[toSessionId]
	if !ok {
		okChan <- false
		return
	}

	connInfo.Conn.Send(data)
	okChan <- true
}
