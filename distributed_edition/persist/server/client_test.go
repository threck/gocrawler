package main

import (
	"fmt"
	"gocrawler/concurrent_edition/engine"
	"gocrawler/concurrent_edition/model"
	"gocrawler/distributed_edition/config"
	"gocrawler/distributed_edition/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	item := engine.Item{
		Url: "http://test url",
		Id:  "123",
		Payload: model.User{
			Name:  "测试姓名",
			Group: "test group",
		},
	}

	// start ItemSaverServer
	go serveRpc(":1234", "test1")
	time.Sleep(time.Millisecond)

	// start ItemSaverClient
	client, err := rpcsupport.ClientRpc(fmt.Sprintf(":%d", config.PersistServerPort))
	if err != nil {
		panic(err)
	}
	// call save

	result := ""
	err = client.Call(config.PersistServiceRPC, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}

}
