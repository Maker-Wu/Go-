package main

import (
	"errors"
	"fmt"
	"strings"
)

// 定义函数类型
type calculation func(int, int) int

// 函数定义
func intSum(x int, y int) int {
	return x+y
}

/*
   函数的参数中如果相邻变量的类型相同
   则可以省略类型
 */
func intSub(x, y int) int {
	return x - y
}

// 可变参数
func intSum2(x ...int) int {
	fmt.Println(x)		//x是一个切片
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// 多返回值,必须用括号包裹起来
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 函数可以作为参数
func calc3(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数可以作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return intSum, nil
	case "-":
		return intSub, nil
	default:
		err := errors.New("无法识别的操作")
		return nil, err
	}
}

func someFunc(x string) map[string]int {
	if x == "" {
		return nil		//没必要返回[]int{}
	}
	return map[string]int{}
}

// 闭包=函数+引用环境
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

// 闭包应用添加后缀
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// defer后面的语句入栈时，也会将相关值拷贝同时入栈
func sum(n1, n2 int) int {
	defer fmt.Println("ok1 n1 =", n1)	//ok1 n1 = 10
	defer fmt.Println("ok2 n2 =", n2)	//ok2 n1 = 20
	n1++
	n2++
	res := n1 + n2 //32
	fmt.Println("ok3 res =", res)
	return res
}

func main() {
	res := intSum(10, 20)
	fmt.Println(res)
	// 调用可变参数函数
	ret1 := intSum2()
	ret2 := intSum2(10)
	ret3 := intSum2(10, 20)
	fmt.Println(ret1, ret2, ret3)	//10, 20, 30
	ret4, ret5 := calc2(50, 40)
	fmt.Println(ret4, ret5)
	fmt.Println(someFunc("b"))

	// 声明函数变量
	var c calculation							//声明一个calculaiton类型的变量c
	c = intSum									//把intSum赋值给c
	fmt.Printf("type of c:%T\n", intSum)	//type of c:func(int, int) int
	fmt.Printf("type of c:%T\n", c)		//type of c:main.go.calculation
	fmt.Println(c(1, 2))						//3

	f := intSum									//将函数intSum函数给变量f
	fmt.Printf("type of f:%T\n", f)		//type of c:func(int, int) int
	fmt.Println(f(1, 2))					//3

	// 函数可以作为参数
	ret6 := calc3(10, 20, intSum)
	fmt.Println(ret6)							//30

	// 匿名函数
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) 							//通过变量调用匿名函数
	fmt.Printf("type of add:%T\n", add)	//type of add:func(int, int)

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	f1 := adder()
	fmt.Println(f1(10))		//10
	fmt.Println(f1(20))		//30
	fmt.Println(f1(30))		//60

	f2 := adder()
	fmt.Println(f2(40))		//40
	fmt.Println(f2(50))		//90

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))			//test.jpg
	fmt.Println(txtFunc("test"))			//test.txt

	fmt.Println(jpgFunc("beautifulGirl"))			//beautifulGirl.jpg
	fmt.Println(txtFunc("beautifulGirl"))			//beautifulGirl.txt

	sum(10, 20)
}
