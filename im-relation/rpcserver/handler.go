package rpcserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"pigeon/im-auth-route/api"
	"pigeon/im-auth-route/rds"
	"pigeon/im-relation/db"
	"pigeon/im-relation/db/model"
	"pigeon/kitex_gen/service/evloopio"
	authroute "pigeon/kitex_gen/service/imauthroute"
	"pigeon/kitex_gen/service/imauthroute/imauthroute"
	"pigeon/kitex_gen/service/imgateway"
	"pigeon/kitex_gen/service/imrelation"
	relay "pigeon/kitex_gen/service/imrelay"
	"pigeon/kitex_gen/service/imrelay/imrelay"
)

type RPCContext struct {
	DB           *db.DB
	Rds          *rds.RdsAction
	RelayCli     imrelay.Client
	AuthRouteCli imauthroute.Client
}

type RPCServer struct {
	RPCContext
}

/*
create group逻辑概要:
begin txn
insert拿到自增主键: groupId TODO: 这里以后优化, 增加id生成器来生成groupId, 目前是直接用mysql自增主键
调用imrelay.CreateChatEventLoop, 创建chat event loop
commit txn
return info
*/
func (s *RPCServer) CreateGroup(ctx context.Context, req *imrelation.CreateGroupReq) (res *imrelation.CreateGroupResp, err error) {
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
		GroupId:  fmt.Sprint(group.Id),
		CreateAt: group.CreatedAt,
	}, nil
}

/*
imrelay在一致性转发消息前得先判断group是否存在, 如果
存在它才转发, 这是为了配合迁移的场景
*/
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

/*
用户登录后, 会查询自己所有关系, 这个接口由im-relay调用
*/
func (s *RPCServer) FetchAllRelations(ctx context.Context, req *imrelation.FetchAllRelationsReq) (res *imrelation.FetchAllRelationsResp, err error) {
	txn := s.DB.Txn()
	defer txn.Rollback()

	data, err := db.GetAllRelationsByUsername(txn, req.Session.Username)
	if err != nil {
		log.Printf("failed to fetch all relations: %v\n", err)
		return nil, err
	}

	relations := make([]*imrelation.RelationEntry, 0, len(data))
	for _, v := range data {
		relations = append(relations, &imrelation.RelationEntry{
			UserId:     v.OwnerId,
			GroupId:    fmt.Sprint(v.GroupId),
			RelationId: v.RelationCounter,
			InGroup:    v.Status == model.RelationStatusInGroup,
			UpdatedAt:  v.UpdatedAt,
		})
	}

	return &imrelation.FetchAllRelationsResp{
		Relations: relations,
	}, nil
}

/*
用户登录后, 会查询自己所有申请, 这个接口由im-relay调用
*/
func (s *RPCServer) FetchAllApplications(ctx context.Context, req *imrelation.FetchAllApplicationsReq) (res *imrelation.FetchAllApplicationsResp, err error) {
	txn := s.DB.Txn()
	defer txn.Rollback()

	data, err := db.GetAllApplicationsByUsername(txn, req.Session.Username)
	if err != nil {
		log.Printf("failed to fetch all applications: %v\n", err)
		return nil, err
	}

	applications := make([]*imrelation.ApplyEntry, 0, len(data))
	for _, v := range data {
		applications = append(applications, &imrelation.ApplyEntry{
			UserId:       v.OwnerId,
			GroupId:      fmt.Sprint(v.GroupId),
			ApplyVersion: v.ApplyCounter,
			ApplyAt:      v.UpdatedAt,
			ApplyMsg:     v.ApplyMsg,
		})
	}

	return &imrelation.FetchAllApplicationsResp{
		Applications: applications,
	}, nil
}

