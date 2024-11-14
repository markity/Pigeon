package rpcserver

import (
	"context"
	"log"

	"pigeon/im-relation/bizpush"
	"pigeon/im-relation/db"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/imrelation"
)

/*
用户登录后, 会查询自己所有申请, 这个接口由im-relay调用
比如此用户是 a b c三个群的群主, 则申请a b c三群的所有请求都会被拉到
*/
func (s *RPCServer) FetchAllApplications(ctx context.Context, req *imrelation.FetchAllApplicationsReq) (res *imrelation.FetchAllApplicationsResp, err error) {
	txn := s.DB.Txn()
	defer txn.Rollback()

	data, err := db.GetAllApplicationsByUsername(txn, req.Session.Username)
	if err != nil {
		log.Printf("failed to fetch all applications: %v\n", err)
		return nil, err
	}

	applications := make([]*base.ApplyEntry, 0, len(data))
	for _, v := range data {
		applications = append(applications, &base.ApplyEntry{
			UserId:       v.OwnerId,
			GroupId:      v.GroupId,
			ApplyVersion: v.ApplyVersion,
			ApplyAt:      v.UpdatedAt,
			ApplyMsg:     v.ApplyMsg,
			Status:       v.Status,
		})
	}

	go func() {
		s.BPush.FetchAllAppliesResp(&bizpush.FetchAllAppliesRespInput{
			Session:      req.Session,
			Applications: applications,
		})
	}()

	return &imrelation.FetchAllApplicationsResp{
		Applications: applications,
	}, nil
}
