package main

import (
	"Golang/micro_service/gRPC_metadata/client/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

var name = flag.String("name", "golang", "input the name:")

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

	//发起普通RPC调用
	//带元数据
	md := metadata.Pairs(
		"token", "app-test-Stephen",
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	//声明两变量
	var header, trailer metadata.MD
	resp, err := cli.SayHello(
		ctx,
		&proto.HelloRequest{Name: *name},
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	//拿到响应之前可以获取header
	fmt.Printf("header: %v\n", header)
	//拿到了RPC响应
	log.Printf("response: %v\n", resp.GetReply())
	//拿到响应数据后可以获取trailer
	fmt.Printf("trailer: %v\n", trailer)
}
