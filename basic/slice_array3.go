package main

import "fmt"



func main(){
	var numbers = make([]int, 3, 5)//开辟长度为3，容量为5,不指定时容量==长度
	fmt.Printf("len = %d, cap = %d, slice = %v\n",len(numbers), cap(numbers), numbers)

	//向切片追加一个元素2，numbers len=4, [0,0,0,2]， cap=5
	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n",len(numbers), cap(numbers), numbers)

	//向切片追加一个元素5，numbers len=5, [0,0,0,2,5]， cap=5
	numbers = append(numbers, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n",len(numbers), cap(numbers), numbers)

	//向切片追加超过cap上限的一个元素100，numbers len=6, [0,0,0,2,5,100]， cap=5*2=10
	numbers = append(numbers, 100)
	fmt.Printf("len = %d, cap = %d, slice = %v\n",len(numbers), cap(numbers), numbers)
}
