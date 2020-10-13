package controller

import (
	"context"
	"gocrawler/concurrent_edition/config"
	"gocrawler/concurrent_edition/engine"
	"gocrawler/concurrent_edition/frontend/model"
	"gocrawler/concurrent_edition/frontend/view"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"fmt"

	"github.com/olivere/elastic/v7"
)

// TODO
// fill in query string
// fill in query string -> kind version
// support search button
// support paging
// add start page

type SearchResultHandler struct {
	View   view.SearchResultView
	Client *elastic.Client
}

// localhost:8888/search?q=男 已购房&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	//fmt.Fprintf(w, "q=%s, from=%d\n", q, from)
	//fmt.Printf("q=%s, from=%d\n", q, from)
	var pageData model.SearchResult
	pageData, err = h.GetSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.View.Render(w, pageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

const pagesize = 10

func (h SearchResultHandler) GetSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := h.Client.Search(config.ElasticSearchIndex).
		Query(elastic.NewQueryStringQuery(
			rewriteString(q))).
		From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Query = q
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))

	// 如果 from - pagesize 等于 -10, 则不显示上一页
	// 如果 0 > from - pagesize > -10, 说明还有不满一页未显示完, 上一页应该统一设置为从 from 0 开始
	newFrom := from - pagesize
	if -10 < newFrom && newFrom < 0 {
		newFrom = 0
	}
	result.PrevFrom = newFrom
	// 如果from + pagesize 大于 查询总数, 则不显示下一页
	result.NextFrom = from + pagesize

	//for _, profile := range resp.Each(reflect.TypeOf(engine.Item{})) {
	//	item := profile.(engine.Item)
	//
	//	actualUser, _ := common.FromJsonObj(item.Payload)
	//	item.Payload = actualUser
	//
	//	result.Items = append(result.Items, item)
	//}

	fmt.Printf("%s\n", result)
	return result, nil
}

func CreatSearchResultHandler(filename string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		View:   view.CreateSearchResultView(filename),
		Client: client,
	}
}

func rewriteString(q string) string {
	re := regexp.MustCompile(`([A-z][a-z]*):`)
	return re.ReplaceAllString(q, "PayLoad.$1:")
}
