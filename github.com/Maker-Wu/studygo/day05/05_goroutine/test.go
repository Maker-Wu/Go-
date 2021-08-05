package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func putNum(out chan<- int) {
	defer wg.Done()
	for i := 1; i <= 8000; i++ {
		out <- i
	}
	close(out)
}

func isPrime(num int) bool {
	for i := 2; i <= num-1; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

func putPrimeNum(in <-chan int, out chan<- int) {
	defer wg.Done()
	for i := range in {
		if isPrime(i) {
			out <- i
		}
	}
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	start := time.Now()
	wg.Add(5)
	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go putPrimeNum(intChan, primeChan)
	}
	wg.Wait()
	close(primeChan)

	for i := range primeChan {
		fmt.Printf("素数 = %d\n", i)
	}
	end := time.Now()
	fmt.Println(end.Sub(start).Milliseconds())
}
