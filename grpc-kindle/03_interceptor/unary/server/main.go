package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/ygongq/grpc-demo/grpc-kindle/03_interceptor/unary/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	addr                    = flag.String("addr", "localhost:5003", "unaryInterceptor server addr")
	UnaryInterceptorService = defaultService{}
)

// 实现接口
type defaultService struct {
	pb.UnimplementedUnaryInterceptorServer
}

func (s *defaultService) GetValue(ctx context.Context, req *pb.UnaryInterceptorRequest) (rsp *pb.UnaryInterceptorResponse, err error) {
	fmt.Println("进入 --> UnaryInterceptorService.GetValue")

	fmt.Println(req)

	rsp = new(pb.UnaryInterceptorResponse)
	rsp.Value = "hello world"

	return
}

// 一元拦截器
// 需要实现UnaryServerInterceptor函数
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 前置处理逻辑
	fmt.Printf("拦截器：前置处理逻辑 --> grpc.UnaryServerInfo: %+v\n", info)

	// 调用RPC方法
	fmt.Println("拦截器：调用RPC方法")
	rsp, err := handler(ctx, req)
	if err != nil {
		log.Fatal("调用RPC方法失败")
	}

	// 后置处理
	fmt.Println("拦截器：后置处理逻辑")
	rspp, ok := rsp.(*pb.UnaryInterceptorResponse)
	if ok {
		// 注意区分值类型和指针类型
		// grpc: server failed to encode response:
		// rpc error: code = Internal desc = grpc: error while marshaling: failed to marshal, message is **proto.UnaryInterceptorResponse, want proto.Message
		rspp.Value = rspp.Value + "已被后置处理"
		return rspp, nil
	}

	return rsp, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
	)

	pb.RegisterUnaryInterceptorServer(
		srv,
		&UnaryInterceptorService,
	)

	// 将服务注册到反射服务中，可以使用grpcurl查询gRPC列表或调用gRPC方法。
	// https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-08-grpcurl.html
	reflection.Register(srv)

	srv.Serve(lis)
}
