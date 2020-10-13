package main

import (
	"flag"
	"fmt"
	"learngo/18_crawler_distribution/final_edition/rpcsupport"
	"learngo/18_crawler_distribution/final_edition/worker"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
