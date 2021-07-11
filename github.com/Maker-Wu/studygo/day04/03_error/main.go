package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("type of f:%T\n", err)
		fmt.Println("error:", err)
		return
	}
	fmt.Println(f.Name(), "open successfully")
}


