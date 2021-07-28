package utils

import "fmt"

type familyAccount struct {
	balance float64
	details []string
}

func NewFamilyAccount(balance float64) *familyAccount {
	return &familyAccount{
		balance: balance,
	}
}

// 显示收支明细
func (f *familyAccount) showDetails() {
	defer fmt.Println()
	fmt.Println("----------当前收支明细记录----------")
	if f.details == nil {
		fmt.Println("当前收支没有明细... 来一笔吧!")
		return
	}
	for _, detail := range f.details {
		fmt.Println(detail)
	}
}

// 登记收入
func (f *familyAccount) incomeRecord()  {
	defer fmt.Println()
	fmt.Print("本次收入的金额：")
	var money float64
	fmt.Scanln(&money)
	f.balance += money
	var note string
	fmt.Print("本次收入说明：")
	fmt.Scanln(&note)
	detail := fmt.Sprintf("收入\t\t%.2f\t\t%.2f\t\t%s", f.balance, money, note)
	f.details = append(f.details, detail)
}

// 登记支出
func (f *familyAccount) payRecord()  {
	defer fmt.Println()
	fmt.Print("本次支出的金额：")
	var money float64
	fmt.Scanln(&money)
	f.balance -= money
	var note string
	fmt.Print("本次支出说明：")
	fmt.Scanln(&note)
	detail := fmt.Sprintf("支出\t\t%.2f\t\t%.2f\t\t%s", f.balance, money, note)
	f.details = append(f.details, detail)
}

func (f *familyAccount) exit() bool {
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

// MainMenu 主菜单
func (f *familyAccount) MainMenu() {
	for {
		fmt.Println("----------家庭收支记账软件----------")
		fmt.Println("            1 收支明细")
		fmt.Println("            2 登记收入")
		fmt.Println("            3 登记支出")
		fmt.Println("            4 退出软件")
		var key string
		fmt.Print("请选择<1-4>:")
		fmt.Scanln(&key)
		fmt.Println()
		switch key {
		case "1":
			f.showDetails()
		case "2":
			f.incomeRecord()
		case "3":
			f.payRecord()
		case "4":
			if f.exit() {
				return
			}
		}
	}
}
