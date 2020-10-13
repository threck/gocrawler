package main

import (
	"gocrawler/concurrent_edition/config"
	"gocrawler/concurrent_edition/frontend/controller"
	"net/http"
)

func main() {
	//http.Handle("/", http.FileServer(
	//	http.Dir("crawler/frontend/view")))
	//http.Handle("/search",
	//	code.CreateSearchResultHandler(
	//		"crawler/frontend/view/template.html"))
	http.Handle("/", http.FileServer(
		http.Dir("18_crawler_distribution/concurrent_edition/frontend/view")))
	http.Handle("/search", controller.CreatSearchResultHandler(config.ActiveTemplate))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
