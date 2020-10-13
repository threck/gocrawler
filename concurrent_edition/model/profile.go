package model

import "encoding/json"

type Profile struct {
	Id         []string
	Name       []string
	BasicInfo  []string
	DetailInfo []string
	ObjectInfo []string
}

type ProfileErshou struct {
	Name       string // 楼盘名 人和逸景
	City       string // 城市 成都
	PerPrice   string // 参考单价 10508元/m
	Area       string // 面积  90.4m
	TotalPrice string // 总价 95万
	HouseType  string // 户型 2室2厅1卫
	Floor      string // 楼层 20/22
	Comment    string // 备注，标题
	Street     string // 城区
	Addr       string // 地址 [金牛-抚琴]营门口路188号
	Developer  string // 开发商
	Tenement   string // 物业类型：住宅、别墅、商业
}

type User struct {
	Name  string
	Group string
}

func FromJsonObj(o interface{}) (User, error) {
	var user User
	bytes, err := json.Marshal(o)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}
