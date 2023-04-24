package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	pb "hello_client/proto"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var name string = "斯蒂芬库里"

func main() {
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	defer conn.Close()
	//创建客户端
	c := pb.NewGreeterClient(conn)
	// 调用RPC方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 普通RPC调用
	// 带元数据
	md := metadata.Pairs(
		"token", "app-test-20230421",
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	// 声明两个变量
	var header, trailer metadata.MD
	resp, err := c.SayHello(ctx,
		&pb.HelloRequest{Name: name},
		grpc.Header(&header),
		grpc.Trailer(&trailer))

	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	// 拿到响应数据之前可以获取header
	fmt.Printf("header:%v\n", header)

	//拿到了RPC响应
	log.Printf("resp:%v\n", resp.GetReply())

	// 拿到响应数据之前可以获取header
	fmt.Printf("trailer:%v\n", trailer)
	// 调用服务端流式的RPC
	//callLotsOfReplies(c)
	//callLotsOfGreetings(c)
	//runBidiHello(c)
}

func callLotsOfReplies(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.LotsOfReplies(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Println(err)
		return
	}
	//依次从流式响应中读取返回的响应数据
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("stream.Recv failed, err:%v\n", err)
			return
		}
		log.Printf("recv :%v\n", res.GetReply())
	}

}

func callLotsOfGreetings(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 客户端要流式的发送请求消息
	stream, err := c.LotsOfGreetings(ctx)
	if err != nil {
		fmt.Printf("c.LotsOfGreetings failed, error:%v\n", err)
		return
	}
	names := []string{"詹姆斯", "乔丹", "科比", "库里"}
	for _, name := range names {
		//发送流式数据
		err = stream.Send(&pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("c.LotsOfGreetings stream.Send(%v) failed, err:%v", name, err)
			return
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("c.LotsOfGreetings failed: error:%v\n", err)
		return
	}
	log.Printf("got reply: %v", res.GetReply())
}

func runBidiHello(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	// 双向流模式
	stream, err := c.BidiHello(ctx)
	if err != nil {
		log.Fatalf("c.BidiHello failed, err: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			// 接收服务端返回的响应
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("c.BidiHello stream.Recv() failed, err: %v", err)
			}
			fmt.Printf("AI：%s\n", in.GetReply())
		}
	}()
	// 从标准输入获取用户输入
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	for {
		cmd, _ := reader.ReadString('\n') // 读到换行
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		if strings.ToUpper(cmd) == "QUIT" {
			break
		}
		// 将获取到的数据发送至服务端
		if err := stream.Send(&pb.HelloRequest{Name: cmd}); err != nil {
			log.Fatalf("c.BidiHello stream.Send(%v) failed: %v", cmd, err)
		}
	}
	stream.CloseSend()
	<-waitc
}
