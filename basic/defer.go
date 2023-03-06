package main

import "fmt"

func main(){
	//写入defer关键字
	defer fmt.Println("main end")
	defer fmt.Println("main2 end") //比上面先执行，因为栈的出入关系，后进先出

	fmt.Println("main:hello go 1")	
	fmt.Println("main:hello go 2")	
}