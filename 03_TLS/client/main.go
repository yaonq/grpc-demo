package main

import (
	"context"
	"log"

	"github.com/ygongq/grpc-demo/02_1_server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	/*
		Country Name (2 letter code) [XX]:demo.grpc.com
		string is too long, it needs to be less than  2 bytes long
		Country Name (2 letter code) [XX]:CN
		State or Province Name (full name) []:ShenZhen
		Locality Name (eg, city) [Default City]:XiXiang
		Organization Name (eg, company) [Default Company Ltd]:YGQ
		Organizational Unit Name (eg, section) []:yanggongqi
		Common Name (eg, your name or your server's hostname) []:ygq
		Email Address []:no.guess@qq.com

		Please enter the following 'extra' attributes
		to be sent with your certificate request
		A challenge password []:
		An optional company name []:demo.grpc.com
	*/
	// 基于服务器的证书和服务器的名字对服务器进行验证；服务器名字在制作证书申请时指定
	creds, err := credentials.NewClientTLSFromFile("../keys/server.crt", "ygq")
	if err != nil {
		log.Fatal("[credentials.NewClientTLSFromFile] error: ", err)
	}

	conn, err := grpc.Dial(":6688", grpc.WithTransportCredentials(creds))
	defer conn.Close()
	if err != nil {
		log.Fatal("[grpc.Dial] error: ", err)
	}

	client := proto.NewDogServiceClient(conn)

	res, err := client.GetDogInfo(context.Background(), &proto.Request{DogId: 6688})
	if err != nil {
		log.Fatal("[client.GetDogInfo] error: ", err)
	}

	log.Println(res)
}
