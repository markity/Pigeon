package rpcserver

import (
	"context"
	"pigeon/kitex_gen/service/imchatevloop"
	"sync"
	"sync/atomic"
)

// 创建群聊
type RPCServer struct {
	CurrentVersionMu sync.Mutex
	CurrentVersion   atomic.Int64

	ChatEventloops sync.Map
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
	if _, ok := s.ChatEventloops.Load(req.GroupId); ok {
		return &imchatevloop.CreateGroupResponse{
			Success: true,
			Version: currentVersion,
		}, nil
	}
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
	return
}
