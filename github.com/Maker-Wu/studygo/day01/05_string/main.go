package main

import "fmt"

func main() {
	// 字符串
	s := "Hello World"
	fmt.Printf("%s\n", s)

	// 单独的字母、汉字和符号表示一个字符
	c1 := 'A'
	c2 := '1'
	c3 := '汉'

	fmt.Printf("%c, %c, %c\n", c1, c2, c3)
	//fmt.Println(c1, c2, c3)

	// go语言定义多行字符串,必须使用反引号``
	var s2 = `第一行
第二行
第三行`
	fmt.Println(s2)

	// 字符串相关操作
	// 字节长度
	fmt.Println(len(s2)) //29

	// 字符串拼接
	s3 := "理想"
	s4 := "大帅比"

	var ss1 = s3 + s4
	fmt.Println(ss1)
	ss2 := fmt.Sprintf("%s是成为%s", s3, s4)
	fmt.Println(ss2)

	/* 字符串其他操作：
	strings.Split							分割
	strings.Contains						判断是否包含
	strings.HasPrefix, strings.HasSuffix	前缀/后缀判断
	strings.Index, strings.LastIndex		子串出现的位置
	strings.Join							join操作

	*/
}
