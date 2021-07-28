package main

import "fmt"

func main() {
	var a map[string]int
	fmt.Println(a == nil)		//true
	a = make(map[string]int, 10)
	fmt.Println(a == nil)		//false

	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 100
	scoreMap["小明"] = 200
	fmt.Println(scoreMap)							//map[小明:200 张三:100]
	fmt.Println(scoreMap["小明"])					//200
	fmt.Printf("type of a:%T\n", scoreMap)	//type of a:map[string]int
	fmt.Println(scoreMap["小陈"])					//0
	// map也支持在声明的时候填充元素
	userInfo := map[string]string{
		"username": "沙河小王子",
		"password": "123456",
	}
	fmt.Println(userInfo)							//map[password:123456 username:沙河小王子]

	// 判断键值是否存在
	value, ok := scoreMap["张三"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("查无此人")
	}

	value2, ok := scoreMap["伍胜强"]
	if ok {
		fmt.Println(value2)
	} else {
		fmt.Println(value2)		//0  map不存在指定key，则value为类型零值
		fmt.Println("查无此人")
	}

	scoreMap["娜扎"] = 60
	// 遍历map时的顺序与添加键值对的顺序无关
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	//删除键值对
	delete(scoreMap, "张三")		//将张三从map中删除
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	delete(scoreMap, "沙河")		//删除不存在的key，则不做任何操作，也不会报错

	// 元素为map类型的切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Println(value == nil)				//true
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	// 值为切片类型的map
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	v, isok := sliceMap[key]
	if !isok {
		v = make([]string, 0, 2)
	}
	v = append(v, "北京", "上海")
	sliceMap[key] = v
	fmt.Println(sliceMap)
}
