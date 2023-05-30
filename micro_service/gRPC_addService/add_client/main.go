package main

import (
	"add_client/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	// 连接rpc server
	conn, err := grpc.Dial("127.0.0.1:8973", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed, error:%v\n", err)
		return
	}
	defer conn.Close()
	// 创建rpc client端
	client := proto.NewCalcServiceClient(conn)
	// 发起RPC调用
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.Add(ctx, &proto.AddRequest{
		X: 217,
		Y: 283,
	})
	if err != nil {
		log.Fatalf("client.Add failed, error:%v\n", err)
		return
	}
	// 打印结果
	log.Printf("相加的结果是: %v \n", resp.GetResult())
}
