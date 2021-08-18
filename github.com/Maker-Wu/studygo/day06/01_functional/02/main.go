package main

import "fmt"

// 闭包实现累加器
func adder() func() int {
	v, sum := 0, 0
	return func() int {
		sum += v
		v++
		return sum
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a())
	}
}
