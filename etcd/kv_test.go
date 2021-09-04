package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/scorpiotzh/toolib"
	"testing"
)

func getClient() (*Client, error) {
	return NewClient(context.Background(), []string{"http://127.0.0.1:2379"})
}

func TestPutKV(t *testing.T) {
	c, err := getClient()
	if err != nil {
		t.Fatal(err)
	}
	if resp, err := c.PutKV("/test/t1", "1234"); err != nil {
		t.Fatal(err)
	} else {
		StringPutResponse(resp)
	}
	if resp, err := c.PutKV("/test/t2", "3214"); err != nil {
		t.Fatal(err)
	} else {
		StringPutResponse(resp)
	}
	if resp, err := c.PutKV("/test/t3", "32145"); err != nil {
		t.Fatal(err)
	} else {
		StringPutResponse(resp)
	}
}

func TestGetKV(t *testing.T) {
	c, err := getClient()
	if err != nil {
		t.Fatal(err)
	}
	if resp, err := c.GetKV("/test/t1"); err != nil {
		t.Fatal(err)
	} else {
		StringGetResponse(resp)
	}

	if resp, err := c.GetKV("/test/", clientv3.WithPrefix()); err != nil {
		t.Fatal(err)
	} else {
		StringGetResponse(resp)
	}
}

func TestDelKV(t *testing.T) {
	c, err := getClient()
	if err != nil {
		panic(err)
	}
	if resp, err := c.DelKV("test"); err != nil {
		t.Fatal(err)
	} else {
		log.Info("resp:", toolib.JsonString(resp))
		StringDeleteResponse(resp)
	}
	if resp, err := c.DelKV("/test/t2", clientv3.WithPrefix()); err != nil {
		t.Fatal(err)
	} else {
		log.Info("resp:", toolib.JsonString(resp))
		StringDeleteResponse(resp)
	}
}
