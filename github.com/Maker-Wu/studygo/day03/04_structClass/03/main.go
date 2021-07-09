package main

import (
	"encoding/json"
	"fmt"
)

// Student 学生
type Student struct {
	ID int	`json:"id"`
	Gender string
	Name string
}

// Class 班级
type Class struct {
	Title string
	Students []Student
}

func main() {
	c := &Class{
		Title: "101",
		Students: make([]Student, 0, 200),
	}

	for i := 0; i < 10; i++ {
		stu := Student{
			ID: i,
			Gender: "男",
			Name: fmt.Sprintf("stu%02d", i),
		}
		c.Students = append(c.Students, stu)
	}
	// JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("type of data:%T\n", data)
	fmt.Printf("json:%s\n", data)

	// JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"id":0,"Gender":"男","Name":"stu00"},{"id":1,"Gender":"男","Name":"stu01"},{"id":2,"Gender":"男","Name":"stu02"},{"id":3,"Gender":"男","Name":"stu03"},{"id":4,"Gender":"男","Name":"stu04"},{"id":5,"Gender":"男","Name":"stu05"},{"id":6,"Gender":"男","Name":"stu06"},{"id":7,"Gender":"男","Name":"stu07"},{"id":8,"Gender":"男","Name":"stu08"},{"id":9,"Gender":"男","Name":"stu09"}]}`
	c2 := &Class{}
	err = json.Unmarshal([]byte(str), c2)
	if err != nil {
		fmt.Println("json unmarshal failed!")
	}
	fmt.Printf("%#v\n", c2)
}


