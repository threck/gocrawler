package zhenai

import (
	"gocrawler/concurrent_edition/engine"
	"gocrawler/concurrent_edition/fetcher"
	"regexp"
)

const regexpCity = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, s string) (engine.ParseResult, error) {
	re := regexp.MustCompile(regexpCity)
	matches := re.FindAllSubmatch(contents, -1)
	parseResult := engine.ParseResult{}
	for _, m := range matches {
		//parseResult.Items = append(parseResult.Items, string(m[2]))
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url:     string(m[1]),
			Fetcher: engine.NewFuncFetcher(fetcher.Fetch, "Fetch"),
			Parser:  engine.NewFuncParser(ParseUserList, "ParseUserList"),
			//SubParseFunc: NilParser,
		})
	}

	return parseResult, nil
}
