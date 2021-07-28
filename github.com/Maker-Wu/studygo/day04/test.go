package main

import (
	"fmt"
)

type Sayer interface {
	Say()
	Move()
}

type Mover interface {
	Move()
}

type Animal interface {
	Sayer
	Mover
	Eat()
}

type Dog struct {

}

func (d *Dog) Say() {
	fmt.Println("Dog can wangwang~~")
}

func (d *Dog) Move() {
	fmt.Println("Dog can run fast")
}

func (d *Dog) Eat() {
	fmt.Println("Dog can eat much")
}

func main() {
	var a Animal = &Dog{}
	a.Move()
}