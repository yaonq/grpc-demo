package main

import (
	"context"
	"log"
	"net"

	"github.com/ygongq/grpc-demo/02_1_server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Service struct{}

func (s *Service) GetDogInfo(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	return &proto.Response{Data: &proto.Dog{Name: "旺财", Age: 3}, Msg: &proto.Msg{Code: 200, Err: "nil"}}, nil
}

func main() {

	// 从文件为服务器构建证书对象
	creds, err := credentials.NewServerTLSFromFile("../keys/server.crt", "../keys/server.key")
	if err != nil {
		log.Fatal("[credentials.NewServerTLSFromFile] error: ", err)
	}

	// 将证书包装为选项作为参数传入server
	srv := grpc.NewServer(grpc.Creds(creds))

	// 注册服务
	proto.RegisterDogServiceServer(srv, new(Service))

	conn, _ := net.Listen("tcp", ":6688")
	defer conn.Close()

	srv.Serve(conn)
}
