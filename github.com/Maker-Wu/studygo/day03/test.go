package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	ID int8
	name string
	age int8
	score float64
}

type Class struct {
	Title string
	Students []*Student
}

func NewClass(title string) *Class {
	return &Class{
		Title: title,
		Students: make([]*Student, 0),
	}
}

func (c *Class) AddStu(ID int8, name string, age int8, score float64) {
	stu := Student{
		ID: ID,
		name: name,
		age: age,
		score: score,
	}
	c.Students = append(c.Students, &stu)
}

func (c *Class) ShowStu() {
	for _, student := range c.Students {
		fmt.Println(student)
	}
}

func main() {
	class := NewClass("101")
	fmt.Printf("%#v\n", class)
	for i := 0; i < 10; i++ {
		id := int8(i)
		name := fmt.Sprintf("stu%02d", i)
		age := int8(25)
		score := rand.Float64()
		class.AddStu(id, name, age, score)
	}
	class.ShowStu()
}