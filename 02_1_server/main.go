package main

import (
	"context"
	"net"

	"github.com/ygongq/grpc-demo/02_1_server/proto"
	"google.golang.org/grpc"
)

type DogService struct {
}

func (d *DogService) GetDogInfo(c context.Context, r *proto.Request) (*proto.Response, error) {
	return &proto.Response{
		Data: &proto.Dog{
			Name: "大黄",
			Age:  8,
		},
		Msg: &proto.Msg{
			Code: 200,
			Err:  "",
		},
	}, nil
}

func main() {
	srv := grpc.NewServer()
	proto.RegisterDogServiceServer(srv, new(DogService))

	licten, _ := net.Listen("tcp", ":6688")

	srv.Serve(licten)
}
