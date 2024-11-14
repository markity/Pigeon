package rpcserver

import (
	"context"
	"fmt"
	"log"

	"pigeon/im-relation/bizpush"
	"pigeon/im-relation/db"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/imrelation"
)

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

	relations := make([]*base.RelationEntry, 0, len(data))
	for _, v := range data {
		relations = append(relations, &base.RelationEntry{
			UserId:          v.OwnerId,
			GroupId:         fmt.Sprint(v.GroupId),
			RelationVersion: v.RelationVersion,
			Status:          v.Status,
			ChangeAt:        v.UpdatedAt,
			ChangeType:      v.ChangeType,
		})
	}

	go func() {
		s.BPush.FetchAllRelationsResp(&bizpush.FetchAllRelationsRespInput{
			Session:   req.Session,
			EchoCode:  req.EchoCode,
			Relations: relations,
		})
	}()

	return &imrelation.FetchAllRelationsResp{
		Relations: relations,
	}, nil
}
