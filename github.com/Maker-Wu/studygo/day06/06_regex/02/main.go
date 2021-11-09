package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`[\s,]+`)
	fmt.Println(re.Split("Helo Gopher, Hello Golang", -1))
}
