package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/Maker-Wu/studygo/day05/06_socket/02/proto"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	num := 0
	for {
		msg, err := proto.Decode(reader)
		fmt.Println(err)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
		if num == 25 {
			break
		}
		num++
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	num := 0
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		num++
		fmt.Printf("开启第%d个协程", num)
		go process(conn)
	}
}
