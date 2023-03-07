package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go writeChan(c, 666)
	time.Sleep(1 * time.Second)
}

func writeChan(c chan int, x int) {
	fmt.Println(x) //打印的是这一行
	c <- x         //函数进入阻塞态
	close(c)
	fmt.Println(x) //不会执行
}
