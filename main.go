package main

import (
	"flag"
	"learngo/18_crawler_distribution/concurrent_edition/engine"
	"learngo/18_crawler_distribution/concurrent_edition/fetcher"
	"learngo/18_crawler_distribution/concurrent_edition/parser/zhenai"
	"learngo/18_crawler_distribution/concurrent_edition/scheduler"
	"learngo/18_crawler_distribution/final_edition/config"
	persistClient "learngo/18_crawler_distribution/final_edition/persist/client"
	"learngo/18_crawler_distribution/final_edition/rpcsupport"
	workerClient "learngo/18_crawler_distribution/final_edition/worker/client"
	"log"
	"net/rpc"
	"strings"
)

//const url = "http://www.zhenai.com/zhenghun/chengdu"

// 增加 对 各个城市用户列表的分析
// 增加 对 iphoneFetch ，用户详情页被禁止爬取, 但是手机页面还没有被禁止

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemSaver, err := persistClient.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	client := createClientPool(strings.Split(*workerHosts, ","))

	processor := workerClient.CreateProcessor(client)

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemSaver,
		Processor:   processor,
	}
	//e := engine.SimpleEngine{}

	e.Run(engine.Request{
		Url:     config.Url,
		Fetcher: engine.NewFuncFetcher(fetcher.Fetch, config.NormalFetcher),
		Parser:  engine.NewFuncParser(zhenai.ParseCityList, config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {

	var clients []*rpc.Client
	for _, host := range hosts {
		client, err := rpcsupport.ClientRpc(host)
		if err != nil {
			log.Printf("connect to %s error : %v", host, err)
			continue
		} else {
			log.Printf("connect to %s success", host)
		}

		clients = append(clients, client)
	}

	clientChan := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				clientChan <- client
			}
		}
	}()
	return clientChan
}
