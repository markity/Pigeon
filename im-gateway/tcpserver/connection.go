package tcpserver

import (
	"log"
	"time"

	goreactor "github.com/markity/go-reactor"
)

func OnConn(conn goreactor.TCPConnection) {
	log.Printf("new conn received\n")
	conn.SetKeepAlive(false)
	conn.SetNoDelay(true)
	conn.SetDisConnectedCallback(OnDisConn)

	InitContext(conn)
	// connCtx := GetConnContextFromConn(conn)
	// connCtx.State.CommState

	// 发送心跳包循环
	conn.GetEventLoop().RunAt(time.Now(), time.Second, func(timerID int) {
		// conn.Send()
	})
}

func OnDisConn(conn goreactor.TCPConnection) {
	log.Printf("conn disconnected\n")
}
