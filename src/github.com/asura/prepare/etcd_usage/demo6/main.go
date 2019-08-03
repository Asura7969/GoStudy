package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"time"
)

func main() {
	var (
		config         clientv3.Config
		client         *clientv3.Client
		err            error
		kv             clientv3.KV
		lease          clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId        clientv3.LeaseID
		putResp        *clientv3.PutResponse
		getResp        *clientv3.GetResponse
		keepResp       *clientv3.LeaseKeepAliveResponse
		keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
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

	// 申请一个租约
	lease = clientv3.NewLease(client)

	// 申请一个10s的租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}

	// put 一个kv，与租约关联
	leaseId = leaseGrantResp.ID
	//ctx,_ := context.WithTimeout(context.TODO(),5 * time.Second)

	if keepRespChan, err = lease.KeepAlive(context.TODO(), leaseId); err != nil {
		fmt.Println(err)
		return
	}

	// 处理续租应答的协程
	go func() {
		for true {
			select {
			case keepResp = <-keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约已经失效")
					goto END
				} else { // 每秒会续租一次,所以就会收到一次应答
					fmt.Println("收到续租应答", keepResp.ID)
				}

			}
		}
	END:
	}()

	// 用于读写etcd的键值对
	kv = clientv3.NewKV(client)

	if putResp, err = kv.Put(context.TODO(), "/cron/lock/job1", "lock", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success", putResp.Header.Revision)

	// 定期的看一下key过期没
	for true {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		}

		if getResp.Count == 0 {
			fmt.Println("kv过期了")
			break
		}
		fmt.Println("还没过期", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}

}
