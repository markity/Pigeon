package tcpserver

import (
	"log"
	"time"

	"pigeon/im-gateway/protocol"

	goreactor "github.com/markity/go-reactor"
)

// 处理心跳包, 直接重置timer即可
func handleHb(conn goreactor.TCPConnection, _ *protocol.HeartbeatPacket) {
	connState := MustGetConnStateFromConn(conn)
	evloopCtx := MustLoadEvLoopContext(conn.GetEventLoop())

	conn.GetEventLoop().CancelTimer(connState.HeartbeatTimeoutTimerId)
	heartbeatTimeoutTimerId := conn.GetEventLoop().RunAt(time.Now().
		Add(evloopCtx.HeartbeatTimeout), 0, func(timerID int) {
		log.Printf("timeout: force close\n")
		conn.ForceClose()
	})
	connState.HeartbeatTimeoutTimerId = heartbeatTimeoutTimerId
}
