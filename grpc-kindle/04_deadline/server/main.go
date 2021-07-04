package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/ygongq/grpc-demo/grpc-kindle/04_deadline/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	addr                    = flag.String("addr", "localhost:5004", "unaryInterceptor server addr")
	UnaryInterceptorService = defaultService{}
)

// 实现接口
type defaultService struct {
	pb.UnimplementedUnaryInterceptorServer
}

func (s *defaultService) GetValue(ctx context.Context, req *pb.UnaryInterceptorRequest) (rsp *pb.UnaryInterceptorResponse, err error) {
	time.Sleep(7 * time.Second)

	// 判断是否已截止
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("截止时间到: ", ctx.Err())
		return nil, ctx.Err()
	}

	rsp = new(pb.UnaryInterceptorResponse)
	rsp.Value = "hello world 未到截止时间"

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
