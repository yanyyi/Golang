package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//用go创建一个形参为空,返回值为空的一个函数
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")

			//return  //退出执行B的函数
			runtime.Goexit() //退出Go程
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
