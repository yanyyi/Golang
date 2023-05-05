package main

import (
	"context"
	"fmt"
	"gRPC_stream/server/proto"
	"google.golang.org/grpc"

	"io"
	"net"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	reply := "Hello " + request.GetName() + ". 你今年 " + fmt.Sprintf("%d", request.GetAge()) + " 岁了."
	return &proto.HelloResponse{Reply: reply}, nil
}

func (s *server) LotsOfReplies(request *proto.HelloRequest, stream proto.Greeter_LotsOfRepliesServer) error {
	words := []string{
		"你好",
		"hello",
		"こんにちは",
		"안녕하세요",
	}
	for _, word := range words {
		data := &proto.HelloResponse{
			Reply: word + request.GetName() + ". You are " + fmt.Sprintf("%d", request.GetAge()) + " years old.",
		}
		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return nil

}

func (s *server) LotofGreetings(stream proto.Greeter_LotofGreetingsServer) error {
	reply := "你好: "
	for {
		//接收客户端发来的消息
		resp, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.HelloResponse{
				Reply: reply,
			})

		}
		if err != nil {
			return err
		}
		reply += resp.GetName() + ". 你今年" + fmt.Sprintf("%d", resp.GetAge()) + "岁了。\n"
	}
	return nil
}

func main() {
	// 启动服务
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen, error:%v\n", err)
		return
	}
	s := grpc.NewServer()
	// 注册服务
	proto.RegisterGreeterServer(s, &server{})
	//启动服务
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("failed to server,error:%v\n", err)
		return
	}

}
