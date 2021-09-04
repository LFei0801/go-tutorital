package main

import (
	"encoding/json"
	"fmt"
)

/**
使用JSON标记结构体中的数据
*/
type colorGroup struct {
	id int `json:"id"`
	name string `json:"name"`
	colors []string `json:"colors"`
}

func jsonBasis() {
	group := []colorGroup{
		{
			1,
			"Red",
			[]string{"Crimson", "Red", "Ruby", "Maroon"},
		},
		{
			id:     2,
			name:   "Green",
			colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		},
	}
	// 将group转为json数据
	data,_ := json.Marshal(group)
	fmt.Printf("%s\n",data)
	// 将json数据 转为 切片数据结构
	var decode []colorGroup
	_ = json.Unmarshal(data,&decode)
	fmt.Printf("%+v\n",decode)
}