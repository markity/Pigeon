package regetcd

// 为kitex准备的etcd服务发现

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// etcd其实就是一个高可用kv
type storeInfo struct {
	Tags    map[string]string
	Network string
	Addr    string
	Weight  int
}

type etcdRegistry struct {
	etcdClient *clientv3.Client
	addrString string
	exitChan   chan struct{}
	exitOKChan chan struct{}
}

func NewEtcdRegistry(endpoints []string) (registry.Registry, error) {
	return NewEtcdRegistryWithAuth(endpoints, "", "")
}

func NewEtcdRegistryWithAuth(endpoints []string, username, password string) (registry.Registry, error) {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
		Username:  username,
		Password:  password,
	})
	if err != nil {
		return nil, err
	}

	s, err := getLocalIPv4Host()
	if err != nil {
		return nil, err
	}

	return &etcdRegistry{
		etcdClient: etcdClient,
		addrString: s,
		exitChan:   make(chan struct{}),
		exitOKChan: make(chan struct{}),
	}, nil
}

func (r *etcdRegistry) Register(info *registry.Info) error {
	_, p, err := net.SplitHostPort(info.Addr.String())
	if err != nil {
		panic(err)
	}
	putContent_, _ := json.Marshal(&storeInfo{
		Tags:    info.Tags,
		Network: info.Addr.Network(),
		Addr:    r.addrString + ":" + p,
		Weight:  info.Weight,
	})

	putContent := string(putContent_)

	// 租期15s
	lease, err := r.etcdClient.Grant(context.Background(), 15)
	if err != nil {
		log.Printf("[kitex-registry] failed to Register(grant lease), retrying: %s\n", err)
		return err
	}

	_, err = r.etcdClient.Put(context.Background(), fmt.Sprintf("pigeon-registry/%s/%d", info.ServiceName,
		lease.ID), putContent, clientv3.WithLease(lease.ID))
	if err != nil {
		log.Printf("[kitex-registry] failed to Registry(put key), retrying: %s\n", err)
		return err
	}

	// 保证第一次注册成功, 将内容写入etcd, 后面起一个goroutine持续保持lease
	go func() {
		name := info.ServiceName
		leaseId := lease.ID
		log.Println("[kitex-registry] succeed to register, now keeping lease")

		// 退出时尝试把之前注册的key都销毁, 不能销毁也没关系, etcd会保底的
		defer func() {
			r.etcdClient.Delete(context.Background(), fmt.Sprintf("kitex-registry/%s/%v", name, leaseId))
		}()

		for {
			_, err := r.etcdClient.KeepAliveOnce(context.Background(), lease.ID)
			if err != nil {
				log.Printf("[kitex-registry] failed to Registry(keepalive), retrying: %s\n", err)
				// 租期15s
				for {
					lease, err = r.etcdClient.Grant(context.Background(), 15)
					if err != nil {
						log.Printf("[kitex-registry] failed to Register(grant lease), retrying: %s\n", err)
						continue
					}

					_, err = r.etcdClient.Put(context.Background(), fmt.Sprintf("pigeon-registry/%s/%d", info.ServiceName,
						lease.ID), putContent, clientv3.WithLease(lease.ID))
					if err != nil {
						log.Printf("[kitex-registry] failed to Registry(put key), retrying: %s\n", err)
					}
				}
			}

			select {
			case <-r.exitChan:
				// 让KeepAlive停止
				r.exitOKChan <- struct{}{}
				return
				// 3s续约一次
			case <-time.After(time.Second * 3):
			}
		}
	}()

	return nil

}

// 关闭相关的goroutine, 让etcd自动清理
// Deregister不是重点, 不必太过关心, 只需要保证最终key被删除即可
func (r *etcdRegistry) Deregister(info *registry.Info) error {
	r.exitChan <- struct{}{}
	<-r.exitOKChan
	return nil
}

func getLocalIPv4Host() (string, error) {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addr {
		ipNet, isIpNet := addr.(*net.IPNet)
		if isIpNet && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			if ipv4 != nil {
				return ipv4.String(), nil
			}
		}
	}
	return "", fmt.Errorf("not found ipv4 address")
}

type etcdResolver struct {
	etcdClient *clientv3.Client
}

func NewEtcdResolver(endpoints []string) (discovery.Resolver, error) {
	return NewEtcdResolverWithAuth(endpoints, "", "")
}

func NewEtcdResolverWithAuth(endpoints []string, username string, password string) (discovery.Resolver, error) {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
		Username:  username,
		Password:  password,
	})
	if err != nil {
		return nil, err
	}

	return &etcdResolver{
		etcdClient: etcdClient,
	}, nil
}

func (r *etcdResolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) (description string) {
	return target.ServiceName()
}

func (r *etcdResolver) Resolve(ctx context.Context, desc string) (discovery.Result, error) {
	result := discovery.Result{}

	resp, err := r.etcdClient.Get(ctx, fmt.Sprintf("kitex-registry/%s", desc), clientv3.WithPrefix())
	if err != nil {
		return result, err
	}

	for _, v := range resp.Kvs {
		info := storeInfo{}
		err := json.Unmarshal(v.Value, &info)
		if err != nil {
			log.Printf("fail to unmarshal with err: %v, ignore key: %v\n", err, v.Key)
			continue
		}

		w := info.Weight
		if w <= 0 {
			w = discovery.DefaultWeight
		}

		ins := discovery.NewInstance(info.Network, info.Addr, info.Weight, info.Tags)
		result.Instances = append(result.Instances, ins)
	}

	if len(result.Instances) == 0 {
		return result, errors.New("no instance remains for: " + desc)
	}

	result.CacheKey = desc
	result.Cacheable = true
	return result, nil
}

func (r *etcdResolver) Diff(cacheKey string, prev, next discovery.Result) (discovery.Change, bool) {
	return discovery.DefaultDiff(cacheKey, prev, next)
}

// Name implements the Resolver interface.
func (r *etcdResolver) Name() string {
	return "etcd"
}
