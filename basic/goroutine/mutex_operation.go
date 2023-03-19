package main

import (
	"fmt"
	"sync"
)

var total struct {
	sync.Mutex
	num int
}

func sum(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		total.Lock()
		total.num += 1
		fmt.Println(total.num)
		total.Unlock()
	}

}

func substract(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		total.Lock()
		total.num -= 1
		fmt.Println(total.num)
		total.Unlock()
	}

}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go sum(&wg)       //执行完
	go substract(&wg) //再执行

	wg.Wait()
}
