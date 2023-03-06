package main

import "fmt"

type myint int
//定义一个结构体
type Book struct{
	title string
	author string
}

func changeBook1(book Book){
	//传递一个book的副本
	book.author = "Lisi" 
}

func changeBook2(book *Book){
	//指针传递
	book.author = "Lisi" 
}

func main(){
	// var a myint = 10
	// fmt.Println("a = ",a)
	// fmt.Printf("type of a = %T\n",a)
	var book1 Book
	book1.title = "Golang"
	book1.author = "Zhangsan"

	fmt.Printf("%v\n",book1)

	changeBook1(book1)
	fmt.Printf("%v\n",book1)

	changeBook2(&book1)
	fmt.Printf("%v\n",book1)
}
