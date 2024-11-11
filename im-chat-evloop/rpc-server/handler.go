package rpcserver

import (
	"context"
	"fmt"
	"pigeon/common/keylock"
	"pigeon/im-chat-evloop/api"
	"pigeon/im-chat-evloop/bizpush"
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
	BPush    *bizpush.BizPusher

	// 读写current version, 需要加锁
	CurrentVersionMu sync.Mutex
	CurrentVersion   atomic.Int64

	// key: group id, value: chat eventloop
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
	lp, createdAt := evloop.NewChatEvLoopAndStart(&evloop.NewChatEvLoopInput{
		ChatId:  req.GroupId,
		OwnerId: req.GroupOwnerId,
		PushMan: s.BPush,
	})
	s.ChatEventloops.Store(req.GroupId, lp)
	return &imchatevloop.CreateGroupResponse{
		Success:   true,
		Version:   currentVersion,
		CreatedAt: createdAt,
	}, nil
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

	fmt.Println("universal group request in")

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
				GroupId: req.GroupId,
			})
			if err != nil {
				unlockFunc()
				return nil, err
			}
			lp = evloop.NewMigrateEvLoop(migrateResp, s.BPush)

			// todo这里可以幂等重试
			_, err = evloopCli.MigrateDone(context.Background(), &imchatevloop.MigrateDoneReq{
				GroupId: req.GroupId,
			})
			if err != nil {
				unlockFunc()
				return nil, err
			}

			s.ChatEventloops.Store(req.GroupId, lp)
			unlockFunc()
		}
	}

	// 有loop, 开始干活
	evl := lp.(*evloop.ChatEvLoop)
	// err != nil, 可能状态是stop或者migrating, 此时发生了迁移
	// 那么需要拿新version返回过去
	output, err := evl.QueueMessage(req.Input)
	if err != nil {
		newVersion := s.CurrentVersion.Load()
		return &imchatevloop.UniversalGroupEvloopRequestResp{
			Success: false,
			Version: newVersion,
		}, nil
	}

	fmt.Println("universal group request out")

	// 信息发送成功
	return &imchatevloop.UniversalGroupEvloopRequestResp{
		Success: true,
		Version: currentVersion,
		Output:  output.Output,
	}, nil
}

func (s *RPCServer) DoMigrate(ctx context.Context, req *imchatevloop.DoMigrateReq) (res *imchatevloop.DoMigrateResp, err error) {
	// 找到eventloop
	evlAny, ok := s.ChatEventloops.Load(req.GroupId)
	if !ok {
		return &imchatevloop.DoMigrateResp{
			Ok: false,
		}, nil
	}

	out, err := evlAny.(*evloop.ChatEvLoop).Move()
	if err != nil {
		return &imchatevloop.DoMigrateResp{
			Ok: false,
		}, nil
	}

	rlations := make(map[string]*imchatevloop.DoMigrateResp_RelationInfo)
	for k, v := range out.Relations {
		rlations[k] = &imchatevloop.DoMigrateResp_RelationInfo{
			Relation: v,
		}
	}
	subs := make(map[string]*imchatevloop.DoMigrateResp_UserSubscribeEntry)
	for k, v := range out.Subs {
		subs[k] = &imchatevloop.DoMigrateResp_UserSubscribeEntry{
			Entries: make([]*imchatevloop.DoMigrateResp_UserSubscribeEntry_SubscribeEntry, 0),
		}
		for _, s := range v {
			subs[k].Entries = append(subs[k].Entries, &imchatevloop.DoMigrateResp_UserSubscribeEntry_SubscribeEntry{
				OnSubRelationVersion: s.SubRelationVersion,
				Session:              s.Entry,
			})
		}

	}
	return &imchatevloop.DoMigrateResp{
		Ok:          true,
		GroupId:     req.GroupId,
		OwnerId:     out.OwnerId,
		SeqId:       out.SeqId,
		Relations:   rlations,
		Subscribers: subs,
	}, nil

}

func (s *RPCServer) MigrateDone(ctx context.Context, req *imchatevloop.MigrateDoneReq) (
	*imchatevloop.MigrateDoneResp, error) {

	// 找到eventloop
	evlAny, ok := s.ChatEventloops.Load(req.GroupId)
	// 如果没找到, 返回true
	if !ok {
		return &imchatevloop.MigrateDoneResp{
			Ok: true,
		}, nil
	}

	ev := evlAny.(*evloop.ChatEvLoop)
	err := ev.Stop()
	s.ChatEventloops.Delete(req.GroupId)

	return &imchatevloop.MigrateDoneResp{
		Ok: err != nil,
	}, nil
}
