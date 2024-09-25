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

// forceclose 或者连接意外关闭都会触发OnDisConn
func OnDisConn(conn goreactor.TCPConnection) {
	log.Printf("conn disconnected\n")
	MustLoadEvLoopContext(conn.GetEventLoop()).ConnMetrics.Add(-1)
	ReleaseConn(conn)
}
