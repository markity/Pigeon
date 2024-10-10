package rpcserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"pigeon/im-auth-route/rds"
	"pigeon/im-relation/db"
	"pigeon/im-relation/db/model"
	"pigeon/kitex_gen/service/imrelation"
	relay "pigeon/kitex_gen/service/imrelay"
	"strconv"
	"time"

	"pigeon/kitex_gen/service/imrelay/imrelay"
)

type RPCContext struct {
	DB       *db.DB
	Rds      *rds.RdsAction
	RelayCli imrelay.Client
}

type RPCServer struct {
	RPCContext
}

func (s *RPCServer) CreateGroup(ctx context.Context, req *imrelation.CreateGroupReq) (res *imrelation.CreateGroupResp, err error) {
	ownerId := req.Session.Username
	txn := s.DB.Txn()
	defer txn.Rollback()
	var group = model.GroupModel{
		OwnerId:   ownerId,
		CreatedAt: time.Now().UnixMilli(),
	}
	err = db.CreateGroup(txn, &group)
	if err != nil {
		log.Printf("failed to create group: %v\n", err)
		return nil, err
	}
	resp, err := s.RelayCli.CreateChatEventLoop(context.Background(), &relay.CreateChatEventLoopReq{
		GroupId: fmt.Sprint(group.Id),
		OwnerId: ownerId,
	})
	if err != nil {
		log.Printf("failed to create chat event loop: %v\n", err)
		return nil, err
	}
	if !resp.Success {
		log.Printf("failed to create chat event loop, not success\n")
		return nil, errors.New("unexpected error")
	}
	err = txn.Commit().Error
	if err != nil {
		log.Printf("failed to commit transaction: %v\n", err)
		return nil, err
	}

	return &imrelation.CreateGroupResp{
		GroupId: fmt.Sprint(group.Id),
	}, nil
}

func (s *RPCServer) GetGroupInfo(ctx context.Context, req *imrelation.GetGroupInfoReq) (res *imrelation.GetGroupInfoResp, err error) {
	txn := s.DB.Txn()
	defer txn.Rollback()
	groupId, err := strconv.ParseInt(req.GroupId, 10, 64)
	if err != nil {
		log.Printf("failed to parse group id: %v\n", err)
		return nil, err
	}
	group, err := db.GetGroupInfo(txn, groupId)
	if err != nil {
		log.Printf("failed to get group info: %v\n", err)
		return nil, err
	}
	if group == nil {
		return &imrelation.GetGroupInfoResp{
			Exists: false,
			Info:   nil,
		}, nil
	}

	return &imrelation.GetGroupInfoResp{
		Exists: true,
		Info: &imrelation.GroupInfo{
			GroupId:    req.GroupId,
			OwnerId:    group.OwnerId,
			CreateAt:   group.CreatedAt,
			Disbanded:  group.Disbaned,
			DisbanedAt: group.DisbanedAt,
		},
	}, nil
}
func (s *RPCServer) FetchAllRelations(ctx context.Context, req *imrelation.FetchAllRelationsReq) (res *imrelation.FetchAllRelationsResp, err error) {
	return
}
func (s *RPCServer) FetchAllApplications(ctx context.Context, req *imrelation.FetchAllApplicationsReq) (res *imrelation.FetchAllApplicationsResp, err error) {
	return
}
func (s *RPCServer) ApplyGroup(ctx context.Context, req *imrelation.ApplyGroupReq) (res *imrelation.ApplyGroupReqResp, err error) {
	return
}
func (s *RPCServer) HandleApply(ctx context.Context, req *imrelation.HandleApplyReq) (res *imrelation.HandleApplyResp, err error) {
	return
}
func (s *RPCServer) QuitGroup(ctx context.Context, req *imrelation.QuitGroupReq) (res *imrelation.QuitGroupResp, err error) {
	return
}
