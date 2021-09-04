package etcd

import "github.com/coreos/etcd/clientv3"

func (c *Client) OpPut(k, v string, opts ...clientv3.OpOption) (clientv3.OpResponse, error) {
	kv := clientv3.NewKV(c.client)
	putOp := clientv3.OpPut(k, v, opts...)
	return kv.Do(c.ctx, putOp)
}

func (c *Client) OpGet(k string, opts ...clientv3.OpOption) (clientv3.OpResponse, error) {
	kv := clientv3.NewKV(c.client)
	putOp := clientv3.OpGet(k, opts...)
	return kv.Do(c.ctx, putOp)
}

func (c *Client) OpDel(k string, opts ...clientv3.OpOption) (clientv3.OpResponse, error) {
	kv := clientv3.NewKV(c.client)
	putOp := clientv3.OpDelete(k, opts...)
	return kv.Do(c.ctx, putOp)
}
