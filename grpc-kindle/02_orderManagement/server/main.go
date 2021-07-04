package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	data "github.com/ygongq/grpc-demo/grpc-kindle/02_orderManagement/server/dataMap"
	pb "github.com/ygongq/grpc-demo/grpc-kindle/02_orderManagement/server/proto"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	addr                   = flag.String("addr", "localhost:5002", "oredManagement server addr")
	OrderManagementService = orderManagementService{orderMap: data.NewDataMap()}
)

type orderManagementService struct {
	pb.UnimplementedOrderManagementServer
	orderMap data.DataMap
}

func (o *orderManagementService) AddOrder(ctx context.Context, order *pb.Order) (*wrapperspb.StringValue, error) {
	fmt.Printf("Add order: %v\n", order)
	o.orderMap.Set(order)
	return &wrapperspb.StringValue{Value: "Add Order: " + order.Id}, nil
}

func (o *orderManagementService) GetOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order, error) {

	return nil, nil
}

func (o *orderManagementService) SearchOrders(value *wrapperspb.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	for _, order := range o.orderMap.List() {
		for _, item := range order.Items {
			if strings.Contains(item, value.Value) {
				if err := stream.Send(order); err != nil {
					return fmt.Errorf("error sending message to stream : %v", err)
				}
			}
		}
	}
	return nil
}

func (o *orderManagementService) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {

	ordersStr := "Updated Order IDs : "

	for {
		order, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&wrapperspb.StringValue{Value: "Orders processed " + ordersStr})
		}

		if err != nil {
			return err
		}

		o.orderMap.Set(order)

		ordersStr = fmt.Sprintf("%s %s", ordersStr, order.Id)
	}
}

func (o *orderManagementService) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {
	return nil
}

// 判断orderManagementService是否实现了pb.OrderManagementService接口
var _ pb.OrderManagementServer = new(orderManagementService)

func main() {
	var (
		lis net.Listener
		err error
	)

	if lis, err = net.Listen("tcp", *addr); err != nil {
		log.Fatalf("create socket listen error: %v\n", err)
	}

	var srv *grpc.Server
	srv = grpc.NewServer()

	pb.RegisterOrderManagementServer(srv, &OrderManagementService)

	srv.Serve(lis)
}
