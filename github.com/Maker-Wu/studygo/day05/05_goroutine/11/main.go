package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var cond = sync.NewCond(new(sync.Mutex))

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("协程", i, "启动。。。")
			cond.L.Lock()
			fmt.Println("协程", i, "加锁。。。")
			defer cond.L.Unlock()
			cond.Wait()
			fmt.Println("协程", i, "解锁。。。")
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("主协程发送信号量。。。")
	cond.Signal()

	time.Sleep(time.Second)
	fmt.Println("主协程发送信号量。。。")
	cond.Signal()

	time.Sleep(time.Second)
	fmt.Println("主协程发送信号量。。。")
	cond.Signal()
	wg.Wait()
}
