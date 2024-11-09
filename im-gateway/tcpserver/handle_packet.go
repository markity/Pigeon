package tcpserver

import (
	"pigeon/im-gateway/protocol"

	goreactor "github.com/markity/go-reactor"
)

// protect mode 严格模式, 当客户端出了一点小小的纰漏
// 都直接关闭连接, 要求客户端做完全完备的检查
const ProtectMode = false

func handleC2SPacket(conn goreactor.TCPConnection, packet interface{}) {
	switch pack := packet.(type) {
	case *protocol.HeartbeatPacket:
		handleHb(conn, pack)
	case *protocol.C2SQueryStatusPacket:
		handleC2SQueryStatus(conn, pack)
	case *protocol.C2SLoginPacket:
		handleC2SLogin(conn, pack)
	// 主动登出的包
	case *protocol.C2SLogoutPacket:
		handleC2SLogout(conn, pack)
	// 客户端发送主动踢掉其他设备的包
	case *protocol.C2SKickOtherDevicePacket:
		handleC2SKickOtherDevice(conn, pack)
	case *protocol.C2SBizMessagePacket:
		handleC2SBizMessage(conn, pack)
	}
}
