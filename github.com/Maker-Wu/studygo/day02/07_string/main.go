package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("字符串测试")

	// 字符串转换为int型
	n1, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误, err:", err)
	}
	fmt.Printf("type of n1:%T\n", n1)		//type of n1:int

	// 字符串转换为float64型
	f1, err := strconv.ParseFloat("100.55", 64)
	if err != nil {
		fmt.Println("转换错误, err:", err)
	}
	fmt.Printf("type of f1:%T\n", f1)		//type of f1:float64

	// int类型转换为字符串
	str1 := strconv.Itoa(99)
	fmt.Println(str1)

	// 判断字符串是否包括某字符串
	str2 := "hello,world"
	isCon := strings.Contains(str2, "hello")
	fmt.Println(isCon)		//true

	// 字符串比较
	str3 := "hello,World"
	fmt.Println(str2 == str3)		//false
	isEq := strings.EqualFold(str2, str3)
	fmt.Println(isEq)				//true

	// 查找子串的位置,返回第一次出现的位置值,找不到返回-1
	index := strings.Index(str3, "l")
	fmt.Println(index)				//2
	index = strings.LastIndex(str3, "l")
	fmt.Println(index)				//9

	// 替换, n表示替换的次数，-1表示全部替换
	str4 := strings.Replace(str3, "World", "中兴", -1)
	fmt.Println(str4)				//hello,中兴

	// 字符串按照指定分隔符分割
	str5 := "foo,bar,baz"
	fmt.Println(strings.Split(str5, ","))	//[foo bar baz]

	// 根据空白符分割,不限定中间间隔几个空白符
	str6 := "  hello   it's  a  nice day today    "
	strSlice := strings.Fields(str6)
	fmt.Println(strSlice)			//[hello it's a nice day today]

	// 大小写
	str7 := "Hello hao"
	str7 = strings.ToUpper(str7)
	fmt.Println(str7)		//HELLO HAO
	str7 = strings.ToLower(str7)
	fmt.Println(str7)		//hello hao

	// 删除字符串的开头和尾部的空白符
	str8 := strings.TrimSpace(str6)
	fmt.Println(str8)		//hello   it's  a  nice day today

	// 删除字符串的开头和尾部指定的字符串
	str9 := "/Users/Documents/GOPatch/src/MyGO/config/TestString/"
	fmt.Println(strings.Trim(str9, "/ "))		// Users/Documents/GOPatch/src/MyGO/config/TestString
	// 此外还有strings.TrimLeft
	fmt.Println(strings.TrimRight(str9, "/"))	// /Users/Documents/GOPatch/src/MyGO/config/TestString

	// 统计子串出现的次数
	num := strings.Count("cheeseeee", "ee")
	fmt.Println(num)		//3

	// 前缀后缀
	fmt.Println(strings.HasPrefix("Gopher", "Go"))	//true
	fmt.Println(strings.HasSuffix("Gopher", "Go"))	//false

	// 拼接元素类型为string的slice
	str10 := strings.Join(strSlice, " ")
	fmt.Println(str10)			//hello it's a nice day today

}
