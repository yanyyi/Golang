package main

import (
	"context"
	"fmt"
	"gRPC_interceptor/client/proto"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"log"
	"time"
)

type wrappedStream struct {
	grpc.ClientStream
}

// RecvMsg 重写嵌入的接口类型的RecvMsg方法
func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("Receive a message (Type: %T) at %v\n", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

// SendMsg 重写嵌入的接口类型的SendMsg方法
// 先执行自定义操作
// 再执行原来stream的操作
func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("Send a message (Type: %T) at %v\n", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var credsConfigured bool
	for _, o := range opts {
		_, ok := o.(grpc.PerRPCCredsCallOption)
		if ok {
			credsConfigured = true
			break
		}
	}
	if !credsConfigured {
		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: "some-secret-token",
		})))
	}
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	end := time.Now()
	fmt.Printf("RPC: %s, start time: %s, end time: %s, err: %v\n", method, start.Format("Basic"), end.Format(time.RFC3339), err)
	return err
}

func main() {
	// 加载证书
	creds, err := credentials.NewClientTLSFromFile("certs/server.crt", "")
	if err != nil {
		fmt.Printf("credentials.NewClientTLSFromFile failed, error:%v\n", err)
		return
	}
	conn, err := grpc.Dial("127.0.0.1:8972",
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(unaryInterceptor),
	)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	defer conn.Close()
	cli := proto.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	name := "斯蒂芬库里"
	resp, err := cli.SayHello(ctx, &proto.HelloRequest{
		Name: name,
	})
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	log.Printf("response: %v\n", resp.GetReply())
}
