package main

import "fmt"



func main(){
	s := []int{1, 2, 3}
	s1 := s[:2]
	fmt.Println("s1 = ",s1) //s1=[1 2]
	s[0] = 100 //修改s，则s1也会修改
	fmt.Println("s1 = ",s1) //s1=[100 2]

	s2 := make([]int, 3)
	copy(s2, s)
	fmt.Println("s2 = ",s2) //s2=[100 2 3]
	s[0] = 99 //修改s，由于s2指向了另一个内存空间,故s2不会修改
	fmt.Println("s2 = ",s2) //s2=[100 2 3]
}
