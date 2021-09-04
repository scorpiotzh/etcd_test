package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/scorpiotzh/mylog"
	"time"
)

var log = mylog.NewLogger("etcd", mylog.LevelDebug)

type Client struct {
	ctx    context.Context
	client *clientv3.Client
}

func NewClient(ctx context.Context, endpoints []string) (*Client, error) {
	var c Client
	config := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Second * 5,
	}
	if client, err := clientv3.New(config); err != nil {
		return nil, err
	} else {
		c.client = client
		c.ctx = ctx
		return &c, nil
	}
}
