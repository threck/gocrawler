package fetcher

import (
	"fmt"
	"testing"
)

func TestIphoneFetch(t *testing.T) {
	var url = "https://album.zhenai.com/u/1322421714"
	bytes, err := IphoneFetch(url)
	if err != nil {
		t.Errorf("fetch error : %s", err)
	} else {
		fmt.Printf("fetch success: %s", bytes)
	}
}
