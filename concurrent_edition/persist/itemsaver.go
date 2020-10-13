package persist

import (
	"context"

	"gocrawler/concurrent_edition/engine"
	"log"

	"github.com/olivere/elastic/v7"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		// Will use port 9200 default if have not this option
		//elastic.SetURL("http://localhost:9200"),
		// Must turn off sniff in docker. sniff是go客户端维护集群状态的的，但是集群在docker image的内网里, go客户端所在的服务器的外网看不到
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++

			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item #%d: %v %v", itemCount, item, err)
			} else {
				log.Printf("Item Saver: item #%d saved: %v", itemCount, item)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) (err error) {
	item.Index = index //给 item 加上index
	_, err = client.Index().
		Index(item.Index).
		//Type("zhenai"). // elastic search 7 里 已经弃用type; (deprecated)
		Id(item.Id).
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return err
	}
	//fmt.Println(resp)
	//fmt.Printf("%+v\n", resp) // %+v 打印结构体的时候会把字段名打出来
	return nil

}
