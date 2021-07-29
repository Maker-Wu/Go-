package main

import (
	"encoding/json"
	"fmt"
)

// User 在tag中添加omitempty忽略空值
// 注意这里 hobby,omitempty 合起来是json tag值，中间用英文逗号分隔
type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email,omitempty"`
	Hobby []string `json:"hobby,omitempty"`
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
	fmt.Printf("str:%s\n", b)	//str:{"name":"七米"}
}

func main() {
	jsonMarsh1()
}

