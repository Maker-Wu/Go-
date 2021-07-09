package main

import "fmt"

// Address 地址结构体
type Address struct {
	Province string
	City string
}

// User 用户结构体
type User struct {
	Name string
	Gender string
	Address //匿名字段
}

// Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动!\n", a.name)
}

type Dog struct {
	Feet int8
	*Animal		//通过嵌套匿名结构体实现继承
}

type student struct {
	ID int8
	Name string
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	user1 := User{
		Name: "小王子",
		Gender: "男",
		Address: Address{
			Province: "山东",
			City: "威海",
		},
	}
	fmt.Printf("user1=%#v\n", user1)	//user1=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}

	// 访问匿名字段默认使用类型名作为字段名
	fmt.Println(user1.Address.Province)
	// 匿名字段可以省略
	fmt.Println(user1.City)

	d1 := Dog{
		Feet: 4,
		Animal: &Animal{		//注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang()
	d1.move()
}