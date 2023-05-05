package main

import (
	"client/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

var name = flag.String("name", "default", "input the name: ")

func main() {
	flag.Parse()
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	defer conn.Close()
	cli := proto.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := cli.SayHello(ctx, &proto.HelloRequest{
		Name: *name,
	})
	if err != nil {
		//收到带detail的error
		s := status.Convert(err)
		for _, d := range s.Details() {
			switch info := d.(type) {
			case *errdetails.QuotaFailure:
				fmt.Printf("QuoraFailure:%v\n", info.String())
			default:
				fmt.Printf("unexpected type: %v\n", info)
			}
		}
		fmt.Printf("c.SayHello() failed, error:%v\n", err)
		return
	}
	log.Printf("response: %v\n", resp)
}
