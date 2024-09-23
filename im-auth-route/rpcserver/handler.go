package rpcserver

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"pigeon/im-auth-route/db"
	"pigeon/im-auth-route/rds"
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
	result, err := server.Rds.Login(&rds.SessionEntry{
		LoginAt:      time.Now().Unix(),
		Username:     req.Username,
		SessionId:    sessionId,
		DeviceDesc:   req.DeviceDesc,
		GwAdAddrPort: req.GwAdvertiseAddrPort,
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

	sessions := make([]*imauthroute.SessionEntry, 0, len(result.AllSessions))
	for _, v := range result.AllSessions {
		sessions = append(sessions, &imauthroute.SessionEntry{
			LoginAt:             v.LoginAt,
			Username:            v.Username,
			SessionId:           sessionId,
			DeviceDesc:          v.DeviceDesc,
			GwAdvertiseAddrPort: v.GwAdAddrPort,
		})
	}

	return &imauthroute.LoginResp{
		Code:      imauthroute.LoginResp_SUCCESS,
		SessionId: sessionId,
		Version:   result.Version,
		Sessions:  sessions,
	}, nil
}

func (server *RPCServer) Logout(ctx context.Context, req *imauthroute.LogoutReq) (res *imauthroute.LogoutResp, err error) {
	return nil, nil
}
func (server *RPCServer) ForceOffline(ctx context.Context, req *imauthroute.ForceOfflineReq) (res *imauthroute.ForceOfflineResp, err error) {
	return nil, nil
}
func (server *RPCServer) QuerySessionRoute(ctx context.Context, req *imauthroute.QuerySessionRouteReq) (res *imauthroute.QuerySessionRouteResp, err error) {
	return nil, nil
}
func (server *RPCServer) QueryUserRoute(ctx context.Context, req *imauthroute.QueryUserRouteReq) (res *imauthroute.QueryUserRouteResp, err error) {
	return nil, nil
}
