package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
	"time"
)

type EtcdMutex struct {
	Ttl     int64
	Conf    clientv3.Config
	Key     string
	cancel  context.CancelFunc
	lease   clientv3.Lease
	leaseId clientv3.LeaseID
	txn     clientv3.Txn
	client  *clientv3.Client
}

func (e *EtcdMutex) init() error {
	var ctx context.Context

	client, err := clientv3.New(e.Conf)
	if err != nil {
		return err
	}
	e.client = client

	//创建租约和事务
	e.lease = clientv3.NewLease(client)
	e.txn = clientv3.NewKV(client).Txn(context.TODO())

	grant, err := e.lease.Grant(context.TODO(), e.Ttl)
	if err != nil {
		return err
	}
	ctx, e.cancel = context.WithCancel(context.TODO())

	//生成租约id
	e.leaseId = grant.ID
	//只发一次心态，即到期不续，租约过期后，锁就会自动失效
	//要自动续期，使用KeepAlive
	_, err = e.lease.KeepAliveOnce(ctx, e.leaseId)
	return err
}

func (e *EtcdMutex) Lock() (ok bool, err error) {
	err = e.init()
	if err != nil {
		return false, err
	}

	//先比较再设置，使用事务保证原子性
	e.txn.If(clientv3.Compare(clientv3.CreateRevision(e.Key), "=", 0)).Then(clientv3.OpPut(e.Key, "", clientv3.WithLease(e.leaseId))).Else()
	res, err := e.txn.Commit()
	if err != nil {
		return false, err
	}

	if res.Succeeded {
		return true, nil
	}

	//未获取锁成功，监听key被删除后，再抢锁
	watch := e.client.Watch(context.TODO(), e.Key)
	for w := range watch {
		for _, ev := range w.Events {
			if ev.Type == clientv3.EventTypeDelete {
				return e.Lock()
			}
		}
	}

	return false, nil
}

func (e *EtcdMutex) Unlock() {
	e.cancel()
	//释放租约
	_, err := e.lease.Revoke(context.TODO(), e.leaseId)
	if err != nil {
		panic(err)
	}
}

func main() {
	var conf = clientv3.Config{
		Endpoints:   []string{"http://139.196.105.2:2379"},
		DialTimeout: 5 * time.Second,
	}
	mx1 := &EtcdMutex{
		Ttl:  10,
		Conf: conf,
		Key:  "test_01",
	}

	mx2 := &EtcdMutex{
		Ttl:  10,
		Conf: conf,
		Key:  "test_01",
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		ok, err := mx1.Lock()
		defer mx1.Unlock()
		defer wg.Done()
		if err != nil {
			fmt.Println("first:", err)
		}
		if !ok {
			fmt.Println("first get lock failed")
			return
		}

		fmt.Println("first get lock success")
		time.Sleep(1 * time.Second)
		fmt.Println("first keep lock")
		time.Sleep(3 * time.Second)
		fmt.Println("first release lock")
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ok, err := mx2.Lock()
		defer mx2.Unlock()
		defer wg.Done()
		if err != nil {
			fmt.Println("next:", err)
		}
		if !ok {
			fmt.Println("next get lock failed")
			return
		}
		fmt.Println("next get lock success")
	}()
	wg.Wait()

	fmt.Println("ending")
}
