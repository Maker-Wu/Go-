package main

import (
	"fmt"
	"strconv"
)

var (
	studentName string	// 默认值""
	age  int	// 默认值 0
	isOk bool	// 默认值false
)

// 基本数据类型转成字符串
func basicToStr() {
	// 方式1：fmt.Sprintf
	num1 := 99
	num2 := 23.456
	b := true
	myChar := 'h'

	str := fmt.Sprintf("%v", num1)
	fmt.Println(str)
	str = fmt.Sprintf("%f", num2)
	fmt.Println(str)
	str = fmt.Sprintf("%t", b)
	fmt.Println(str)
	str = fmt.Sprintf("%c", myChar)
	fmt.Println(str)
}

func basicToStr2() {
	var num1 = 99
	var num2 = 23.456
	var b = true

	str := strconv.Itoa(num1)
	fmt.Println(str)
	str = strconv.FormatInt(int64(num1), 10)
	fmt.Println(str)
	str = strconv.FormatFloat(num2, 'f', 10, 64)
	fmt.Println(str)
	str = strconv.FormatBool(b)
	fmt.Println(str)
}

// string类型转基本数据类型
func strToBasic() {
	str1 := "true"
	var b bool
	b, _ = strconv.ParseBool(str1)
	fmt.Printf("type of b:%T, val:%t\n", b, b)	//type of b:bool, val:true
	str2 := "12345"
	num, _ := strconv.ParseInt(str2, 10, 32)
	fmt.Printf("type of num:%T, val:%d\n", num, num)	//type of num:int64, val:12345
	str3 := "-3.141592678"
	fnum, _ := strconv.ParseFloat(str3, 32)
	fmt.Printf("type of fnum:%T, val:%f\n", fnum, fnum)	//type of fnum:float64, val:-3.141593

	/*
	要确保string类型能够转换成有效的数据，否则会被转换成对应类型的默认值
	比如：字符串"today"被转换为整数后值为0，转换成布尔值为false
	 */
	st4 := "700today"
	num2, _ := strconv.ParseInt(st4, 10, 32)
	fmt.Printf("type of num2:%T, val:%d\n", num2, num2)	//type of num2:int64, val:0
}

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
	fmt.Println("此时的心情:"+mood)

	basicToStr()
	strToBasic()
}
