package persist

import (
	"context"
	"encoding/json"
	"gocrawler/concurrent_edition/engine"
	"gocrawler/concurrent_edition/model"
	"testing"

	"github.com/olivere/elastic/v7"
)

func TestSave(t *testing.T) {
	expect := engine.Item{
		Index: "dating_test",
		Url:   "http://test url",
		Id:    "123",
		Payload: model.User{
			Name:  "测试姓名",
			Group: "test group",
		},
	}
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	err = Save(client, expect.Index, expect)
	if err != nil {
		panic(err)
	}

	result, err := client.Get().Index(expect.Index).Id(expect.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("elastic search data: %s", result.Source)

	var actual engine.Item
	json.Unmarshal(result.Source, &actual) //  json 反序列化到 actual

	actualUser, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualUser

	t.Logf("expect: %v", expect)
	t.Logf("actual: %v", actual)
	if expect != actual {
		t.Errorf("got %v; expect %v", actual, expect)
	}

}
