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
	"pigeon/kitex_gen/service/evloopio"
	"pigeon/kitex_gen/service/imrelation"
	relay "pigeon/kitex_gen/service/imrelay"
)

// 接受/拒绝申请
func (s *RPCServer) HandleApply(ctx context.Context, req *imrelation.HandleApplyReq) (
	res *imrelation.HandleApplyResp, err error) {
	now := time.Now()
	groupIdInt, err := strconv.ParseInt(req.GroupId, 10, 64)
	if err != nil {
		log.Printf("failed to parse group id: %v\n", err)
		return nil, err
	}

	txn := s.DB.Txn()
	defer txn.Rollback()

	groupInfo, err := db.GetGroupInfo(txn, groupIdInt)
	if err != nil {
		log.Printf("failed to get group info: %v\n", err)
		return nil, err
	}
	if groupInfo.OwnerId != req.Session.Username {
		go func() {
			push.HandleApplyResp(req.Session, &push.HandleApplyRespInput{
				EchoCode: req.EchoCode,
				Code:     imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_NO_PERMISSION,
			})
		}()
		return &imrelation.HandleApplyResp{
			Code: imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_NO_PERMISSION,
		}, nil
	}

	relation, err := db.InsertOrSelectForUpdateRelationByUsernameGroupId(txn, &model.RelationModel{
		OwnerId:         req.UserId,
		GroupId:         groupIdInt,
		Status:          base.RelationStatus_RELATION_STATUS_NOT_IN_GROUP,
		ChangeType:      base.RelationChangeType_RELATION_CHNAGE_TYPE_NONE,
		RelationCounter: 0,
		CreatedAt:       now.UnixMilli(),
		UpdatedAt:       now.UnixMilli(),
	})
	if err != nil {
		log.Printf("failed to insert or lock relation entry: %v\n", err)
		return nil, err
	}

	apply, err := db.InsertOrSelectForUpdateApplyByUsernameGroupId(txn, &model.ApplyModel{
		OwnerId:      req.UserId,
		GroupId:      groupIdInt,
		ApplyCounter: 0,
		ApplyMsg:     "",
		CreatedAt:    now.UnixMilli(),
		UpdatedAt:    now.UnixMilli(),
		Status:       base.ApplyStatus_APPLY_STATUS_NONE,
		GroupOwnerId: groupInfo.OwnerId,
	})
	if err != nil {
		log.Printf("failed to insert or lock apply entry: %v\n", err)
		return nil, err
	}
	applyAt := apply.UpdatedAt

	fmt.Println(apply.Status.String())
	if apply.Status != base.ApplyStatus_APPLY_STATUS_PENDING {
		go func() {
			push.HandleApplyResp(req.Session, &push.HandleApplyRespInput{
				EchoCode: req.EchoCode,
				Code:     imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_NO_APPLY,
			})
		}()
		return &imrelation.HandleApplyResp{
			Code: imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_NO_APPLY,
		}, nil
	}

	// 如果已经disbaned, 无论接受还是拒绝, 都返回ok, 但是不更新关系
	if groupInfo.Disbaned {
		apply.Status = base.ApplyStatus_APPLY_STATUS_GROUP_DISBANDED
		apply.ApplyCounter++
		apply.UpdatedAt = now.UnixMilli()

		err := db.UpdateApply(txn, apply)
		if err != nil {
			log.Printf("failed to update apply: %v\n", err)
			return nil, err
		}

		err = txn.Commit().Error
		if err != nil {
			log.Printf("failed to commit apply: %v\n", err)
			return nil, err
		}

		go func() {
			push.HandleApplyResp(req.Session, &push.HandleApplyRespInput{
				EchoCode:     req.EchoCode,
				Code:         imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_OK,
				ApplyMsg:     apply.ApplyMsg,
				ApplyAt:      applyAt,
				ApplyVersion: apply.ApplyCounter,
				ApplyStatus:  apply.Status,
				HandleAt:     now.UnixMilli(),

				RelationVersion: 0,
			})
		}()

		return &imrelation.HandleApplyResp{
			Code:            imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_OK,
			RelationVersion: 0,
			ApplyMsg:        apply.ApplyMsg,
			ApplyAt:         applyAt,
			HandleAt:        now.UnixMilli(),
			ApplyVersion:    apply.ApplyCounter,
			ApplyStatus:     apply.Status,
		}, nil
	}

	// 群没有disbanded

	if !req.Accept {
		apply.Status = base.ApplyStatus_APPLY_STATUS_REJECTED
		apply.ApplyCounter++
		apply.UpdatedAt = now.UnixMilli()
		err := db.UpdateApply(txn, apply)
		if err != nil {
			log.Printf("failed to update apply: %v\n", err)
			return nil, err
		}
		err = txn.Commit().Error
		if err != nil {
			log.Printf("failed to commit apply: %v\n", err)
			return nil, err
		}

		// 拒绝更新ok, 需要把消息推送group owner
		go func() {
			push.HandleApplyNotify(&push.HandleApplyNotifyInput{
				AuthRoute:    s.AuthRouteCli,
				OwnerId:      groupInfo.OwnerId,
				Username:     apply.OwnerId,
				GroupId:      req.GroupId,
				ApplyVersion: apply.ApplyCounter,
				ApplyMsg:     apply.ApplyMsg,
				ApplyStatus:  apply.Status,
				ApplyAt:      applyAt,
				HandleAt:     apply.UpdatedAt,
			})
		}()

		go func() {
			push.HandleApplyResp(req.Session, &push.HandleApplyRespInput{
				EchoCode: req.EchoCode,

				Code: imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_OK,

				RelationVersion: relation.RelationCounter,

				ApplyMsg:     apply.ApplyMsg,
				ApplyAt:      applyAt,
				ApplyVersion: apply.ApplyCounter,

				HandleAt: now.UnixMilli(),
			})
		}()

		return &imrelation.HandleApplyResp{
			Code:            imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_OK,
			RelationVersion: relation.RelationCounter,

			ApplyMsg:     apply.ApplyMsg,
			ApplyAt:      applyAt,
			ApplyVersion: apply.ApplyCounter,

			HandleAt: now.UnixMilli(),
		}, nil
	}

	// accept, 申请成功, 尝试将信息打入evloop

	// pending状态, 直接打入evloop
	resp, err := s.RelayCli.RedirectToChatEventLoop(context.Background(),
		&relay.RedirectToChatEventLoopReq{
			GroupId: fmt.Sprint(apply.GroupId),
			Input: &evloopio.UniversalGroupEvloopInput{
				Input: &evloopio.UniversalGroupEvloopInput_AlterGroupMember{
					AlterGroupMember: &evloopio.AlterGroupMemberRequest{
						GroupId:         req.GroupId,
						IsAdd:           true,
						MemberId:        apply.OwnerId,
						RelationVersion: apply.ApplyCounter,
					},
				},
			},
		})
	if err != nil {
		log.Printf("failed to RedirectToChatEventLoop: %v\n", err)
		return nil, err
	}
	out := resp.Output.Output.(*evloopio.UniversalGroupEvloopOutput_AlterGroupMember)
	switch out.AlterGroupMember.Code {
	case evloopio.AlterGroupMemberResponse_GROUP_DISBANDED:
		apply.Status = base.ApplyStatus_APPLY_STATUS_GROUP_DISBANDED
		apply.ApplyCounter++
		apply.UpdatedAt = now.UnixMilli()
		err := db.UpdateApply(txn, apply)
		if err != nil {
			log.Printf("failed to update apply: %v\n", err)
			return nil, err
		}

		err = txn.Commit().Error
		if err != nil {
			log.Printf("failed to commit apply: %v\n", err)
			return nil, err
		}

		go func() {
			push.HandleApplyResp(req.Session, &push.HandleApplyRespInput{
				EchoCode:     req.EchoCode,
				Code:         imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_OK,
				ApplyMsg:     apply.ApplyMsg,
				ApplyAt:      applyAt,
				ApplyVersion: apply.ApplyCounter,
				ApplyStatus:  apply.Status,
				HandleAt:     now.UnixMilli(),

				RelationVersion: 0,
			})
		}()
		return &imrelation.HandleApplyResp{
			Code:         imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_OK,
			ApplyMsg:     apply.ApplyMsg,
			ApplyAt:      applyAt,
			ApplyVersion: apply.ApplyCounter,
			ApplyStatus:  apply.Status,
			HandleAt:     now.UnixMilli(),

			RelationVersion: 0,
		}, nil
	case evloopio.AlterGroupMemberResponse_OK:
		// accept, 申请成功
		err := db.UpdateApply(txn, apply)
		if err != nil {
			log.Printf("failed to update apply: %v\n", err)
			return nil, err
		}

		relation.Status = base.RelationStatus_RELATION_STATUS_NOT_IN_GROUP
		relation.ChangeType = base.RelationChangeType_RELATION_CHANGE_TYPE_OWNER_ACCEPT
		relation.RelationCounter++
		relation.UpdatedAt = now.UnixMilli()
		err = db.UpdateRelation(txn, relation)
		if err != nil {
			log.Printf("failed to update relation: %v\n", err)
			return nil, err
		}

		err = txn.Commit().Error
		if err != nil {
			log.Printf("failed to commit apply: %v\n", err)
			return nil, err
		}

		// 发送relation变更notify
		go func() {
			push.RelationChangeNotify(&push.RelationChangeNotifyInput{
				AuthRoute: s.AuthRouteCli,

				Username:   relation.OwnerId,
				GroupId:    fmt.Sprint(relation.GroupId),
				Version:    relation.RelationCounter,
				Status:     base.RelationStatus_RELATION_STATUS_MEMBER,
				ChangeType: relation.ChangeType, // relation变更场景, 是因为群主接受申请
				ChangeAt:   now.UnixMilli(),
			})
		}()

		// 给group owner发处理apply的通知
		go func() {
			push.HandleApplyNotify(&push.HandleApplyNotifyInput{
				AuthRoute:    s.AuthRouteCli,
				OwnerId:      groupInfo.OwnerId,
				Username:     apply.OwnerId,
				GroupId:      req.GroupId,
				ApplyVersion: apply.ApplyCounter,
				ApplyMsg:     apply.ApplyMsg,
				ApplyStatus:  apply.Status,
				ApplyAt:      applyAt,
				HandleAt:     now.UnixMilli(),
			})
		}()

		// 推消息, 告诉group owner以及新relation的建立
		go func() {
			push.HandleApplyResp(req.Session, &push.HandleApplyRespInput{
				EchoCode:        req.EchoCode,
				Code:            imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_OK,
				RelationVersion: relation.RelationCounter,
				ApplyMsg:        apply.ApplyMsg,
				ApplyAt:         applyAt,
				ApplyVersion:    apply.ApplyCounter,
				HandleAt:        now.UnixMilli(),
			})
		}()

		return &imrelation.HandleApplyResp{
			Code:            imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_OK,
			RelationVersion: apply.ApplyCounter,
			ApplyMsg:        apply.ApplyMsg,
			ApplyAt:         applyAt,
			ApplyVersion:    apply.ApplyCounter,
			ApplyStatus:     apply.Status,
			HandleAt:        now.UnixMilli(),
		}, nil
	default:
		panic("unexpected")
	}
}
