package main

import (
	"context"
	"flag"
	"fmt"
	"pigeon/im-auth-route/config"
	distributelock "pigeon/im-auth-route/distribute_lock"

	"github.com/redis/go-redis/v9"
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
	cli := redis.NewClient(&redis.Options{
		Addr: rdsAddrPort,
	})
	if err := cli.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	disLock := distributelock.NewDisLockClient(cli, cfg.RedisConfig.KeyPrefix+"/dis_lock/")
}
