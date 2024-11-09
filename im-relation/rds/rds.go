package rds

import (
	"context"
	"fmt"
	distributelock "pigeon/common/distribute_lock"
	"time"

	"github.com/redis/go-redis/v9"
)

type RdsAction struct {
	cli        *redis.Client
	dataPrefix string

	lock       *distributelock.DisLockClient
	lockPrefix string
}

func NewRdsAction(cli *redis.Client, prefix string) *RdsAction {
	lockPrefix := prefix + "lock/"
	return &RdsAction{
		cli:        cli,
		lock:       distributelock.NewDisLockClient(cli, lockPrefix),
		lockPrefix: lockPrefix,
		dataPrefix: prefix + "data/",
	}
}

// 10000000
const GroupIDStart = 100000000

func (act *RdsAction) GenerateGroupId() (string, error) {
	script := `
local prefix = KEYS[1]
local groupIdStart = KEYS[2]

local keyGroupGen = prefix.."group_gen"

local groupId = redis.call('GET', keyGroupGen)
if groupId == false then
	redis.call('SET', keyGroupGen, groupIdStart)
	return groupIdStart
else
	redis.call('INCR', keyGroupGen)
	return redis.call('GET', keyGroupGen)
end

-- unexpected
return nil
`

	cmd := act.cli.Eval(context.Background(), script, []string{act.dataPrefix, fmt.Sprint(GroupIDStart)})
	err := cmd.Err()
	if err != nil {
		return "", err
	}
	r, err := cmd.Text()
	if err != nil {
		return "", err
	}
	return r, nil
}

// 锁群, relation会用到这个锁, 防止并发冲突
func (act *RdsAction) LockGroup(groupId string, ttl time.Duration) (*distributelock.LockEntry, error) {
	return act.lock.Lock(groupId, ttl)
}
