package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age int
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type of v:%T\n", v)		//type of v:*reflect.rtype
}

func main() {
	var a float32 = 3.14
	reflectType(a)			//type:float32
	var b int64 = 100
	reflectType(b)			//type:int64
}
