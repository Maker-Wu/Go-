package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	logFile, err := os.OpenFile("./0802.log",
		os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.SetPrefix("[小王子]")
}

func main() {
	log.Println("这是一条很普通的日志。")
}
