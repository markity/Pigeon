package rpcserver

import (
	"context"
	"fmt"
	"log"

	"pigeon/im-relation/db"
	"pigeon/im-relation/push"
	"pigeon/kitex_gen/service/imrelation"
)

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
			ApplyVersion: v.ApplyVersion,
			ApplyAt:      v.UpdatedAt,
			ApplyMsg:     v.ApplyMsg,
			Status:       v.Status,
		})
	}

	go func() {
		push.FetchAllAppliesResp(req.Session, &push.FetchAllAppliesRespInput{
			EchoCode: req.EchoCode,
			Applies:  applications,
		})
	}()

	return &imrelation.FetchAllApplicationsResp{
		Applications: applications,
	}, nil
}
