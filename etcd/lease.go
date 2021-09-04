package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
)

// LeaseGrant 租约
func (c *Client) LeaseGrant(ttl int64) (*clientv3.LeaseGrantResponse, error) {
	lease := clientv3.NewLease(c.client)
	return lease.Grant(c.ctx, ttl)
}

// KeepAlive 续租
func (c *Client) KeepAlive(ctx context.Context, leaseId clientv3.LeaseID) (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	lease := clientv3.NewLease(c.client)
	return lease.KeepAlive(ctx, leaseId)
}

func (c *Client) StringKeepAlive(ka <-chan *clientv3.LeaseKeepAliveResponse) {
	go func() {
		for {
			select {
			case resp := <-ka:
				if resp == nil {
					log.Info("租约失效")
					goto END
				} else { //每秒会续租一次，所以就会收到一次应答
					log.Info("自动续租：", resp.ID, resp.TTL)
				}
			case <-c.ctx.Done():
				goto END
			}
		}
	END:
	}()
}

// Revoke 回收租约
func (c *Client) Revoke(leaseId clientv3.LeaseID) (*clientv3.LeaseRevokeResponse, error) {
	lease := clientv3.NewLease(c.client)
	return lease.Revoke(c.ctx, leaseId)
}
