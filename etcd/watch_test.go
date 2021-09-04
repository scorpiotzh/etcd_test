package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"testing"
	"time"
)

func TestWatch(t *testing.T) {
	c, err := getClient()
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		for {
			c.PutKV("test", time.Now().String())
			c.DelKV("test")
			time.Sleep(1 * time.Second)
		}
	}()
	resp, err := c.GetKV("test")
	if err != nil {
		t.Fatal(err)
	}
	// 当前etcd集群事务ID, 单调递增的
	watchStartRevision := resp.Header.Revision + 1
	log.Info("Revision:", resp.Header.Revision)
	ctx, cancelFunc := context.WithCancel(context.Background())
	time.AfterFunc(5*time.Second, func() {
		cancelFunc()
	})
	wc := c.Watch(ctx, "test", clientv3.WithRev(watchStartRevision))
	StringWatch(wc, nil, nil)
}
