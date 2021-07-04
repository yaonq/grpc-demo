package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/gofrs/uuid"
	_ "github.com/gofrs/uuid"
	pb "github.com/ygongq/grpc-demo/grpc-kindle/01_productInfo/server/proto"

	"google.golang.org/grpc"
)

type productInfoService struct {
	// 如果要实现向前兼容，必须嵌入该结构体
	pb.UnimplementedProductlnfoServer

	productInfoMap map[string]*pb.Product
}

func (s *productInfoService) AddProduct(ctx context.Context, request *pb.Product) (response *pb.ProductID, err error) {

	fmt.Println("新增商品信息")

	uuidOut, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	if s.productInfoMap == nil {
		s.productInfoMap = make(map[string]*pb.Product)
	}

	request.Id = uuidOut.String()

	s.productInfoMap[uuidOut.String()] = request

	return &pb.ProductID{Value: request.Id}, nil
}

func (s *productInfoService) GetProduct(ctx context.Context, request *pb.ProductID) (response *pb.Product, err error) {
	fmt.Println("获取商品信息")

	return s.productInfoMap[request.Value], nil
}

var (
	addr               = flag.String("addr", "localhost:5001", "server addr")
	ProductInfoService = productInfoService{}
)

func main() {
	flag.Parse()

	var (
		lis net.Listener
		err error
	)

	if lis, err = net.Listen("tcp", *addr); err != nil {
		log.Fatalf("grpc: create  socket listen error: %v\n", err)
	}

	var srv *grpc.Server
	srv = grpc.NewServer()

	pb.RegisterProductlnfoServer(srv, &ProductInfoService)

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
