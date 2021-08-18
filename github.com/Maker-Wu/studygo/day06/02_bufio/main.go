package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("abcdefghijklmn")
	fmt.Printf("%T\n", r) //*strings.Reader
	fmt.Println(r.Len())  //14 初始时，未读长度等于字符串长度
	buf := make([]byte, 5)
	readLen, err := r.Read(buf)
	fmt.Println("读取到的长度:", readLen) //读取到的长度5
	if err != nil {
		fmt.Println("错误:", err)
	}
	fmt.Println(buf)      //[97 98 99 100 101]
	fmt.Println(r.Len())  //9 读取到了5个 剩余未读是14-5
	fmt.Println(r.Size()) //字符串的长度

	scanner := bufio.NewScanner(
		strings.NewReader("ABCDEDF\nHIJKLM"),
	)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
