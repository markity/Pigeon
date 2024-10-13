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
	"pigeon/kitex_gen/service/evloopio"
	"pigeon/kitex_gen/service/imrelation"
	relay "pigeon/kitex_gen/service/imrelay"
)

// 接受/拒绝申请
func (s *RPCServer) HandleApply(ctx context.Context, req *imrelation.HandleApplyReq) (res *imrelation.HandleApplyResp, err error) {
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

	relation, err := db.InsertOrSelectForUpdateRelationByUsernameGroupId(txn, &model.RelationModel{
		OwnerId:         req.Session.Username,
		GroupId:         groupIdInt,
		Status:          imrelation.RelationEntry_NOT_IN_GROUP,
		RelationCounter: 0,
		CreatedAt:       now.UnixMilli(),
		UpdatedAt:       now.UnixMilli(),
	})
	if err != nil {
		log.Printf("failed to insert or lock relation entry: %v\n", err)
		return nil, err
	}

	apply, err := db.InsertOrSelectForUpdateApplyByUsernameGroupId(txn, &model.ApplyModel{
		OwnerId:      req.Session.Username,
		GroupId:      groupIdInt,
		ApplyCounter: 0,
		ApplyMsg:     "",
		CreatedAt:    now.UnixMilli(),
		UpdatedAt:    now.UnixMilli(),
		Status:       imrelation.ApplyEntry_NONE,
	})
	if err != nil {
		log.Printf("failed to insert or lock apply entry: %v\n", err)
		return nil, err
	}
	applyAt := apply.UpdatedAt

	if apply.Status != imrelation.ApplyEntry_PENDING {
		go func() {
			push.HandleApplyResp(req.Session, &push.HandleApplyRespInput{
				EchoCode: req.EchoCode,
				Code:     imrelation.HandleApplyResp_NO_APPLY,
			})
		}()
		return &imrelation.HandleApplyResp{
			Code: imrelation.HandleApplyResp_NO_APPLY,
		}, nil
	}

	if !req.Accept {
		apply.Status = imrelation.ApplyEntry_REJECTED
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

				Code: imrelation.HandleApplyResp_OK,

				RelationVersion: relation.RelationCounter,

				ApplyMsg:     apply.ApplyMsg,
				ApplyAt:      applyAt,
				ApplyVersion: apply.ApplyCounter,

				HandleAt: now.UnixMilli(),
			})
		}()

		return &imrelation.HandleApplyResp{
			Code:            imrelation.HandleApplyResp_OK,
			RelationVersion: relation.RelationCounter,

			ApplyMsg:     apply.ApplyMsg,
			ApplyAt:      applyAt,
			ApplyVersion: apply.ApplyCounter,

			HandleAt: now.UnixMilli(),
		}, nil
	}

	// pending状态, 直接打入evloop
	resp, err := s.RelayCli.RedirectToChatEventLoop(context.Background(),
		&relay.RedirectToChatEventLoopReq{
			GroupId: fmt.Sprint(apply.GroupId),
			Input: &evloopio.UniversalGroupEvloopInput{
				Input: &evloopio.UniversalGroupEvloopInput_AlterGroupMember{
					AlterGroupMember: &evloopio.AlterGroupMemberRequest{
						GroupId:         req.GroupId,
						IsAdd:           true,
						MemberId:        req.UserId,
						RelationVersion: res.RelationVersion,
					},
				},
			},
		})
	if err != nil {
		log.Printf("failed to RedirectToChatEventLoop: %v\n", err)
		return nil, err
	}
	out := resp.Output.Output.(*evloopio.UniversalGroupEvloopOutput_AlterGroupMember)
	code := out.AlterGroupMember.Code
	switch code {
	case evloopio.AlterGroupMemberResponse_GROUP_DISBANDED:
		go func() {
			push.HandleApplyResp(req.Session, &push.HandleApplyRespInput{
				EchoCode: req.EchoCode,
				Code:     imrelation.HandleApplyResp_GROUP_DISBANDED,
			})
		}()
		return &imrelation.HandleApplyResp{
			Code: imrelation.HandleApplyResp_GROUP_DISBANDED,
		}, nil
	case evloopio.AlterGroupMemberResponse_OK:
		// accept, 申请成功
		apply.Status = imrelation.ApplyEntry_ACCEPTED
		apply.ApplyCounter++
		apply.UpdatedAt = now.UnixMilli()
		err = db.UpdateApply(txn, apply)
		if err != nil {
			log.Printf("failed to update apply: %v\n", err)
			return nil, err
		}

		relation.Status = imrelation.RelationEntry_NOT_IN_GROUP
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
				Code:            imrelation.HandleApplyResp_OK,
				RelationVersion: relation.RelationCounter,
				ApplyMsg:        apply.ApplyMsg,
				ApplyAt:         applyAt,
				ApplyVersion:    apply.ApplyCounter,
				HandleAt:        now.UnixMilli(),
			})
		}()

		return &imrelation.HandleApplyResp{
			Code:            imrelation.HandleApplyResp_OK,
			RelationVersion: apply.ApplyCounter,
			ApplyMsg:        apply.ApplyMsg,
			ApplyAt:         applyAt,
			ApplyVersion:    apply.ApplyCounter,
			HandleAt:        now.UnixMilli(),
		}, nil
	default:
		panic("unexpected")
	}

}
