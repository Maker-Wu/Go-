package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Hello Gopher, Hello Golang"
	re := regexp.MustCompile(`Go(\w+)`)
	fmt.Println(re.FindAllStringSubmatch(text, -1))
	// [[Gopher pher] [Golang lang]]
}