package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ygongq/grpc-demo/grpc-kindle/03_interceptor/unary/server/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5003", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("client create conn err: ", err)
	}
	defer conn.Close()

	client := proto.NewUnaryInterceptorClient(conn)

	rsp, err := client.GetValue(context.Background(), &proto.UnaryInterceptorRequest{ID: "101"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rsp)
}
