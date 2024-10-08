package chatevloopconfig

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"
	"stathat.com/c/consistent"
)

// chat eventloop配置
/*
	chat eventloop配置是有版本号的, 比如v1可能只有三个节点, 那么v1的key是: chatevloop_config/v1
	如果新增一个v2的版本, 那么就是chatevloop_config/v2
	key对应的value是json序列化的配置信息, 见下面结构体
	且还有一个key, 保存当前的version, chatevloop_config/version
*/

type Config struct {
	Version int64          `json:"version"`
	Nodes   []*ConfigEntry `json:"nodes"`
}

type ConfigEntry struct {
	NodeName string `json:"node_name"`
	IPPort   string `json:"ipport"`
}

// TODO: remove me, 写死最初的config
func InitConfig(cli *clientv3.Client) {
	cfg := Config{
		Version: 1,
		Nodes: []*ConfigEntry{
			{
				NodeName: "node1",
				IPPort:   "127.0.0.1:8001",
			},
			{
				NodeName: "node2",
				IPPort:   "127.0.0.1:8002",
			},
			{
				NodeName: "node3",
				IPPort:   "127.0.0.1:8002",
			},
		},
	}

	jsonBs, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}

	_, err = cli.KV.Txn(context.Background()).
		If(clientv3.Compare(clientv3.ModRevision("/chatevloop_config/version"), "=", 0)).
		Then(clientv3.OpPut("/chatevloop_config/version", "1"),
			clientv3.OpPut("/chatevloop_config/node_info/1", string(jsonBs))).Commit()
	if err != nil {
		panic(err)
	}
}

type updateChan struct {
	version  int64
	respChan chan int64
}

type waitUpdateChans struct {
	version int64
	ch      chan int64
}

type ChatevWatcher struct {
	mu sync.Mutex

	stopChan chan struct{}
	stopped  bool

	forceUpdateChan chan *updateChan

	currentVersion int64
	configs        map[int64]*Config

	lastConfig     map[string]*ConfigEntry
	lastConsistent *consistent.Consistent

	currentConfig     map[string]*ConfigEntry
	currentConsistent *consistent.Consistent

	waitUpdateChans []*waitUpdateChans
}

