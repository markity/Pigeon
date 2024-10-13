package rpcserver

import (
	"context"
	"log"
	"strconv"

	"pigeon/im-relation/db"
	"pigeon/kitex_gen/service/imrelation"
)

/*
imrelay在一致性转发消息前得先判断group是否存在, 如果
存在它才转发, 这是为了配合迁移的场景, 不会将消息直接send给一个不存在的chateventloop
*/
func (s *RPCServer) GetGroupInfo(ctx context.Context, req *imrelation.GetGroupInfoReq) (
	res *imrelation.GetGroupInfoResp, err error) {
	groupId, err := strconv.ParseInt(req.GroupId, 10, 64)
	if err != nil {
		log.Printf("failed to parse group id: %v\n", err)
		return nil, err
	}

	txn := s.DB.Txn()
	defer txn.Rollback()

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
