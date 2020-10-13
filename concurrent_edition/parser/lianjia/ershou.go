package lianjia

import (
	"gocrawler/concurrent_edition/engine"
	"gocrawler/concurrent_edition/fetcher"
	"regexp"
)

//const regexpErshou = `<a  class="" href="(https://[^.]+(.fang){0,1}.lianjia.com/ershoufang)">([^<]+)</a>`
const regexpErshou = `https://[^.]+.lianjia.com/ershoufang`

func ParseErshou(contents []byte, url string) (engine.ParseResult, error) {
	re := regexp.MustCompile(regexpErshou)
	matches := re.FindAllSubmatch(contents, 1)
	parseResult := engine.ParseResult{}
	for _, m := range matches {
		//parseResult.Items = append(parseResult.Items, string(m[2]))
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url:     string(m[0]),
			Fetcher: engine.NewFuncFetcher(fetcher.Fetch, "Fetch"),
			Parser:  engine.NewFuncParser(ParseProfileList, "ParseProfileList"),
			//Parser: engine.NewFuncParser(NilParser, "NilParser"),
		})
	}

	return parseResult, nil
}
