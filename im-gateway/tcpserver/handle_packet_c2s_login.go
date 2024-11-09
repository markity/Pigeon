package tcpserver

import (
	"context"
	"log"

	"pigeon/im-gateway/protocol"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/imauthroute"

	goreactor "github.com/markity/go-reactor"
)

// 处理心跳包, 直接重置timer即可
func handleC2SLogin(conn goreactor.TCPConnection, pack *protocol.C2SLoginPacket) {
	connState := MustGetConnStateFromConn(conn)
	evloopCtx := MustLoadEvLoopContext(conn.GetEventLoop())

	if !protocol.IsUsernameValid(pack.Username) || !protocol.IsPasswordValid(pack.Password) ||
		!protocol.IsDeviceDescValid(pack.DeviceDesc) {
		// 如果输入错误, 直接force close，防止攻击
		if ProtectMode {
			conn.ForceClose()
			return
		}
	}

	send := &protocol.S2CLoginRespPacket{}
	send.SetEchoCode(pack.EchoCode())

	if connState.StateCode == StateCodeLogin {
		if ProtectMode {
			conn.ForceClose()
		} else {
			// 如果已经登录, 这里给一个特殊语意, 已登录就query路由信息
			send.Code = protocol.LoginRespCodeAlreadyLogin
			send.SessionId = connState.LoginSession.SessionId
			queryResp, err := evloopCtx.AuthRouteCli.QueryUserRoute(context.Background(), &imauthroute.QueryUserRouteReq{
				Username: connState.LoginSession.Username,
			})
			if err != nil {
				log.Printf("failed to call auth route: %v\n", err)
				conn.ForceClose()
				return
			}
			sessions := make([]*protocol.DeviceSessionEntry, 0, len(queryResp.Routes))
			for _, v := range queryResp.Routes {
				sessions = append(sessions, &protocol.DeviceSessionEntry{
					SessionId:  v.SessionId,
					LoginAt:    v.LoginAt,
					DeviceDesc: v.DeviceDesc,
				})
			}
			send.Version = queryResp.Version
			send.Sessions = sessions
			conn.Send(protocol.PackData(protocol.MustEncodePacket(send)))
		}
		return
	}

	// 调用登录接口
	loginResp, err := evloopCtx.AuthRouteCli.Login(context.Background(), &imauthroute.LoginReq{
		GwAdvertiseAddrPort: evloopCtx.RPCAdAddr,
		Username:            pack.Username,
		Password:            pack.Password,
		DeviceDesc:          pack.DeviceDesc,
	})
	if err != nil {
		log.Printf("failed to call auth route: %v\n", err)
		conn.ForceClose()
		return
	}

	sessions := make([]*protocol.DeviceSessionEntry, 0, len(loginResp.Sessions))
	var selfSession *base.SessionEntry
	for _, v := range loginResp.Sessions {
		sessions = append(sessions, &protocol.DeviceSessionEntry{
			SessionId:  v.SessionId,
			LoginAt:    v.LoginAt,
			DeviceDesc: v.DeviceDesc,
		})
		if v.SessionId == loginResp.SessionId {
			selfSession = v
		}
	}
	send.Version = loginResp.Version
	send.Sessions = sessions
	switch loginResp.Code {
	case imauthroute.LoginResp_SUCCESS:
		connState.LoginSession = selfSession
		connState.StateCode = StateCodeLogin
		evloopCtx.LoginedConnInfo[loginResp.SessionId] = connState
		evloopCtx.EvloopRoute.Store(loginResp.SessionId, conn.GetEventLoop())
		send.SessionId = loginResp.SessionId
		send.Code = protocol.LoginRespCodeSuccess
	case imauthroute.LoginResp_AUTH_ERROR:
		send.Code = protocol.LoginRespCodeAuthError
	case imauthroute.LoginResp_DEVICE_NUM_LIMIT:
		send.Code = protocol.LoginRespCodeDeviceNumLimit
	default:
		panic("check me")
	}
	conn.Send(protocol.PackData(protocol.MustEncodePacket(send)))
}
