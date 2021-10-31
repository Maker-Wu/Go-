package main

import (
	"fmt"
	"sync"
)

func main() {
	var pool sync.Pool

	pool.Put(1)
	pool.Put("hello")
	pool.Put(true)
	pool.Put(3.14)

	for {
		value := pool.Get()
		if value == nil {
			break
		}
		fmt.Println(value)
	}
}
