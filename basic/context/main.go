package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50000*time.Millisecond)
	defer cancel()
	printNum(ctx)

	time.Sleep(2 * time.Second)
}

func printNum(ctx context.Context) {

	for i := 1; i < 1001; i++ {

		time.Sleep(10 * time.Millisecond)
		fmt.Printf("%d ", i)
		select {
		case <-ctx.Done():
			fmt.Println("err-------------------:", ctx.Err())
			return
		}
	}

}
