package main

import (
	"Golang/micro_service/gRPC_metadata/server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
	"strconv"
	"time"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	defer func() {
		trailer := metadata.Pairs(
			"timestamp", strconv.Itoa(int(time.Now().Unix())),
		)
		grpc.SetTrailer(ctx, trailer)
	}()
	//在执行业务逻辑之前要check metadata中是否包含token
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok { //没有元数据则拒接
		return nil, status.Error(codes.Unauthenticated, "无效请求")
	}
	valList := md.Get("token")
	if len(valList) < 1 || valList[0] != "app-test-Stephen" {
		// 无效的请求
		return nil, status.Error(codes.Unauthenticated, "无效token")
	}

	reply := "hello " + request.GetName()
	//发送数据前发送header
	header := metadata.New(map[string]string{
		"location": "Shenzhen",
	})
	grpc.SendHeader(ctx, header)
	return &proto.HelloResponse{Reply: reply}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
}
