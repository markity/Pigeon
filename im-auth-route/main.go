package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	distributelock "pigeon/common/distribute_lock"
	regetcd "pigeon/common/kitex-registry/etcd"
	"pigeon/im-auth-route/config"
	"pigeon/im-auth-route/rpcserver"
	"pigeon/kitex_gen/service/imauthroute/imauthroute"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	rdsAddrPort := fmt.Sprintf("%v:%v", cfg.RedisConfig.Host, cfg.RedisConfig.Port)
	rdsCli := redis.NewClient(&redis.Options{
		Addr: rdsAddrPort,
	})
	if err := rdsCli.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MysqlConfig.User, cfg.MysqlConfig.Pwd, cfg.MysqlConfig.Host, cfg.MysqlConfig.Port, cfg.MysqlConfig.Db)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var eps []string
	for _, v := range cfg.EtcdConfig {
		eps = append(eps, fmt.Sprintf("%v:%v", v.Host, v.Port))
	}

	reg, err := regetcd.NewEtcdRegistry(eps)
	if err != nil {
		panic(err)
	}
	listenAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%v:%v", cfg.RPCServerConfig.Host, cfg.RPCServerConfig.Port))
	if err != nil {
		panic(err)
	}
	adAddr, err := net.ResolveTCPAddr("tcp", cfg.AppConfig.RPCAdvertiseAddrport)
	if err != nil {
		panic(err)
	}
	server := imauthroute.NewServer(&rpcserver.RPCServer{
		RPCContext: rpcserver.RPCContext{
			DB:   db,
			Lock: distributelock.NewDisLockClient(rdsCli, cfg.RedisConfig.KeyPrefix+"lock/"),
		},
	}, server.WithRegistry(reg), server.WithServiceAddr(listenAddr),
		server.WithRegistryInfo(&registry.Info{
			ServiceName: "im-auth-route",
			Addr:        adAddr,
		}))
	err = server.Run()
	if err != nil {
		panic(err)
	}
}
