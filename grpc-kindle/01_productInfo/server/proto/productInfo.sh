#!/bin/bash

# --go_out: 指定go文件生成目录
# --go_opt=paths=source_relative: productInfo.proto文件中定义的go的package为github.com/ygongq/grpc-demo/grpc-kindle/proto/productInfo，这个路径是项目路径，如果不指定source_relative则会在当前目录下生成项目路径
# --go-grpc_out: 自动生成的grpc服务框架代码目录

# 需要注意的--go_out指定的.的实际目录并不是当前脚本执行的目录(grpc-demo/grpc-kindle/server/)，而是./proto，这是因为指定的proto文件是proto/productInfo.proto，protoc会拼接
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/productInfo.proto

# 在server目录下执行