package main

import (
	"client/proto"
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var name = flag.String("name", "default", "input your name")

func main() {
	flag.Parse()
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	defer conn.Close()
	cli := proto.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := cli.SayHello(ctx, &proto.HelloRequest{
		Name: *name,
	})
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	log.Printf("response: %v\n", resp)
}
