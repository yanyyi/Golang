package main

import "fmt"

func linearSearch(arr []int, target int) int {
	i := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return i
}

func main() {
	arr := []int{7, 10, 13, 16, 19, 29, 32, 33, 37, 41, 43}
	target := 33
	i := linearSearch(arr, target)
	fmt.Println("i=", i)
}
