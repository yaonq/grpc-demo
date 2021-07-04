#!/bin/bash

# 1. 该脚本在/home/grpc-demo/grpc-kindle/02_orderManagement/server下执行
# 2. --proto_path指定从何处导入import中指定的包

protoc --proto_path=/home/grpc-demo/grpc-kindle:. \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/orderManagement.proto