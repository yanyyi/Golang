package main

import (
	"Golang/micro_service/gRPC_TLS/server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	reply := "hello " + request.GetName()
	return &proto.HelloResponse{Reply: reply}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	//加载证书信息
	creds, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")
	if err != nil {
		fmt.Printf("credentials.NewServerTLSFromFile failed, error:%v\n", err)
		return
	}
	s := grpc.NewServer(grpc.Creds(creds)) //创建grpc服务
	// 注册服务
	proto.RegisterGreeterServer(s, &server{})
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
}
