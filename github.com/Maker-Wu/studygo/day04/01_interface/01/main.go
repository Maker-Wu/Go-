package main

import "fmt"

// 接口是一种类型，一种抽象的类型
type cat struct{}

// say 值接收者实现接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}

type dog struct{}

// Say 指针接收者实现接口
func (d *dog) say() {
	fmt.Println("汪汪汪")
}

// Sayer 接口
type Sayer interface {
	say()
}

func main() {
	var x Sayer //声明一个Sayer类型的变量x
	jerry := cat{}
	tom := &cat{}
	x = jerry //x可以直接接收cat类型
	x = tom   //x可以接收*cat类型
	x.say()
	wangcai := &dog{}
	//x = wangcai			//不可以把dog实例直接赋值给x
	x = wangcai
	x.say()
}
