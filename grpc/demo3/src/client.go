package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

const PORT = "9091"

func main() {
	conn, err := grpc.Dial("localhost:"+PORT, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	defer conn.Close()

	client := NewPubsubServiceClient(conn)

	_, _ = client.Publish(context.Background(), &String{Value: "golang: hello Go"})
	_, _ = client.Publish(context.Background(), &String{Value: "docker: hello Docker"})

}
