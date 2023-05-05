package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime" // 注意v2版本
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"server/proto"
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

	//创建grpc服务
	s := grpc.NewServer()
	// 注册服务
	proto.RegisterGreeterServer(s, &server{})
	//启动服务
	//err = s.Serve(l)
	//if err != nil {
	//	fmt.Printf("error:%v\n", err)
	//	return
	//}
	go func() { //开启goroutine启动grpc server
		err = s.Serve(l)
		if err != nil {
			fmt.Printf("error:%v\n", err)
			return
		}
	}()

	// 创建一个连接到我们刚刚启动的 gRPC 服务器的客户端连接
	// gRPC-Gateway 就是通过它来代理请求（将HTTP请求转为RPC请求）
	conn, err := grpc.DialContext(
		context.Background(),
		"127.0.0.1:8972",
		grpc.WithBlock(), // 阻塞,直到连接成功
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// 注册Greeter
	err = proto.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	// 定义HTTP server配置
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	// 8090端口提供gRPC-Gateway服务
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
