package rpcserver

import (
	"context"
	"pigeon/common/keylock"
	"pigeon/im-chat-evloop/api"
	"pigeon/im-chat-evloop/evloop"
	"pigeon/kitex_gen/service/imchatevloop"
	relay "pigeon/kitex_gen/service/imrelay"
	"pigeon/kitex_gen/service/imrelay/imrelay"

	"sync"
	"sync/atomic"
)

// 创建群聊
type RPCServer struct {
	RelayCli imrelay.Client

	// 读写current version, 需要加锁
	CurrentVersionMu sync.Mutex
	CurrentVersion   atomic.Int64

	// key: group id, value: chat eventloo
	CreateChatEventloopMu sync.Mutex
	ChatEventloops        sync.Map

	// 迁移锁, 防止并发迁移eventloop
	MoveLock keylock.KeyedMutex
}

// 创建群聊rpc
func (s *RPCServer) CreateGroup(ctx context.Context, req *imchatevloop.CreateGroupRequest) (
	res *imchatevloop.CreateGroupResponse, err error) {
	currentVersion := s.CurrentVersion.Load()
	if req.Version > currentVersion {
		s.CurrentVersionMu.Lock()
		currentVersion = s.CurrentVersion.Load()
		if req.Version > currentVersion {
			s.CurrentVersion.Store(req.Version)
		}
		s.CurrentVersionMu.Unlock()
	}

	// 可以尝试创建chat eventloop

	// 下面检查一次是因为有可能rpc失败, 连续两次创建event loop, 防止出现多个event loop
	if _, ok := s.ChatEventloops.Load(req.GroupId); ok {
		return &imchatevloop.CreateGroupResponse{
			Success: true,
			Version: currentVersion,
		}, nil
	}

	s.CreateChatEventloopMu.Lock()
	defer s.CreateChatEventloopMu.Unlock()
	if _, ok := s.ChatEventloops.Load(req.GroupId); ok {
		return &imchatevloop.CreateGroupResponse{
			Success: true,
			Version: currentVersion,
		}, nil
	}
	lp := evloop.NewChatEvLoopAndStart(&evloop.NewChatEvLoopInput{
		ChatId:  req.GroupId,
		OwnerId: req.GroupOwnerId,
	})
	s.ChatEventloops.Store(req.GroupId, lp)
	return
}
func (s *RPCServer) UniversalGroupEvloopRequest(ctx context.Context,
	req *imchatevloop.UniversalGroupEvloopRequestReq) (res *imchatevloop.UniversalGroupEvloopRequestResp, err error) {
	currentVersion := s.CurrentVersion.Load()
	if req.Version > currentVersion {
		s.CurrentVersionMu.Lock()
		currentVersion = s.CurrentVersion.Load()
		if req.Version > currentVersion {
			s.CurrentVersion.Store(req.Version)
		}
		s.CurrentVersionMu.Unlock()
	}

	lp, ok := s.ChatEventloops.Load(req.GroupId)
	if !ok {
		// 群不存在, 此时发生迁移, 加迁移锁
		unlockFunc := s.MoveLock.Lock(req.GroupId)
		// double check
		lp, ok = s.ChatEventloops.Load(req.GroupId)
		if ok {
			unlockFunc()
		} else {
			// 还是不存在, 不释放迁移lock, 做迁移, todo: 重试流程, 这里重试是幂等的, 理论可以无限重试
			resp, err := s.RelayCli.GetLastVersionConfig(context.Background(), &relay.GetLastVersionConfigReq{
				GroupId: req.GroupId,
			})
			if err != nil {
				unlockFunc()
				return nil, err
			}
			addrPort := resp.EvloopServerAddrPort
			evloopCli := api.MustNewChatEvLoopCliFromAdAddr(addrPort)
			// todo: 这个是可以幂等重试的
			migrateResp, err := evloopCli.DoMigrate(context.Background(), &imchatevloop.DoMigrateReq{
				Version: currentVersion,
				GroupId: req.GroupId,
			})
			if err != nil {
				unlockFunc()
				return nil, err
			}
			lp = evloop.NewMigrateEvLoop(migrateResp)
			s.ChatEventloops.Store(req.GroupId, lp)
		}
	}

	// 有loop, 开始干活
	evl := lp.(*evloop.ChatEvLoop)

	return
}

func (s *RPCServer) DoMigrate(ctx context.Context, req *imchatevloop.DoMigrateReq) (res *imchatevloop.DoMigrateResp, err error) {
	currentVersion := s.CurrentVersion.Load()
	if req.Version > currentVersion {
		s.CurrentVersionMu.Lock()
		currentVersion = s.CurrentVersion.Load()
		if req.Version > currentVersion {
			s.CurrentVersion.Store(req.Version)
		}
		s.CurrentVersionMu.Unlock()
	}

}

func (s *RPCServer) MigrateDone(ctx context.Context, req *imchatevloop.MigrateDoneReq) (res *imchatevloop.MigrateDoneResp, err error) {

}
