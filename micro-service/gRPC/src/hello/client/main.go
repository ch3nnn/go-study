package main

import (
	pb "go-study/micro-service/gRPC/src/proto/hello" // 引入proto包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"log"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)
	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		grpclog.Fatalln(err)
	}

	log.Printf("Greeting: %s", res.GetMessage())
}
