package lianjia

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseErshou(t *testing.T) {
	file, err := ioutil.ReadFile("ershou_test_data.html")
	if err != nil {
		panic(err)
	}
	result, _ := ParseErshou(file, "ershou_test_data.html")

	for _, v := range result.Requests {
		fmt.Printf("request is :%s\n", v)
	}

	const resultSize = 1
	testUrl := []string{"https://aq.lianjia.com/ershoufang"}

	for i, url := range testUrl {
		if result.Requests[i].Url != url {
			t.Errorf("Url should be %s; but now is %s", url, result.Requests[i].Url)
		}
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}

}
