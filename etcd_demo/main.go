package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	fmt.Println("test")
	config := clientv3.Config{
		Endpoints:   []string{"10.0.8.12:2379"},
		DialTimeout: 5 * time.Second,
	}
	cli, err := clientv3.New(config)
	if err != nil {
		log.Println("connect to etcd err", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.KV.Get(ctx, "name")
	// cli.KV.Get()
	cancel()
	if err != nil {
		log.Println("get from etcd err", err)
	}
	for _, kv := range resp.Kvs {
		log.Println(kv)
	}
}
