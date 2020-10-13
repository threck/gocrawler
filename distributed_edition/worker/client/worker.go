package client

import (
	"gocrawler/concurrent_edition/engine"
	"gocrawler/distributed_edition/config"
	"gocrawler/distributed_edition/worker"
	"net/rpc"
)

//func Worker(re engine.Request) (engine.ParseResult, error) {
//
//	client, err := rpcsupport.ClientRpc(fmt.Sprintf(":%d", config.WorkerPort0))
//	if err != nil {
//		panic(err)
//	}
//
//	sRequest := worker.SerializeRequest(re)
//
//	sResult := worker.ParseResult{}
//	err = client.Call(config.CrawlServiceRPC, sRequest, &sResult)
//	if err != nil {
//		return engine.ParseResult{}, err
//	}
//
//	return worker.DeSerializeParseResult(sResult), nil
//}

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(re engine.Request) (engine.ParseResult, error) {
		sRequest := worker.SerializeRequest(re)

		sResult := worker.ParseResult{}
		client := <-clientChan
		err := client.Call(config.CrawlServiceRPC, sRequest, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeSerializeParseResult(sResult), nil
	}
}
