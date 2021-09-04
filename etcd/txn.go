package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
)

// 分布式锁
func (c *Client) Txn(ctx context.Context, k, v string, leaseId clientv3.LeaseID) (*clientv3.TxnResponse, error) {
	kv := clientv3.NewKV(c.client)
	txn := kv.Txn(ctx)
	txn.If(clientv3.Compare(clientv3.CreateRevision(k), "=", 0)).
		Then(clientv3.OpPut(k, v, clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet(k))
	return txn.Commit()
}
