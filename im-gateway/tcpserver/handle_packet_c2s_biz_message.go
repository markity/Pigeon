package tcpserver

import (
	"context"
	"log"

	"pigeon/im-gateway/protocol"
	"pigeon/kitex_gen/service/imrelay"

	goreactor "github.com/markity/go-reactor"
)

// 处理心跳包, 直接重置timer即可
func handleC2SBizMessage(conn goreactor.TCPConnection, pack *protocol.C2SBizMessagePacket) {
	connState := MustGetConnStateFromConn(conn)
	evloopCtx := MustLoadEvLoopContext(conn.GetEventLoop())

	if connState.StateCode == StateCodeUnLogin {
		// 严格模式, 不允许非登录状态下发送这个消息
		if ProtectMode {
			conn.ForceClose()
		}
		return
	}

	_, err := evloopCtx.RelayCli.BizMessage(context.Background(), &imrelay.BizMessageReq{
		Session:  connState.LoginSession,
		Biz:      pack.BizType,
		EchoCode: pack.EchoCode(),
		Data:     pack.Data.([]byte),
	})
	if err != nil {
		log.Printf("failed to call relay: %v\n", err)
		conn.ForceClose()
	}
}
