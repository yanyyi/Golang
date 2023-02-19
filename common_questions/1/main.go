package main

import "fmt"

func main() {
	var arr [5]int = [5]int{0, 1, 2, 3, 4}
	setArrval(&arr)
	fmt.Println(arr)

}

func setArrval(arr *[5]int) {
	arr[2] = 100
	fmt.Println(*arr)
}
