package main

import (
	"fmt"
)

type Reader interface{
	ReadBook()
}

type Writer interface{
	WriteBook()
}

//具体类型
type Book struct{
}

func (this *Book) ReadBook(){
	fmt.Println("Read a Book")
}

func (this *Book) WriteBook(){
	fmt.Println("Write a Book")
}

func main(){
	
	//b: pair<type:Book value:book{}的地址>
	b:= &Book{}

	//r pair<type;, value>
	var r Reader
	//r :pair<type:Book, value:book{}的地址>
	r = b

	r.ReadBook()

	var w Writer
	w = r.(Writer) // 此处的断言为什么会成功?因为w r具体的type是一致的
	w.WriteBook()
}
