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
)

// 申请加入群
func (s *RPCServer) ApplyGroup(ctx context.Context, req *imrelation.ApplyGroupReq) (res *imrelation.ApplyGroupResp, err error) {
	now := time.Now()

	// 先检查群信息, 如果群不存在直接返回
	group, err := db.GetGroupInfo(s.DB.DB(), req.GroupId)
	if err != nil {
		log.Printf("failed to get group info: %v\n", err)
		return nil, err
	}
	if group == nil {
		go func() {
			push.ApplyGroupResp(req.Session, &push.ApplyGroupRespInput{
				Code:     imrelation.ApplyGroupResp_APPLY_GROUP_RESP_CODE_NO_GROUP,
				EchoCode: req.EchoCode,
			})
		}()
		return &imrelation.ApplyGroupResp{
			Code: imrelation.ApplyGroupResp_APPLY_GROUP_RESP_CODE_NO_GROUP,
		}, nil
	}

	// 一个用户关于一个群只有一条ApplyModel, 这里有个并发问题, 这里选择直接群维度加锁
	// 锁group
	le, err := s.RdsAct.LockGroup(req.GroupId, time.Second*1)
	if err != nil {
		return nil, err
	}
	defer le.UnLock()

	txn := s.DB.Txn()
	defer txn.Rollback()

	// TODO: 目前先忽略群解散的逻辑, 先跑起来大流程再加退群的细节, 现在不支持退群
	// // 如果群已解散, 不能申请
	// if group.Disbaned {
	// 	go func() {
	// 		push.ApplyGroupResp(req.Session, &push.ApplyGroupRespInput{
	// 			Code:     imrelation.ApplyGroupResp_APPLY_GROUP_RESP_CODE_GROUP_DISBANDED,
	// 			EchoCode: req.EchoCode,
	// 		})
	// 	}()
	// 	return &imrelation.ApplyGroupResp{
	// 		Code: imrelation.ApplyGroupResp_APPLY_GROUP_RESP_CODE_GROUP_DISBANDED,
	// 	}, nil
	// }

	relation, err := db.GetRelationByUsernameGroupId(txn, req.Session.Username, group.GroupId)
	if err != nil {
		log.Printf("failed to GetRelationByUsernameGroupId: %v\n", err)
		return nil, err
	}

	fmt.Println(relation)

	// 如果是member或owner, 则不能apply, user in group错误
	if relation != nil && (relation.Status == base.RelationStatus_RELATION_STATUS_MEMBER ||
		relation.Status == base.RelationStatus_RELATION_STATUS_OWNER) {
		go func() {
			push.ApplyGroupResp(req.Session, &push.ApplyGroupRespInput{
				Code:     imrelation.ApplyGroupResp_APPLY_GROUP_RESP_CODE_USER_IN_GROUP,
				EchoCode: req.EchoCode,
			})
		}()
		return &imrelation.ApplyGroupResp{
			Code: imrelation.ApplyGroupResp_APPLY_GROUP_RESP_CODE_USER_IN_GROUP,
		}, nil
	}

	apply, err := db.GetApplyByUsernameAndGroupId(txn, req.Session.Username, group.GroupId)
	if err != nil {
		log.Printf("failed to GetApplyByUsernameAndGroupId: %v\n", err)
		return nil, err
	}

	fmt.Println(apply)

	if apply == nil {
		apply = &model.ApplyModel{
			OwnerId:      req.Session.Username,
			GroupId:      group.GroupId,
			ApplyVersion: 0,
			ApplyMsg:     "",
			CreatedAt:    now.UnixMilli(),
			UpdatedAt:    now.UnixMilli(),
			Status:       0,
			GroupOwnerId: group.OwnerId,
		}
		inserted, err := db.InsertApply(txn, apply)
		if err != nil || !inserted {
			if err == nil {
				err = errors.New("dupkey")
			}
			log.Printf("failed to insert apply: %v\n", err)
			return nil, err
		}
	}
	if err != nil {
		log.Printf("failed to insert or lock apply entry: %v\n", err)
		return nil, err
	}

	// none, pending reject三种状态, 更新为pendding状态
	apply.ApplyVersion++
	apply.ApplyMsg = req.ApplyMsg
	apply.Status = base.ApplyStatus_APPLY_STATUS_PENDING
	apply.UpdatedAt = now.UnixMilli()
	err = db.UpdateApply(txn, apply)
	if err != nil {
		log.Printf("failed to update apply: %v\n", err)
		return nil, err
	}
	err = txn.Commit().Error
	if err != nil {
		log.Printf("failed to commit apply: %v\n", err)
		return nil, err
	}

	// 给申请方, 推多端notify
	go func() {
		push.ApplyGroupNotify(&push.ApplyGroupNotifyInput{
			AuthRoute:    s.AuthRouteCli,
			OwnerId:      group.OwnerId,
			Username:     apply.OwnerId,
			GroupId:      fmt.Sprint(apply.GroupId),
			ApplyMsg:     apply.ApplyMsg,
			ApplyVersion: apply.ApplyVersion,
			ApplyAt:      apply.UpdatedAt,
		})
	}()

	// 推带有echo code的resp
	go func() {
		push.ApplyGroupResp(req.Session, &push.ApplyGroupRespInput{
			EchoCode: req.EchoCode,
			Code:     imrelation.ApplyGroupResp_APPLY_GROUP_RESP_CODE_OK,
		})
	}()

	return &imrelation.ApplyGroupResp{
		Code: imrelation.ApplyGroupResp_APPLY_GROUP_RESP_CODE_OK,
	}, nil
}
