package main

import "fmt"

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a)		//a:10 ptr:0xc00000a0a8
	fmt.Printf("b:%p type:%T\n", b, b)		//b:0xc00000a0a8 type:*int
	fmt.Println(&b)									//0xc000006028

	modify1(a)
	fmt.Println(10)								//10
	modify2(&a)
	fmt.Println(a)									//100

	// new()函数申请内存空间
	c := new(int)
	d := new(bool)
	fmt.Printf("%T\n", c)					//*int
	fmt.Printf("%T\n", d)					//*bool

	// make()函数申请内存空间
	// e是一个map类型的变量，需要使用make函数进行初始化操作后，才能对齐进行赋值
	var e map[string]int
	e = make(map[string]int, 10)
	e["沙河娜扎"] = 100
	fmt.Println(e)
}
