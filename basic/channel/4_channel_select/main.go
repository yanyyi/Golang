package main

import "fmt"

func main() {
	pipeline1 := make(chan string, 1)
	pipeline2 := make(chan string, 2)
	pipeline1 <- "hello"
	pipeline2 <- "world"
	//select随机选取一个执行
	select {
	case data1 := <-pipeline1:
		fmt.Println("pipeline1 接收的数据:", data1)
	case data2 := <-pipeline2:
		fmt.Println("pipeline2 接收的数据:", data2)
	default: //写default避免死锁
		fmt.Println("两个信道均无数据")

	}
}
