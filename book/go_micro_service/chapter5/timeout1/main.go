package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	rand.Seed(time.Now().UnixNano())
	no := rand.Intn(6)
	no = no * 1000
	du := time.Duration(int32(no)) * time.Millisecond
	fmt.Println("timeout duration is: ", du)
	wg.Done()
	if isTiemout(&wg, du) {
		fmt.Println("Time out!")
	} else {
		fmt.Println("Not time out!")
	}
}

func isTiemout(wg *sync.WaitGroup, du time.Duration) bool {
	ch1 := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		defer close(ch1)
		wg.Wait()
	}()
	select {
	case <-ch1:
		return false
	case <-time.After(du):
		return true
	}
}
