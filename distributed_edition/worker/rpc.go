package worker

import (
	"gocrawler/concurrent_edition/engine"
)

type CrawlService struct{}

func (CrawlService) Process(req Request, result *ParseResult) error {
	engineRequest, err := DeSerializeRequest(req)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineRequest)
	if err != nil {
		return err
	}

	*result = SerializeParseResult(engineResult)

	return nil

}
