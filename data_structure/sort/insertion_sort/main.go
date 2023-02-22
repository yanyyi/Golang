package main

import (
	"fmt"
	"time"
)

func insertSort(arr []int) {

	for i := 0; i < len(arr); i++ {
		// 从后面第二个开始，取出找到合适的位置（插入）
		for j := i; j > 0; j-- {
			// 比较
			if arr[j] < arr[j-1] {
				// 交换
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}
}

func main() {
	start := time.Now()
	a := []int{12, 45, 33, 78, 9, 908, 14, 98, 765}
	insertSort(a)
	fmt.Println(a) //
	end := time.Since(start)
	fmt.Println("Insertion Sort排序时间为:", end) //
}
