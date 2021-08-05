package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a string = "123"
	v := reflect.ValueOf(&a)
	v.Elem().SetString("234")
	fmt.Println(a)
}
