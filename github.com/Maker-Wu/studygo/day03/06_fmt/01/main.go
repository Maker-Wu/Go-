package main

import (
	"fmt"
	"os"
)

func main() {
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./Maker-Wu/studygo/day03/06_fmt/01/text.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer fileObj.Close()
	name := "沙河小王子"
	fmt.Fprintf(fileObj, "往文件中写如下信息：%s", name)
}
