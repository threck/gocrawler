package lianjia

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var (
	cssSel = totalPriceCssSel
	expReg = totalPriceReg
)

func TestCssSelect(t *testing.T) {
	file, err := ioutil.ReadFile("profileinfo_test_data.html")
	if err != nil {
		panic(err)
	}

	s, err := cssSelect(file, cssSel)
	fmt.Println(s)

}

func TestExtractString(t *testing.T) {
	file, err := ioutil.ReadFile("profileinfo_test_data.html")
	if err != nil {
		panic(err)
	}

	matches, err := extractString(file, cssSel, expReg)
	fmt.Println(matches)
	//	for _, m := range matches {
	//		fmt.Println(m)
	//	}
}

func TestGetInfoString(t *testing.T) {
	file, err := ioutil.ReadFile("profileinfo_test_data.html")
	if err != nil {
		panic(err)
	}

	s, err := getInfoString(file, cssSel, expReg)
	if err != nil {
		t.Errorf("error : %s", err)
	} else {
		fmt.Println(s)
	}
}
