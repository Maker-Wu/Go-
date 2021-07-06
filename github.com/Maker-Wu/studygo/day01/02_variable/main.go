package main

import "fmt"

var (
	studentName string	// 默认值""
	age  int	// 默认值 0
	isOk bool	// 默认值false
)

func main() {
	studentName = "伍胜强"
	age = 16
	isOk = true		
	// 局部变量声明后必须使用，否则编译器报错
	var score float32 = 98
	// 类型推导（根据右边的值推导变量的类型）
	var grade = 3
	// 简短变量声明，只能在函数中使用
	mood := "哈哈哈"
	fmt.Print(isOk)
	fmt.Println()
	fmt.Println(age)
	fmt.Printf("name:%s\ngrade:%d\nscore:%f\n", studentName, grade, score)
	fmt.Print("此时的心情:"+mood)
}
