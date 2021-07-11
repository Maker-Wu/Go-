package main

import "fmt"

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// dog既可以实现Sayer接口，也可以实现Mover接口
type dog struct {
	name string
}

// 实现Sayer接口
func (d dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

// WashingMachiner 洗衣机
type WashingMachiner interface {
	wash()
	dry()
}

// 甩干器
type dryer struct {}

func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	/*
		一个接口的方法并不需要由一个类型完全实现，
		接口的方法可以通过在类型中嵌入其他类型或者结构体实现
	*/
	dryer
}

func (h haier) wash() {
	fmt.Println("洗刷刷")
}

// 接口与接口间可以通过嵌套创建出新的接口
type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Printf("%s会喵喵喵\n", c.name)
}

func (c cat) move() {
	fmt.Printf("%s会跳\n", c.name)
}

func main() {
	var x Sayer
	var y Mover

	d := dog{name: "旺财"}
	x = d
	y = d
	x.say()
	y.move()

	var w WashingMachiner
	h := haier{}
	w = h
	w.wash()
	w.dry()

	var al animal
	c := cat{name: "Jerry"}
	al = c
	al.say()
	al.move()
}