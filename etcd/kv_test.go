package etcd

import (
	"context"
	"github.com/scorpiotzh/toolib"
	"testing"
)

func TestKV(t *testing.T) {
	c, err := NewClient(context.Background(), []string{"http://127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}
	if resp, err := c.PutKV("test", "123"); err != nil {
		panic(err)
	} else {
		log.Info(toolib.JsonString(resp.PrevKv))
		if resp.PrevKv != nil {
			log.Info(string(resp.PrevKv.Key), string(resp.PrevKv.Value))
		}
	}

	if resp, err := c.GetKV("test"); err != nil {
		panic(err)
	} else {
		log.Info(toolib.JsonString(resp.Kvs))
		for _, v := range resp.Kvs {
			log.Info(string(v.Key), string(v.Value))
		}
	}
}
