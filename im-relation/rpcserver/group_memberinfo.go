package rpcserver

import (
	"context"
	"fmt"
	"pigeon/im-relation/db"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/imrelation"
)

func (s *RPCServer) GetGroupMemberInfo(ctx context.Context, req *imrelation.GetGroupMemberInfoReq) (
	*imrelation.GetGroupMemberInfoResp, error) {
	r, err := db.GetRelationByUsernameAndGroupId(s.DB.DB(), req.Username, req.GroupId)
	if err != nil {
		return nil, err
	}
	fmt.Println(r)
	if r == nil {
		return &imrelation.GetGroupMemberInfoResp{
			Stauts:          base.RelationStatus_RELATION_STATUS_NOT_IN_GROUP,
			RelationVersion: 0,
		}, nil
	}
	return &imrelation.GetGroupMemberInfoResp{
		Stauts:          r.Status,
		RelationVersion: r.RelationVersion,
		ChangeType:      r.ChangeType,
		UpdatedAt:       r.UpdatedAt,
	}, nil
}
