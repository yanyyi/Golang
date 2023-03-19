package main

import (
	"fmt"
	"sync"
)

func print1(wg *sync.WaitGroup, num1 chan struct{}, num2 chan struct{}) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		<-num1
		fmt.Printf("%d", 3*i+1)
		num2 <- struct{}{}
	}
	<-num1
}

func print2(wg *sync.WaitGroup, num2 chan struct{}, num3 chan struct{}) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		<-num2
		fmt.Printf("%d", 3*i+2)
		num3 <- struct{}{}
	}

}

func print3(wg *sync.WaitGroup, num3 chan struct{}, num1 chan struct{}) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		<-num3
		fmt.Printf("%d", 3*i+3)
		num1 <- struct{}{}
	}

}

func main() {
	var wg sync.WaitGroup
	num1 := make(chan struct{})
	num2 := make(chan struct{})
	num3 := make(chan struct{})
	wg.Add(3)
	go print1(&wg, num1, num2)
	go print2(&wg, num2, num3)
	go print3(&wg, num3, num1)
	num1 <- struct{}{}

	wg.Wait()
}
