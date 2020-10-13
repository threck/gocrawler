package fetcher

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

var tm = time.Duration(RandInt(500, 1500)) * time.Millisecond
var rateLimiter = time.Tick(tm)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	//log.Println("Fetching URL:", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// bytes, err := httputil.DumpResponse(resp, true) // dumpResponse 里面包含了头部信息，这里不需要头部的信息
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	// 直接获取页面内容
	return ioutil.ReadAll(resp.Body)
}
