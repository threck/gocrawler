package lianjia

import (
	"gocrawler/concurrent_edition/engine"
	"gocrawler/concurrent_edition/model"
	"io/ioutil"
	"testing"
)

func TestParseProfileInfo(t *testing.T) {
	file, err := ioutil.ReadFile("profileinfo_test_data_forbidden.html")
	if err != nil {
		panic(err)
	}
	result, _ := ParseProfileInfo(file, "https://cq.lianjia.com/ershoufang/106105989710.html")

	//testUserInfo := []string{"阿坝", "阿克苏", "阿拉善盟"}
	var profileInfos []engine.Item
	profileInfos = append(profileInfos, engine.Item{
		Id:  "106105989710",
		Url: "https://cq.lianjia.com/ershoufang/106105989710.html",
		Payload: model.ProfileErshou{
			Name:       "首地江山赋",
			City:       "重庆二手房",
			PerPrice:   "11534 元/平米",
			Area:       "92.6平米",
			TotalPrice: "106.8",
			HouseType:  "3室2厅",
			Floor:      "高楼层/共33层",
			Comment:    "首地江山赋正三房业主诚心出售+全款证在手方便交易",
			Street:     "渝北 悦来",
			Addr:       "近10号线悦来站",
		},
	})

	for i, profile := range profileInfos {
		if result.Items[i] != profile {
			t.Errorf("city should be %s; \nbut now is %s\n", profile, result.Items[i])
		}
	}
}
