package main

import "fmt"

/*
切片是基于数组类型做的一层封装。他非常灵活，支持自动扩容
切片是一个引用类型，它的内部结构包含地址、长度和容量。
 */
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
	fmt.Printf("c:%v len(c):%d cap(c):%d\n", c, len(c), cap(c))	//c:[false true] len(c):2 cap(c):2

	// 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s := a1[:4]
	/*
	切片指向了一个底层的数组
	切片的长度就是它元素的个数
	切片的容量是底层数组从切片的第一个元素到最后的元素数量
	 */
	fmt.Printf("s:%v len(s):%d cap(s):%d\n", s, len(s), cap(s))	//s:[1 3 5 7] len(s):4 cap(s):7

	// 使用make()函数构造切片
	m := make([]int, 5, 10)
	fmt.Printf("m:%v len(m):%d cap(m)%d\n", m, len(m), cap(m))	//m:[0 0 0 0 0] len(m):5 cap(m)10

	//cities := []string{"北京", "上海", "深圳"}
	//cities[3] = "广州"		//错误的写法，会导致编译错误:索引越界
	// append()方法为切片添加元素
	var s1 []int
	s1 = append(s1, 1)		//[1]
	s1 = append(s1, 2, 3, 4)
	s2 := []int{5, 6, 7}
	s1 = append(s1, s2...)
	fmt.Println(s1)					//[1 2 3 4 5 6 7]

	//copy()函数复制切片
	a2 := []int{1, 2, 3, 4, 5}
	c2 := make([]int, 5, 5)		//提前申请足够的空间,
	copy(c2, a2)
	fmt.Println(a2)				//[1 2 3 4 5]
	fmt.Println(c2)				//[1 2 3 4 5]

	c3 := make([]int , 1)
	copy(c3, a2)			//copy()不会自动扩容
	fmt.Println(c3)				//[1]

	//从切片中删除元素
	a3 := []int{31, 32, 33, 34, 35, 36, 37}
	a3 = append(a3[:2], a3[3:]...)
	fmt.Println(a3)		//[31 32 34 35 36 37]
}
