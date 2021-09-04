package etcd

import (
	"github.com/coreos/etcd/clientv3"
)

func (c *Client) PutKV(k, v string) (*clientv3.PutResponse, error) {
	kv := clientv3.NewKV(c.client)
	return kv.Put(c.ctx, k, v, clientv3.WithPrevKV())
}

func (c *Client) GetKV(k string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	kv := clientv3.NewKV(c.client)
	return kv.Get(c.ctx, k, opts...)
}

func StringPutResponse(resp *clientv3.PutResponse) {
	if resp.PrevKv != nil {
		log.Info("put prev: ", string(resp.PrevKv.Key), string(resp.PrevKv.Value))
	} else {
		log.Info("put prev: nil")
	}
}

func StringGetResponse(resp *clientv3.GetResponse) {
	log.Info("get len:", len(resp.Kvs))
	for _, v := range resp.Kvs {
		log.Info(string(v.Key), string(v.Value))
	}
}
