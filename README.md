## rpc(远程过程调用)

## rpc框架
```shell
go get google.golang.org/grpc
```
## 序列化库
```shell
# 安装序列化库
go get -u github.com/golang/protobuf/proto

# 下载protoc并添加到环境变量，protoc是protobuff文件的编译器，将.proto文件转义成各种编程语言对应的源码
wget https://github.com/protocolbuffers/protobuf/releases

# protoc-gen-go，是proto的golang插件，根据.proto文件转换成*.pb.go；如果设置GOBIN的话会自动安装到$GOBIN
go get -u github.com/golang/protobuf/protoc-gen-go

# 根据.proto文件转换成*_grpc.pb.go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## 自签证书
```shell
# 生成私钥
openssl genrsa -out server.key 1024
# 根据私钥生成证书申请文件csr
openssl req -new -key server.key -out server.csr
# 使用私钥对证书申请进行签名从而生成证书
openssl x509 -req -in server.csr -out server.crt -signkey server.key -days 3650
```
## 单项证书和双向证书