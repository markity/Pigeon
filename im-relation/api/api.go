package api

import (
	"pigeon/kitex_gen/service/imauthroute/imauthroute"
	"pigeon/kitex_gen/service/imgateway/imgateway"
	"pigeon/kitex_gen/service/imrelay/imrelay"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
)

func NewGatewayClientFromAdAddr(adAddr string) imgateway.Client {
	return imgateway.MustNewClient("im-gateway", client.WithHostPorts(adAddr))
}
func MustNewIMRelayClient(resolver discovery.Resolver) imrelay.Client {
	return imrelay.MustNewClient("im-relay", client.WithResolver(resolver))
}
func MustNewIMAuthRouteClient(resolver discovery.Resolver) imauthroute.Client {
	return imauthroute.MustNewClient("im-authroute", client.WithResolver(resolver))
}
