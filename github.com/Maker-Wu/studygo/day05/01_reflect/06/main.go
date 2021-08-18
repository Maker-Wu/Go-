package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name string		`json:"name"`
	Age int			`json:"monster_age"`
	Score float32
	Sex string
}

// Print 方法，显示s的值
func (m Monster) Print() {
	fmt.Println("---start----")
	fmt.Println(m)
	fmt.Println("---end----")
}

func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}




func TestStruct(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	if t.Kind() != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := v.NumField()
	fmt.Printf("struct has %d fields\n", num)	//
	for i := 0; i < num; i++ {
		fieldStruct := t.Field(i)
		// v.Field(i)返回字段的reflect.Value
		fmt.Printf("Field %d: 值为=%v\n", i, v.Field(i))
		tagVal := fieldStruct.Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}

	numOfMethod := v.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	method := t.Method(1)
	fmt.Println(method.Name)	//Print
	fmt.Println(method.Type)	//func(main.Monster)

	// 调用Print方法
	v.Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := v.Method(0).Call(params)
	fmt.Println("res =", res[0].Int())	//res = 50
}

func ModifyField(x interface{}) {
	v := reflect.ValueOf(x)
	v.Elem().Field(0).SetString("小地鼠")
}

func main() {
	a := Monster{
		Name: "黄鼠狼精",
		Age: 400,
		Score: 30.8,
	}
	TestStruct(a)
	ModifyField(&a)
	fmt.Println(a)		//{小地鼠 400 30.8 }
}
