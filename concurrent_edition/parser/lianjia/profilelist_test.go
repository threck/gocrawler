package lianjia

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseProfileList(t *testing.T) {
	file, err := ioutil.ReadFile("profilelist_test_data.html")
	if err != nil {
		panic(err)
	}
	result, _ := ParseProfileList(file, "https://aq.lianjia.com/ershoufang")

	fmt.Println(result)
	for _, v := range result.Requests {
		fmt.Printf("request is :%s\n", v)
	}

	const resultSize = 31
	testUrl := []string{"https://aq.lianjia.com/ershoufang/103110637096.html",
		"https://aq.lianjia.com/ershoufang/103111131721.html",
		"https://aq.lianjia.com/ershoufang/103110151907.html",
		"https://aq.lianjia.com/ershoufang/103111025154.html",
	}

	for i, url := range testUrl {
		if result.Requests[i].Url != url {
			t.Errorf("Url should be %s ; but now is %s", url, result.Requests[i].Url)
		}
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests ; but had %d", resultSize, len(result.Requests))
	}

}
