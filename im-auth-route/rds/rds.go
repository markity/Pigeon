package rds

import (
	"context"
	"encoding/json"
	distributelock "pigeon/common/distribute_lock"
	"time"

	"github.com/redis/go-redis/v9"
)

type RdsAction struct {
	cli        *redis.Client
	dataPrefix string

	lock       *distributelock.DisLockClient
	lockPrefix string

	// 限制设备数
	deviceLimit *int
}

func NewRdsAction(cli *redis.Client, prefix string, deviceLimit int) *RdsAction {
	lockPrefix := prefix + "lock/"
	var devLimit *int
	if deviceLimit > 0 {
		devLimit = new(int)
		*devLimit = deviceLimit
	}
	return &RdsAction{
		cli:        cli,
		lock:       distributelock.NewDisLockClient(cli, lockPrefix),
		lockPrefix: lockPrefix,
		dataPrefix: prefix + "data/",
		// deveiceLimit为nil时不限制数量
		deviceLimit: devLimit,
	}
}

func (act *RdsAction) LockUsername(username string, ttl time.Duration) (*distributelock.LockEntry, error) {
	return act.lock.Lock(username, ttl)
}

type SessionEntry struct {
	LoginAt      int64  `json:"login_at"`
	Username     string `json:"username"`
	SessionId    string `json:"session_id"`
	DeviceDesc   string `json:"device_desc"`
	GwAdAddrPort string `json:"gw_ad_addr_port"`
}

type LoginResult struct {
	Success     bool
	Version     int64
	AllSessions map[string]*SessionEntry
}

/*
登录流程:
lock user
查mysql数据看, 检查登录信息
创建route并获取所有session(lua脚本也就是下面的登录流程)
unlock user
*/

/*
version存储:
key: prefix/data/route/version/{username}
val: int类型的版本号

路由存储:
key: prefix/data/route/user/{username}
val: map[sessionId]json编码

key: prefix/data/route/session/{sessionId}
val: json编码
*/

func (act *RdsAction) Login(session *SessionEntry) (*LoginResult, error) {
	script := `
-- Keys[1]: username
-- Keys[2]: sessionId
-- Keys[3]: prefix
-- ARGV[1]: json编码session信息
-- ARGV[2]: 可选, limit, 为空则不限制数量
-- Result[1]: 一个boolean, true表示登录成功, false表示设备限制而登录失败
-- Result[2]: version
-- Result[3,4,5,6...]: 现在的所有设备的json编码

local keyUser = KEYS[3].."route/user/"..KEYS[1]
local keySession = KEYS[3].."route/session/"..KEYS[2]
local keyVersion = KEYS[3].."route/version/"..KEYS[1]

local result = {}

-- 如果limit不为空, 则检查设备数量
if ARGV[2] ~= nil and redis.call('HLEN', keyUser) >= tonumber(ARGV[2]) then
	table.insert(result,0)
	return result
end

-- 说明可以登录, 将session信息写入
-- hset keyUser, sessionId, json编码
redis.call('HSET', keyUser, KEYS[2], ARGV[1])
-- set keySession, json编码
redis.call('SET', keySession, ARGV[1])
-- version++
redis.call('INCR', keyVersion)

local version = redis.call('GET', keyVersion)

table.insert(result,1)
table.insert(result,version)

local keys = redis.call('HGETALL', keyUser)
-- 遍历HGETALL返回的列表, 把value存入result列表
for i = 1, #keys, 2 do
    table.insert(result, keys[i+1])
end
-- 返回包含所有值的列表
return result
`
	data, err := json.Marshal(session)
	if err != nil {
		// 这里通常不会失败, 把事情做绝一点
		panic(err)
	}

	var cmd *redis.Cmd
	if act.deviceLimit == nil {
		cmd = act.cli.Eval(context.Background(), script,
			[]string{session.Username, session.SessionId, act.dataPrefix},
			string(data))
	} else {
		cmd = act.cli.Eval(context.Background(), script,
			[]string{session.Username, session.SessionId, act.dataPrefix},
			string(data), *act.deviceLimit)
	}
	if err := cmd.Err(); err != nil {
		return nil, err
	}
	result, err := cmd.Slice()
	if err != nil {
		return nil, err
	}

	ok := result[0].(int64) == 1
	if !ok {
		return &LoginResult{
			Success: false,
		}, nil
	}

	version, _ := result[1].(int64)
	sessions := make(map[string]*SessionEntry)
	for _, v := range result[2:] {
		var entry SessionEntry
		err := json.Unmarshal([]byte(v.(string)), &entry)
		if err != nil {
			return nil, err
		}

		sessions[entry.SessionId] = &entry
	}

	return &LoginResult{
		Success:     ok,
		Version:     version,
		AllSessions: sessions,
	}, nil
}
