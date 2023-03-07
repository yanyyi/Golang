package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 0

	var juice int
	// myMap := make(map[int]int)
	fmt.Scanln(&a, &b)
	myMap := make(map[int]int)

	for i := 0; i < a; i++ {
		fmt.Scan(&juice)
		myMap[juice-1]++
	}
	fmt.Println("len(myMap)=", len(myMap))
	for i := 0; i < len(myMap); i++ {
		fmt.Print(" ", myMap[i])
	}
	//sum := 0
	//for i := 0; i < len(myMap); i++ {
	//	sum += (myMap[i] + 1) / 2
	//}
	//fmt.Println(sum)

}
