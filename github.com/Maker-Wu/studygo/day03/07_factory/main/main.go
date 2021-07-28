package main

import (
	"fmt"
	"github.com/Maker-Wu/studygo/day03/07_factory/model"
)

func main() {
	var stu = model.NewStudent("tom", 98)
	fmt.Printf("stu:%#v\n", stu)		//stu:&model.student{Name:"tom", score:98}
	//fmt.Println(stu.score) 				//不能直接访问score
	fmt.Println(stu.GetScore())				//98
}
