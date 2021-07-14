package main

import "fmt"

func main() {
	/*
	var 数组变量名 [元素数量]T
	数组的长度必须是常量，并且长度是数组类型的一部分
	 */
	var a1 [3]bool		//a1:[3]bool
	var a2 [4]bool		//a1:[4]bool

	fmt.Printf("a1:%T\n", a1)
	fmt.Printf("a2:%T\n", a2)

	// 数组的初始化
	var testArray [3]int
	var numArray = [3]int{1, 2}		//根据指定的初始值完成初始化，未指定初始值的元素取默认值
	var cityArray = [3]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)		//[0 0 0]
	fmt.Println(numArray)		//[1 2 0]
	fmt.Println(cityArray)		//[北京 上海 深圳]

	// 根据初始值的个数自行推断数组的长度
	var testArray2 [3]int
	var numArray2 = [...]int{1, 2}
	var cityArray2 = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray2)		//[0 0 0]
	fmt.Println(numArray2)		//[1 2]
	fmt.Printf("type of numArray2:%T\n", numArray2)		//type of numArray2:[2]int
	fmt.Println(cityArray2)		//[北京 上海 深圳]
	fmt.Printf("type of cityArray2:%T\n", cityArray2)	//type of cityArray2:[3]string

	// 指定索引值
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)		//[0 1 0 5]
	fmt.Printf("type of a:%T\n", a)		//type of a:[4]int

	// 多维数组
	var a11 [3][2]int
	a11 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a11)

	// 多维数组只有第一层可以是使用...来让编译器推导数组长度
	cities := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(cities)

	// 遍历二维数组
	for _, city := range cities {
		for _, s := range city {
			fmt.Println(s)
		}
	}

	nums := [...]int{1, 3, 5, 7, 8}

	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] + nums[j] == 8 {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
}
