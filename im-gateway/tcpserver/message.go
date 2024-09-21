package tcpserver

import (
	"pigeon/im-gateway/protocol"

	goreactor "github.com/markity/go-reactor"
	"github.com/markity/go-reactor/pkg/buffer"
)

func OnMessage(conn goreactor.TCPConnection, buf buffer.Buffer) {
	for {
		pack, ok, err := protocol.UnpackDataFromBuffer(buf)
		if err != nil {
			conn.ForceClose()
			return
		}
		if !ok {
			return
		}

		c2sPacket, ok := protocol.ParseC2SPacket(pack)
		if !ok {
			conn.ForceClose()
			return
		}

		handleC2SPacket(conn, c2sPacket)
	}
}
