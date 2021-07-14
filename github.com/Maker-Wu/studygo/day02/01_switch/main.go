package main

import "fmt"

func main() {
	switch n := 3; n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的数字")
	}

	// 分支可以有多个值
	switch n := 7; n{
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	}

	// 分支还可以使用表达式，这时switch语句后面不需要再跟变量
	var age, _ = fmt.Sscan("请输入年龄：")
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age < 35:
		fmt.Println("好好工作吧")
	case age < 55:
		fmt.Println("准备退休了")

	}
}
