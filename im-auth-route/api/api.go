package api

import (
	"pigeon/kitex_gen/service/imgateway/imgateway"

	"github.com/cloudwego/kitex/client"
)

func NewGatewayClientFromAdAddr(adAddr string) imgateway.Client {
	return imgateway.MustNewClient(adAddr, client.WithHostPorts(adAddr))
}
