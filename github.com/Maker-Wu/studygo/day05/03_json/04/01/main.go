package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string		`json:"name"`
	Email string	`json:"email,omitempty"`
	Hobby []string	`json:"hobby,omitempty"`
	Profile			`json:"profile"`
}

type Profile struct {
	Website string	`json:"website"`
	Slogan string	`json:"slogan"`
}

// 嵌套结构体序列化后的json串为单层的：
func main() {
	u1 := User{
		Name: "七米",
		Hobby: []string{"足球", "双色球"},
	}
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)	//str:{"name":"七米","hobby":["足球","双色球"],"website":"","slogan":""}
}
