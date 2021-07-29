package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Hobby []string `json:"hobby"`
}

// 序列化不忽略空值
func jsonMarsh1()  {
	u1 := User{
		Name: "七米",
	}
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)	//str:{"name":"七米","email":"","hobby":null}
}

func main() {
	jsonMarsh1()
}
