package lianjia

import (
	"gocrawler/concurrent_edition/engine"
	"gocrawler/concurrent_edition/fetcher"
	"regexp"
)

const regexpCityList = `<a href="(https://[^.]+.lianjia.com/)">[^<]+</a>`

func ParseCityList(contents []byte, url string) (engine.ParseResult, error) {
	re := regexp.MustCompile(regexpCityList)
	matches := re.FindAllSubmatch(contents, -1)
	parseResult := engine.ParseResult{}
	for _, m := range matches {
		//parseResult.Items = append(parseResult.Items, string(m[2]))
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url:     string(m[1]),
			Fetcher: engine.NewFuncFetcher(fetcher.Fetch, "Fetch"),
			Parser:  engine.NewFuncParser(ParseErshou, "ParseErshou"),
			//Parser: engine.NewFuncParser(NilParser, "NilParser"),
		})
	}
	//parseResult.Items = append(parseResult.Items, engine.Item{Url: "aaa"})

	return parseResult, nil
}
