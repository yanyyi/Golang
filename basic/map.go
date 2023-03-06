package main

import "fmt"

func main(){
	//===> 第一种声明方式

	//声明myMap1是一种map类型 key是string, value也是string
	var myMap1 map[string]string
	if myMap1 == nil{
		fmt.Println("myMap1是一个空map!")
	}

	//
	myMap1 = make(map[string]string, 10)
	myMap1["one"]="java"
	myMap1["two"]="C++"
	myMap1["three"]="python"
	fmt.Println("myMap1 = ",myMap1)

	//====> 第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1]="java"
	myMap2[2]="C++"
	myMap2[3]="python"
	fmt.Println("myMap2 = ",myMap2)

	//====> 第三种声明方式
	myMap3 := map[string]string{
		"one": "php",
		"two": "c++",
		"three": "python",
	}
	fmt.Println("myMap3 = ",myMap3)
}
