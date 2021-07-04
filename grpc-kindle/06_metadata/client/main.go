package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ygongq/grpc-demo/grpc-kindle/06_metadata/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("localhost:5005", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("client create conn err: ", err)
	}
	defer conn.Close()

	client := proto.NewUnaryInterceptorClient(conn)

	// 创建元数据
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"kn", "vn",
	)

	// 基于元数据创建新的上下文
	mdCtx := metadata.NewOutgoingContext(context.Background(), md)

	// 往已有的上下文中追加元数据
	mdCtx = metadata.AppendToOutgoingContext(mdCtx, "k1", "v1")

	// 用来接收返回的头信息和trailer
	var header, trailer metadata.MD

	rsp, err := client.GetValue(mdCtx, &proto.UnaryInterceptorRequest{ID: "101"}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-------------------------")
	for k, v := range header {
		fmt.Println(k, "==>", v)
	}
	fmt.Println("-------------------------")
	for _, v := range trailer {
		fmt.Println(v)
	}
	fmt.Println("-------------------------")

	fmt.Println(rsp)
}
