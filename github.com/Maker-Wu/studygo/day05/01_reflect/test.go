package main

import "fmt"

type student struct {
	Name string
}

func (s *student) SetField(name string) {
	s.Name = name
}

func main() {
	stu := student{
		Name: "wsq",
	}
	stu.SetField("12")
	fmt.Println(stu)
}
