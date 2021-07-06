package main

import "fmt"

const pi = 3.14159

// 批量声明常量，如果某一行没有赋值，默认取上一行的值
const (
	n1 = 100
	n2
	n3
)

// iota只能在常量表达式中使用。
/* iota在const关键字出现时被重置为0，
   const中每新增一行iota计数一次(iota可理解为const语句块的行索引)*/
const(
	a1 = iota // 0
	a2 // 1
	a3 // 2
)

const(
	b1 = iota // 0
	b2 // 1
	_  // 2
	b3 // 3
)

// 插队
const(
	c1 = iota //0
	c2 = 100  //100
	c3 = iota // 2
	c4 // 3
)

// 多个常量声明在一行
const(
	d1, d2 = iota+1, iota+2 //1, 2
	d3, d4 = iota+1, iota+2 //2, 3
)

func main() {
	// 定义了常量之后就不能修改
	//pi = 3
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("a1:", a1)  // 0
	fmt.Println("a2:", a2)	// 1
	fmt.Println("a3:", a3)	// 2

	fmt.Println("b1:", b1)  // 0
	fmt.Println("b2:", b2)	// 1
	fmt.Println("b3:", b3)	// 3

	fmt.Println("c1:", c1)  // 0
	fmt.Println("c2:", c2)	// 100
	fmt.Println("c3:", c3)	// 2
	fmt.Println("c4:", c4)	// 3

	fmt.Println("d1:", d1)  // 1
	fmt.Println("d2:", d2)	// 2
	fmt.Println("d3:", d3)	// 2
	fmt.Println("d4:", d4)	// 3
}