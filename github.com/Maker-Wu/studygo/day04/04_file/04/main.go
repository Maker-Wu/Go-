package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeFile() {
	file, err := os.OpenFile("./read.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "Hello 中兴\n"
	file.Write([]byte(str))  //写入字节切片数据
	file.WriteString("Hello 伍胜强")	//直接写入字符串数据
}

//bufio写入文件
func bufioWriteFile() {
	file, err := os.OpenFile("./read.txt", os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello沙河\n")
	}
	writer.Flush()		//将缓冲中的内容写入文件
}

//ioutil写入文件
func ioutilWriteFile() {
	str := "hello 长沙"
	err := ioutil.WriteFile("./read.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func main() {
	writeFile()
	bufioWriteFile()
	// ioutilWriteFile()
}
