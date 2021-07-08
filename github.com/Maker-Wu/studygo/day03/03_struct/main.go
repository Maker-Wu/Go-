package main

import "fmt"

// 结构体的定义
type person struct {
	name, city string
	age int
}

// 构造函数
func newPerson(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age: age,
	}
}

func main() {
	var p person
	p.name = "沙河娜扎"
	p.city = "北京"
	p.age = 18
	fmt.Printf("p=%v\n", p)			//p={沙河娜扎 北京 18}
	fmt.Printf("p=%#v\n", p)			//p=main.person{name:"沙河娜扎", city:"北京", age:18}

	// 匿名结构体
	var user struct{Name string; Age int}
	user.Name = "小王子"
	user.Age = 18
	fmt.Printf("%#v\n", user)		//struct { Name string; Age int }{Name:"小王子", Age:18}

	// 创建指针类型结构体
	var p2 = new(person)
	fmt.Printf("%T\n", p2)			//*main.person
	fmt.Printf("%#v\n", p2)			//&main.person{name:"", city:"", age:0}

	// go语言支持对结构体指针直接使用.来访问结构体的成员
	p2.name = "小王子"
	p2.age = 28
	p2.city = "上海"
	fmt.Printf("p2=%#v\n", p2)			//p2=&main.person{name:"小王子", city:"上海", age:28}

	// 取结构体的地址实例化
	p3 := &person{}
	fmt.Printf("%T\n", p3)				//*main.person
	fmt.Printf("p3=%#v\n", p3)			//p3=&main.person{name:"", city:"", age:0}
	p3.age = 30
	p3.city = "成都"
	p3.name = "七米"				//p3.name 其实在底层是(*p3).name
	fmt.Printf("p3=%#v\n", p3)			//&main.person{name:"七米", city:"成都", age:30}

	// 使用键值对对结构体进行初始化
	p4 := person{
		name: "小王子",
		city: "北京",
		age: 18,
	}
	fmt.Printf("p4=%#v\n", p4)

	// 调用构造函数
	p5 :=newPerson("张三", "沙河", 18)
	fmt.Printf("%#v\n", p5)		//&main.person{name:"张三", city:"沙河", age:18}
}
