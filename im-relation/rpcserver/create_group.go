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

	ownerId := req.Session.Username
	txn := s.DB.Txn()
	defer txn.Rollback()
	var group = model.GroupModel{
		OwnerId:    ownerId,
		CreatedAt:  time.Now().UnixMilli(),
		Disbaned:   false,
		DisbanedAt: 0,
	}
	err = db.CreateGroup(txn, &group)
	if err != nil {
		log.Printf("failed to create group: %v\n", err)
		return nil, err
	}

	_, err = db.InsertOrSelectForUpdateRelationByUsernameGroupId(txn, &model.RelationModel{
		OwnerId:         req.Session.Username,
		GroupId:         group.Id,
		Status:          imrelation.RelationStatus_RELATION_STATUS_OWNER,
		ChangeType:      imrelation.RelationChangeType_RELATION_CHNAGE_TYPE_NONE,
		RelationCounter: 1,
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
