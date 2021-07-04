package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ygongq/grpc-demo/grpc-kindle/03_interceptor/unary/server/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:5004", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("client create conn err: ", err)
	}
	defer conn.Close()

	client := proto.NewUnaryInterceptorClient(conn)

	ctx, cFunc := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cFunc()

	rsp, err := client.GetValue(ctx, &proto.UnaryInterceptorRequest{ID: "-1"})
	if err != nil {
		// got := status.Code(err)
		errorStatus := status.Convert(err)
		for _, v := range errorStatus.Details() {
			switch info := v.(type) {
			case *errdetails.BadRequest_FieldViolation:
				log.Printf("Request Field Invalid: %s", info)
			default:
				log.Printf("Unexpected error type: %s", info)
			}
		}
	}

	fmt.Println(rsp)
}
