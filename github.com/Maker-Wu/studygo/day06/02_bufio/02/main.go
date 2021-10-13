package main

import (
	"bufio"
	"fmt"
)

type Write int

// 实现io.Writer接口
func (*Write) Write(p []byte) (int, error) {
	fmt.Println(len(p))
	return len(p), nil
}

func main() {
	fmt.Println("Unbuffered I/O")
	w := new(Write)
	w.Write([]byte{'a'})
	w.Write([]byte{'b'})
	w.Write([]byte{'c'})
	w.Write([]byte{'d'})
	fmt.Println("Buffered I/O")
	bw := bufio.NewWriterSize(w, 3)
	bw.Write([]byte{'a'})
	bw.Write([]byte{'b'})
	bw.Write([]byte{'c'})
	bw.Write([]byte{'d'})
	err := bw.Flush()
	if err != nil {
		panic(err)
	}

}
