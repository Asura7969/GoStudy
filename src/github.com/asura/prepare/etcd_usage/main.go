package main

import (
	"context"
	"fmt"
	"time"
)

// go get go.etcd.io/etcd/clientv3
func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      = clientv3.KV
		putResp = *clientv3.PutResponse
	)

	// 客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.0:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// 用于读写etcd的键值对
	kv = clientv3.NewKV(client)

	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job1", "hello", clientv3.WithPrevKV); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Revision", putResp.Header.Revision)
		if putResp.PrevKV != nill {
			fmt.Println("PreValue", string(putResp.PrevKV.Value))
		}
	}
}
