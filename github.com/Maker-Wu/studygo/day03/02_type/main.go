package main

import "fmt"

// NewInt 类型定义
type NewInt int

// MyInt 类型别名
type MyInt = int

func main() {
	var a NewInt
	var b MyInt

	/*
		main.NewInt表示main包下定义的NewInt类型。
		b的类型是int说明MyInt只会在代码中存在，编译完成时并不会有MyInt类型
	 */
	fmt.Printf("type of a:%T\n", a)			//type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b)			//type of b:int
}
