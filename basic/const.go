package main

import "fmt"

//const来定义不可更改的枚举类型
const(
	//可以在const()添加一个关键字 iota，每行的iota都会累加1，第一行的iota的默认值是0,记住iota只能配合const()进行使用
	BEIJING = iota*10+1  //iota=0
	SHANGHAI        //iota=1
	SHENZHEN        //iota=2
)
const(
	a,b = iota+1, iota+2 //iota逐行累加
	c,d
	e,f 

	g,h = iota*2, iota*3
	i,k
)

func main(){
	//常量(只读属性),不允许修改
	const length int = 10
	fmt.Println("length = ",length) 
	fmt.Println("Beijing = ",BEIJING)
	fmt.Println("Shanghai = ",SHANGHAI)
	fmt.Println("Shenzhen = ",SHENZHEN)
	fmt.Println("a=",a)
	fmt.Println("b=",b)
	fmt.Println("c=",c)
	fmt.Println("d=",d)
	fmt.Println("e=",e)
	fmt.Println("f=",f)
	fmt.Println("g=",g)
	fmt.Println("h=",h)
	fmt.Println("i=",i)
	fmt.Println("k=",k)
}