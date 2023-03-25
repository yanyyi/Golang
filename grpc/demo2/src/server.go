package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
	reply := &String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func (p *HelloServiceImpl) mustEmbedUnimplementedHelloServiceServer() {}

func main() {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	grpcServer.Serve(lis)
}
