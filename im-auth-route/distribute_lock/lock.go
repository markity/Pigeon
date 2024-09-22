package distributelock

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type disLockClient struct {
	cli    *redis.Client
	prefix string
}

func NewDisLockClient(cli *redis.Client, prefix string) *disLockClient {
	return &disLockClient{
		cli:    cli,
		prefix: prefix,
	}
}

func (cli *disLockClient) Lock(key string) error {
	for {
		ok, err := cli.cli.SetNX(context.Background(), cli.prefix+key, 1, 0).Result()
		if err != nil {
			return err
		}
		if !ok {
			time.Sleep(time.Millisecond * 50)
			continue
		}
		return nil
	}
}

func (cli *disLockClient) UnLock(key string) error {
	return cli.cli.Del(context.Background(), cli.prefix+key).Err()
}
