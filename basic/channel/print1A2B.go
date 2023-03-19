package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// 创建两个 channel 用于数据传输
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 启动第一个协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 26; i++ {
			// 循环打印 A~Z
			fmt.Printf("%c", 'A'+i)
			ch1 <- i
			// 等待 ch2 发送数据
			<-ch2
		}
	}()

	// 启动第二个协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 26; i++ {
			// 等待 ch1 发送数据
			<-ch1
			// 循环打印 1~26
			fmt.Printf("%d", i+1)
			// 发送数据给 ch2
			ch2 <- i
		}
	}()

	// 等待两个协程执行完毕
	wg.Wait()
}
