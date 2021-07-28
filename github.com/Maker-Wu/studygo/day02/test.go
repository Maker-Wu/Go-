package main

import (
	"fmt"
)

func peach(day int) int {
	if day == 10 {
		return 1
	}
	return 2*(peach(day+1) + 1)
}

func fbn(n int) []int {
	sli := make([]int, 0)
	if n == 1 {
		return append(sli, 1)
	} else if n == 2 {
		return append(sli, 1, 1)
	} else {
		sli = append(sli, fbn(n-1)...)
		sli = append(sli, sli[len(sli)-1]+sli[len(sli)-2])
		return sli
	}
}

func BubbleSort(arr *[5]int) {
	for i := 0; i < len(arr)-1; i++ {
		flag := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				flag = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !flag {
			break
		}
		fmt.Printf("第%d次排序，arr:%v\n", i+1, arr)
	}
}

// BinaryFind 二分查找
func BinaryFind(arr [6]int, left, right, value int) {
	if left > right {
		fmt.Println("找不到")
		return
	}
	middle := (left + right) / 2
	if arr[middle] == value {
		fmt.Println("找到了，坐标值为", middle)
	} else if arr[middle] < value {
		left = middle + 1
		BinaryFind(arr, left, right, value)
	} else if arr[middle] > value {
		right = middle - 1
		BinaryFind(arr, left, right, value)
	}
}

func modifyUser(users map[string]map[string]string, name string) {
	_, ok := users[name]
	if ok {
		fmt.Println("用户名存在")
		users[name]["password"] = "888888"
	} else {
		user := make(map[string]string, 2)
		user["nickname"] = name
		user["password"] = "888888"
		users[name] = user
	}
}

func main() {
	users := make(map[string]map[string]string, 10)
	modifyUser(users, "stu01")
	fmt.Println(users)
}

func sum(n1, n2, n3 int) int{
	return n1 + n2
}

//func calc(index string, a, b int) int {
//	ret := a + b
//	fmt.Println(index, a, b, ret)
//	return ret
//}
//
//func main() {
//	x := 1
//	y := 2
//	defer calc("AA", x, calc("A", x, y))
//	x = 10
//	defer calc("BB", x, calc("B", x, y))
//	y = 20
//}