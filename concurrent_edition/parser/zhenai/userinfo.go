package zhenai

import (
	"bytes"
	"fmt"
	"gocrawler/concurrent_edition/engine"
	"gocrawler/concurrent_edition/model"
	"io"
	"regexp"
	"strings"

	"github.com/ericchiang/css"
	"golang.org/x/net/html"
)

var (
	nameCssSelector       = css.MustCompile(".nickName")
	basicInfoCssSelector  = css.MustCompile(".purple-btns")
	detailInfoCssSelector = css.MustCompile(".pink-btns")
	objectInfoCssSelector = css.MustCompile(".gray-btns")
	userInfoRe            = regexp.MustCompile(`<div [^>]*>([^<]+)</div>`)
	nameRe                = regexp.MustCompile(`<h1 [^>]*>([^<]+)</h1>`)
)

func cssSelect(contents []byte, sel *css.Selector) (string, error) {
	node, err := html.Parse(strings.NewReader(string(contents)))
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	w := io.Writer(&buf)
	for _, ele := range sel.Select(node) {
		html.Render(w, ele)
	}
	return buf.String(), nil
}

func extractString(contents []byte, sel *css.Selector, re *regexp.Regexp) ([]string, error) {
	sBasic, err := cssSelect(contents, sel)
	var s []string
	if err != nil {
		return s, err
	}

	matches := re.FindAllStringSubmatch(sBasic, -1)
	var Q []string
	for _, m := range matches {
		Q = append(Q, m[1])
	}
	return Q, nil

}
func ParseUserInfo(contents []byte, url string) (engine.ParseResult, error) {
	// get id
	var UserIdReg = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
	id := UserIdReg.FindAllSubmatch([]byte(url), 1)
	// get basic info
	fmt.Printf("%s", string(contents))
	name, err := extractString(contents, nameCssSelector, nameRe)
	if err != nil {
		return engine.ParseResult{}, err
	}
	//fmt.Println(name)

	basicInfo, err := extractString(contents, basicInfoCssSelector, userInfoRe)
	if err != nil {
		return engine.ParseResult{}, err
	}
	//fmt.Println(basicInfo)

	detailInfo, err := extractString(contents, detailInfoCssSelector, userInfoRe)
	if err != nil {
		return engine.ParseResult{}, err
	}
	//fmt.Println(detailInfo)

	objectInfo, err := extractString(contents, objectInfoCssSelector, userInfoRe)
	if err != nil {
		return engine.ParseResult{}, err
	}
	//fmt.Println(objectInfo)

	payload := model.Profile{}

	payload.Name = append(payload.Name, name...)
	payload.BasicInfo = append(payload.BasicInfo, basicInfo...)
	payload.DetailInfo = append(payload.DetailInfo, detailInfo...)
	payload.ObjectInfo = append(payload.ObjectInfo, objectInfo...)

	profile := engine.Item{
		Url:     url,
		Id:      string(id[0][1]),
		Payload: payload,
	}

	parseResult := engine.ParseResult{}
	parseResult.Items = append(parseResult.Items, profile)

	return parseResult, nil
}
