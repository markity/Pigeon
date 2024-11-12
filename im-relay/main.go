package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	chatevloopconfig "pigeon/common/chatevloop_config"
	regetcd "pigeon/common/kitex_registry/etcd"
	"pigeon/common/push"
	"pigeon/im-relation/db/model"
	"pigeon/im-relay/api"
	"pigeon/im-relay/bizpush"
	"pigeon/im-relay/config"
	"pigeon/im-relay/rpcserver"
	"pigeon/kitex_gen/service/imrelay/imrelay"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var cfgFilePath = flag.String("cfg", "../config/im-relay/config.yaml", "config file path")

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

	debugModeLogger := logger.Default.LogMode(logger.Info)
	if !cfg.AppConfig.Debug {
		debugModeLogger = nil
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MysqlConfig.User, cfg.MysqlConfig.Pwd, cfg.MysqlConfig.Host, cfg.MysqlConfig.Port, cfg.MysqlConfig.Db)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: debugModeLogger,
	})
	if err != nil {
		panic(err)
	}
	err = model.Migrate(gormDB)
	if err != nil {
		panic(err)
	}

	pushMan := push.NewPushManager(time.Millisecond*50, nil)
	bp := bizpush.NewBisPusher(pushMan)

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
				BPush:        bp,
				DB:           gormDB,
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
