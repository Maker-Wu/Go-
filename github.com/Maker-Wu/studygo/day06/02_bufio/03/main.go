package main

import (
	"bufio"
	"fmt"
)

type Write1 int

func (w *Write1) Write(p []byte) (n int, err error) {
	fmt.Printf("writer#1: %q\n", p)
	return len(p), nil
}

type Write2 int

func (w *Write2) Write(p []byte) (n int, err error) {
	fmt.Printf("writer#2: %q\n", p)
	return len(p), nil
}

func main() {
	w1 := new(Write1)
	bw := bufio.NewWriterSize(w1, 2)
	bw.Write([]byte("ab"))
	bw.Write([]byte("cd"))
	w2 := new(Write2)
	bw.Reset(w2)
	bw.Write([]byte("ef"))
	bw.Flush()

}
