package main

import (
	"Golang/micro_service/gRPC_TLS/client/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)

func main() {
	//加载证书
	creds, err := credentials.NewClientTLSFromFile("certs/server.crt", "")
	if err != nil {
		fmt.Printf("credentials.NewClientTLSFromFile failed, error:%v\n", err)
		return
	}
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(creds))
	//conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials())) //没有证书无法运行了
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	defer conn.Close()
	cli := proto.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	name := "Golang"
	resp, err := cli.SayHello(ctx, &proto.HelloRequest{
		Name: name,
	})
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	log.Printf("response: %v\n", resp)
}
