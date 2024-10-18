package api

import (
	"pigeon/kitex_gen/service/imchatevloop/imchatevloop"
	"pigeon/kitex_gen/service/imrelay/imrelay"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
)

func MustNewIMRelayClient(resolver discovery.Resolver) imrelay.Client {
	return imrelay.MustNewClient("im-relay", client.WithResolver(resolver))
}

func MustNewChatEvLoopCliFromAdAddr(adAddr string) imchatevloop.Client {
	return imchatevloop.MustNewClient("im-chatevloop", client.WithHostPorts(adAddr))
}
