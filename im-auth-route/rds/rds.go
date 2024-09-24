package rds

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	distributelock "pigeon/common/distribute_lock"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/imauthroute"

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

type LoginResult struct {
	Success     bool
	Version     int64
	AllSessions []*base.SessionEntry
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

func (act *RdsAction) Login(session *base.SessionEntry) (*LoginResult, error) {
	script := `
-- Keys[1]: username
-- Keys[2]: sessionId
-- Keys[3]: prefix
-- ARGV[1]: json编码session信息
-- ARGV[2]: 可选, limit, 为空则不限制数量
-- Result[1]: 一个int64, 1表示登录成功, 0表示设备限制而登录失败
-- Result[2]: version
-- Result[3,4,5,6...]: 现在的所有设备的json编码

local username = KEYS[1]
local sessionId = KEYS[2]
local prefix = KEYS[3]
local jsonData = ARGV[1]
local limit = ARGV[2]

local keyUser = prefix.."route/user/"..username
local keySession = prefix.."route/session/"..sessionId
local keyVersion = prefix.."route/version/"..username

local result = {}

-- 如果limit不为空, 则检查设备数量
if limit ~= nil and redis.call('HLEN', keyUser) >= tonumber(limit) then
	table.insert(result,0)
else
	table.insert(result,1)
end

if result[1] == 1 then
	-- 说明可以登录, 将session信息写入
	-- hset keyUser, sessionId, json编码
	redis.call('HSET', keyUser, sessionId, jsonData)
	-- set keySession, json编码
	redis.call('SET', keySession, jsonData)
	-- version++
	redis.call('INCR', keyVersion)
end

local version = redis.call('GET', keyVersion)
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
	version_, err := strconv.ParseInt(result[1].(string), 10, 64)
	if err != nil {
		panic(err)
	}
	version := version_
	sessions := make([]*base.SessionEntry, 0, len(result[2:]))
	for _, v := range result[2:] {
		var entry base.SessionEntry
		err := json.Unmarshal([]byte(v.(string)), &entry)
		if err != nil {
			return nil, err
		}

		sessions = append(sessions, &entry)
	}

	return &LoginResult{
		Success:     ok,
		Version:     version,
		AllSessions: sessions,
	}, nil
}

type LogoutResult struct {
	// 理论上不会失败, 这里理论上肯定是true, 因为gateway是有登录状态的
	Success     bool
	Version     int64
	AllSessions []*base.SessionEntry
}

/*
登出流程: 直接lua脚本尝试登出就行, 登入要加锁是为了防止用户使用旧密码登入而做的同步
	登出就不需要这些了
*/

func (act *RdsAction) Logout(username string, sessionId string) (*LogoutResult, error) {
	script := `
-- Keys[1]: username
-- Keys[2]: sessionId
-- Keys[3]: prefix
-- Result[1]: 一个int64, 1表示登录成功, 0表示设备限制而登录失败
-- Result[2]: version
-- Result[3,4,5,6...]: 现在的所有设备的json编码

local username = KEYS[1]
local sessionId = KEYS[2]
local prefix = KEYS[3]

local keyUser = prefix.."route/user/"..username
local keySession = prefix.."route/session/"..sessionId
local keyVersion = prefix.."route/version/"..username


local result = {}

local val = redis.call('HGET', keyUser, sessionId)
if val == false then
	table.insert(result, 0)
else
	table.insert(result, 1)
	-- 做操作, 删除session信息
	redis.call('HDEL', keyUser, sessionId)
	redis.call('DEL', keySession)
	redis.call('INCR', keyVersion)
end

local version = redis.call('GET', keyVersion)
if version == false then
	table.insert(result, 0)
else
	table.insert(result, version)
end


local keys = redis.call('HGETALL', keyUser)
-- 遍历HGETALL返回的列表, 把value存入result列表
for i = 1, #keys, 2 do
    table.insert(result, keys[i+1])
end
-- 返回包含所有值的列表
return result
`

	results, err := act.cli.Eval(context.Background(), script,
		[]string{username, sessionId, act.dataPrefix}).Slice()

	if err != nil {
		return nil, err
	}

	ok := results[0].(int64) == 1
	if !ok {
		return &LogoutResult{
			Success: false,
		}, nil
	}

	version, _ := results[1].(int64)
	sessions := make([]*base.SessionEntry, 0, len(results[2:]))
	for _, v := range results[2:] {
		var entry base.SessionEntry
		err := json.Unmarshal([]byte(v.(string)), &entry)
		if err != nil {
			return nil, err
		}

		sessions = append(sessions, &entry)
	}

	return &LogoutResult{
		Success:     ok,
		Version:     version,
		AllSessions: sessions,
	}, nil
}

type ForceOfflineResult struct {
	Code        imauthroute.ForceOfflineResp_ForceOfflineRespCode
	Version     int64
	AllSessions map[string]*base.SessionEntry
}

func (act *RdsAction) ForceOffline(username string, fromSessionId string, targetSessionId string) (*ForceOfflineResult, error) {

	/*
		SUCCESS = 0;
		FROM_SESSION_NOT_FOUND = 1;
		TO_SESSION_NOT_FOUND = 2
	*/
	script := `
	-- Keys[1]: username
	-- Keys[2]: from_sessionId
	-- Keys[3]: target_sessionId
	-- Keys[4]: prefix
	-- Result[1]: 一个int64, 0表示登录成功, 1表示没有fromsession, 2代表有fromsession但是没有targetsession
	-- Result[2]: version
	-- Result[3,4,5,6...]: 现在的所有设备的json编码
	
	local username = KEYS[1]
	local fromSessionId = KEYS[2]
	local targetSessionId = KEYS[3]
	local prefix = KEYS[4]
	
	local keyUser = prefix.."route/user/"..username
	local keySession = prefix.."route/session/"..sessionId
	local keyVersion = prefix.."route/version/"..username
	
	local result = {}

	local fromSessionVal = redis.call('HGET', keyUser, fromSessionId)
	local toSessionVal = redis.call('HGET', keyUser, targetSessionId)
	if fromSessionVal == nil then
		table.insert(result, 1)
	elseif toSessionVal == nil then
		table.insert(result, 2)
	else
		-- 都存在, 那么删除targetSessionId的session信息
		redis.call('HDEL', keyUser, targetSessionId)
		redis.call('DEL', keySession)
		table.insert(result, 0)
		-- 更新版本号
		redis.call('INCR', keyVersion)
	end

	local version = redis.call('GET', keyVersion)
	table.insert(result,1)
	table.insert(result,version)

	local keys = redis.call('HGETALL', keyUser)
	-- 遍历HGETALL返回的列表, 把value存入result列表
	for i = 1, #keys, 2 do
		table.insert(result, keys[i+1])
	end
	return result
`
	results, err := act.cli.Eval(context.Background(), script, []string{username, fromSessionId, targetSessionId, act.dataPrefix}).Slice()
	if err != nil {
		return nil, err
	}

	code := results[0].(int64)
	version, _ := results[1].(int64)
	sessions := make(map[string]*base.SessionEntry)
	for _, v := range results[2:] {
		var entry base.SessionEntry
		err := json.Unmarshal([]byte(v.(string)), &entry)
		if err != nil {
			return nil, err
		}

		sessions[entry.SessionId] = &entry
	}

	return &ForceOfflineResult{
		Code:        imauthroute.ForceOfflineResp_ForceOfflineRespCode(code),
		Version:     version,
		AllSessions: sessions,
	}, nil
}

// 如果sessionId不存在, 返回nil, nil
func (act *RdsAction) QuerySessionRoute(sessionId string) (*base.SessionEntry, error) {
	script := `
-- Keys[1]: sessionId
-- Keys[2]: prefix
local sessionId = KEYS[1]
local prefix = KEYS[2]

local keySession = prefix.."route/session/"..sessionId

local val = redis.call('GET', keySession)
return val
`
	cmd, err := act.cli.Eval(context.Background(), script, []string{sessionId, act.dataPrefix}).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	s := cmd.(string)
	var entry base.SessionEntry
	if err := json.Unmarshal([]byte(s), &entry); err != nil {
		return nil, err
	}
	return &entry, nil
}

// 如果sessionId不存在, 返回nil, nil
func (act *RdsAction) QueryUserRoute(username string) ([]*base.SessionEntry, error) {
	script := `
-- Keys[1]: username
-- Keys[2]: prefix
local username = KEYS[1]
local prefix = KEYS[2]

local keyUser = prefix.."route/user/"..username

local result = {}

local keys = redis.call('HGETALL', keyUser)
-- 遍历HGETALL返回的列表, 把value存入result列表
for i = 1, #keys, 2 do
	table.insert(result, keys[i+1])
end

return result
`
	cmd := act.cli.Eval(context.Background(), script, []string{username, act.dataPrefix})
	if err := cmd.Err(); err != nil {
		return nil, err
	}

	s, err := cmd.Slice()
	if err != nil {
		return nil, err
	}

	var result []*base.SessionEntry
	for _, entry := range s {
		var ent base.SessionEntry
		if err := json.Unmarshal([]byte(entry.(string)), &ent); err != nil {
			return nil, err
		}
		result = append(result, &ent)
	}

	return result, nil
}
