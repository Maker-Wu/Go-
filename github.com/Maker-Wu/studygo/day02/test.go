package main

import "fmt"

func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3, 4)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", m["q1mi"])
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