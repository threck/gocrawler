package zhenai

import (
	"io/ioutil"
	"testing"
)

func TestParseUserInfo(t *testing.T) {
	file, err := ioutil.ReadFile("userinfo_test_data.html")
	if err != nil {
		panic(err)
	}
	result, _ := ParseUserInfo(file, "userinfo_test_data.html")

	testUserInfo := []string{"阿坝", "阿克苏", "阿拉善盟"}

	for i, user := range testUserInfo {
		if result.Items[i] != user {
			t.Errorf("city should be %s; but now is %s", user, result.Items[i])
		}
	}
}
