package main

import "fmt"

// Person 结构体
type Person struct {
	name string
	age int8
}

// NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age: age,
	}
}

// Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言!\n", p.name)
}

// SetAge 设置p的年龄
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

type MyInt int

// SayHello 为MyInt添加一个SayHello()方法
func (m MyInt) SayHello() {
	fmt.Println("Hello")
}

func main() {
	p := NewPerson("伍胜强", 26)
	p.Dream()
	fmt.Println(p.age)
	p.SetAge(25)
	fmt.Println(p.age)

	var m1 MyInt
	m1 = 10
	m1.SayHello()
}
