package distributelock

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type disLockClient struct {
	cli    *redis.Client
	prefix string
}

type lockEntry struct {
	cli     *redis.Client
	prefix  string
	key     string
	uuidPwd string
}

func NewDisLockClient(cli *redis.Client, prefix string) *disLockClient {
	return &disLockClient{
		cli:    cli,
		prefix: prefix,
	}
}

func (cli *disLockClient) Lock(key string, autoDeleteDur time.Duration) (*lockEntry, error) {
	id := uuid.NewString()
	k := cli.prefix + key
	for {
		ok, err := cli.cli.SetNX(context.Background(), k, id, autoDeleteDur).Result()
		if err != nil {
			return nil, err
		}
		if !ok {
			time.Sleep(time.Millisecond * 50)
			continue
		}
		return &lockEntry{
			cli:     cli.cli,
			prefix:  cli.prefix,
			key:     k,
			uuidPwd: id,
		}, nil
	}
}

func (cli *lockEntry) UnLock() error {
	script := `if redis.call("get", KEYS[1]) == ARGV[1] then  
    redis.call("del", KEYS[1])  
    return 1  
else  
    return 0  
end`
	v, err := cli.cli.Eval(context.Background(), script, []string{cli.key}, cli.uuidPwd).Result()
	if err != nil {
		return err
	}
	i, ok := v.(int64)
	if !ok {
		return errors.New("check lua script")
	}
	if i == 0 {
		return errors.New("key is removed")
	}

	return nil
}
