package main

import (
	"context"
	"flag"
	"fmt"
	"gRPC_stream/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

type members struct {
	Name string
	Age  int32
}

var name = flag.String("name", "DefaultUser", "通过-name告诉server端你是谁")
var age = flag.Int64("age", 0, "通过-age告诉server端你的年龄")

func main() {
	flag.Parse() //解析命令行参数
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc.Dial() failed, error:%v\n", err)
		return
	}
	defer conn.Close()
	//创建客户端
	c := proto.NewGreeterClient(conn) //使用生成的Go代码

	//调用流式RPC
	//callLotsOfReplies(c)
	sendLotOfGreetings(c)

}

func callLotsOfReplies(c proto.GreeterClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.LotsOfReplies(ctx, &proto.HelloRequest{Name: *name, Age: int32(*age)})
	if err != nil {
		fmt.Printf("c.LotsOfReplies failed(), error:%v\n", err)
		return
	}
	// 依次从流式响应中读取返回的响应数据
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("stream.Recv failed, error:%v\n", err)
			return
		}
		log.Printf("recv: %v\n", res.GetReply())
	}
}

func sendLotOfGreetings(c proto.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 客户端流式发送请求消息
	stream, err := c.LotofGreetings(ctx)
	if err != nil {
		fmt.Printf("c.LotofGreetings() failed, error:%v\n", err)
		return
	}

	member1 := &members{
		Name: "勒布朗詹姆斯",
		Age:  39,
	}
	member2 := &members{
		Name: "斯蒂芬库里",
		Age:  35,
	}
	member3 := &members{
		Name: "凯文杜兰特",
		Age:  35,
	}
	players := []*members{member1, member2, member3}

	for _, player := range players {
		stream.Send(&proto.HelloRequest{
			Name: player.Name,
			Age:  player.Age})
	}
	//流式发送结束之后要关闭流
	resp, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Printf("stream.CloseAndRecv() failed, error:%v\n", err)
		return
	}
	fmt.Println("resp:")
	fmt.Printf("%v\n", resp.GetReply())
}
