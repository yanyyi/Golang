package main

import (
	"fmt"
	"time"
)

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	start := time.Now()
	a := []int{12, 45, 33, 78, 9, 908, 14, 98, 765}
	bubbleSort(a)
	fmt.Println(a) //
	end := time.Since(start)
	fmt.Println("Bubble Sort排序时间为:", end) //
}
