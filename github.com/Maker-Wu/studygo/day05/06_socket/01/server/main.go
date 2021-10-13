package main

import (
	"bufio"
	"fmt"
	"net"
)

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() //关闭连接
	for {
		//fmt.Println(conn.RemoteAddr().String()) //客户端ip地址
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) //读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据", recvStr)
		conn.Write([]byte(recvStr))
	}
}

func main() {
	// 在本地监听20000端口
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept() // 等待客户端建立链接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
