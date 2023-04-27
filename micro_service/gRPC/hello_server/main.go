package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"hello_server/pb"
	"net"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := "hello " + request.GetName()
	return &pb.HelloResponse{Reply: reply}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
}
