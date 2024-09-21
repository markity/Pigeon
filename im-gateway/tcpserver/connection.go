package tcpserver

import (
	"log"

	goreactor "github.com/markity/go-reactor"
)

func OnConn(conn goreactor.TCPConnection) {
	log.Printf("new conn received\n")
	MustLoadEvLoopContext(conn.GetEventLoop()).ConnMetrics.Add(1)
	conn.SetKeepAlive(false)
	conn.SetNoDelay(true)
	conn.SetDisConnectedCallback(OnDisConn)

	SetUpConn(conn)
}

func OnDisConn(conn goreactor.TCPConnection) {
	MustLoadEvLoopContext(conn.GetEventLoop()).ConnMetrics.Add(-1)
	ReleaseConn(conn)
	log.Printf("conn disconnected\n")
}
