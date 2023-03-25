package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

const PORT = "9091"

func main() {
	conn, err := grpc.Dial("localhost:"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := NewHelloServiceClient(conn)
	reply, _ := client.Hello(context.Background(), &String{Value: "Hello SZU"})

	fmt.Println(reply.GetValue())
}
