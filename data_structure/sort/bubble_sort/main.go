package main

import (
	"fmt"
	"time"
)

func bubbleSort(a []int) {
	llen := len(a)

	for i := 0; i < llen-1; i++ {
		for j := llen - 1; j > i; j-- {
			if a[j] < a[j-1] {
				temp := a[j]
				a[j] = a[j-1]
				a[j-1] = temp
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
