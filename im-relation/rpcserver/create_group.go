package rpcserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"pigeon/im-relation/db"
	"pigeon/im-relation/db/model"
	"pigeon/im-relation/push"
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
		CreatedAt: time.Now().UnixMilli(),
	}
	err = db.CreateGroup(txn, &group)
	if err != nil {
		log.Printf("failed to create group: %v\n", err)
		return nil, err
	}

	err = db.InsertRelation(txn, &model.RelationModel{
		OwnerId:         req.Session.Username,
		GroupId:         groupId,
		Status:          base.RelationStatus_RELATION_STATUS_OWNER,
		ChangeType:      base.RelationChangeType_RELATION_CHANGE_TYPE_CREATE_GROUP,
		RelationVersion: 1,
		CreatedAt:       now.UnixMilli(),
		UpdatedAt:       now.UnixMilli(),
	})
	if err != nil {
		log.Printf("failed to insert or select for update relation: %v\n", err)
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

	go func() {
		// push里面自带重试逻辑
		push.CreateGroupResp(req.Session, &push.CreateGroupRespInput{
			EchoCode:  req.EchoCode,
			OwnerId:   group.OwnerId,
			GroupId:   fmt.Sprint(group.Id),
			CreatedAt: group.CreatedAt,
		})
	}()

	return &imrelation.CreateGroupResp{
		GroupId:  fmt.Sprint(group.Id),
		CreateAt: group.CreatedAt,
	}, nil
}
