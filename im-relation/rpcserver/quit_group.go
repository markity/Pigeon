package rpcserver

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"pigeon/im-relation/db"
	"pigeon/im-relation/push"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/evloopio"
	"pigeon/kitex_gen/service/imrelation"
	relay "pigeon/kitex_gen/service/imrelay"
)

// 主动退群, 锁relation记录, 转移状态
func (s *RPCServer) QuitGroup(ctx context.Context, req *imrelation.QuitGroupReq) (res *imrelation.QuitGroupResp, err error) {
	now := time.Now()

	groupIdInt, err := strconv.ParseInt(req.GroupId, 10, 64)
	if err != nil {
		log.Printf("failed to parse group id: %v\n", err)
		return nil, err
	}
	group, err := db.GetGroupInfo(s.DB.DB(), groupIdInt)
	if err != nil {
		log.Printf("failed to get group info: %v\n", err)
		return nil, err
	}
	if group == nil {
		return &imrelation.QuitGroupResp{
			Code: imrelation.QuitGroupResp_QUIT_GROUP_NOT_MEMBER,
		}, nil
	}

	if group.Disbaned {
		return &imrelation.QuitGroupResp{
			Code: imrelation.QuitGroupResp_QUIT_GROUP_DISBANED,
		}, nil
	}

	// 先查询一次, 看是否有这个relation entry
	relation, err := db.GetRelationByUsernameGroupId(s.DB.DB(), req.Session.Username, groupIdInt)
	if err != nil {
		log.Printf("failed to get relation by username and group id: %v\n", err)
		return nil, err
	}

	if relation == nil || relation.Status != base.RelationStatus_RELATION_STATUS_MEMBER {
		return &imrelation.QuitGroupResp{
			Code: imrelation.QuitGroupResp_QUIT_GROUP_NOT_MEMBER,
		}, nil
	}

	txn := s.DB.Txn()
	defer txn.Rollback()
	relation, err = db.SelectForUpdateRelationByUsernameGroupId(txn, req.Session.Username, groupIdInt)
	if err != nil {
		log.Printf("failed to select for update relation by username and group id: %v\n", err)
		return nil, err
	}

	relation.UpdatedAt = now.UnixMilli()
	relation.Status = base.RelationStatus_RELATION_STATUS_NOT_IN_GROUP
	relation.ChangeType = base.RelationChangeType_RELATION_CHANGE_TYPE_MEMBER_QUIT
	relation.RelationVersion++

	// pending状态, 直接打入evloop
	resp, err := s.RelayCli.RedirectToChatEventLoop(context.Background(),
		&relay.RedirectToChatEventLoopReq{
			GroupId: fmt.Sprint(groupIdInt),
			Input: &evloopio.UniversalGroupEvloopInput{
				Input: &evloopio.UniversalGroupEvloopInput_AlterGroupMember{
					AlterGroupMember: &evloopio.AlterGroupMemberRequest{
						Relation: &base.RelationEntry{
							GroupId:         fmt.Sprint(group.Id),
							UserId:          req.Session.Username,
							Status:          base.RelationStatus_RELATION_STATUS_NOT_IN_GROUP,
							ChangeType:      base.RelationChangeType_RELATION_CHANGE_TYPE_MEMBER_QUIT,
							RelationVersion: relation.RelationVersion,
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
	case evloopio.AlterGroupMemberResponse_OK:
		err := db.UpdateRelation(txn, relation)
		if err != nil {
			log.Printf("failed to update relation: %v\n", err)
			return nil, err
		}
		err = txn.Commit().Error
		if err != nil {
			log.Printf("failed to commit: %v\n", err)
			return nil, err
		}

		// push relation change
		go func() {
			push.RelationChangeNotify(&push.RelationChangeNotifyInput{
				AuthRoute:  s.AuthRouteCli,
				Username:   req.Session.Username,
				GroupId:    fmt.Sprint(groupIdInt),
				Version:    relation.RelationVersion,
				Status:     relation.Status,
				ChangeAt:   now.UnixMilli(),
				ChangeType: relation.ChangeType,
			})
		}()

		// push echo
		go func() {
			push.QuitGroupResp(req.Session, &push.QuitGroupRespInput{
				EchoCode:        req.EchoCode,
				Code:            imrelation.QuitGroupResp_QUIT_GROUP_OK,
				RelationVersion: relation.RelationVersion,
			})
		}()

		return &imrelation.QuitGroupResp{
			Code:            imrelation.QuitGroupResp_QUIT_GROUP_OK,
			RelationVersion: relation.RelationVersion,
		}, nil
	case evloopio.AlterGroupMemberResponse_GROUP_DISBANDED:
		go func() {
			push.QuitGroupResp(req.Session, &push.QuitGroupRespInput{
				EchoCode:        req.EchoCode,
				Code:            imrelation.QuitGroupResp_QUIT_GROUP_DISBANED,
				RelationVersion: relation.RelationVersion,
			})
		}()
		// push echo
		return &imrelation.QuitGroupResp{
			Code: imrelation.QuitGroupResp_QUIT_GROUP_DISBANED,
		}, nil
	default:
		panic("unexpected")
	}
}
