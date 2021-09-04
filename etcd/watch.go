package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

// Watch 监听
func (c *Client) Watch(ctx context.Context, k string, opts ...clientv3.OpOption) clientv3.WatchChan {
	w := clientv3.NewWatcher(c.client)
	return w.Watch(ctx, k, opts...)
}

func StringWatch(resp clientv3.WatchChan, funcPut, funcDel func()) {
	for r := range resp {
		for _, e := range r.Events {
			switch e.Type {
			case mvccpb.PUT:
				log.Info("put:", string(e.Kv.Key), string(e.Kv.Value), e.Kv.CreateRevision, e.Kv.ModRevision)
				if funcPut != nil {
					funcPut()
				}
			case mvccpb.DELETE:
				log.Info("del:", string(e.Kv.Key), string(e.Kv.Value), e.Kv.CreateRevision, e.Kv.ModRevision)
				if funcDel != nil {
					funcDel()
				}
			}
		}
	}
}
