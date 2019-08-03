package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"time"
)

func main() {

	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
		kv     clientv3.KV
		putOp  clientv3.Op
		getOp  clientv3.Op
		opResp clientv3.OpResponse
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

	kv = clientv3.NewKV(client)

	// 创建Op:operation
	putOp = clientv3.OpPut("/cron/jobs/job8", "job8")
	// 执行Op
	if opResp, err = kv.Do(context.TODO(), putOp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("put Revision：", opResp.Put().Header.Revision)

	// 创建Op
	getOp = clientv3.OpGet("/cron/jobs/job8")
	// 执行Op
	if opResp, err = kv.Do(context.TODO(), getOp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("get Revision：", opResp.Get().Kvs[0].ModRevision)
	fmt.Println("get value：", string(opResp.Get().Kvs[0].Value))
}
