package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s1 := strings.NewReader(strings.Repeat("a", 20))
	reader := bufio.NewReaderSize(s1, 16)
	b, err := reader.Peek(3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%q\n", b)
	b, err = reader.Peek(17)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%q\n", b)
	s2 := strings.NewReader("aaa")
	reader.Reset(s2)
	b, err = reader.Peek(10)
	if err != nil {
		fmt.Println(err)
	}
}
