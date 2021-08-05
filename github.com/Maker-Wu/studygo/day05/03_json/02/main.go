package main

import (
	"encoding/json"
	"fmt"
)

// Person 使用json tag指定序列化与反序列化时的行为
type Person struct {
	Name   string `json:"name"` // 指定json序列化/反序列化时使用小写name
	Age    int64
	Weight float64
}

func main() {
	p1 := Person{
		Name: "七米",
		Age: 18,
		Weight: 71.5,
	}
	// struct -> json string
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)			//str:{"name":"七米","Age":18,"Weight":71.5}

	// json string -> struct
	var p2 Person
	str := `{"name":"七米","Age":18,"Weight":71.5}`
	err = json.Unmarshal([]byte(str), &p2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)			//p2:main.go.Person{Name:"七米", Age:18, Weight:71.5}
}
