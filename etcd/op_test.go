package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"testing"
)

func TestOp(t *testing.T) {
	c, err := getClient()
	if err != nil {
		t.Fatal(err)
	}
	if resp, err := c.OpPut("/test/t1", "1"); err != nil {
		t.Fatal(err)
	} else {
		StringPutResponse(resp.Put())
	}
	if resp, err := c.OpPut("/test/t2", "2"); err != nil {
		t.Fatal(err)
	} else {
		StringPutResponse(resp.Put())
	}
	if resp, err := c.OpGet("/test", clientv3.WithPrefix()); err != nil {
		t.Fatal(err)
	} else {
		StringGetResponse(resp.Get())
	}

	if resp, err := c.OpDel("/test", clientv3.WithPrefix()); err != nil {
		t.Fatal(err)
	} else {
		StringDeleteResponse(resp.Del())
	}
}
