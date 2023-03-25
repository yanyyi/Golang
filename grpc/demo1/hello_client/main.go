package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc/hello_server/proto"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("连接错误:%v", err)
		return
	}
	defer conn.Close()

	// 建立连接
	client := pb.NewSayHelloClient(conn)

	// 执行rpc调用(这个方法在服务器端来实现并返回结果)
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "I am learning gRPC"})
	fmt.Println(resp.GetResponseMsg())
}
