package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//生成10到100之间的数
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(91) + 10
	fmt.Println("num=", num)
}
