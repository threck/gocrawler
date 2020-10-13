package lianjia

//type ProfileErshou struct {
//--Name       string // 楼盘名 人和逸景
//--City       string // 城市 成都
//--Street     string
//--Addr       string // 地址 [金牛-抚琴]营门口路188号
//--PerPrice   string // 参考单价 10508元/m
//	Area       string // 面积  90.4m
//	TotalPrice string // 总价 95万
//	HouseType  string // 户型 2室2厅1卫
//	Floor      string // 楼层 20/22
//--Comment    string // 备注，标题
//	Developer  string // 开发商
//	Tenement   string // 物业类型：住宅、别墅、商业
//}

import (
	"gocrawler/concurrent_edition/engine"
	"gocrawler/concurrent_edition/model"
	"regexp"

	"github.com/ericchiang/css"
)

var (
	nameCssSel       = css.MustCompile(".communityName>a")
	cityCssSel       = css.MustCompile(".container>div>a")
	streetCssSel     = css.MustCompile(".areaName>span>a")
	addrCssSel       = css.MustCompile(".supplement")
	perPriceCssSel   = css.MustCompile(".unitPriceValue")
	areaCssSel       = css.MustCompile(".area>.mainInfo")
	totalPriceCssSel = css.MustCompile(".total")
	houseTypeCssSel  = css.MustCompile(".room>.mainInfo")
	floorCssSel      = css.MustCompile(".room>.subInfo")
	commentCssSel    = css.MustCompile(".main")

	eleText       = regexp.MustCompile(`>([^<]+)<`)
	nameRe        = regexp.MustCompile("<a .*?class=\"info[^>]+>([^<]+)</a>")
	cityRe        = regexp.MustCompile("<a href=\"/ershoufang/\">([^<]+)</a>")
	streetReg     = eleText
	addrReg       = eleText
	perPriceReg   = eleText
	areaReg       = eleText
	totalPriceReg = eleText
	houseTypeReg  = eleText
	floorReg      = eleText
	commentRe     = regexp.MustCompile(`<h1 class="main" title="([^"]+)">[^<]+</h1>`)
)

func ParseProfileInfo(contents []byte, url string) (engine.ParseResult, error) {
	// get id
	var UserIdReg = regexp.MustCompile(`https://[^.]+.lianjia.com/ershoufang/([\d]+).html`)
	id := UserIdReg.FindAllStringSubmatch(url, 1)
	//fmt.Printf("%s\n", id)
	// get basic info
	name, err := getInfoString(contents, nameCssSel, nameRe)
	if err != nil {
		return engine.ParseResult{}, err
	}

	city, err := getInfoString(contents, cityCssSel, cityRe)
	if err != nil {
		return engine.ParseResult{}, err
	}

	street, err := getInfoString(contents, streetCssSel, streetReg)
	if err != nil {
		return engine.ParseResult{}, err
	}

	addr, err := getInfoString(contents, addrCssSel, addrReg)
	if err != nil {
		return engine.ParseResult{}, err
	}

	perPrice, err := getInfoString(contents, perPriceCssSel, perPriceReg)
	if err != nil {
		return engine.ParseResult{}, err
	}

	area, err := getInfoString(contents, areaCssSel, areaReg)
	if err != nil {
		return engine.ParseResult{}, err
	}

	totalPrice, err := getInfoString(contents, totalPriceCssSel, totalPriceReg)
	if err != nil {
		return engine.ParseResult{}, err
	}

	houseType, err := getInfoString(contents, houseTypeCssSel, houseTypeReg)
	if err != nil {
		return engine.ParseResult{}, err
	}

	floor, err := getInfoString(contents, floorCssSel, floorReg)
	if err != nil {
		return engine.ParseResult{}, err
	}

	comment, err := getInfoString(contents, commentCssSel, commentRe)
	if err != nil {
		return engine.ParseResult{}, err
	}

	payload := model.ProfileErshou{}
	payload.Name = name
	payload.City = city
	payload.Street = street
	payload.Addr = addr
	payload.PerPrice = perPrice
	payload.Area = area
	payload.TotalPrice = totalPrice
	payload.HouseType = houseType
	payload.Floor = floor
	payload.Comment = comment

	profile := engine.Item{
		Url:     url,
		Id:      id[0][1],
		Payload: payload,
	}

	parseResult := engine.ParseResult{}
	parseResult.Items = append(parseResult.Items, profile)

	return parseResult, nil
}
