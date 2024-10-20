package rpcserver

import (
	"bytes"
	"context"
	"log"
	"time"

	"pigeon/im-auth-route/api"
	"pigeon/im-auth-route/db"
	"pigeon/im-auth-route/rds"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/imauthroute"
	"pigeon/kitex_gen/service/imgateway"

	"github.com/google/uuid"
)

type RPCContext struct {
	DB  *db.DB
	Rds *rds.RdsAction
}

type RPCServer struct {
	RPCContext
}

func (server *RPCServer) Login(ctx context.Context, req *imauthroute.LoginReq) (*imauthroute.LoginResp, error) {
	lk, err := server.Rds.LockUsername(req.Username, time.Second*15)
	if err != nil {
		log.Printf("lock username error: %v\n", err)
		return nil, err
	}
	defer func() {
		if err := lk.UnLock(); err != nil {
			log.Printf("unlock username error: %v\n", err)
		}
	}()

	user, err := server.DB.GetUserByUsername(req.Username)
	if err != nil {
		log.Printf("get user error: %v\n", err)
		return nil, err
	}

	if user == nil || !bytes.Equal(db.ToSha256([]byte(req.Password)), user.PasswordSha256) {
		return &imauthroute.LoginResp{
			Code: imauthroute.LoginResp_AUTH_ERROR,
		}, nil
	}

	// redis尝试登录
	// 7709f7ba-8a06-4ed3-a730-d7aacd22b87e
	// 7c771677-017a-4f83-bf07-1da642b45814
	// 237fca45-5e3f-430a-b9df-8a5a8372d78e
	sessionId := uuid.NewString()
	result, err := server.Rds.Login(&base.SessionEntry{
		LoginAt:             time.Now().Unix(),
		Username:            req.Username,
		SessionId:           sessionId,
		DeviceDesc:          req.DeviceDesc,
		GwAdvertiseAddrport: req.GwAdvertiseAddrPort,
	})
	if err != nil {
		log.Printf("redis login error: %v\n", err)
		return nil, err
	}

	var code imauthroute.LoginResp_LoginRespCode
	if result.Success {
		code = imauthroute.LoginResp_SUCCESS
	} else {
		code = imauthroute.LoginResp_DEVICE_NUM_LIMIT
	}

	for _, v := range result.AllSessions {
		// 除了当前的session, 其余的都广播通知
		if v.SessionId != sessionId {
			go func() {
				for {
					_, err := api.NewGatewayClientFromAdAddr(v.GwAdvertiseAddrport).BroadcastDeviceInfo(context.Background(), &imgateway.BroadcastDeviceInfoReq{
						SessionId: v.SessionId,
						Version:   result.Version,
						Sessions:  result.AllSessions,
					})
					if err != nil {
						log.Printf("broadcast device info error, retry: %v\n", err)
						time.Sleep(time.Millisecond * 50)
						continue
					}
					return
				}
			}()
		}
	}
	return &imauthroute.LoginResp{
		Code:      code,
		SessionId: sessionId,
		Version:   result.Version,
		Sessions:  result.AllSessions,
	}, nil
}

// 当用户主动下线 或者 网关处连接断开调用, 用来下线路由, im-auth-route会直接删除路由
func (server *RPCServer) Logout(ctx context.Context, req *imauthroute.LogoutReq) (res *imauthroute.LogoutResp, err error) {
	result, err := server.Rds.Logout(req.Username, req.SessionId)
	if err != nil {
		log.Printf("redis logout error: %v\n", err)
		return nil, err
	}

	for _, v := range result.AllSessions {
		go func() {
			for {
				_, err := api.NewGatewayClientFromAdAddr(v.GwAdvertiseAddrport).BroadcastDeviceInfo(context.Background(), &imgateway.BroadcastDeviceInfoReq{
					SessionId: v.SessionId,
					Version:   result.Version,
					Sessions:  result.AllSessions,
				})
				if err != nil {
					log.Printf("broadcast device info error, retry: %v\n", err)
					time.Sleep(time.Millisecond * 50)
					continue
				}
				break
			}
		}()
	}

	return &imauthroute.LogoutResp{
		Success:  result.Success,
		Version:  result.Version,
		Sessions: result.AllSessions,
	}, nil
}

// 用户使用踢人命令时调用此接口
func (server *RPCServer) ForceOffline(ctx context.Context, req *imauthroute.ForceOfflineReq) (*imauthroute.ForceOfflineResp, error) {
	result, err := server.Rds.ForceOffline(req.Username, req.SelfSessionId, req.RemoteSessionId)
	if err != nil {
		log.Printf("redis force offline error: %v\n", err)
		return nil, err
	}

	var send = imauthroute.ForceOfflineResp{
		Code:        result.Code,
		FromSession: result.FromSession,
		ToSession:   result.ToSession,
		Version:     result.Version,
		Sessions:    result.AllSessions,
	}
	if result.Code == imauthroute.ForceOfflineResp_SUCCESS {
		for _, v := range result.AllSessions {
			go func() {
				for {
					_, err := api.NewGatewayClientFromAdAddr(v.GwAdvertiseAddrport).BroadcastDeviceInfo(context.Background(), &imgateway.BroadcastDeviceInfoReq{
						SessionId: v.SessionId,
						Version:   result.Version,
						Sessions:  result.AllSessions,
					})
					if err != nil {
						log.Printf("broadcast device info error, retry: %v\n", err)
						time.Sleep(time.Millisecond * 50)
						continue
					}
					break
				}
			}()
		}
		// 对targetSession发退出消息
		go func() {
			for {
				_, err = api.NewGatewayClientFromAdAddr(result.ToSession.GwAdvertiseAddrport).OtherDeviceKick(context.Background(), &imgateway.OtherDeviceKickReq{
					FromSession:     req.SelfSessionId,
					FromSessionDesc: result.FromSession.DeviceDesc,
					ToSession:       req.RemoteSessionId,
				})
				if err != nil {
					log.Printf("send other device kick error, retry: %v\n", err)
					time.Sleep(time.Millisecond * 50)
					continue
				}
				return
			}
		}()
	}

	return &send, nil
}

func (server *RPCServer) QuerySessionRoute(ctx context.Context, req *imauthroute.QuerySessionRouteReq) (*imauthroute.QuerySessionRouteResp, error) {
	result, err := server.Rds.QuerySessionRoute(req.SessionId)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return &imauthroute.QuerySessionRouteResp{
			Success: false,
		}, nil
	}
	return &imauthroute.QuerySessionRouteResp{
		Success: true,
		Route: &base.SessionEntry{
			LoginAt:             result.LoginAt,
			Username:            result.Username,
			SessionId:           result.SessionId,
			DeviceDesc:          result.DeviceDesc,
			GwAdvertiseAddrport: result.GwAdvertiseAddrport,
		},
	}, nil
}

func (server *RPCServer) QueryUserRoute(ctx context.Context, req *imauthroute.QueryUserRouteReq) (res *imauthroute.QueryUserRouteResp, err error) {
	version, result, err := server.Rds.QueryUserRoute(req.Username)
	if err != nil {
		return nil, err
	}
	return &imauthroute.QueryUserRouteResp{
		Version: version,
		Routes:  result,
	}, nil
}
