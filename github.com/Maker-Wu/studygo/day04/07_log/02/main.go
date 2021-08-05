package main

import (
	"log"
)

func main() {
	// 日志输出配置
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("这是一条很普通的日志。")
	// 配置日志前缀
	log.SetPrefix("[小王子]")
	log.Println("这是一条很普通的日志。")
}

