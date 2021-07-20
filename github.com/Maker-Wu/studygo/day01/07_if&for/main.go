package main

import (
	"fmt"
	"unicode"
)

func main()  {
	//if age > 18 {
	//	fmt.Println("澳门首家线上赌场开业了！")
	//} else {
	//	fmt.Println("回去写暑假作业！")
	//}

	// 多个判断条件
	if age := 19; age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("大为青年")
	} else {
		fmt.Println("好好学习！")
	}

	// 普通的for循环
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}

	// go语言没有while循环，for循环实现while
	j := 0
	for ; j < 10; j++ {
		fmt.Printf("j = %d\n", j)
		j++
	}

	// for range循环
	s := "Hello World"
	for index, value := range s {
		fmt.Printf("%d = %c\n", index, value)
	}

	// for循环打印乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

	// 判断字符串中汉字的数量
	s1 := "Hello伍胜强"
	var count int
	for _, c := range s1 {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Printf("字符串中出现汉字的个数:%d\n", count)

	// break-label
	label2:
	for i := 0; i < 4; i++ {
		//label1:
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label2
			}
			fmt.Println("j =", j)
		}
	}
}
