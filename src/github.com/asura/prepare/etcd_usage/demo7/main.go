package main

import (
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"
	"time"
)

func main() {
	var (
		config            clientv3.Config
		client            *clientv3.Client
		err               error
		kv                clientv3.KV
		getResp           *clientv3.GetResponse
		watchStartRevison int64
		watchRespChan     <-chan clientv3.WatchResponse
		watchResp         clientv3.WatchResponse
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

	// 模拟etcd中kv的变化
	go func() {
		for true {
			kv.Put(context.TODO(), "/cron/jobs/job7", "i am job7")

			kv.Delete(context.TODO(), "/cron/jobs/job7")

			time.Sleep(1 * time.Second)
		}
	}()

	// 先 get 当前的值,并监听后续变化
	getResp, _ = kv.Get(context.TODO(), "/cron/jobs/job7")

	if len(getResp.Kvs) != 0 {
		// k 存在
		fmt.Println("当前值：", string(getResp.Kvs[0].Value))
	}

	// 当前 etcd 集群事务ID,单调递增
	watchStartRevison = getResp.Header.Revision + 1

	// 创建一个watcher
	watcher := clientv3.NewWatcher(client)

	// 启动监听
	fmt.Println("从改版本向后监听:", watchStartRevison)

	ctx, cancelFunc := context.WithCancel(context.TODO())
	time.AfterFunc(5*time.Second, func() {
		cancelFunc()
	})

	watchRespChan = watcher.Watch(ctx, "/cron/jobs/job7", clientv3.WithRev(watchStartRevison))

	// 处理kv变化事件
	for watchResp = range watchRespChan {
		for _, event := range watchResp.Events {
			switch event.Type.String() {
			case mvccpb.PUT.String():
				fmt.Println("修改为:", string(event.Kv.Value), "Revision:", event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE.String():
				fmt.Println("删除了 Revision:", event.Kv.ModRevision)
			}
		}
	}

}
