package zhenai

import (
	"gocrawler/concurrent_edition/engine"
	"gocrawler/concurrent_edition/fetcher"
	"gocrawler/concurrent_edition/model"
	"regexp"
)

var UserListReg = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
var NextPageReg = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
var UserIdReg = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseUserList(contents []byte, s string) (engine.ParseResult, error) {
	//log.Printf("Parsing URL:%s", s)
	matches := UserListReg.FindAllSubmatch(contents, -1)
	parseResult := engine.ParseResult{}
	user := engine.Item{}
	for _, m := range matches {
		url := string(m[1])
		id := UserIdReg.FindAllSubmatch([]byte(url), 1)
		user.Id = string(id[0][1])
		user.Url = url
		user.Payload = model.User{
			Name:  string(m[2]),
			Group: s,
		}
		parseResult.Items = append(parseResult.Items, user)
	}

	pageMatches := NextPageReg.FindAllSubmatch(contents, -1)
	for _, m := range pageMatches {
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url:     string(m[1]),
			Fetcher: engine.NewFuncFetcher(fetcher.Fetch, "Fetch"), // 用iphoneFetch来 fetch 用户详情地址
			Parser:  engine.NewFuncParser(ParseUserList, "ParseUserList"),
		})
	}

	return parseResult, nil
}
