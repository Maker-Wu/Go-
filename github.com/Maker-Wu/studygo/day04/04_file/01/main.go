package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 只读方式打开当前目录下的note.md文件
	file, err := os.Open("./Maker-Wu/studygo/day04/04_file/值类型和引用类型.md")		//返回一个*File和一个error
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	// 使用defer注册文件关闭语句
	defer file.Close()

	// 使用Read方法读取数据
	var content []byte
	var tmp = make([]byte, 128)
	// 循环读取文件
	for {
		// 接收字节切片，返回读取的字节数和可能的具体错误
		n, err := file.Read(tmp)
		if err == io.EOF {		//End of File
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Printf("读取了%d字节数据\n", n)
		content = append(content, tmp[:n]...)		//append(content, tmp)最后一次会读取多余的内容
	}
	fmt.Println(string(content))
}
