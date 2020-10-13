package lianjia

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestRegExp(t *testing.T) {
	reg := regexp.MustCompile(`<h1 class="main" title="([^"]+)">[^<]+</h1>`)
	file, err := ioutil.ReadFile("profileinfo_test_data.html")
	if err != nil {
		panic(err)
	}

	matches := reg.FindAllSubmatch(file, -1)
	for _, m := range matches {
		fmt.Printf("%s\n", m)
	}

}
