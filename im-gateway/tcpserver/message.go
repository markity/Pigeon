package tcpserver

import (
	goreactor "github.com/markity/go-reactor"
	"github.com/markity/go-reactor/pkg/buffer"
)

func OnMessage(conn goreactor.TCPConnection, buf buffer.Buffer) {
}
