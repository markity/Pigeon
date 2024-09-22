package main

import (
	"flag"
	"fmt"
	"log"
	regetcd "pigeon/common/kitex-registry/etcd"
	"pigeon/im-gateway/api"
	"pigeon/im-gateway/config"
	metrics "pigeon/im-gateway/mertics"
	"pigeon/im-gateway/tcpserver"
	"sync"
	"time"

	goreactor "github.com/markity/go-reactor"
	eventloop "github.com/markity/go-reactor/pkg/event_loop"
	clientv3 "go.etcd.io/etcd/client/v3"
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

	tcpAddrPort := fmt.Sprintf("%v:%v", cfg.TCPServerConfig.Host, cfg.TCPServerConfig.Port)
	// rpcAddrPort := fmt.Sprintf("%v:%v", cfg.RPCServerConfig.Host, cfg.RPCServerConfig.Port)
	numThread := cfg.TCPServerConfig.WorkerNum
	log.Printf("starting reactor server, addrport: %v num thread: %v\n", tcpAddrPort, numThread)

	eps := make([]string, 0, len(cfg.EtcdConfig))
	for _, entry := range cfg.EtcdConfig {
		addrPort := fmt.Sprintf("%v:%v", entry.Host, entry.Port)
		eps = append(eps, addrPort)
	}

	resolver, err := regetcd.NewEtcdResolver(eps)
	if err != nil {
		panic(err)
	}

	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints: eps,
	})
	if err != nil {
		panic(err)
	}

	relayCli := api.MustNewIMRelayClient(resolver)
	authRouteCli := api.MustNewIMAuthRouteClient(resolver)
	evloopRoute := sync.Map{}

	// 跑rpc server
	// imgateway.NewServer()

	// 跑reactor event loop
	mainEvLoop := eventloop.NewEventLoop()
	tcpServer := goreactor.NewTCPServer(mainEvLoop, tcpAddrPort, numThread, goreactor.LeastConnection())
	tcpServer.GetAllLoops()
	tcpServer.SetConnectionCallback(tcpserver.OnConn)
	tcpServer.SetMessageCallback(tcpserver.OnMessage)
	tcpServer.Start()
	mainLoop, others := tcpServer.GetAllLoops()
	evloopCtx := tcpserver.EvloopContext{
		RelayCli:          relayCli,
		AuthRouteCli:      authRouteCli,
		ConnMetrics:       &metrics.Conns,
		HeartbeatInterval: time.Millisecond * time.Duration(cfg.AppConfig.HeartbeatIntervalMs),
		HeartbeatTimeout:  time.Millisecond * time.Duration(cfg.AppConfig.CloseConnIntervalMs),
		EvloopRoute:       &evloopRoute,
	}
	// 注意这里是值传递, 有个拷贝的过程
	tcpserver.SetUpEvLoopContext(mainLoop, evloopCtx)
	for _, otherLoop := range others {
		tcpserver.SetUpEvLoopContext(otherLoop, evloopCtx)
	}
	// 启动metrics上报
	mainEvLoop.DoOnLoop(func(el eventloop.EventLoop) {
		go metrics.GoLoopUpdateMetrics(etcdCli, &metrics.MetricsEtcdData{
			Name:                 cfg.AppConfig.Name,
			Conns:                0,
			TCPAdvertiseAddrPort: cfg.AppConfig.TCPAdvertiseAddrport,
			RPCAdvertiseAddrPort: cfg.AppConfig.RPCAdvertiseAddrport,
		}, time.Millisecond*time.Duration(cfg.AppConfig.MetricsUpdateIntervalMs))
	})
	mainEvLoop.Loop()
}
