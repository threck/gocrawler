package view

import (
	"gocrawler/concurrent_edition/engine"
	"gocrawler/concurrent_edition/frontend/model"
	common "gocrawler/concurrent_edition/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html") // 建立一个文件
	if err != nil {
		panic(err)
	}
	defer out.Close()
	page := model.SearchResult{}
	page.Hits = 123
	page.Start = 2
	item := engine.Item{
		Url:   "http://test url",
		Index: "dating_zhenai",
		Id:    "123",
		Payload: common.User{
			Name:  "测试姓名",
			Group: "test group",
		},
	}
	for i := 0; i < 20; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		t.Error(err)
	}
	// TODO: verify contents in template.test.html
}
