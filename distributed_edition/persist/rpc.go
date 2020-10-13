package persist

import (
	"gocrawler/concurrent_edition/engine"
	"gocrawler/concurrent_edition/persist"
	"log"

	"github.com/olivere/elastic/v7"
)

// 定义 ITEMSaver 的 service
// 将service包装成 serve client 结构
type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (r *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(r.Client, r.Index, item)
	if err == nil {
		log.Printf("item saved ok: %s\n", item)
		*result = "ok"
	} else {
		log.Printf("item %s saved failed.\n-> error: %s\n", item, err)
	}
	return err
}
