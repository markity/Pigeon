package tcpserver

import (
	"log"
	"pigeon/im-gateway/protocol"
	"time"

	goreactor "github.com/markity/go-reactor"
)

func handleC2SPacket(conn goreactor.TCPConnection, packet interface{}) {
	connState := MustGetConnStateFromConn(conn)
	switch packet.(type) {
	case *protocol.C2SLoginPacket:
	case *protocol.C2SLogoutPacket:
	case *protocol.HeartbeatPacket:
		conn.GetEventLoop().CancelTimer(connState.HeartbeatTimeoutTimerId)
		heartbeatTimeoutTimerId := conn.GetEventLoop().RunAt(time.Now().
			Add(MustLoadEvLoopContext(conn.GetEventLoop()).HeartbeatTimeout), 0, func(timerID int) {
			conn.ForceClose()
			log.Printf("timeout: force close\n")
		})
		connState.HeartbeatTimeoutTimerId = heartbeatTimeoutTimerId
	}
}
