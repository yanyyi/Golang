package main

import (
	"fmt"
	"time"
)

func main() {
	//多线程问题
	workcount := 2
	for i := 0; i < workcount; i++ {
		go doIt(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("all done!")
}

func doIt(worderID int) {
	fmt.Printf("[%v] is running\n", worderID)
	time.Sleep(3 * time.Second) //模拟goroutine正在执行
	fmt.Printf("[%v] is done\n", worderID)

}
