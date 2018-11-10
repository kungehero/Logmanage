package models

import (
	"beego-master/logs"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

var Client *clientv3.Client

func EtcdNew(endpoints []string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	Client = cli
	fmt.Println("connect succ")
	//defer cli.Close()

}

var EtcdValue []CollectValue

func EtcdGet(key string) error {
	EtcdValue = nil
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	resp, err := Client.Get(ctx, key, clientv3.WithPrefix())
	cancel()
	if err != nil {
		return err
	}
	conf := CollectConf{}
	for _, ev := range resp.Kvs {
		json.Unmarshal(ev.Value, &conf)
		EtcdValue = append(EtcdValue, CollectValue{string(ev.Key), CollectConf{conf.LogPath, conf.Topic}})
	}
	return nil
}

type CollectValue struct {
	Key string `json:"key"`
	CollectConf
}

type CollectConf struct {
	LogPath string `json:"path"`
	Topic   string `json:"topic"`
}

func Watch(key string) {
	for {
		var collectConf []CollectConf
		rch := Client.Watch(context.Background(), key)
		var getConfSucc = true

		for wresp := range rch {
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					logs.Warn("key[%s] 's config deleted", key)
					continue
				}

				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err := json.Unmarshal(ev.Kv.Value, collectConf)
					if err != nil {
						logs.Error("key [%s], Unmarshal[%s], err:%v ", err)
						getConfSucc = false
						continue
					}
				}
				logs.Debug("get config from etcd, %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
			if getConfSucc {
				logs.Debug("get config from etcd succ, %v", collectConf)
			}
		}
	}
}
