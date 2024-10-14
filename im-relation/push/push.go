package push

import (
	"context"
	"encoding/json"
	"log"
	"pigeon/im-relation/api"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/imauthroute"
	authroute "pigeon/kitex_gen/service/imauthroute/imauthroute"

	"pigeon/kitex_gen/service/imgateway"
	"pigeon/kitex_gen/service/imrelation"
	"time"
)

type CreateGroupRespInput struct {
	EchoCode  string
	OwnerId   string
	GroupId   string
	CreatedAt int64
}

func CreateGroupResp(session *base.SessionEntry, input *CreateGroupRespInput) {
	cli := api.NewGatewayClientFromAdAddr(session.GwAdvertiseAddrPort)
	for {
		_, err := cli.PushMessage(context.Background(), &imgateway.PushMessageReq{
			SessionId: session.SessionId,
			PushType:  "push-create-group-resp",
			EchoCode:  input.EchoCode,
			Data: mustMarshal(map[string]interface{}{
				"owner_id":   input.OwnerId,
				"group_id":   input.GroupId,
				"created_at": input.CreatedAt,
			}),
		})
		if err != nil {
			log.Printf("push CreateGroupResp: %v\n", err)
			time.Sleep(time.Millisecond * 50)
			continue
		}
		break
	}
}

type FetchAllRelationsRespInput struct {
	EchoCode  string
	Relations []*imrelation.RelationEntry
}

type relationEntry struct {
	Username        string                    `json:"username"`
	GroupId         string                    `json:"group_id"`
	RelationVersion int64                     `json:"version"`
	RelationStatus  imrelation.RelationStatus `json:"relation_status"`
	UpdatedAt       int64                     `json:"updated_at"`
}

func FetchAllRelationsResp(session *base.SessionEntry, input *FetchAllRelationsRespInput) {
	cli := api.NewGatewayClientFromAdAddr(session.GwAdvertiseAddrPort)
	myRelations := make([]*relationEntry, 0, len(input.Relations))
	for _, v := range input.Relations {
		myRelations = append(myRelations, &relationEntry{
			Username:        v.UserId,
			GroupId:         v.GroupId,
			RelationVersion: v.RelationVersion,
			RelationStatus:  v.Status,
			UpdatedAt:       v.UpdatedAt,
		})
	}
	for {
		_, err := cli.PushMessage(context.Background(), &imgateway.PushMessageReq{
			SessionId: session.SessionId,
			PushType:  "push-fetch-relations-resp",
			EchoCode:  input.EchoCode,
			Data: mustMarshal(map[string]interface{}{
				"relations": myRelations,
			}),
		})
		if err != nil {
			log.Printf("push FetchAllRelationsResp: %v\n", err)
			time.Sleep(time.Millisecond * 50)
			continue
		}
		break
	}
}

type FetchAllAppliesRespInput struct {
	EchoCode string
	Applies  []*imrelation.ApplyEntry
}

type applyEntry struct {
	Username     string `json:"username"`
	GroupId      string `json:"group_id"`
	ApplyVersion int64  `json:"version"`
	ApplyStatus  imrelation.ApplyStatus
	UpdatedAt    int64 `json:"updated_at"`
}

func FetchAllAppliesResp(session *base.SessionEntry, input *FetchAllAppliesRespInput) {
	cli := api.NewGatewayClientFromAdAddr(session.GwAdvertiseAddrPort)
	myRelations := make([]*applyEntry, 0, len(input.Applies))
	for _, v := range input.Applies {
		myRelations = append(myRelations, &applyEntry{
			Username:     v.UserId,
			GroupId:      v.GroupId,
			ApplyVersion: v.ApplyVersion,
			ApplyStatus:  v.Status,
			UpdatedAt:    v.ApplyAt,
		})
	}
	for {
		_, err := cli.PushMessage(context.Background(), &imgateway.PushMessageReq{
			SessionId: session.SessionId,
			PushType:  "push-fetch-applies-resp",
			EchoCode:  input.EchoCode,
			Data: mustMarshal(map[string]interface{}{
				"applies": myRelations,
			}),
		})
		if err != nil {
			log.Printf("push FetchAllAppliesResp: %v\n", err)
			time.Sleep(time.Millisecond * 50)
			continue
		}
		break
	}
}

type ApplyGroupRespInput struct {
	EchoCode string
	Code     imrelation.ApplyGroupResp_ApplyGroupRespCode
}

func ApplyGroupResp(session *base.SessionEntry, input *ApplyGroupRespInput) {
	cli := api.NewGatewayClientFromAdAddr(session.GwAdvertiseAddrPort)
	for {
		_, err := cli.PushMessage(context.Background(), &imgateway.PushMessageReq{
			SessionId: session.SessionId,
			PushType:  "push-apply-group-resp",
			EchoCode:  input.EchoCode,
			Data: mustMarshal(map[string]interface{}{
				"code": input.Code,
			}),
		})
		if err != nil {
			log.Printf("push ApplyGroupResp: %v\n", err)
			time.Sleep(time.Millisecond * 50)
			continue
		}
		break
	}
}

type ApplyGroupNotifyInput struct {
	AuthRoute authroute.Client

	OwnerId      string
	Username     string
	GroupId      string
	ApplyMsg     string
	ApplyVersion int64
	ApplyAt      int64
}

