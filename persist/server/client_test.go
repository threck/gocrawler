package main

import (
	"fmt"
	"learngo/18_crawler_distribution/concurrent_edition/engine"
	"learngo/18_crawler_distribution/concurrent_edition/model"
	"learngo/18_crawler_distribution/final_edition/config"
	"learngo/18_crawler_distribution/final_edition/rpcsupport"
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
