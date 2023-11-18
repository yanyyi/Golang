package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	a := []int{12, 45, 33, 78, 9, 908, 14, 98, 765}
	sort.Ints(a)
	fmt.Println(a) //[9 12 14 31 33 45 52 53 56 65 77 78 98]
	end := time.Since(start)
	fmt.Println("Go语言标准库Sort排序时间为:", end) //515.1微秒
}
