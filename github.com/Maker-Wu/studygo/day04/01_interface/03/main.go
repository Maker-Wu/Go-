package main

import "fmt"

func main() {
	var x interface{}
	x = "hello"
	fmt.Printf("type:%T value:%v\n", x, x)	//type:string value:hello
	ret, ok := x.(string)
	if !ok {
		fmt.Println("不是字符串类型")
	}
	fmt.Println(ret)	//hello

	// x.(type)只能用在switch语句
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int，value is %v\n", v)
	case bool:
		fmt.Printf("x is a bool，value is %v\n", v)
	default:
		fmt.Println("猜不到了")
	}
}
