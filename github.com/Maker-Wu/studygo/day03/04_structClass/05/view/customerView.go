package view

import (
	"fmt"
	"github.com/Maker-Wu/studygo/day03/04_structClass/05/model"
)

type customerView struct {
	customers []*model.Customer
	no int
}

func NewCustomerView() *customerView{
	return &customerView{}
}

func (c *customerView) AddCustomers() {
	defer fmt.Println()
	var (
		name string
		gender string
		age int
		phone string
		email string
	)
	fmt.Println("----------添加客户信息明细----------")
	fmt.Print("姓名：")
	fmt.Scanln(&name)
	fmt.Print("性别：")
	fmt.Scanln(&gender)
	fmt.Print("年龄：")
	fmt.Scanln(&age)
	fmt.Print("电话：")
	fmt.Scanln(&phone)
	fmt.Print("邮箱：")
	fmt.Scanln(&email)
	c.no++
	customer := model.NewCustomer(c.no, name, gender, age, phone, email)
	c.customers = append(c.customers, customer)
}

func (c *customerView) MainMenu() {
	for {
		fmt.Println("----------客户信息管理软件----------")
		fmt.Println("            1 添加客户")
		fmt.Println("            2 修改客户")
		fmt.Println("            3 删除客户")
		fmt.Println("            4 客户列表")
		fmt.Println("            5 退   出")
		var key string
		fmt.Print("请选择<1-5>:")
		fmt.Scanln(&key)
		fmt.Println()
		switch key {
		case "1":
			c.AddCustomers()
		case "2":
			//f.incomeRecord()
		case "3":
			c.deleteCustomer()
		case "4":
			c.showDetails()
		case "5":
			if c.exit() {
				return
			}
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
}

func (c *customerView) showDetails() {
	defer fmt.Println()
	fmt.Println("----------当前客户信息明细----------")
	if c.customers == nil {
		fmt.Println("当前没有客户信息... 添加一个吧!")
		return
	}
	for _, customer := range c.customers {
		fmt.Println(customer)
	}
}

func (c *customerView) deleteCustomer() {
	defer fmt.Println()
	fmt.Println("----------删除客户信息明细----------")
	fmt.Printf("请选择待删除客户编号（-1退出）:")
	var no int
	fmt.Scanln(&no)
	for index, customer := range c.customers {
		if no == customer.Id {
			if c.exit() {
				c.customers = append(c.customers[:index], c.customers[index+1:]...)
			}
			return
		}
	}
	fmt.Println("没有这个用户")
}

func (c *customerView) exit() bool{
	var res string
	defer fmt.Println()
	fmt.Println("你确定要退出吗? y/n")
	for {
		fmt.Scanln(&res)
		if res == "y" {
			return true
		} else if res == "n" {
			return false
		} else {
			fmt.Println("输入错误请重新输入 y/n")
		}
	}
}