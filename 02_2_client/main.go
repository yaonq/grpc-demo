package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ygongq/grpc-demo/02_1_server/proto"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	var err error
	// grpc是建立在http/2之上的；grpc.WithInsecure()跳过对服务器证书的验证
	if conn, err = grpc.Dial(":6688", grpc.WithInsecure()); err != nil {
		log.Fatal("[grpc.Dial] error: ", err)
	}
	defer conn.Close()

	client := proto.NewDogServiceClient(conn)

	var res *proto.Response
	if res, err = client.GetDogInfo(context.Background(), &proto.Request{DogId: 6688}); err != nil {
		log.Fatal("[client.GetDogInfo] error: ", err)
	}

	fmt.Println(res)
}
