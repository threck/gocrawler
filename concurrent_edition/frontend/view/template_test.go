package view

import (
	"gocrawler/concurrent_edition/engine"
	"gocrawler/concurrent_edition/frontend/model"
	common "gocrawler/concurrent_edition/model"
	"html/template"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	//tp := template.Must(template.ParseFiles("template.html"))
	tp, _ := template.ParseFiles("template.html")

	out, err := os.Create("template.test.html") // 建立一个文件
	page := model.SearchResult{}
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

	err = tp.Execute(out, page)
	if err != nil {
		panic(err)
	}

}
