package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func bufferWrite() {
	fmt.Println("===========以下通过Write把swift写入Learning缓冲器尾部=========")
	newBytes := []byte("swift")
	// 创建一个内容为Learning的缓冲器
	buf := bytes.NewBuffer([]byte("Learning"))
	fmt.Println(buf)          //Learning
	fmt.Println(buf.String()) //Learning
	//将newBytes这个slice写到buf的尾部
	buf.Write(newBytes)
	fmt.Println(buf.String()) //Learningswift

	buf.WriteString(" and go")
	fmt.Println(buf.String()) //Learningswift and go
}

func IntToBytes(n int) []byte {
	x := int32(n)
	// 创建缓冲器
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

func main() {
	// 整形转换成字节
	var n int = 1000
	intToBytes := IntToBytes(n)
	fmt.Println(intToBytes) //[232 3 0 0]

	buf1 := bytes.NewBufferString("swift")
	buf2 := bytes.NewBuffer([]byte("swift"))
	buf3 := bytes.NewBuffer([]byte{'s', 'w', 'i', 'f', 't'})

	fmt.Println("buf1:", buf1) //buf1: swift
	fmt.Println("buf2:", buf2) //buf2: swift
	fmt.Println("buf3:", buf3) //buf3: swift
	fmt.Println("===========以下创建空的缓冲器等效=========")
	buf4 := bytes.NewBufferString("")
	buf5 := bytes.NewBuffer([]byte{})
	fmt.Println("buf4:", buf4) //buf4:
	fmt.Println("buf5:", buf5) //buf5:
	bufferWrite()
}
