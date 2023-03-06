package main

import "fmt"

func printArray(myArray []int){
	//下划线表示匿名变量
	for _, value := range myArray{
		fmt.Println("value = ",value)
	}

	myArray[0] = 100  //对传进来的myArray是产生影响的，因为动态数组本身是指向数组内存的指针
}

func main(){
	//固定长度的数组
	myArray := []int{1,2,3,4} //动态数组，切片 slice
	fmt.Printf("myArray type is %T\n",myArray)

	printArray(myArray)

	for _, value := range myArray{
		fmt.Println("value = ",value)
	}
}
