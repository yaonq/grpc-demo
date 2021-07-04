package main

import (
	"context"
	"log"
	"time"

	"github.com/ygongq/grpc-demo/grpc-kindle/02_orderManagement/server/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewOrderManagementClient(conn)

	ctx, cFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cFunc()

	res, _ := client.AddOrder(ctx, &proto.Order{Id: "101", Items: []string{"iPhone XS", "Mac Book Pro"}, Destination: "San Jose, CA", Price: 2300.00})

	if res != nil {
		log.Println("AddOrder Response -> ", res)
	}

	res, _ = client.AddOrder(ctx, &proto.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA", Price: 1800.00})

	if res != nil {
		log.Println("AddOrder Response -> ", res)
	}

	res, _ = client.AddOrder(ctx, &proto.Order{Id: "103", Items: []string{"Apple Watch S4"}, Destination: "San Jose, CA", Price: 400.00})

	if res != nil {
		log.Println("AddOrder Response -> ", res)
	}
}
