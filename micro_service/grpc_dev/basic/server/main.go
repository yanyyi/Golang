package main

import (
	"context"
	"log"
	"net"
	"server/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	reply := "hello, " + req.Name
	return &proto.HelloResponse{
		Reply: reply,
	}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		log.Printf("net.Listen() failed,err:%v\n", err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	if err = s.Serve(l); err != nil {
		log.Printf("s.Serve(l) failed", err)
		return
	}
}
