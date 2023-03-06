package main

import "fmt"

//interface{}是万能数据类型
func myFunc(arg interface{}){
	fmt.Println("myFunc is called..")
	fmt.Println(arg)

	//interface{} 该如何区分 此时引用类型的底层数据类型到底是什么？

	//给 interface{} 提供"断言"的机制
	value, ok := arg.(string)
	if !ok{
		fmt.Println("arg is not a string")
	}else{
		fmt.Println("arg is string type, value = ",value)
		fmt.Printf("value type is %T\n",value)
	}

}

type Book struct{
	author string
	language string
}

func main(){
	book := Book{"Mike","Golang"}
	myFunc(book)
	myFunc("abc")
	myFunc(3.14)
}
