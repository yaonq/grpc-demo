package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/ygongq/grpc-demo/grpc-kindle/06_metadata/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

var (
	addr                    = flag.String("addr", "localhost:5005", "unaryInterceptor server addr")
	UnaryInterceptorService = defaultService{}
)

// 实现接口
type defaultService struct {
	pb.UnimplementedUnaryInterceptorServer
}

func (s *defaultService) GetValue(ctx context.Context, req *pb.UnaryInterceptorRequest) (rsp *pb.UnaryInterceptorResponse, err error) {

	// 读取元数据
	md, _ := metadata.FromIncomingContext(ctx)
	for k, v := range md {
		fmt.Println(k, "==>", v)
	}

	// 发送元数据
	head := metadata.New(map[string]string{"md_test": "hello world"})
	grpc.SendHeader(ctx, head)

	rsp = new(pb.UnaryInterceptorResponse)
	rsp.Value = "元数据"

	return
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}

	srv := grpc.NewServer()

	pb.RegisterUnaryInterceptorServer(
		srv,
		&UnaryInterceptorService,
	)

	// 将服务注册到反射服务中，可以使用grpcurl查询gRPC列表或调用gRPC方法。
	// https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-08-grpcurl.html
	reflection.Register(srv)

	srv.Serve(lis)
}
