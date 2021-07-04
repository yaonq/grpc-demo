package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	pb "github.com/ygongq/grpc-demo/grpc-kindle/04_deadline/server/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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
	id, err := strconv.Atoi(req.ID)
	// 非法请求
	if err != nil || id < 1 {
		// 创建一个新的错误状态
		errStatus := status.New(codes.InvalidArgument, "Invalid information received（无效的参数）")
		// 包装
		ds, err := errStatus.WithDetails(&errdetails.BadRequest_FieldViolation{
			Field:       "ID",
			Description: fmt.Sprintf("Order ID received is not valid %s", req.ID),
		})

		if err != nil {
			return nil, err
		}
		// 返回错误
		return nil, ds.Err()
	}

	rsp = new(pb.UnaryInterceptorResponse)
	rsp.Value = "错误处理"

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
