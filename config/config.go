package config

const (
	// start page
	UrlZhenai  = "http://www.zhenai.com/zhenghun"
	UrlSoufang = "https://cd.sofang.com/country.html"
	UrlLianjia = "https://www.lianjia.com/city/"

	// ElasticSearch
	ElasticSearchIndexZhenai  = "dating_zhenai"
	ElasticSearchIndexXcar    = "dating_xcar"
	ElasticSearchIndexLianjia = "dating_lianjia"

	// PORT
	PersistServerPort = 1234
	WorkerPort0       = 9000

	// RPC Endpoints
	PersistServiceRPC = "ItemSaverService.Save"
	CrawlServiceRPC   = "CrawlService.Process"

	// Parser names
	ParseCityList = "ParseCityList"
	ParseUserList = "ParseUserList"
	ParseUserInfo = "ParseUserInfo"
	NilParser     = "NilParser"

	// Fetcher names
	NormalFetcher = "Fetch"
	IphoneFetch   = "IphoneFetch"
)

// activity elastic Index
var Url = UrlZhenai
var ElasticSearchIndex = ElasticSearchIndexZhenai
