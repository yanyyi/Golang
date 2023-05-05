package main

import (
	"Golang/micro_service/gRPC_gateway/client/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	defer conn.Close()
	cli := proto.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	name := "斯蒂芬库里"
	resp, err := cli.SayHello(ctx, &proto.HelloRequest{
		Name: name,
	})
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	log.Printf("response: %v\n", resp)
}
