package main

import "fmt"

// 键盘输入

func main() {
	// 从控制台接收用户信息[姓名，年龄，薪水，是否通过考试]
	var name string
	var age int
	var salary float32
	var isPass bool

	fmt.Println("请输入姓名:")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄:")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水:")
	fmt.Scanln(&salary)
	fmt.Println("请输入是否通过考试:")
	fmt.Scanln(&isPass)

	// 姓名：伍胜强, 年龄：26, 薪水：28500.00, 是否通过考试:true
	fmt.Printf("姓名：%s, 年龄：%d, 薪水：%.2f, 是否通过考试:%t", name, age, salary, isPass)

	// fmt.Scanf 按照指定的格式输入
	fmt.Println("请输入姓名，年龄，薪水，是否通过考试（使用,隔开)")
	fmt.Scanf("%s,%d,%.2f,%t", &name, &age, &salary, &isPass)
	fmt.Printf("姓名：%s, 年龄：%d, 薪水：%.2f, 是否通过考试:%t", name, age, salary, isPass)

}
