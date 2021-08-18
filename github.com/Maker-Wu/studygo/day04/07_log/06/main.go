package main

import (
	"fmt"
	"path"
	"runtime"
)

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file) // // Base函数返回路径的最后一个元素
	return
}

func main() {
	fmt.Println(getInfo(0))
	fmt.Println(getInfo(1))
}
