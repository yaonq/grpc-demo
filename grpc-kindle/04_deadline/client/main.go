package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ygongq/grpc-demo/grpc-kindle/03_interceptor/unary/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:5004", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("client create conn err: ", err)
	}
	defer conn.Close()

	client := proto.NewUnaryInterceptorClient(conn)

	ctx, cFunc := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cFunc()

	rsp, err := client.GetValue(ctx, &proto.UnaryInterceptorRequest{ID: "101"})
	if err != nil {
		got := status.Code(err)
		log.Fatal(got)
	}

	fmt.Println(rsp)
}
