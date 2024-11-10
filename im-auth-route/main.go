package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	regetcd "pigeon/common/kitex_registry/etcd"
	"pigeon/im-auth-route/config"
	"pigeon/im-auth-route/db"
	"pigeon/im-auth-route/db/model"
	"pigeon/im-auth-route/rds"
	"pigeon/im-auth-route/rpcserver"
	"pigeon/kitex_gen/service/imauthroute/imauthroute"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var cfgFilePath = flag.String("cfg", "../config/im-auth-route/config.yaml", "config file path")

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
	if err := model.Migrate(gormDB); err != nil {
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
			DB:  db.NewDB(gormDB),
			Rds: rds.NewRdsAction(rdsCli, cfg.RedisConfig.KeyPrefix, cfg.AppConfig.DeviceNumLimit),
		},
	}, server.WithRegistry(reg), server.WithServiceAddr(listenAddr),
		server.WithRegistryInfo(&registry.Info{
			ServiceName: "im-authroute",
			Addr:        adAddr,
		}))
	err = server.Run()
	if err != nil {
		panic(err)
	}
}
