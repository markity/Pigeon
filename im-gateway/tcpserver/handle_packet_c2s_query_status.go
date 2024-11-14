package tcpserver

import (
	"context"
	"log"

	"pigeon/common/protocol"
	"pigeon/kitex_gen/service/imauthroute"

	goreactor "github.com/markity/go-reactor"
)

// 处理c2s query status
func handleC2SQueryStatus(conn goreactor.TCPConnection, pack *protocol.C2SQueryStatusPacket) {
	connState := MustGetConnStateFromConn(conn)
	evloopCtx := MustLoadEvLoopContext(conn.GetEventLoop())
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
	// 状态unlogin
	if connState.StateCode == StateCodeUnLogin {
		resp.Sessions = make([]*protocol.DeviceSessionEntry, 0)
		resp.Version = 0
	} else {
		// 状态login, 查询route, 拿到其它的session
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
}
