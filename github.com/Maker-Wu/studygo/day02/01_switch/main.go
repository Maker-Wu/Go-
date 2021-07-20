package main

import (
	"fmt"
)

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

	// switch语句后面可以不跟变量，类似if-else分支来使用
	var age, _ = fmt.Sscan("请输入年龄：")
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age < 35:
		fmt.Println("好好工作吧")
	case age < 55:
		fmt.Println("准备退休了")
	}

	var letter rune
	fmt.Println("请输入一个字符")
	fmt.Scanf("%c", &letter)
	switch letter {
	case 'a':
		fmt.Println("星期一")
	case 'b':
		fmt.Println("星期二")
	case 'c':
		fmt.Println("星期三")
	case 'e':
		fmt.Println("星期四")
	case 'f':
		fmt.Println("星期五")
	default:
		fmt.Println("输入错误。。")
	}

	// type-switch来判断某个interface变量中实际指向的变量类型
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type){
	case nil:
		fmt.Printf("type of x:%T\n", i)
	case int:
		fmt.Printf("type of x:%T\n", i)
	case float64:
		fmt.Printf("type of x:%T\n", i)
	case func(int) float64:
		fmt.Printf("type of x:%T\n", i)
	case bool, string:
		fmt.Printf("type of x:%T\n", i)
	}
}
