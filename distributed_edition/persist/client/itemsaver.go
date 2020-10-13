package client

import (
	"gocrawler/concurrent_edition/engine"
	"gocrawler/distributed_edition/config"
	"gocrawler/distributed_edition/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	// start ItemSaverClient
	client, err := rpcsupport.ClientRpc(host)
	if err != nil {
		panic(err)
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++

			// call save
			result := ""
			err = client.Call(config.PersistServiceRPC, item, &result)
			if err != nil || result != "ok" {
				log.Printf("Item Saver: error saving item #%d: %v %v", itemCount, item, err)
			} else {

				log.Printf("Item Saver: item #%d saved: %v", itemCount, item)
			}
		}
	}()
	return out, nil
}
