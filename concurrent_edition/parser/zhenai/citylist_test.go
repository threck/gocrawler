package zhenai

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	file, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result, _ := ParseCityList(file, "citylist_test_data.html")

	const resultSize = 470
	testUrl := []string{"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng"}

	for i, url := range testUrl {
		if result.Requests[i].Url != url {
			t.Errorf("Url should be %s; but now is %s", url, result.Requests[i].Url)
		}
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}

}
