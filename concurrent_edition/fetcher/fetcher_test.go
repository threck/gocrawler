package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	var url = "https://gl.lianjia.com/ershoufang/105105251178.html"
	bytes, err := Fetch(url)
	if err != nil {
		t.Errorf("fetch error : %s", err)
	} else {
		fmt.Printf("fetch success: %s", bytes)
	}
}
