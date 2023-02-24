package main

import "fmt"

func main() {
	//定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束")

		fmt.Println("goroutine 正在运行...")
		c <- 666 //将666写入c中
	}()

	num := <-c //从c中接收数据,并赋值给num
	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")

}
