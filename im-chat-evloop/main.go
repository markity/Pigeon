package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	regetcd "pigeon/common/kitex-registry/etcd"
	"pigeon/im-chat-evloop/api"
	"pigeon/im-chat-evloop/config"
	rpcserver "pigeon/im-chat-evloop/rpc-server"
	"pigeon/kitex_gen/service/imchatevloop/imchatevloop"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
)

var cfgFilePath = flag.String("cfg", "../config/im-chat-evloop/config.yaml", "config file path")

func main() {
	// 加载配置文件
	var cfg *config.Config
	flag.Parse()
	if *cfgFilePath == "" {
		fmt.Println("config file must be specified, use --help")
		return
	}
	cfg = config.MustGetConfigFromFile(*cfgFilePath)

	// 启动服务
	rpcAddrPort := fmt.Sprintf("%v:%v", cfg.RPCServerConfig.Host, cfg.RPCServerConfig.Port)
	log.Printf("rpc server start at %v...\n", rpcAddrPort)

	eps := make([]string, 0, len(cfg.EtcdConfig))
	for _, entry := range cfg.EtcdConfig {
		addrPort := fmt.Sprintf("%v:%v", entry.Host, entry.Port)
		eps = append(eps, addrPort)
	}

	res, err := regetcd.NewEtcdResolver(eps)
	if err != nil {
		panic(err)
	}

	// 跑rpc server
	reg, err := regetcd.NewEtcdRegistry(eps)
	if err != nil {
		panic(err)
	}
	rpcserverAddr, err := net.ResolveTCPAddr("tcp", rpcAddrPort)
	if err != nil {
		panic(err)
	}
	server := imchatevloop.NewServer(
		&rpcserver.RPCServer{
			RelayCli: api.MustNewIMRelayClient(res),
		},
		server.WithServiceAddr(rpcserverAddr),
		server.WithRegistry(reg),
		server.WithRegistryInfo(&registry.Info{
			ServiceName: "im-chatevloop",
			Addr:        rpcserverAddr,
		}))
	err = server.Run()
	if err != nil {
		panic(err)
	}
}
