package main

// 接口是一种类型，一种抽象的类型
type cat struct {

}

func (c *cat) Say() string {
	return "喵喵喵"
}

type dog struct {

}

func (d *dog) Say() string {
	return "汪汪汪"
}

type writer interface {
	Writer([]byte) error
}


