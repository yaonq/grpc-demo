package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ygongq/grpc-demo/grpc-kindle/02_orderManagement/server/proto"
	"google.golang.org/grpc"
)

// 客户端流
func main() {
	updOrder1 := proto.Order{Id: "101", Items: []string{"iPhone XS Up", "Mac Book Pro Up"}, Destination: "San Jose, CA", Price: 1100.00}
	updOrder2 := proto.Order{Id: "102", Items: []string{"Google Pixel 3A Up", "Mac Book Pro Up"}, Destination: "Mountain View, CA", Price: 2800.00}
	updOrder3 := proto.Order{Id: "103", Items: []string{"Apple Watch S4 Up"}, Destination: "San Jose, CA", Price: 2200.00}
	conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewOrderManagementClient(conn)

	ctx, cFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cFunc()

	stream, _ := client.UpdateOrders(ctx)

	res1 := stream.Send(&updOrder1)
	fmt.Println(res1)
	res2 := stream.Send(&updOrder2)
	fmt.Println(res2)
	res3 := stream.Send(&updOrder3)
	fmt.Println(res3)

	// 发送完成后关闭
	ws, err := stream.CloseAndRecv()
	if err == nil {
		fmt.Printf("Update Orders Res : %s\n", ws.String())
	}
}
