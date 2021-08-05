package main

import (
	"fmt"
	"sync"
)

// 使用sync.WaitGroup来实现goroutine的同步
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello goroutine!", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)	// 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait()
}
