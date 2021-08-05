package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "<New> ",
		log.Ldate | log.Lmicroseconds | log.Llongfile)
	logger.Println("这是一个自定义的logger记录的日志")
}

