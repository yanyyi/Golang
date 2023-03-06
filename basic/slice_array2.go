package main

import "fmt"



func main(){
	//----------  方式1  ----------//
	//声明slice1是一个切片、默认值是1、2、3，长度len是3
	// slice1 := []int{1,2,3}

	//----------  方式2  ----------//
	//声明slice1是一个切片，但是并没有给slice1分配空间
	var slice1 []int
	//通过make给slice1开辟内存空间，默认值是0
	// slice1 = make([]int, 3)
	// slice1[0] = 129

	//----------  方式3  ----------//
	//声明slice1是一个切片，同时给slice1分配空间，3个空间，初始值是0
	// var slice1 []int = make([]int, 3)
	// slice1[0] = 111

	//----------  方式4  ----------//
	//声明slice1是一个切片，同时给slice1分配空间,3个空间，初始值是0，通过:=推导出slice是一个切片
	// slice1 := make([]int, 3)
	// slice1[0] = 999

	fmt.Printf("len = %d, slice = %v\n",len(slice1), slice1)

	if slice1 == nil{
		fmt.Println("slice是一个空切片")
	}else{
		fmt.Println("slice是由空间的")
	}
}
