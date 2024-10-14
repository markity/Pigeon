package main

import (
	"flag"
	"fmt"
	"net"

	chatevloopconfig "pigeon/common/chatevloop-config"
	regetcd "pigeon/common/kitex-registry/etcd"
	"pigeon/im-relay/api"
	"pigeon/im-relay/config"
	"pigeon/im-relay/rpcserver"
	"pigeon/kitex_gen/service/imrelay/imrelay"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
)

var cfgFilePath = flag.String("cfg", "", "config file path")

func main() {
	// 加载配置文件
	var cfg *config.Config
	flag.Parse()
	if *cfgFilePath == "" {
		fmt.Println("config file must be specified, use --help")
		return
	}
	cfg = config.MustGetConfigFromFile(*cfgFilePath)

	var etcdEndpoints []string
	for _, v := range cfg.EtcdConfig {
		etcdEndpoints = append(etcdEndpoints, fmt.Sprintf("%v:%v", v.Host, v.Port))
	}

	reg, err := regetcd.NewEtcdRegistry(etcdEndpoints)
	if err != nil {
		panic(err)
	}

	res, err := regetcd.NewEtcdResolver(etcdEndpoints)
	if err != nil {
		panic(err)
	}

	relationCli := api.MustNewIMRelationClient(res)

	listenAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%v:%v", cfg.RPCServerConfig.Host, cfg.RPCServerConfig.Port))
	if err != nil {
		panic(err)
	}
	adAddr, err := net.ResolveTCPAddr("tcp", cfg.AppConfig.RPCAdvertiseAddrport)
	if err != nil {
		panic(err)
	}
	server := imrelay.NewServer(
		&rpcserver.RPCServer{
			RPCContext: rpcserver.RPCContext{
				EvCfgWatcher: chatevloopconfig.NewWatcher(etcdEndpoints),
				RelationCli:  relationCli,
			},
		},
		server.WithRegistry(reg), server.WithServiceAddr(listenAddr),
		server.WithRegistryInfo(&registry.Info{
			ServiceName: "im-relay",
			Addr:        adAddr,
		}))

	if err := server.Run(); err != nil {
		panic(err)
	}
}
