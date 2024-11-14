package tcpserver

import (
	"context"
	"log"

	"pigeon/common/protocol"
	"pigeon/kitex_gen/service/imauthroute"

	goreactor "github.com/markity/go-reactor"
)

// 处理心跳包, 直接重置timer即可
func handleC2SLogout(conn goreactor.TCPConnection, pack *protocol.C2SLogoutPacket) {
	connState := MustGetConnStateFromConn(conn)
	evloopCtx := MustLoadEvLoopContext(conn.GetEventLoop())

	send := &protocol.S2CLogoutRespPacket{}
	send.SetEchoCode(pack.EchoCode())
	send.Sessions = make([]*protocol.DeviceSessionEntry, 0)
	// 没有登录, 直接返回
	if connState.StateCode == StateCodeUnLogin {
		if ProtectMode {
			conn.ForceClose()
		} else {
			send.Success = false
			send.Version = 0
			conn.Send(protocol.PackData(protocol.MustEncodePacket(send)))
		}
		return
	}

	// 已经登录,调用rpc
	logoutResp, err := evloopCtx.AuthRouteCli.Logout(context.Background(), &imauthroute.LogoutReq{
		SessionId: connState.LoginSession.SessionId,
		Username:  connState.LoginSession.Username,
	})
	if err != nil {
		log.Printf("failed to call auth route: %v\n", err)
		conn.ForceClose()
		return
	}

	if !logoutResp.Success {
		log.Printf("unexpected logout failed: logoutResp: %v\n", logoutResp)
		conn.ForceClose()
		return
	}

	// 路由删除成功了, 删除本地状态
	evloopCtx.EvloopRoute.Delete(connState.LoginSession.SessionId)
	connState.StateCode = StateCodeUnLogin
	connState.LoginSession = nil
	send.Success = true
	send.Version = logoutResp.Version
	send.Sessions = make([]*protocol.DeviceSessionEntry, 0, len(logoutResp.Sessions))
	for _, v := range logoutResp.Sessions {
		send.Sessions = append(send.Sessions, &protocol.DeviceSessionEntry{
			SessionId:  v.SessionId,
			LoginAt:    v.LoginAt,
			DeviceDesc: v.DeviceDesc,
		})
	}
	conn.Send(protocol.PackData(protocol.MustEncodePacket(send)))
}
