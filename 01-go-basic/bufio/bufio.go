package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		bufio:高效io读写
			buffer缓存
			io: input/output
		将io包下的Reader,Writer对象进行封装，带缓存的包装，提高读写效率
			ReadBytes()
			ReadString()
			ReadLine()
	*/
	fileName := "D:\\桌面\\projects\\goprojects\\Golang仓库\\Golang\\01-go-basic\\bufio\\rpc_intro.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	defer file.Close()

	// 创建reader对象
	b1 := bufio.NewReader(file)
	p := make([]byte, 1024)
	n1, err := b1.Read(p)
	fmt.Println(n1)
	fmt.Println(string(p[:n1]))
}
