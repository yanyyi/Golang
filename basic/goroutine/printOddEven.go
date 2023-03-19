package main

import (
	"fmt"
	"sync"
	"time"
)

func printOdd(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("奇数:", 2*i+1)
		time.Sleep(50 * time.Millisecond)
	}
}

func printEven(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("偶数:", 2*i)
		time.Sleep(50 * time.Millisecond)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printOdd(&wg)
	go printEven(&wg)
	wg.Wait()

}
