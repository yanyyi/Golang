package main

import (
	"fmt"
	"time"
)

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		var j int
		for j = i - 1; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
	return arr
}

func main() {
	start := time.Now()
	a := []int{12, 45, 33, 78, 9, 908, 14, 98, 765}

	insertionSort(a)
	fmt.Println(a) //
	end := time.Since(start)
	fmt.Println("Insertion Sort排序时间为:", end) //
}
