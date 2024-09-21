package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"pigeon/im-ipconf/config"
	"sort"
	"sync"

	"github.com/gin-gonic/gin"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var cfgFilePath = flag.String("cfg", "", "config file path")

type MetricsEtcdData struct {
	Name                 string
	Conns                uint64
	TCPAdvertiseAddrPort string
	RPCAdvertiseAddrPort string
}

type Entry struct {
	Addr string `json:"addr"`
}

type JSONResp struct {
	Code int `json:"code"`
	Data []Entry
}

var mu sync.RWMutex
var info map[string]*MetricsEtcdData

func main() {
	// 加载配置文件
	var cfg *config.Config
	flag.Parse()
	if *cfgFilePath == "" {
		fmt.Println("config file must be specified, use --help")
		return
	}
	cfg = config.MustGetConfigFromFile(*cfgFilePath)

	eps := make([]string, 0, len(cfg.EtcdConfig))
	for _, v := range cfg.EtcdConfig {
		eps = append(eps, fmt.Sprintf("%v:%v", v.Host, v.Port))
	}

	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints: eps,
	})
	if err != nil {
		panic(err)
	}

	resp, err := etcdCli.Get(context.Background(), "im-gateway-metrics/", clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	data := make(map[string]*MetricsEtcdData)
	for _, kv := range resp.Kvs {
		var m MetricsEtcdData
		err := json.Unmarshal(kv.Value, &m)
		if err != nil {
			log.Printf("failed to unmarshal watch im-gateway-metrics node info key:%v val:%v\n", string(kv.Key), string(kv.Value))
			continue
		}
		data[string(kv.Key)] = &m
	}
	mu.Lock()
	info = data
	mu.Unlock()

	go func() {
		rev := resp.Header.Revision
		c := etcdCli.Watch(clientv3.WithRequireLeader(context.Background()),
			"im-gateway-metrics/", clientv3.WithPrefix(), clientv3.WithRev(rev+1))
		for {
			entry := <-c
			if err := entry.Err(); err == rpctypes.ErrNoLeader {
				log.Printf("leader partitioned, retry watch")
				resp, err := etcdCli.Get(context.Background(), "im-gateway-metrics/", clientv3.WithPrefix())
				if err != nil {
					panic(err)
				}
				data := make(map[string]*MetricsEtcdData)
				for _, kv := range resp.Kvs {
					var m MetricsEtcdData
					err := json.Unmarshal(kv.Value, &m)
					if err != nil {
						log.Printf("failed to unmarshal watch im-gateway-metrics node info key:%v val:%v\n", string(kv.Key), string(kv.Value))
						continue
					}
					data[string(kv.Key)] = &m
				}
				mu.Lock()
				info = data
				mu.Unlock()
				c = etcdCli.Watch(clientv3.WithRequireLeader(context.Background()),
					"im-gateway-metrics/", clientv3.WithPrefix(), clientv3.WithRev(rev))
				continue
			}
			newMap := make(map[string]*MetricsEtcdData)
			for k, v := range info {
				newMap[k] = v
			}
			for _, ev := range entry.Events {
				var m MetricsEtcdData
				if ev.Type == mvccpb.PUT {
					err := json.Unmarshal(ev.Kv.Value, &m)
					if err != nil {
						log.Printf("failed to unmarshal watch im-gateway-metrics node info key:%v val:%v\n", string(ev.Kv.Key), string(ev.Kv.Value))
						continue
					}
					newMap[string(ev.Kv.Key)] = &m
				} else if ev.Type == mvccpb.DELETE {
					delete(newMap, string(ev.Kv.Key))
				}
			}
			mu.Lock()
			info = newMap
			mu.Unlock()
		}
	}()

	tcpAddrPort := fmt.Sprintf("%v:%v", cfg.HTTPServerConfig.Host, cfg.HTTPServerConfig.Port)
	engine := gin.New()
	engine.GET("/ipconf", func(ctx *gin.Context) {
		mu.RLock()
		data := info
		mu.RUnlock()
		// 排序
		s := make([]*MetricsEtcdData, 0, len(data))
		for _, v := range data {
			s = append(s, v)
		}
		sort.Slice(s, func(i, j int) bool {
			return s[i].Conns < s[j].Conns
		})
		var resp JSONResp
		resp.Data = make([]Entry, 0)
		resp.Code = 0
		for i := 0; i < 3 && i < len(s); i++ {
			resp.Data = append(resp.Data, Entry{
				Addr: s[i].TCPAdvertiseAddrPort,
			})
		}
		ctx.JSON(200, resp)
	})
	engine.Run(tcpAddrPort)
}
