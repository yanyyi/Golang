package main


import "fmt"

//声明全局变量
var gA  int = 100
var gB = 200



func main(){
	var a int
	fmt.Println("a = ",a)
	fmt.Printf("type of a = %T\n",a)
	var b int = 100
	fmt.Println("b = ",b)
	fmt.Printf("type of b = %T\n",b)
	var c = 200.06
	fmt.Println("c = ",c)
	fmt.Printf("type of c = %T\n",c)
	var d = "abcd"
	fmt.Println("d = ",d)
	fmt.Printf("type of d = %T\n",d)

	e := 100  //:=只能用在函数体内
	fmt.Println("e = ",e)
	fmt.Printf("type of e = %T\n",e)

	f := "abcd"
	fmt.Println("f = ",f)
	fmt.Printf("type of f = %T\n",f) 

	fmt.Println("gA = ",gA, "gB = ",gB)

	//声明多个变量
	var xx, yy int = 100,200
	fmt.Println("xx = ",xx,"yy = ",yy)
	var kk, ll = 3.14159, "ABCDEFG"
	fmt.Println("kk = ",kk, "ll = ",ll)

	//多行的多变量声明
	var(
		vv int = 200
		jj bool = true

	)
	fmt.Println("vv = ",vv, "bool = ",jj)
	
}
