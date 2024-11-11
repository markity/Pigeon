package push

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"pigeon/kitex_gen/service/base"
	authroute "pigeon/kitex_gen/service/imauthroute"
	"pigeon/kitex_gen/service/imauthroute/imauthroute"
	gateway "pigeon/kitex_gen/service/imgateway"
	"pigeon/kitex_gen/service/imgateway/imgateway"

	"github.com/cloudwego/kitex/client"
)

func newGatewayClientFromAdAddr(adAddr string) imgateway.Client {
	return imgateway.MustNewClient("im-gateway", client.WithHostPorts(adAddr))
}

type PushManager struct {
	retryWaitTime *time.Duration
	routeCli      imauthroute.Client
}

// 0 means no retry
func NewPushManager(retryWaitTime time.Duration, routeCli imauthroute.Client) *PushManager {
	t := &retryWaitTime
	if retryWaitTime == 0 {
		t = nil
	}
	return &PushManager{
		retryWaitTime: t,
		routeCli:      routeCli,
	}
}

func (pm *PushManager) PushToSessionByMap(session *base.SessionEntry, pushType string, echoCode string, data map[string]interface{}) {
	cli := newGatewayClientFromAdAddr(session.GwAdvertiseAddrport)
	for {
		_, err := cli.PushMessage(context.Background(), &gateway.PushMessageReq{
			SessionId: session.SessionId,
			PushType:  pushType,
			EchoCode:  echoCode,
			Data:      mustMarshal(data),
		})
		if err != nil {
			log.Printf("push: %v\n", err)
			if pm.retryWaitTime != nil {
				time.Sleep(*pm.retryWaitTime)
				continue
			}
			return
		}
		break
	}
}

// TODO: 改造推送, 让推送下发服务化, 支持合并推送, 并发push
func (pm *PushManager) PushToUserByMap(username string, pushType string, echoCode string, data map[string]interface{}) {
	var queryResp *authroute.QueryUserRouteResp
	for {
		var err error
		queryResp, err = pm.routeCli.QueryUserRoute(context.Background(),
			&authroute.QueryUserRouteReq{
				Username: username,
			})
		if err != nil {
			log.Printf("failed to query user route: %v, retrying\n", err)
			if pm.retryWaitTime != nil {
				time.Sleep(*pm.retryWaitTime)
				continue
			}
			return
		}
		break
	}

	for _, v := range queryResp.Routes {
		gwCli := newGatewayClientFromAdAddr(v.GwAdvertiseAddrport)
		// 此处retry是因为网络原因, 比如网断了, 重试是安全的, 保证客户端至少收到消息一次
		// 客户端也得做自己的幂等, 防止消息重放导致副作用
	retry:
		_, err := gwCli.PushMessage(context.Background(), &gateway.PushMessageReq{
			SessionId: v.SessionId,
			PushType:  pushType,
			EchoCode:  echoCode,
			Data:      mustMarshal(data),
		})
		if err != nil {
			log.Printf("failed to push message: %v, retrying\n", err)
			if pm.retryWaitTime != nil {
				time.Sleep(*pm.retryWaitTime)
				goto retry
			}
			return
		}
	}
}

func mustMarshal(obj interface{}) []byte {
	bs, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	return bs
}
