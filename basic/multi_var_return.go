package main

import "fmt"

func foo1(a string, b int) int{
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	c:=100
	return c
}
//返回多个值，匿名的
func foo2(a string, b int)(int, int){
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	return 666, 777
}

//返回多个返回值，有形参名称的
func foo3(a string, b int)(r1 int, r2 int){
	fmt.Println("---foo3---")
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}

func foo4(a string, b int)(r1,r2 int){
	fmt.Println("---foo4---")
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	//给有名称的返回值变量赋值
	r1 = 999999999
	r2 = 777777777
	return
}

func main(){
	c := foo1("abc", 555)
	fmt.Println("c = ",c)

	ret1, ret2 := foo2("hahaha",999)
	fmt.Println("ret1 = ",ret1, "ret2 = ",ret2)

	ret1, ret2 = foo3("foo3",333)
	fmt.Println("ret1 = ",ret1, "ret2 = ",ret2)

	ret1, ret2 = foo4("foo4",333)
	fmt.Println("ret1 = ",ret1, "ret2 = ",ret2)
}