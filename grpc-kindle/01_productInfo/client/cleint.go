package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/ygongq/grpc-demo/grpc-kindle/01_productInfo/server/proto"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	var err error
	// grpc是建立在http/2之上的；grpc.WithInsecure()跳过对服务器证书的验证
	if conn, err = grpc.Dial("localhost:5001", grpc.WithInsecure()); err != nil {
		log.Fatal("[grpc.Dial] error: ", err)
	}
	defer conn.Close()

	var client pb.ProductlnfoClient
	client = pb.NewProductlnfoClient(conn)

	var (
		ctx context.Context
		cfc context.CancelFunc
	)
	ctx, cfc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cfc()

	var v1 *pb.ProductID
	if v1, err = client.AddProduct(ctx, &pb.Product{Name: "苹果", Description: "又红又甜的苹果"}); err != nil {
		log.Fatal("v1 err :", err)
	}

	fmt.Printf("v1: %v\n", v1)

	var v2 *pb.Product
	if v2, err = client.GetProduct(context.TODO(), v1); err != nil {
		log.Fatal("v2 err :", err)
	}

	fmt.Printf("v2: %v\n", v2)
}
