package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/ygongq/grpc-demo/grpc-kindle/02_orderManagement/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// 服务端流
func main() {
	conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewOrderManagementClient(conn)

	ctx, cFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cFunc()

	stream, _ := client.SearchOrders(ctx, &wrapperspb.StringValue{Value: "Mac"})

	for {

		order, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}

		if err == nil {
			fmt.Println("Search Result : ", order)
		} else {
			fmt.Println(err)
		}

	}
}
