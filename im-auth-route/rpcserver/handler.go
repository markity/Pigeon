package rpcserver

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"pigeon/im-auth-route/db"
	"pigeon/im-auth-route/rds"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/imauthroute"
	"time"

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
	fmt.Println("login")
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
	sessionId := uuid.New().String()
	result, err := server.Rds.Login(&base.SessionEntry{
		LoginAt:             time.Now().Unix(),
		Username:            req.Username,
		SessionId:           sessionId,
		DeviceDesc:          req.DeviceDesc,
		GwAdvertiseAddrPort: req.GwAdvertiseAddrPort,
	})
	if err != nil {
		log.Printf("redis login error: %v\n", err)
		return nil, err
	}
	if !result.Success {
		return &imauthroute.LoginResp{
			Code: imauthroute.LoginResp_DEVICE_NUM_LIMIT,
		}, nil
	}

	sessions := make([]*base.SessionEntry, 0, len(result.AllSessions))
	for _, v := range result.AllSessions {
		sessions = append(sessions, &base.SessionEntry{
			LoginAt:             v.LoginAt,
			Username:            v.Username,
			SessionId:           sessionId,
			DeviceDesc:          v.DeviceDesc,
			GwAdvertiseAddrPort: v.GwAdvertiseAddrPort,
		})
	}

	return &imauthroute.LoginResp{
		Code:      imauthroute.LoginResp_SUCCESS,
		SessionId: sessionId,
		Version:   result.Version,
		Sessions:  sessions,
	}, nil
}

// 当用户主动下线 或者 网关处连接断开调用, 用来下线路由, im-auth-route会直接删除路由
func (server *RPCServer) Logout(ctx context.Context, req *imauthroute.LogoutReq) (res *imauthroute.LogoutResp, err error) {
	result, err := server.Rds.Logout(req.Username, req.SessionId)
	if err != nil {
		log.Printf("redis logout error: %v\n", err)
		return nil, err
	}
	return &imauthroute.LogoutResp{
		Success: result.Success,
	}, nil
}

// 用户使用踢人命令时调用此接口
func (server *RPCServer) ForceOffline(ctx context.Context, req *imauthroute.ForceOfflineReq) (res *imauthroute.ForceOfflineResp, err error) {
	result, err := server.Rds.ForceOffline(req.Username, req.SelfSessionId, req.RemoteSessionId)
	if err != nil {
		log.Printf("redis force offline error: %v\n", err)
		return nil, err
	}
	s := make([]*base.SessionEntry, 0, len(result.AllSessions))
	for _, v := range result.AllSessions {
		s = append(s, v)
	}
	return &imauthroute.ForceOfflineResp{
		Code:     result.Code,
		Version:  result.Version,
		Sessions: s,
	}, nil
}

func (server *RPCServer) QuerySessionRoute(ctx context.Context, req *imauthroute.QuerySessionRouteReq) (res *imauthroute.QuerySessionRouteResp, err error) {
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
			GwAdvertiseAddrPort: result.GwAdvertiseAddrPort,
		},
	}, nil
}

func (server *RPCServer) QueryUserRoute(ctx context.Context, req *imauthroute.QueryUserRouteReq) (res *imauthroute.QueryUserRouteResp, err error) {
	result, err := server.Rds.QueryUserRoute(req.Username)
	if err != nil {
		return nil, err
	}
	return &imauthroute.QueryUserRouteResp{
		Routes: result,
	}, nil
}
