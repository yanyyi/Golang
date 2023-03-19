package main

import "fmt"

func main() {
	num := make(chan int, 2)

	go func() {
		fmt.Println("Hello World")
		num <- 1
	}()
	<-num
}
