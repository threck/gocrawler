package main

import (
	"gocrawler/concurrent_edition/config"
	"gocrawler/concurrent_edition/engine"
	"gocrawler/concurrent_edition/fetcher"
	"gocrawler/concurrent_edition/parser/lianjia"
	"gocrawler/concurrent_edition/persist"
	"gocrawler/concurrent_edition/scheduler"
)

//const url = "http://www.zhenai.com/zhenghun/chengdu"

// 增加 对 各个城市用户列表的分析
// 增加 对 iphoneFetch ，用户详情页被禁止爬取, 但是手机页面还没有被禁止
func main() {
	itemSaver, err := persist.ItemSaver(config.ElasticSearchIndex)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemSaver,
		Processor:   engine.Worker,
	}
	//e := engine.SimpleEngine{}

	e.Run(engine.Request{
		Url:     config.Url,
		Fetcher: engine.NewFuncFetcher(fetcher.Fetch, "Fetch"),
		Parser:  engine.NewFuncParser(lianjia.ParseCityList, "ParseCityList"),
	})
}
