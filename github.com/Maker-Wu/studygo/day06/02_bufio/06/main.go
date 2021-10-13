package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s1 := strings.NewReader(strings.Repeat("a", 16) + strings.Repeat("b", 16))
	r := bufio.NewReaderSize(s1, 16)
	b, _ := r.Peek(3)
	fmt.Printf("%q\n", b) //"aaa"
	r.Read(make([]byte, 16))
	r.Read(make([]byte, 15))
	fmt.Printf("%q\n", b) //"bbb"
}
