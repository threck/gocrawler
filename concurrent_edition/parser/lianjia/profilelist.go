package lianjia

import (
	"gocrawler/concurrent_edition/engine"
	"gocrawler/concurrent_edition/fetcher"
	"regexp"
	"strconv"

	"github.com/ericchiang/css"
)

var (
	profileTitileCssSelector = css.MustCompile(".title")
	ProfileListReg           = regexp.MustCompile(`<a class="title" href="(https://[a-z0-9]+.lianjia.com/ershoufang/[0-9]+.html)".*?data-housecode="([0-9]+)"[^>]+>([^<]+)</a>`)

	totalPageReg = regexp.MustCompile(`"totalPage":([\d]+),`)
	curPageReg   = regexp.MustCompile(`"curPage":([\d]+)`)
)

func ParseProfileList(contents []byte, url string) (engine.ParseResult, error) {
	profiles, err := extractString(contents, profileTitileCssSelector, ProfileListReg)
	if err != nil {
		return engine.ParseResult{}, err
	}

	parseResult := engine.ParseResult{}
	for _, m := range profiles {
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url:     m[1],
			Fetcher: engine.NewFuncFetcher(fetcher.Fetch, "Fetch"),
			Parser:  engine.NewFuncParser(ParseProfileInfo, "ParseProfileInfo"),
		})
	}

	totalPa := totalPageReg.FindAllSubmatch(contents, -1)
	curlPa := curPageReg.FindAllSubmatch(contents, -1)
	total, err := strconv.Atoi(string(totalPa[0][1]))
	if err != nil {
		return parseResult, nil
	}
	cur, err := strconv.Atoi(string(curlPa[0][1]))
	if err != nil {
		return parseResult, nil
	}
	expReg := regexp.MustCompile(`/pg.*`)
	url = expReg.ReplaceAllString(url, "")
	if cur < total {
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url:     url + "/pg" + strconv.Itoa(cur+1),
			Fetcher: engine.NewFuncFetcher(fetcher.Fetch, "Fetch"),
			Parser:  engine.NewFuncParser(ParseProfileList, "ParseUserList"),
		})
	}

	return parseResult, nil
}
