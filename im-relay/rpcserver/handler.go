package rpcserver

import (
	"context"
	"errors"
	"log"
	"strings"

	chatevloopconfig "pigeon/common/chatevloop_config"
	"pigeon/im-relay/api"
	"pigeon/im-relay/handle"
	"pigeon/im-relay/handle/chat"
	"pigeon/im-relay/handle/echo"
	"pigeon/kitex_gen/service/imchatevloop"
	"pigeon/kitex_gen/service/imrelation/imrelation"
	"pigeon/kitex_gen/service/imrelay"
)

type RPCContext struct {
	EvCfgWatcher *chatevloopconfig.ChatevWatcher
	RelationCli  imrelation.Client
}

type RPCServer struct {
	RPCContext
}

func (s *RPCServer) BizMessage(ctx context.Context,
	req *imrelay.BizMessageReq) (*imrelay.BizMessageResp, error) {
	// 请求直接异步出去
	go func() {
		s.handleBizMessage(req)
	}()
	return &imrelay.BizMessageResp{}, nil
}

func (s *RPCServer) handleBizMessage(req *imrelay.BizMessageReq) {
	splits := strings.Split(req.Biz, "-")
	if len(splits) < 1 {
		return
	}
	switch splits[0] {
	case "echo":
		echo.HandleEcho(&handle.HandleContext{RelationCli: s.RelationCli, EvCfgWatcher: s.EvCfgWatcher}, req)
	case "chat":
		chat.HandleChat(&handle.HandleContext{RelationCli: s.RelationCli, EvCfgWatcher: s.EvCfgWatcher}, req)
	default:
	}
}

// 由im-relation调用, 创造群聊event loop
func (s *RPCServer) CreateChatEventLoop(ctx context.Context,
	req *imrelay.CreateChatEventLoopReq) (*imrelay.CreateChatEventLoopResp, error) {
	for {
		nodeEntry, version := s.EvCfgWatcher.GetNode(req.GroupId)
		if nodeEntry == nil {
			log.Printf("consistent ring get node failed, no node available\n")
			return nil, errors.New("no node availabel")
		}

		evloopCli := api.MustNewChatEvLoopCliFromAdAddr(nodeEntry.IPPort)
		resp, err := evloopCli.CreateGroup(context.Background(), &imchatevloop.CreateGroupRequest{
			Version:      version,
			GroupId:      req.GroupId,
			GroupOwnerId: req.OwnerId,
		})
		if err != nil {
			log.Printf("create group failed, err: %v\n", err)
			return nil, errors.New("create group failed")
		}
		if !resp.Success || resp.Version != version {
			s.EvCfgWatcher.ForceUpdate(resp.Version)
			log.Printf("version not match, current version: %d, server version: %d, retrying...\n", version, resp.Version)
			continue
		}
		return &imrelay.CreateChatEventLoopResp{
			Success: true,
		}, nil
	}
}

func (s *RPCServer) RedirectToChatEventLoop(ctx context.Context,
	req *imrelay.RedirectToChatEventLoopReq) (*imrelay.RedirectToChatEventLoopResp, error) {

	for {
		evloopSpec, version := s.EvCfgWatcher.GetNode(req.GroupId)
		cli := api.MustNewChatEvLoopCliFromAdAddr(evloopSpec.IPPort)
		resp, err := cli.UniversalGroupEvloopRequest(context.Background(), &imchatevloop.UniversalGroupEvloopRequestReq{
			Version: version,
			GroupId: req.GroupId,
			Input:   req.Input,
		})
		if err != nil {
			return nil, err
		}
		if !resp.Success {
			s.EvCfgWatcher.ForceUpdate(resp.Version)
			continue
		}
		return &imrelay.RedirectToChatEventLoopResp{
			Success: true,
			Output:  resp.Output,
		}, nil
	}
}

func (s *RPCServer) GetLastVersionConfig(ctx context.Context,
	req *imrelay.GetLastVersionConfigReq) (res *imrelay.GetLastVersionConfigResp, err error) {
	n, v := s.EvCfgWatcher.GetLastVersionNode(req.GroupId)
	if v != req.Version {
		s.EvCfgWatcher.ForceUpdate(req.Version)
		n, _ = s.EvCfgWatcher.GetLastVersionNode(req.GroupId)
	}
	return &imrelay.GetLastVersionConfigResp{
		EvloopServerAddrPort: n.IPPort,
	}, nil
}
