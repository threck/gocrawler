package main

import (
	"flag"
	"fmt"
	"gocrawler/distributed_edition/config"
	"gocrawler/distributed_edition/persist"
	"gocrawler/distributed_edition/rpcsupport"

	"github.com/olivere/elastic/v7"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	//err := serveRpc(fmt.Sprintf(":%d", config.PersistServerPort), config.ElasticSearchIndex)
	err := serveRpc(fmt.Sprintf(":%d", *port), config.ElasticSearchIndex)
	if err != nil {
		panic(err)
	}
	//log.Fatal(serveRpc(":1234", "test1"))   // 或者 这样写
}

func serveRpc(host string, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