func NewWatcher(cli *clientv3.Client) *ChatevWatcher {
	InitConfig(cli)
	resp, err := cli.KV.Get(context.Background(), "/chatevloop_config", clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}

	if len(resp.Kvs) == 0 {
		panic("kv's size 0")
	}

	var version int64
	var configs = make(map[int64]*Config)

	var versionKVModReversion int64

	for _, kv := range resp.Kvs {
		if string(kv.Key) == "/chatevloop_config/version" {
			var err error
			version, err = strconv.ParseInt(string(kv.Value), 10, 64)
			if err != nil {
				panic(err)
			}
			versionKVModReversion = kv.ModRevision
		} else if strings.HasPrefix(string(kv.Key), "/chatevloop_config/node_info") {
			versionSplits := strings.Split(string(kv.Key), "/")
			if len(versionSplits) != 4 {
				log.Printf("ignore unexpected node info, key:%v, val:%v\n", string(kv.Key), string(kv.Value))
				continue
			}
			versionStr := versionSplits[len(versionSplits)-1]
			version, err := strconv.ParseInt(versionStr, 10, 64)
			if err != nil {
				log.Printf("parse version failed: %v\n", versionStr)
				continue
			}
			var cfg Config
			err = json.Unmarshal(kv.Value, &cfg)
			if err != nil {
				log.Printf("unexpected json: %v\n", string(kv.Value))
				continue
			}
			configs[version] = &cfg
		}
	}

	if version == 0 {
		panic("unexpected")
	}
	if configs[version] == nil {
		panic("unexpected")
	}

	keys := make([]int, 0)
	for k, _ := range configs {
		keys = append(keys, int(k))
	}

	sort.Ints(keys)

	for i := range keys {
		if i == 0 {
			continue
		}

		if keys[i] != keys[i-1]+1 {
			panic("unexpected")
		}
	}

	currentNodes := make(map[string]*ConfigEntry)

	c := consistent.New()
	for _, v := range configs[version].Nodes {
		c.Add(v.NodeName)
		currentNodes[v.NodeName] = v
	}

	lastNodes := make(map[string]*ConfigEntry)
	c2 := consistent.New()
	if version > 1 {
		for _, v := range configs[version-1].Nodes {
			c2.Add(v.NodeName)
			lastNodes[v.NodeName] = v
		}
	}

	ewc := &ChatevWatcher{
		currentVersion:    version,
		configs:           configs,
		lastConfig:        lastNodes,
		lastConsistent:    c2,
		currentConfig:     currentNodes,
		currentConsistent: c,
		stopChan:          make(chan struct{}, 1),
		forceUpdateChan:   make(chan *updateChan),

		waitUpdateChans: make([]*waitUpdateChans, 0),
	}

	go func() {
		wc := cli.Watch(clientv3.WithRequireLeader(context.Background()), "/chatevloop_config/version", clientv3.WithRev(versionKVModReversion+1))
		for {
			select {
			case rsp := <-wc:
				if rsp.Err() != nil {
					wc = cli.Watch(clientv3.WithRequireLeader(context.Background()), "/chatevloop_config/version", clientv3.WithRev(versionKVModReversion+1))
					continue
				}
				for _, event := range rsp.Events {
					value := string(event.Kv.Value)

					i64, err := strconv.ParseInt(value, 10, 64)
					if err != nil {
						panic("unexpected")
					}

					resp, err := cli.Get(context.Background(), "/chatevloop_config/node_info/"+value)
					if err != nil || len(resp.Kvs) != 1 {
						panic("unexpected")
					}

					var ent Config
					err = json.Unmarshal(resp.Kvs[0].Value, &ent)
					if err != nil {
						fmt.Println(string(resp.Kvs[0].Value))
						panic(err)
					}

					newConsistent := consistent.New()
					newCurrentConfig := make(map[string]*ConfigEntry)
					for _, v := range ent.Nodes {
						newCurrentConfig[v.NodeName] = v
						newConsistent.Add(v.NodeName)
					}

					ewc.mu.Lock()
					ewc.lastConsistent = ewc.currentConsistent
					ewc.lastConfig = ewc.currentConfig
					if i64 != ewc.currentVersion+1 {
						panic("unexpected")
					}
					ewc.configs[i64] = &ent
					ewc.currentConfig = newCurrentConfig
					ewc.currentVersion = i64
					ewc.currentConsistent = newConsistent
					ewc.mu.Unlock()
					newWcs := make([]*waitUpdateChans, 0)
					// ewc.waitUpdateChans仅仅由这个loop读写， 所以不需要加锁
					for _, v := range ewc.waitUpdateChans {
						if i64 >= version {
							v.ch <- i64
							continue
						}
						newWcs = append(newWcs, v)
					}
					ewc.waitUpdateChans = newWcs
					versionKVModReversion = event.Kv.Version
					log.Printf("chatevloop_config, version changed: %v\n", i64)
				}
			case <-ewc.stopChan:
				return
			case v := <-ewc.forceUpdateChan:
				ewc.mu.Lock()
				curVersion := ewc.currentVersion
				if curVersion >= v.version {
					ewc.mu.Unlock()
					v.respChan <- curVersion
					continue
				}

				ewc.waitUpdateChans = append(ewc.waitUpdateChans, &waitUpdateChans{
					version: v.version,
					ch:      v.respChan,
				})

				go ewc.mu.Unlock()
			}
		}
	}()

	return ewc
}

func (cwc *ChatevWatcher) Stop() {
	cwc.mu.Lock()
	defer cwc.mu.Unlock()
	if cwc.stopped {
		return
	}
	cwc.stopped = true
	cwc.stopChan <- struct{}{}
}

func (cwc *ChatevWatcher) GetNode(groupId string) (*ConfigEntry, int64) {
	cwc.mu.Lock()
	c := cwc.currentConsistent
	n, err := c.Get(groupId)
	if err != nil {
		panic(err)
	}
	m := cwc.currentConfig[n]
	v := cwc.currentVersion
	cwc.mu.Unlock()
	return m, v
}

func (cwc *ChatevWatcher) GetLastVersionNode(groupId string) (*ConfigEntry, int64) {
	cwc.mu.Lock()
	defer cwc.mu.Unlock()
	c := cwc.lastConsistent
	n, err := c.Get(groupId)
	if err != nil {
		return nil, 0
	}
	m := cwc.lastConfig[n]
	v := cwc.currentVersion - 1
	return m, v
}

func (cwc *ChatevWatcher) ForceUpdate(leastVersion int64) int64 {
	cwc.mu.Lock()
	resp := make(chan int64, 1)
	cwc.forceUpdateChan <- &updateChan{
		version:  leastVersion,
		respChan: resp,
	}
	cwc.mu.Unlock()
	v := <-resp
	return v
}
