package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"testing"
)

func TestKV(t *testing.T) {
	c, err := NewClient(context.Background(), []string{"http://127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}
	if resp, err := c.PutKV("/test/t3", "1234"); err != nil {
		panic(err)
	} else {
		StringPutResponse(resp)
	}
	if resp, err := c.PutKV("/test/t2", "3214"); err != nil {
		panic(err)
	} else {
		StringPutResponse(resp)
	}

	if resp, err := c.GetKV("/test/t1"); err != nil {
		panic(err)
	} else {
		StringGetResponse(resp)
	}

	if resp, err := c.GetKV("/test/", clientv3.WithPrefix()); err != nil {
		panic(err)
	} else {
		StringGetResponse(resp)
	}
}
