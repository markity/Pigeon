package api

import (
	authroute "pigeon/kitex_gen/service/imauthroute/imauthroute"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
)

func MustNewIMRelayClient(resolver discovery.Resolver) interface{} {
	return nil
}

func MustNewIMAuthRouteClient(resolver discovery.Resolver) (authroute.Client, error) {
	return authroute.NewClient("im-auth-route", client.WithResolver(resolver))
}
