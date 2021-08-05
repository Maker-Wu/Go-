package main

import (
	"fmt"
	"time"
)

func recv(c chan int) {
	ret := <- c
	fmt.Println("接收成功", ret)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
			fmt.Println("channel1", i)
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <- ch1	// 通道关闭后再取值ok=false
			if !ok {
				break
			}
			fmt.Println("channel2", i)
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 {	// 通道关闭后会退出for range循环
		fmt.Println(i)
	}
	time.Sleep(time.Second)
}
