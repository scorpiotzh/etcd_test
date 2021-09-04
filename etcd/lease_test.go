package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/scorpiotzh/toolib"
	"testing"
	"time"
)

// 租约
func TestLeaseGrant(t *testing.T) {
	c, err := getClient()
	if err != nil {
		t.Fatal(err)
	}
	lg, err := c.LeaseGrant(10)
	if err != nil {
		t.Fatal(err)
	} else {
		log.Info("lg:", lg.ID)
	}
	if resp, err := c.PutKV("test", "1234", clientv3.WithLease(lg.ID)); err != nil {
		t.Fatal(err)
	} else {
		StringPutResponse(resp)
	}
	//定时的看一下key过期了没有
	for {
		if resp, err := c.GetKV("test"); err != nil {
			t.Fatal(err)
		} else if resp.Count > 0 {
			StringGetResponse(resp)
		} else {
			log.Info("过期了")
			break
		}
		time.Sleep(time.Second * 2)
		if resp, err := c.Revoke(lg.ID); err != nil {
			t.Fatal(err)
		} else {
			log.Info(toolib.JsonString(resp))
		}
	}
}

// 自动续租
func TestKeepAlive(t *testing.T) {
	c, err := getClient()
	if err != nil {
		t.Fatal(err)
	}
	lg, err := c.LeaseGrant(10)
	if err != nil {
		t.Fatal(err)
	} else {
		log.Info("lg:", lg.ID)
	}
	// 续租
	ka, err := c.KeepAlive(lg.ID)
	if err != nil {
		t.Fatal(err)
	}
	c.StringKeepAlive(ka)

	//
	if resp, err := c.PutKV("test", "1234", clientv3.WithLease(lg.ID)); err != nil {
		t.Fatal(err)
	} else {
		StringPutResponse(resp)
	}
	//定时的看一下key过期了没有
	for {
		if resp, err := c.GetKV("test"); err != nil {
			t.Fatal(err)
		} else if resp.Count > 0 {
			StringGetResponse(resp)
		} else {
			log.Info("过期了")
			break
		}
		time.Sleep(time.Second * 2)
	}
}
