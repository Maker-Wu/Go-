package main

import (
	"fmt"
	"reflect"
)

func reflectValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)	//修改的是副本，reflect包会引发panic
	}
}

func reflectValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a int64 = 100
	// reflectValue1(a)	////panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectValue2(&a)
	fmt.Println(a)
}
