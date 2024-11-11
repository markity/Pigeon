package rpcserver

import (
	"context"
	"fmt"
	"log"
	"time"

	"pigeon/im-relation/bizpush"
	"pigeon/im-relation/db"
	"pigeon/im-relation/db/model"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/evloopio"
	"pigeon/kitex_gen/service/imrelation"
	relay "pigeon/kitex_gen/service/imrelay"
)

// 接受/拒绝申请
func (s *RPCServer) HandleApply(ctx context.Context, req *imrelation.HandleApplyReq) (
	res *imrelation.HandleApplyResp, err error) {
	now := time.Now()

	groupInfo, err := db.GetGroupInfo(s.DB.DB(), req.GroupId)
	if err != nil {
		log.Printf("failed to get group info: %v\n", err)
		return nil, err
	}
	if groupInfo == nil || groupInfo.OwnerId != req.Session.Username {
		go func() {
			s.BPush.HandleApplyResp(&bizpush.HandleApplyRespInput{
				Session:  req.Session,
				EchoCode: req.EchoCode,
				Code:     imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_NO_PERMISSION,
			})
		}()
		return &imrelation.HandleApplyResp{
			Code: imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_NO_PERMISSION,
		}, nil
	}

	// 锁group
	le, err := s.RdsAct.LockGroup(req.GroupId, time.Second*1)
	if err != nil {
		return nil, err
	}
	defer le.UnLock()

	txn := s.DB.Txn()
	defer txn.Rollback()

	relation, err := db.GetRelationByUsernameAndGroupId(txn, req.UserId, req.GroupId)
	if err != nil {
		log.Printf("failed to get relation entry: %v\n", err)
		return nil, err
	}

	apply, err := db.GetApplyByUsernameAndGroupId(txn, req.UserId, req.GroupId)
	if err != nil {
		log.Printf("failed to get apply entry: %v\n", err)
		return nil, err
	}

	if apply == nil || apply.Status != base.ApplyStatus_APPLY_STATUS_PENDING {
		go func() {
			s.BPush.HandleApplyResp(&bizpush.HandleApplyRespInput{
				Session:  req.Session,
				EchoCode: req.EchoCode,
				Code:     imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_NO_APPLY,
			})
		}()
		return &imrelation.HandleApplyResp{
			Code: imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_NO_APPLY,
		}, nil
	}

	applyAt := apply.UpdatedAt

	// TODO: 先简化逻辑, 目前没有disband状态
	// // 如果已经disbaned, 无论接受还是拒绝, 都返回ok, 但是不更新关系
	// if groupInfo.Disbaned {
	// 	apply.Status = base.ApplyStatus_APPLY_STATUS_GROUP_DISBANDED
	// 	apply.ApplyVersion++
	// 	apply.UpdatedAt = now.UnixMilli()

	// 	err := db.UpdateApply(txn, apply)
	// 	if err != nil {
	// 		log.Printf("failed to update apply: %v\n", err)
	// 		return nil, err
	// 	}

	// 	err = txn.Commit().Error
	// 	if err != nil {
	// 		log.Printf("failed to commit apply: %v\n", err)
	// 		return nil, err
	// 	}

	// 	go func() {
	// 		push.HandleApplyResp(req.Session, &push.HandleApplyRespInput{
	// 			EchoCode:     req.EchoCode,
	// 			Code:         imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_OK,
	// 			ApplyMsg:     apply.ApplyMsg,
	// 			ApplyAt:      applyAt,
	// 			ApplyVersion: apply.ApplyVersion,
	// 			ApplyStatus:  apply.Status,
	// 			HandleAt:     now.UnixMilli(),

	// 			RelationVersion: 0,
	// 		})
	// 	}()

	// 	return &imrelation.HandleApplyResp{
	// 		Code:            imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_OK,
	// 		RelationVersion: 0,
	// 		ApplyMsg:        apply.ApplyMsg,
	// 		ApplyAt:         applyAt,
	// 		HandleAt:        now.UnixMilli(),
	// 		ApplyVersion:    apply.ApplyVersion,
	// 		ApplyStatus:     apply.Status,
	// 	}, nil
	// }

	// // 群没有disbanded

	if !req.Accept {
		apply.Status = base.ApplyStatus_APPLY_STATUS_REJECTED
		apply.ApplyVersion++
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
			s.BPush.HandleApplyNotify(&bizpush.HandleApplyNotifyInput{
				OwnerId:      groupInfo.OwnerId,
				Username:     apply.OwnerId,
				GroupId:      req.GroupId,
				ApplyVersion: apply.ApplyVersion,
				ApplyMsg:     apply.ApplyMsg,
				ApplyStatus:  apply.Status,
				ApplyAt:      applyAt,
				HandleAt:     apply.UpdatedAt,
			})
		}()

		go func() {
			s.BPush.HandleApplyResp(&bizpush.HandleApplyRespInput{
				Session: req.Session,

				EchoCode: req.EchoCode,

				Code: imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_OK,

				RelationVersion: relation.RelationVersion,

				ApplyMsg:     apply.ApplyMsg,
				ApplyAt:      applyAt,
				ApplyVersion: apply.ApplyVersion,

				HandleAt: now.UnixMilli(),
			})
		}()

		return &imrelation.HandleApplyResp{
			Code:            imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_OK,
			RelationVersion: relation.RelationVersion,

			ApplyMsg:     apply.ApplyMsg,
			ApplyAt:      applyAt,
			ApplyVersion: apply.ApplyVersion,

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
						Relation: &base.RelationEntry{
							GroupId:         groupInfo.GroupId,
							UserId:          req.UserId,
							Status:          base.RelationStatus_RELATION_STATUS_MEMBER,
							ChangeType:      base.RelationChangeType_RELATION_CHANGE_TYPE_OWNER_ACCEPT,
							RelationVersion: apply.ApplyVersion,
							ChangeAt:        0,
						},
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
	// // TODO: 这里先不考虑disband流程, 先简化
	// case evloopio.AlterGroupMemberResponse_GROUP_DISBANDED:
	// 	apply.Status = base.ApplyStatus_APPLY_STATUS_GROUP_DISBANDED
	// 	apply.ApplyVersion++
	// 	apply.UpdatedAt = now.UnixMilli()
	// 	err := db.UpdateApply(txn, apply)
	// 	if err != nil {
	// 		log.Printf("failed to update apply: %v\n", err)
	// 		return nil, err
	// 	}

	// 	err = txn.Commit().Error
	// 	if err != nil {
	// 		log.Printf("failed to commit apply: %v\n", err)
	// 		return nil, err
	// 	}

	// 	go func() {
	// 		push.HandleApplyResp(req.Session, &push.HandleApplyRespInput{
	// 			EchoCode:     req.EchoCode,
	// 			Code:         imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_OK,
	// 			ApplyMsg:     apply.ApplyMsg,
	// 			ApplyAt:      applyAt,
	// 			ApplyVersion: apply.ApplyVersion,
	// 			ApplyStatus:  apply.Status,
	// 			HandleAt:     now.UnixMilli(),

	// 			RelationVersion: 0,
	// 		})
	// 	}()
	// 	return &imrelation.HandleApplyResp{
	// 		Code:         imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_OK,
	// 		ApplyMsg:     apply.ApplyMsg,
	// 		ApplyAt:      applyAt,
	// 		ApplyVersion: apply.ApplyVersion,
	// 		ApplyStatus:  apply.Status,
	// 		HandleAt:     now.UnixMilli(),

	// 		RelationVersion: 0,
	// 	}, nil
	case evloopio.AlterGroupMemberResponse_OK:
		apply.Status = base.ApplyStatus_APPLY_STATUS_ACCEPTED
		apply.ApplyVersion++
		apply.UpdatedAt = now.UnixMilli()
		err := db.UpdateApply(txn, apply)
		if err != nil {
			log.Printf("failed to update apply: %v\n", err)
			return nil, err
		}

		if relation == nil {
			relation = &model.RelationModel{
				OwnerId:         apply.OwnerId,
				GroupId:         apply.GroupId,
				Status:          base.RelationStatus_RELATION_STATUS_UNUSED,
				ChangeType:      base.RelationChangeType_RELATION_CHANGE_TYPE_UNUSED,
				RelationVersion: 0,
				CreatedAt:       now.UnixMilli(),
				UpdatedAt:       now.UnixMilli(),
			}
			err := db.InsertRelation(txn, relation)
			if err != nil {
				log.Printf("failed to insert relation: %v\n", err)
				return nil, err
			}
		}

		relation.Status = base.RelationStatus_RELATION_STATUS_MEMBER
		relation.ChangeType = base.RelationChangeType_RELATION_CHANGE_TYPE_OWNER_ACCEPT
		relation.RelationVersion++
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
			s.BPush.RelationChangeNotify(&bizpush.RelationChangeNotifyInput{
				Username:   relation.OwnerId,
				GroupId:    fmt.Sprint(relation.GroupId),
				Version:    relation.RelationVersion,
				Status:     base.RelationStatus_RELATION_STATUS_MEMBER,
				ChangeType: relation.ChangeType, // relation变更场景, 是因为群主接受申请
				UpdatedAt:  now.UnixMilli(),
			})
		}()

		// 给group owner发处理apply的通知
		go func() {
			s.BPush.HandleApplyNotify(&bizpush.HandleApplyNotifyInput{
				OwnerId:      groupInfo.OwnerId,
				Username:     apply.OwnerId,
				GroupId:      req.GroupId,
				ApplyVersion: apply.ApplyVersion,
				ApplyMsg:     apply.ApplyMsg,
				ApplyStatus:  apply.Status,
				ApplyAt:      applyAt,
				HandleAt:     now.UnixMilli(),
			})
		}()

		// 推消息, 告诉group owner以及新relation的建立
		go func() {
			s.BPush.HandleApplyResp(&bizpush.HandleApplyRespInput{
				Session:         req.Session,
				Code:            imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_OK,
				RelationVersion: relation.RelationVersion,
				ApplyMsg:        apply.ApplyMsg,
				ApplyAt:         applyAt,
				ApplyVersion:    apply.ApplyVersion,
				HandleAt:        now.UnixMilli(),
			})
		}()

		return &imrelation.HandleApplyResp{
			Code:            imrelation.HandleApplyResp_HANDLE_APPLY_RESP_CODE_OK,
			RelationVersion: apply.ApplyVersion,
			ApplyMsg:        apply.ApplyMsg,
			ApplyAt:         applyAt,
			ApplyVersion:    apply.ApplyVersion,
			ApplyStatus:     apply.Status,
			HandleAt:        now.UnixMilli(),
		}, nil
	default:
		panic("unexpected")
	}
}
