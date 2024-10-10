package api

import (
	"pigeon/kitex_gen/service/imchatevloop/imchatevloop"
	"pigeon/kitex_gen/service/imgateway/imgateway"
	"pigeon/kitex_gen/service/imrelation/imrelation"

	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
)

func MustNewChatEvLoopCliFromAdAddr(adAddr string) imchatevloop.Client {
	return imchatevloop.MustNewClient("im-chatevloop", client.WithHostPorts(adAddr))
}

func MustNewIMGatewayClientFromAdAddr(adAddr string) imgateway.Client {
	return imgateway.MustNewClient("im-gateway", client.WithHostPorts(adAddr))
}

var relationOnce sync.Once
var relation imrelation.Client

// 这里做个懒汉模式的单例
func MustNewIMRelationClient(resolver discovery.Resolver) imrelation.Client {
	relationOnce.Do(func() {
		relation = imrelation.MustNewClient("im-relation", client.WithResolver(resolver))
	})
	return relation
}
