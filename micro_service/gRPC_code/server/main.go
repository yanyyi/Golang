package main

import (
	"context"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"server/proto"
	"sync"
)

type server struct {
	proto.UnimplementedGreeterServer
	count map[string]int
	mu    sync.Mutex
}

func (s *server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	name := request.GetName()
	s.count[name]++
	if s.count[name] > 1 {
		// 返回请求次数限制的错误
		st := status.New(codes.ResourceExhausted, "请求限制(request limit)......")
		//添加错误详情信息,需要接收返回的status
		ds, err := st.WithDetails(
			&errdetails.QuotaFailure{
				Violations: []*errdetails.QuotaFailure_Violation{
					{
						Subject:     fmt.Sprintf("name:%s", name),
						Description: "每个name只能调用一次SayHello",
					},
				},
			},
		)
		if err != nil {
			fmt.Printf("error:%v\n", err)
			return nil, st.Err()
		}
		return nil, ds.Err()
	}
	//正常执行
	reply := "hello " + request.GetName()
	return &proto.HelloResponse{Reply: reply}, nil

}

func main() {
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{count: make(map[string]int)})
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
}
