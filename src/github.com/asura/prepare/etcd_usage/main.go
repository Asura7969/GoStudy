package main

import (
	"context"
	"fmt"
	//"github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3"
	"time"
)

// https://www.cnblogs.com/xuejiale/p/10665845.html
// go get go.etcd.io/etcd/clientv3
func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
		kv     clientv3.KV
		//putResp *clientv3.PutResponse
		getResp *clientv3.GetResponse
	)

	// 客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"47.99.61.133:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	// 用于读写etcd的键值对
	kv = clientv3.NewKV(client)

	//if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job1", "bye", clientv3.WithPrevKV()); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("Revision", putResp.Header.Revision)
	//	if putResp.PrevKv != nil {
	//		fmt.Println("PreValue", string(putResp.PrevKv.Value))
	//	}
	//}

	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job1", clientv3.WithCountOnly()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getResp.Kvs, getResp.Count)
	}
}
