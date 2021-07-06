package main

import "fmt"

func main() {
	s := "Hello沙河"
	fmt.Println(len(s)) //len()是求byte字节的数量

	for _, c := range s {
		fmt.Printf("%c\n", c)
	}

	// 字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[0] = '红'
	fmt.Printf("s2 = %s\n", s2)
	fmt.Println(string(s3)) //红萝卜

}
