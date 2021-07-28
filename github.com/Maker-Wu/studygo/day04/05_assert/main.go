package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {

}

func (p *Phone) Start() {
	fmt.Println("手机开始工作...")
}

func (p *Phone) Stop() {
	fmt.Println("手机结束工作...")
}

func (p *Phone) Call() {
	fmt.Println("手机通话中...")
}

type Camera struct {

}

func (c *Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c *Camera) Stop() {
	fmt.Println("相机结束工作...")
}

type Computer struct {

}

func (c *Computer) Working(u Usb) {
	u.Start()
	if phone, ok := u.(*Phone); ok {
		phone.Call()
	}
	u.Stop()
}

// 判断输入的参数是什么类型
func judgeType(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%d个参数是bool类型，值是%v\n", index, x)
		case float32, float64:
			fmt.Printf("第%d个参数是float类型，值是%v\n", index, x)
		case int:
			fmt.Printf("第%d个参数是int类型，值是%v\n", index, x)
		case string:
			fmt.Printf("第%d个参数是string类型，值是%v\n", index, x)
		}
	}
}

func main() {
	var usbArr [3]Usb
	usbArr[0] = &Phone{}
	usbArr[1] = &Phone{}
	usbArr[2] = &Camera{}

	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
	}
}
