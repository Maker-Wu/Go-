package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibinacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 为闭包创建一个类型，为其实现一个io.Reader接口
type intGen func() int

// 为函数实现Reader接口，把斐波那契数列写进p中
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	// 斐波那契数列，永远读不完
	if next > 10000 {
		return 0, io.EOF
	}
	// 借助strings.NewReader实现功能
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibinacci()
	printFileContents(f)
}