// 申请加入群
func (s *RPCServer) ApplyGroup(ctx context.Context, req *imrelation.ApplyGroupReq) (res *imrelation.ApplyGroupResp, err error) {
	now := time.Now()

	// 一个用户关于一个群只有一条ApplyModel, 这里有个并发问题, 这里使用mysql 1062错误码解决这个问题
	txn := s.DB.Txn()
	defer txn.Rollback()

	groupIdInt, err := strconv.ParseInt(req.GroupId, 10, 64)
	if err != nil {
		return &imrelation.ApplyGroupResp{
			Code: imrelation.ApplyGroupResp_NO_GROUP,
		}, nil
	}

	// 查询group信息
	group, err := db.GetGroupInfo(txn, groupIdInt)
	if err != nil {
		log.Printf("failed to get group info: %v\n", err)
		return nil, err
	}

	if group == nil {
		return &imrelation.ApplyGroupResp{
			Code: imrelation.ApplyGroupResp_NO_GROUP,
		}, nil
	}

	// 如果用户是群主, 不能申请, 返回已在群中的错误码
	if group.OwnerId == req.Session.Username {
		return &imrelation.ApplyGroupResp{
			Code: imrelation.ApplyGroupResp_USER_IN_GROUP,
		}, nil
	}

	// 如果群已解散, 不能申请
	if group.Disbaned {
		return &imrelation.ApplyGroupResp{
			Code: imrelation.ApplyGroupResp_GROUP_DISBANDED,
		}, nil
	}

	// 先尝试插入
	m := &model.ApplyModel{
		OwnerId:      req.Session.Username,
		GroupId:      groupIdInt,
		ApplyCounter: 0,
		ApplyMsg:     "",
		CreatedAt:    now.UnixMilli(),
		UpdatedAt:    now.UnixMilli(),
		Status:       model.ApplyStatusNone,
	}
	ok, err := db.InsertApply(txn, m)
	if err != nil {
		log.Printf("failed to insert apply: %v\n", err)
		return nil, err
	}
	if !ok {
		txn.Rollback()
	} else {
		err := txn.Commit().Error
		if err != nil {
			log.Printf("failed to commit apply: %v\n", err)
		}
		return nil, err
	}

	// txn2, 锁住apply记录
	txn2 := s.DB.Txn()
	defer txn2.Rollback()
	apply, err := db.SelectForUpdateApplyByUsername(txn2, req.Session.Username)
	if err != nil {
		log.Printf("failed to select for update apply: %v\n", err)
		return nil, err
	}

	if apply.Status == model.ApplyStatusPass {
		return &imrelation.ApplyGroupResp{
			Code: imrelation.ApplyGroupResp_USER_IN_GROUP,
		}, nil
	}
	// none, pending reject三种状态, 更新位pendding状态
	m.ApplyCounter++
	m.ApplyMsg = req.ApplyMsg
	m.Status = model.ApplyStatusPendding
	m.UpdatedAt = now.UnixMilli()
	err = db.UpdateApply(txn2, m)
	if err != nil {
		log.Printf("failed to update apply: %v\n", err)
		return nil, err
	}
	err = txn2.Commit().Error
	if err != nil {
		log.Printf("failed to commit apply: %v\n", err)
		return nil, err
	}

	go func() {
		var queryResp *authroute.QueryUserRouteResp
		for {
			var err error
			queryResp, err = s.AuthRouteCli.QueryUserRoute(context.Background(), &authroute.QueryUserRouteReq{
				Username: m.OwnerId,
			})
			if err != nil {
				log.Printf("failed to query user route: %v, retrying\n", err)
				time.Sleep(time.Millisecond * 50)
				continue
			}
			break
		}

		for _, v := range queryResp.Routes {
			gwCli := api.NewGatewayClientFromAdAddr(v.GwAdvertiseAddrPort)
			// 此处retry是因为网络原因, 比如网断了, 重试是安全的, 保证客户端至少收到消息一次
			// 客户端也得做自己的幂等, 防止消息重放导致副作用
		retry:
			_, err := gwCli.PushMessage(context.Background(), &imgateway.PushMessageReq{
				SessionId: v.SessionId,
				PushType:  "biz-apply-notify",
				EchoCode:  req.EchoCode,
				Data: mustMarshalToBytes(map[string]interface{}{
					"user_id":       req.Session.Username,
					"group_id":      req.GroupId,
					"apply_version": m.ApplyCounter,
					"apply_msg":     m.ApplyMsg,
					"apply_at":      m.UpdatedAt,
				}),
			})
			if err != nil {
				log.Printf("failed to push message: %v, retrying\n", err)
				time.Sleep(time.Millisecond * 50)
				goto retry
			}
		}
	}()

	return &imrelation.ApplyGroupResp{
		Code:         imrelation.ApplyGroupResp_OK,
		ApplyVersion: m.ApplyCounter,
		ApplyAt:      m.UpdatedAt,
	}, nil
}

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

	defer txn.Rollback()
	// 先锁住apply记录, 防止并发
	apply, err := db.SelectForUpdateApplyByUsername(txn, req.Session.Username)
	if err != nil {
		log.Printf("failed to select for update apply: %v\n", err)
		return nil, err
	}

	if apply == nil || apply.Status != model.ApplyStatusPendding {
		return &imrelation.HandleApplyResp{
			Code: imrelation.HandleApplyResp_NO_APPLY,
		}, nil
	}

	// 拒绝申请, 更新apply状态, 并且把这个消息下行给group owner以及
	if !req.Accept {
	}

	// pending状态, 直接打入evloop
	resp, err := s.RelayCli.RedirectToChatEventLoop(context.Background(),
		&relay.RedirectToChatEventLoopReq{
			GroupId: fmt.Sprint(apply.GroupId),
			Input: &evloopio.UniversalGroupEvloopInput{
				Input: &evloopio.UniversalGroupEvloopInput_AlterGroupMember{},
			},
		})

	return
}

func (s *RPCServer) QuitGroup(ctx context.Context, req *imrelation.QuitGroupReq) (res *imrelation.QuitGroupResp, err error) {
	return
}

func mustMarshalToBytes(v interface{}) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return data
}
