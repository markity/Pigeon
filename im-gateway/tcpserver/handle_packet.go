package tcpserver

import (
	"context"
	"fmt"
	"log"
	"time"

	"pigeon/im-gateway/protocol"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/imauthroute"
	"pigeon/kitex_gen/service/imrelay"

	goreactor "github.com/markity/go-reactor"
)

func handleC2SPacket(conn goreactor.TCPConnection, packet interface{}) {
	connState := MustGetConnStateFromConn(conn)
	evloopCtx := MustLoadEvLoopContext(conn.GetEventLoop())
	switch pack := packet.(type) {
	case *protocol.HeartbeatPacket:
		conn.GetEventLoop().CancelTimer(connState.HeartbeatTimeoutTimerId)
		heartbeatTimeoutTimerId := conn.GetEventLoop().RunAt(time.Now().
			Add(evloopCtx.HeartbeatTimeout), 0, func(timerID int) {
			log.Printf("timeout: force close\n")
			conn.ForceClose()
		})
		connState.HeartbeatTimeoutTimerId = heartbeatTimeoutTimerId
	case *protocol.C2SQueryStatusPacket:
		var resp protocol.S2CQueryStatusRespPacket
		if connState.StateCode == StateCodeLogin {
			resp.Status = "login"
			resp.Username = connState.LoginSession.Username
			resp.SessionId = connState.LoginSession.SessionId
		} else {
			resp.Status = "unlogin"
		}
		resp.SetEchoCode(pack.EchoCode())
		resp.Sessions = make([]*protocol.DeviceSessionEntry, 0)
		if connState.StateCode == StateCodeUnLogin {
			resp.Sessions = make([]*protocol.DeviceSessionEntry, 0)
			resp.Version = 0
		} else {
			queryuserRouteResp, err := evloopCtx.AuthRouteCli.QueryUserRoute(context.Background(), &imauthroute.QueryUserRouteReq{
				Username: connState.LoginSession.Username,
			})
			if err != nil {
				log.Printf("failed to call auth query user route: %v\n", err)
				conn.ForceClose()
				return
			}
			resp.Version = queryuserRouteResp.Version
			resp.Sessions = make([]*protocol.DeviceSessionEntry, 0, len(queryuserRouteResp.Routes))
			for _, v := range queryuserRouteResp.Routes {
				resp.Sessions = append(resp.Sessions, &protocol.DeviceSessionEntry{
					SessionId:  v.SessionId,
					LoginAt:    v.LoginAt,
					DeviceDesc: v.DeviceDesc,
				})
			}
		}
		conn.Send(protocol.PackData(protocol.MustEncodePacket(&resp)))
	case *protocol.C2SLoginPacket:
		// 如果输入错误, 直接force close，防止攻击
		if !protocol.IsUsernameValid(pack.Username) || !protocol.IsPasswordValid(pack.Password) ||
			!protocol.IsDeviceDescValid(pack.DeviceDesc) {
			conn.ForceClose()
			return
		}

		send := &protocol.S2CLoginRespPacket{}
		send.SetEchoCode(pack.EchoCode())

		// 如果已经登录, 直接返回, 这里给一个特殊语意, 已登录就query路由信息吧
		if connState.StateCode == StateCodeLogin {
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
		if selfSession == nil {
			log.Printf("unexpected: self session is not found in login resp\n")
			conn.ForceClose()
			return
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
	// 主动登出的包
	case *protocol.C2SLogoutPacket:
		send := &protocol.S2CLogoutRespPacket{}
		send.SetEchoCode(pack.EchoCode())
		send.Sessions = make([]*protocol.DeviceSessionEntry, 0)
		// 没有登录, 直接返回
		if connState.StateCode == StateCodeUnLogin {
			send.Success = false
			send.Version = 0
			conn.Send(protocol.PackData(protocol.MustEncodePacket(send)))
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
	// 客户端发送主动踢掉其他设备的包
	case *protocol.C2SKickOtherDevicePacket:
		// 输入非法, 直接force close
		if !protocol.IsSessionIdValid(pack.SessionId) {
			fmt.Println(pack.SessionId)
			conn.ForceClose()
			return
		}

		send := &protocol.S2CKickOhterDeviceRespPacket{
			Sessions: make([]*protocol.DeviceSessionEntry, 0),
		}
		send.SetEchoCode(pack.EchoCode())
		if connState.StateCode == StateCodeUnLogin ||
			connState.LoginSession.SessionId == pack.SessionId {
			conn.ForceClose()
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
	case *protocol.C2SBizMessagePacket:
		if connState.StateCode == StateCodeUnLogin {
			// 不care太多细节了, 直接force close
			conn.ForceClose()
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
}
