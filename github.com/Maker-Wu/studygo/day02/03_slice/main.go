package main

import "fmt"

func main() {
	// 声明切片类型
	var a []string					//声明一个字符串切片
	var b = []int{}					//声明一个整型切片并初始化
	var c = []bool{false, true}
	var d = []bool{false, true}
	fmt.Println(a)					//[]
	fmt.Println(b)					//[]
	fmt.Println(c)					//[false true]
	fmt.Println(d)

	fmt.Println(a == nil)			//true
	fmt.Println(b == nil)			//false
	// fmt.Println(c == d) 			//切片是引用类型，不支持直接比较，只能和nil比较

	// 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s := a1[:4]
	/* 	切片指向了一个底层的数组
		切片的长度就是它元素的个数
		切片的容量是底层数组从切片的第一个元素到最后的元素数量
	 */
	fmt.Printf("s:%v len(s):%d cap(s):%d\n", s, len(s), cap(s))

	// 使用make()函数构造切片
	m := make([]int, 5, 10)
	fmt.Printf("m:%v len(m):%d cap(m)%d\n", m, len(m), cap(m))

	// append()方法为切片添加元素
	var s1 []int
	s1 = append(s1, 1)		//[1]
	s1 = append(s1, 2, 3, 4)
	s2 := []int{5, 6, 7}
	s1 = append(s1, s2...)
	fmt.Println(s1)					//[1 2 3 4 5 6 7]

	//copy()函数复制切片
	a2 := []int{1, 2, 3, 4, 5}
	c2 := make([]int, 5, 5)
	copy(c2, a2)
	fmt.Println(a2)
	fmt.Println(c2)

	//从切片中删除元素
	a3 := []int{31, 32, 33, 34, 35, 36, 37}
	a3 = append(a3[:2], a3[3:]...)
	fmt.Println(a3)
}
