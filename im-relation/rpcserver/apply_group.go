package rpcserver

import (
	"context"
	"fmt"
	"log"
	"strconv"
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

	// 一个用户关于一个群只有一条ApplyModel, 这里有个并发问题, 这里使用mysql 1062错误码解决这个问题
	txn := s.DB.Txn()
	defer txn.Rollback()

	groupIdInt, err := strconv.ParseInt(req.GroupId, 10, 64)
	if err != nil {
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

	// 查询group信息
	group, err := db.GetGroupInfo(txn, groupIdInt)
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

	// 如果群已解散, 不能申请
	if group.Disbaned {
		go func() {
			push.ApplyGroupResp(req.Session, &push.ApplyGroupRespInput{
				Code:     imrelation.ApplyGroupResp_APPLY_GROUP_RESP_CODE_GROUP_DISBANDED,
				EchoCode: req.EchoCode,
			})
		}()
		return &imrelation.ApplyGroupResp{
			Code: imrelation.ApplyGroupResp_APPLY_GROUP_RESP_CODE_GROUP_DISBANDED,
		}, nil
	}

	relation, err := db.InsertOrSelectForUpdateRelationByUsernameGroupId(txn, &model.RelationModel{
		OwnerId:         req.Session.Username,
		GroupId:         groupIdInt,
		Status:          base.RelationStatus_RELATION_STATUS_NOT_IN_GROUP,
		RelationCounter: 0,
		CreatedAt:       now.UnixMilli(),
		UpdatedAt:       now.UnixMilli(),
	})
	if err != nil {
		log.Printf("failed to insert or lock relation entry: %v\n", err)
		return nil, err
	}

	// 如果是member或owner, 则不能apply, user in group错误
	if relation.Status == base.RelationStatus_RELATION_STATUS_MEMBER ||
		relation.Status == base.RelationStatus_RELATION_STATUS_OWNER {
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

	apply, err := db.InsertOrSelectForUpdateApplyByUsernameGroupId(txn, &model.ApplyModel{
		OwnerId:      req.Session.Username,
		GroupId:      groupIdInt,
		ApplyCounter: 0,
		ApplyMsg:     "",
		CreatedAt:    now.UnixMilli(),
		UpdatedAt:    now.UnixMilli(),
		Status:       0,
		GroupOwnerId: group.OwnerId,
	})
	if err != nil {
		log.Printf("failed to insert or lock apply entry: %v\n", err)
		return nil, err
	}

	// none, pending reject三种状态, 更新为pendding状态
	apply.ApplyCounter++
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
			ApplyVersion: apply.ApplyCounter,
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
