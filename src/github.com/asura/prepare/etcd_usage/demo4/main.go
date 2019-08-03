package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"time"
)

func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		putResp *clientv3.PutResponse
		getResp *clientv3.GetResponse
		delResp *clientv3.DeleteResponse
		kvpair  *mvccpb.KeyValue
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

	// set
	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job3", "{....}", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Revision", putResp.Header.Revision)
		if putResp.PrevKv != nil {
			fmt.Println("PreValue", string(putResp.PrevKv.Value))
		}
	}

	// get
	// 读取/cron/jobs/为前缀的所有key
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
	} else {
		// 获取所有
		fmt.Println(getResp.Kvs)
	}

	// delete
	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/job3", clientv3.WithPrevKV()); err != nil {
		// 删除所有
		//if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/",clientv3.WithPrefix()); err != nil{
		fmt.Println(err)
		return
	}
	// 被删之前的kv是什么
	if len(delResp.PrevKvs) != 0 {
		for _, kvpair = range delResp.PrevKvs {
			fmt.Println("删除了:", string(kvpair.Key), string(kvpair.Value))
		}
	}

	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
	} else {
		// 获取所有
		fmt.Println(getResp.Kvs)
	}

}