func ApplyGroupNotify(input *ApplyGroupNotifyInput) {
	var queryResp *imauthroute.QueryUserRouteResp
	for {
		var err error
		queryResp, err = input.AuthRoute.QueryUserRoute(context.Background(),
			&imauthroute.QueryUserRouteReq{
				Username: input.OwnerId,
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
			PushType:  "push-apply-notify",
			EchoCode:  "",
			Data: mustMarshal(map[string]interface{}{
				"user_id":       input.Username,
				"group_id":      input.GroupId,
				"apply_version": input.ApplyVersion,
				"apply_msg":     input.ApplyMsg,
				"apply_at":      input.ApplyAt,
			}),
		})
		if err != nil {
			log.Printf("failed to push message: %v, retrying\n", err)
			time.Sleep(time.Millisecond * 50)
			goto retry
		}
	}
}

type HandleApplyNotifyInput struct {
	AuthRoute authroute.Client

	OwnerId      string
	Username     string
	GroupId      string
	ApplyVersion int64
	ApplyMsg     string
	ApplyStatus  imrelation.ApplyStatus
	ApplyAt      int64
	HandleAt     int64
}

func HandleApplyNotify(input *HandleApplyNotifyInput) {
	var queryResp *imauthroute.QueryUserRouteResp
	for {
		var err error
		queryResp, err = input.AuthRoute.QueryUserRoute(context.Background(),
			&imauthroute.QueryUserRouteReq{
				Username: input.OwnerId,
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
			PushType:  "biz-handle-apply-notify",
			EchoCode:  "",
			Data: mustMarshal(map[string]interface{}{
				"user_id":       input.Username,
				"group_id":      input.GroupId,
				"apply_version": input.ApplyVersion,
				"apply_msg":     input.ApplyMsg,
				"apply_status":  input.ApplyStatus,
				"apply_at":      input.ApplyAt,
				"handle_at":     input.HandleAt,
			}),
		})
		if err != nil {
			log.Printf("failed to push message: %v, retrying\n", err)
			time.Sleep(time.Millisecond * 50)
			goto retry
		}
	}
}

type HandleApplyRespInput struct {
	EchoCode     string
	Code         imrelation.HandleApplyResp_HandleApplyRespCode
	ApplyMsg     string
	ApplyAt      int64
	ApplyVersion int64
	ApplyStatus  imrelation.ApplyStatus
	HandleAt     int64

	// 如果是接受请求, 且ApplyStatus变为已接受, relation_version不为0
	RelationVersion int64
}

func HandleApplyResp(session *base.SessionEntry, input *HandleApplyRespInput) {
	cli := api.NewGatewayClientFromAdAddr(session.GwAdvertiseAddrPort)
	for {
		_, err := cli.PushMessage(context.Background(), &imgateway.PushMessageReq{
			SessionId: session.SessionId,
			PushType:  "push-handle-apply-resp",
			EchoCode:  input.EchoCode,
			Data: mustMarshal(map[string]interface{}{
				"code":             input.Code,
				"relation_version": input.RelationVersion,
				"apply_msg":        input.ApplyMsg,
				"apply_at":         input.ApplyAt,
				"apply_version":    input.ApplyVersion,
				"handle_at":        input.HandleAt,
			}),
		})
		if err != nil {
			log.Printf("push ApplyGroupResp: %v\n", err)
			time.Sleep(time.Millisecond * 50)
			continue
		}
		break
	}
}

// 比如a成功加入了群, 那么a加入群的消息将下发给a的全部设备
type RelationChangeNotifyInput struct {
	AuthRoute authroute.Client

	Username   string
	GroupId    string
	Version    int64
	Status     imrelation.RelationStatus
	ChangeAt   int64
	ChangeType imrelation.RelationChangeType
}

func RelationChangeNotify(input *RelationChangeNotifyInput) {
	var queryResp *imauthroute.QueryUserRouteResp
	for {
		var err error
		queryResp, err = input.AuthRoute.QueryUserRoute(context.Background(),
			&imauthroute.QueryUserRouteReq{
				Username: input.Username,
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
			PushType:  "biz-relation-change-notify",
			EchoCode:  "",
			Data: mustMarshal(map[string]interface{}{
				"user_id":          input.Username,
				"group_id":         input.GroupId,
				"relation_version": input.Version,
				"status":           input.Status,
				"change_type":      input.ChangeType,
				"change_at":        input.ChangeAt,
			}),
		})
		if err != nil {
			log.Printf("failed to push message: %v, retrying\n", err)
			time.Sleep(time.Millisecond * 50)
			goto retry
		}
	}
}

type QuitGroupRespInput struct {
	EchoCode        string
	Code            imrelation.QuitGroupResp_QuitGroupRespCode
	RelationVersion int64
}

func QuitGroupResp(session *base.SessionEntry, input *QuitGroupRespInput) {
	cli := api.NewGatewayClientFromAdAddr(session.GwAdvertiseAddrPort)
	for {
		_, err := cli.PushMessage(context.Background(), &imgateway.PushMessageReq{
			SessionId: session.SessionId,
			EchoCode:  input.EchoCode,
			PushType:  "push-quit-group-resp",
			Data: mustMarshal(map[string]interface{}{
				"code":             input.Code,
				"relation_version": input.RelationVersion,
			}),
		})
		if err != nil {
			log.Printf("push QuitGroupResp: %v\n", err)
			time.Sleep(time.Millisecond * 50)
			continue
		}
		break
	}
}

func mustMarshal(obj interface{}) []byte {
	bs, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	return bs
}
