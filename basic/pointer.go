package main

import "fmt"

func swap(pa *int, pb *int){
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}


func main(){
	var a int = 10
	var b int = 20

	swap(&a, &b)
	//swap
	fmt.Println("a = ",a, "b = ",b)

	var p *int
	p = &a 
	fmt.Println("p = ",p, "&a = ",&a)

	var pp  **int //二级指针
	pp = &p
	fmt.Println("&p = ",&p)
	fmt.Println("pp = ", pp)
	fmt.Println("*pp = ",*pp)
	fmt.Println("**pp = ",**pp)
}