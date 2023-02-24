package main

import "fmt"

func main() {
	c := make(chan string, 3)
	words := []string{"I", "love", "China", "Shenzhen"}
	go func() {
		for _, word := range words {
			c <- word
		}
		close(c)
	}()

	for data := range c {
		fmt.Println(data)
	}
	fmt.Println("-------------Main Finished-----------------")
}
