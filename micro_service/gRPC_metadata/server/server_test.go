package main

import (
	pb "Golang/micro_service/gRPC_metadata/server/proto"
	"context"
	"flag"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	var name = flag.String("name", "golang", "input the name:")
	flag.Parse()
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.Nil(t, err)
	defer conn.Close()
	cli := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//发起普通RPC调用
	//带元数据
	md := metadata.Pairs(
		"token", "app-test-Stephen",
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	var header, trailer metadata.MD
	resp, err := cli.SayHello(ctx,
		&pb.HelloRequest{Name: *name},
		grpc.Header(&header),
		grpc.Trailer(&trailer),
	)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, header.Get("location"), []string{"Shenzhen"})
	assert.Equal(t, resp.Reply, "hello "+(*name))

}
