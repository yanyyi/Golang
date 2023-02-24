package main

import (
	"fmt"
	"net"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	defer listener.Close()
	//阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error:%v\n", err)
			continue
		}
		//接收用户的请求
		buf := make([]byte, 1024) //1024大小的缓冲区
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("error:%v\n", err)
			continue
		}
		fmt.Println("buf = ", string(buf[:n]))
	}

}
