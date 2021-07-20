package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
简单描述：
Go中可以抛出一个panic的异常，
然后在defer中通过recover捕获这个异常，
然后正常处理
 */

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

func arrTest(arr *[3]int) {
	fmt.Println((*arr)[0])
	fmt.Println(arr[0])
}

func main() {

	var intArr = [5]int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		intArr[i] = rand.Intn(100)
	}
	fmt.Println(intArr)
	for i:= 0; i < len(intArr)/2; i++ {
		temp := intArr[i]
		intArr[i] = intArr[len(intArr)-i-1]
		intArr[len(intArr)-i-1] = temp
	}
	fmt.Println(intArr)
}

