package main

import "fmt"

func peach(day int) int {
	if day == 10 {
		return 1
	}
	return 2*(peach(day+1) + 1)
}

func main() {
	fmt.Println(peach(1))
}

func sum(n1, n2, n3 int) int{
	return n1 + n2
}

//func calc(index string, a, b int) int {
//	ret := a + b
//	fmt.Println(index, a, b, ret)
//	return ret
//}
//
//func main() {
//	x := 1
//	y := 2
//	defer calc("AA", x, calc("A", x, y))
//	x = 10
//	defer calc("BB", x, calc("B", x, y))
//	y = 20
//}