package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"hello_server/pb"
	"io"
	"net"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/status"
)

//grpc server

type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello是我们需要实现的方法
// 这个方法是我们对外提供的服务
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	defer func() {
		trailer := metadata.Pairs(
			"timestamp", strconv.Itoa(int(time.Now().Unix())),
		)
		grpc.SetTrailer(ctx, trailer)
	}()
	//在执行业务逻辑之前要check metadata中是否包含token
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok { //没有元数据我拒接
		return nil, status.Error(codes.Unauthenticated, "无效请求")
	}
	v1 := md.Get("token")
	if len(v1) < 1 || v1[0] != "app-test-20230421" {
		return nil, status.Error(codes.Unauthenticated, "无效token")
	}

	// if v1, ok := md["token"]; ok{
	// 	if len(v1)>0 && v1[0]=="app-test-20230421"{
	// 		//有效的请求
	// 	}
	// }
	reply := "Hello " + in.GetName()
	// 发送数据前发送header
	header := metadata.New(map[string]string{
		"location": "Shenzhen",
	})
	grpc.SendHeader(ctx, header)
	return &pb.HelloResponse{Reply: reply}, nil

}

// LotsOfReplies 返回使用多种语言打招呼
func (s *server) LotsOfReplies(in *pb.HelloRequest, stream pb.Greeter_LotsOfRepliesServer) error {
	words := []string{
		"你好",
		"hello",
		"こんにちは",
		"안녕하세요",
	}

	for _, word := range words {
		data := &pb.HelloResponse{
			Reply: word + in.GetName(),
		}
		// 使用Send方法返回多个数据
		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) LotsOfGreetings(stream pb.Greeter_LotsOfGreetingsServer) error {
	reply := "你好: "
	for {
		//接收客户端发来的流式数据
		res, err := stream.Recv()
		if err == io.EOF {
			//最终统一回复
			return stream.SendAndClose(&pb.HelloResponse{
				Reply: reply,
			})
		}
		if err != nil {
			return err
		}
		reply += res.GetName()
	}
}

func (s *server) BidiHello(stream pb.Greeter_BidiHelloServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		reply := magic(in.GetName())
		//返回流式响应
		if err := stream.Send(&pb.HelloResponse{Reply: reply}); err != nil {
			return err
		}
	}
	return nil
}

// magic 一段价值连城的“人工智能”代码
func magic(s string) string {
	s = strings.ReplaceAll(s, "吗", "")
	s = strings.ReplaceAll(s, "吧", "")
	s = strings.ReplaceAll(s, "你", "我")
	s = strings.ReplaceAll(s, "？", "!")
	s = strings.ReplaceAll(s, "?", "!")
	return s
}

func main() {
	//启动服务
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	s := grpc.NewServer() //创建grpc服务
	//注册服务
	pb.RegisterGreeterServer(s, &server{})
	//启动服务
	err = s.Serve(l)
	if err != nil {
		fmt.Printf("failed to serve,err::%v\n,err", err)
		return
	}

}
