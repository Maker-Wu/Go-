package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int


	flag.StringVar(&user, "u", "", "用户名")
	flag.StringVar(&pwd, "pwd", "", "年龄")
	flag.StringVar(&host, "h", "localhost", "主机名")
	flag.IntVar(&port, "P", 3306, "端口号")

	flag.Parse()
	fmt.Printf("user=%v pwd=%v host=%v port=%v\n",
		user, pwd, host, port)
}
