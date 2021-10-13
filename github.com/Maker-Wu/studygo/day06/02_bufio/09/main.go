package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("abcd")
	br := bufio.NewReader(r)
	b, err := br.ReadByte()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", b)
	fmt.Printf("buffered = %d\n", br.Buffered())
	err = br.UnreadByte()
	if err != nil {
		panic(err)
	}
	fmt.Printf("buffered = %d\n", br.Buffered())
	b, err = br.ReadByte()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", b)

}
