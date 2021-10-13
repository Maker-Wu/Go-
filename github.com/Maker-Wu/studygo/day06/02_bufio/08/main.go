package main

import (
	"bufio"
	"fmt"
	"strings"
)

type R struct {
}

func (r *R) Read(p []byte) (n int, err error) {
	fmt.Println("Read")
	copy(p, strings.Repeat("a", len(p)))
	return len(p), nil
}

func main() {
	r := new(R)
	br := bufio.NewReaderSize(r, 16)
	buf := make([]byte, 17)
	n, err := br.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read = %q, n = %d\n", buf[:n], n)
	fmt.Printf("buffered = %d\n", br.Buffered())
}
