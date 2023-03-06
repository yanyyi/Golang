package main

import (
	"fmt"
)

func fabonacii(c, quit chan int){
	x, y, t:= 1, 1, 0
	
	for{
		select{
		case c<-x:
			//如果c可写,则该case就会进来
			t = y
			y = x + y
			x = t
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}

func main(){
	c := make(chan int)
	quit  := make(chan int)

	//sub go
	go func(){
		// quit <- 0 //如果quit直接先给0，则直接输出quit
		for i:= 0 ; i<10; i++{
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	//main go
	fabonacii(c, quit)
	
}

