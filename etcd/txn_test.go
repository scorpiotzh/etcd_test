package etcd

import (
	"context"
	"sync"
	"testing"
	"time"
)

// 分布式锁
func TestTxn(t *testing.T) {
	c, err := getClient()
	if err != nil {
		t.Fatal(err)
	}
	doTxn := func(name string) {
		// 租约
		lg, err := c.LeaseGrant(5)
		if err != nil {
			t.Fatal(err)
		}
		ctx, cancelFunc := context.WithCancel(context.Background())
		defer func() {
			log.Info("结束：", name)
			cancelFunc()
			if _, err := c.Revoke(lg.ID); err != nil {
				log.Error("Revoke err:", err.Error())
			}
		}()
		// 自动续租
		if ka, err := c.KeepAlive(ctx, lg.ID); err != nil {
			t.Fatal(err)
		} else {
			c.StringKeepAlive(ka)
		}
		// 事务
		if txResp, err := c.Txn(context.Background(), "test", "1", lg.ID); err != nil {
			t.Fatal(err)
		} else if txResp.Succeeded {
			log.Info("业务逻辑。。。。", name)
			time.Sleep(time.Second * 6)
		} else {
			log.Info("锁被抢占了", name)
		}
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		doTxn("事务A")
	}()
	go func() {
		defer wg.Done()
		doTxn("事务B")
	}()
	wg.Wait()
}
