package main

import (
	"bufio"
	"errors"
	"fmt"
)

type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	fmt.Printf("Write: %q\n", p)
	return 0, errors.New("boom!")
}
func main() {
	w := new(Writer)
	bw := bufio.NewWriterSize(w, 3)
	fmt.Println(bw.Buffered()) //0
	bw.Write([]byte{'a'})
	fmt.Println(bw.Buffered()) //1
	bw.Write([]byte{'b'})
	fmt.Println(bw.Buffered()) //2
	bw.Write([]byte{'c'})
	fmt.Println(bw.Buffered()) //3
	bw.Write([]byte{'d'})
	fmt.Println(bw.Buffered()) //3
	err := bw.Flush()
	fmt.Println(err)
}
