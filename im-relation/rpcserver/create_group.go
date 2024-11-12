package rpcserver

import (
	"context"
	"errors"
	"log"
	"time"

	"pigeon/im-relation/bizpush"
	"pigeon/im-relation/db"
	"pigeon/im-relation/db/model"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/imrelation"
	relay "pigeon/kitex_gen/service/imrelay"
)

/*
create group逻辑概要:
begin txn
insert拿到自增主键: groupId TODO: 这里以后优化, 增加id生成器来生成groupId, 目前是直接用mysql自增主键
调用imrelay.CreateChatEventLoop, 创建chat event loop
commit txn
return info
*/

// TODO: im-relay可能因为网络错误多次调用create group, 此处是否应该幂?
func (s *RPCServer) CreateGroup(ctx context.Context, req *imrelation.CreateGroupReq) (res *imrelation.CreateGroupResp, err error) {
	now := time.Now()
	// TODO: 此处应该有重试/降级策略, 当redis出故障的时候生成group id怎么办
	groupId, err := s.RdsAct.GenerateGroupId()
	if err != nil {
		log.Printf("failed to generate group id: %v\n", err)
		return nil, err
	}

	// // 锁group
	// le, err := s.RdsAct.LockGroup(groupId, time.Second*1)
	// if err != nil {
	// 	return nil, err
	// }
	// defer le.UnLock()

	ownerId := req.Session.Username
	txn := s.DB.Txn()
	defer txn.Rollback()
	var group = model.GroupModel{
		GroupId:   groupId,
		OwnerId:   ownerId,
		CreatedAt: now.UnixMilli(),
	}
	err = db.CreateGroup(txn, &group)
	if err != nil {
		log.Printf("failed to create group: %v\n", err)
		return nil, err
	}

	relation := &model.RelationModel{
		OwnerId:         req.Session.Username,
		GroupId:         groupId,
		Status:          base.RelationStatus_RELATION_STATUS_OWNER,
		ChangeType:      base.RelationChangeType_RELATION_CHANGE_TYPE_CREATE_GROUP,
		RelationVersion: 1,
		CreatedAt:       now.UnixMilli(),
		UpdatedAt:       now.UnixMilli(),
	}
	err = db.InsertRelation(txn, relation)
	if err != nil {
		log.Printf("failed to insert or select for update relation: %v\n", err)
		return nil, err
	}

	resp, err := s.RelayCli.CreateChatEventLoop(context.Background(), &relay.CreateChatEventLoopReq{
		GroupId: group.GroupId,
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

	go func() {
		// push里面自带重试逻辑
		s.BPush.CreateGroupResp(&bizpush.CreateGroupRespInput{
			Session:   req.Session,
			EchoCode:  req.EchoCode,
			OwnerId:   group.OwnerId,
			GroupId:   group.GroupId,
			CreatedAt: group.CreatedAt,
		})
	}()

	// 给此用户的所有session推送关于这个群聊的关系
	go func() {
		s.BPush.RelationChangeNotify(&bizpush.RelationChangeNotifyInput{
			Username:   req.Session.Username,
			GroupId:    groupId,
			Version:    1,
			Status:     base.RelationStatus_RELATION_STATUS_OWNER,
			UpdatedAt:  relation.UpdatedAt,
			ChangeType: base.RelationChangeType_RELATION_CHANGE_TYPE_CREATE_GROUP,
		})
	}()

	return &imrelation.CreateGroupResp{
		GroupId:  group.GroupId,
		CreateAt: group.CreatedAt,
	}, nil
}
