package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Printf("%d ", 2*i+1)
			ch1 <- i
			<-ch2
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {

			<-ch1
			fmt.Printf("%d ", 2*i+2)
			ch2 <- i
		}
	}()

	wg.Wait()

}
