package main

import (
	"fmt"
	"learngo/18_crawler_distribution/final_edition/config"
	"learngo/18_crawler_distribution/final_edition/rpcsupport"
	"learngo/18_crawler_distribution/final_edition/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.ClientRpc(host)
	if err != nil {
		panic(err)
	}

	request := worker.Request{
		Url: "http://www.zhenai.com/zhenghun/ali",
		Fetcher: worker.SerializedFetcher{
			Name: config.NormalFetcher,
			Args: nil,
		},
		Parser: worker.SerializedParser{
			Name: config.ParseUserList,
			Args: nil,
		},
	}
	var result *worker.ParseResult
	err = client.Call(config.CrawlServiceRPC, request, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}

}
