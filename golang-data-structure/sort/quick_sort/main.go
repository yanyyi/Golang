package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0] // 以第一个元素作为基准值
	var left, right []int
	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}

func main() {
	//start := time.Now()
	arr := []int{12, 45, 33, 78, 9, 908, 14, 98, 765}
	arr = quickSort(arr)
	fmt.Println("快速排序的结果为:", arr)
	//end := time.Since(start)
	//fmt.Println("快速排序时间为:", end)
}
