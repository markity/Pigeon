package api

import (
	authroute "pigeon/kitex_gen/service/imauthroute/imauthroute"
	relay "pigeon/kitex_gen/service/imrelay/imrelay"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
)

func MustNewIMRelayClient(resolver discovery.Resolver) relay.Client {
	return relay.MustNewClient("im-relay", client.WithResolver(resolver))
}

func MustNewIMAuthRouteClient(resolver discovery.Resolver) authroute.Client {
	return authroute.MustNewClient("im-authroute", client.WithResolver(resolver))
}
