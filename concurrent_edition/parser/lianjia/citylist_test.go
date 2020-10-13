package lianjia

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	file, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result, _ := ParseCityList(file, "citylist_test_data.html")

	for _, v := range result.Requests {
		fmt.Printf("request is :%s\n", v)
	}

	const resultSize = 137
	testUrl := []string{"https://bj.lianjia.com/",
		"https://sh.lianjia.com/",
		"https://sz.lianjia.com/",
		"https://aq.lianjia.com/",
		"https://cz.fang.lianjia.com/",
		"https://hf.lianjia.com/"}

	for i, url := range testUrl {
		if result.Requests[i].Url != url {
			t.Errorf("Url should be %s; but now is %s", url, result.Requests[i].Url)
		}
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}

}
