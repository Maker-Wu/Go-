package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 定义一个map
	var a = make(map[string]interface{}, 10)
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	b, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)	//str:{"address":"洪崖洞","age":30,"name":"红孩儿"}

	// 切片序列化
	sli := make([]map[string]interface{}, 0, 10)
	m1 := make(map[string]interface{}, 5)
	m1["name"] = "jack"
	m1["age"] = 7
	m1["address"] = "北京"
	sli = append(sli, m1)

	m2 := make(map[string]interface{}, 5)
	m2["name"] = "tom"
	m2["age"] = 20
	m2["address"] = "东京"
	sli = append(sli, m2)

	b2, err := json.Marshal(sli)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b2)	//str:[{"address":"北京","age":7,"name":"jack"},{"address":"东京","age":20,"name":"tom"}]

	// map反序列化
	str1 := `{"address":"洪崖洞","age":30,"name":"红孩儿"}`
	var m3 map[string]interface{}
	err = json.Unmarshal([]byte(str1), &m3)		//要传指针
	fmt.Println(m3)		//map[address:洪崖洞 age:30 name:红孩儿]
}