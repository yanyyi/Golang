package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "hello_client/proto"
	"log"
	"time"
)

var name string = "斯蒂芬库里"

func main() {
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	defer conn.Close()
	//创建客户端
	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	log.Printf("resp:%v\n", resp.GetReply())

}
