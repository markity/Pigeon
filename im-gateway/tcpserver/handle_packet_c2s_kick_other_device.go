package tcpserver

import (
	"context"
	"log"

	"pigeon/im-gateway/protocol"
	"pigeon/kitex_gen/service/imauthroute"

	goreactor "github.com/markity/go-reactor"
)

// 处理心跳包, 直接重置timer即可
func handleC2SKickOtherDevice(conn goreactor.TCPConnection, pack *protocol.C2SKickOtherDevicePacket) {
	connState := MustGetConnStateFromConn(conn)
	evloopCtx := MustLoadEvLoopContext(conn.GetEventLoop())

	if !protocol.IsSessionIdValid(pack.SessionId) {
		// 严格模式, 校验session id格式, 32位的uuid
		if ProtectMode {
			conn.ForceClose()
			return
		}
	}

	send := &protocol.S2CKickOhterDeviceRespPacket{
		Sessions: make([]*protocol.DeviceSessionEntry, 0),
	}
	send.SetEchoCode(pack.EchoCode())
	if connState.StateCode == StateCodeUnLogin ||
		connState.LoginSession.SessionId == pack.SessionId {
		if ProtectMode {
			conn.ForceClose()
		} else {
			send.Success = false
			send.Version = 0
			conn.Send(protocol.PackData(protocol.MustEncodePacket(send)))
		}
		return
	}

	// 调用rpc
	kickResp, err := evloopCtx.AuthRouteCli.ForceOffline(context.Background(), &imauthroute.ForceOfflineReq{
		Username:        connState.LoginSession.Username,
		SelfSessionId:   connState.LoginSession.SessionId,
		RemoteSessionId: pack.SessionId,
	})
	if err != nil {
		log.Printf("failed to call auth route: %v\n", err)
		conn.ForceClose()
		return
	}
	switch kickResp.Code {
	case imauthroute.ForceOfflineResp_SUCCESS:
		send.KickOK = true
	case imauthroute.ForceOfflineResp_FROM_SESSION_NOT_FOUND:
		send.KickOK = false
		log.Printf("unexpected kick failed: kickResp: %v\n", kickResp)
		conn.ForceClose()
		return
	case imauthroute.ForceOfflineResp_TO_SESSION_NOT_FOUND:
		send.KickOK = false
	default:
		panic("check me")
	}
	send.Version = kickResp.Version
	for _, v := range kickResp.Sessions {
		send.Sessions = append(send.Sessions, &protocol.DeviceSessionEntry{
			SessionId:  v.SessionId,
			LoginAt:    v.LoginAt,
			DeviceDesc: v.DeviceDesc,
		})
	}
	conn.Send(protocol.PackData(protocol.MustEncodePacket(send)))
}
