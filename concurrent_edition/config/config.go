package config

const (
	// URL Seeds
	UrlZhenai  = "http://www.zhenai.com/zhenghun"
	UrlSoufang = "https://cd.sofang.com/country.html"
	UrlLianjia = "https://www.lianjia.com/city/"

	// ElasticSearch
	ElasticSearchIndexZhenai     = "dating_zhenai"
	ElasticSearchIndexXcar       = "dating_xcar"
	ElasticSearchIndexSoufangOld = "dating_soufang"
	ElasticSearchIndexLianjia    = "dating_lianjia"

	// Template
	TemplateZhenai  = "18_crawler_distribution/concurrent_edition/frontend/view/template.html"
	TemplateLianjia = "18_crawler_distribution/concurrent_edition/frontend/view/template_lianjia.html"
)

var Url = UrlLianjia
var ElasticSearchIndex = ElasticSearchIndexLianjia
var ActiveTemplate = TemplateLianjia
