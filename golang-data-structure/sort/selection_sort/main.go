package main

import (
	"fmt"
	"time"
)

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}

		}
	}
}

func main() {
	start := time.Now()
	arr := []int{12, 45, 33, 78, 9, 908, 14, 98, 765}
	selectionSort(arr)
	fmt.Println("选择排序的结果为:", arr)
	end := time.Since(start)
	fmt.Println("选择排序时间为:", end)
}
