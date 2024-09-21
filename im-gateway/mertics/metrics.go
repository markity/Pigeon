package metrics

import (
	"context"
	"encoding/json"
	"log"
	"sync/atomic"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type MetricsEtcdData struct {
	Name                 string
	Conns                uint64
	TCPAdvertiseAddrPort string
	RPCAdvertiseAddrPort string
}

// metrics, event loop每次接收/关闭一个长连接都会原子操作这个变量
var Conns atomic.Uint64

func GoLoopUpdateMetrics(etcdCli *clientv3.Client, cfg *MetricsEtcdData, interval time.Duration) {
	// 创建一个lease, 不断续约
	lease, err := etcdCli.Lease.Grant(context.Background(), 15)
	if err != nil {
		panic(err)
	}
	var leaseID atomic.Int64
	leaseID.Store(int64(lease.ID))
	// 不断续约
	go func() {
		for {
			time.Sleep(time.Second * 3)
			_, err := etcdCli.Lease.KeepAliveOnce(context.Background(), clientv3.LeaseID(leaseID.Load()))
			if err != nil {
				log.Printf("lease lost: %v\n", err)
				lease, err := etcdCli.Lease.Grant(context.Background(), 15)
				if err != nil {
					panic(err)
				}
				leaseID.Store(int64(lease.ID))
			}
		}
	}()
	for {
		time.Sleep(interval)
		cfg.Conns = Conns.Load()
		_stringData, err := json.Marshal(cfg)
		if err != nil {
			panic(err)
		}
		stringData := string(_stringData)
		log.Printf("上传metrics: %v\n", stringData)
		_, err = etcdCli.KV.Put(context.Background(), "im-gateway-metrics/"+cfg.Name, stringData,
			clientv3.WithLease(clientv3.LeaseID(leaseID.Load())))
		if err != nil {
			panic(err)
		}
	}
}
