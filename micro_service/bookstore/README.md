# Bookstore

gRPC&gRPC-Gateway小练习

## bookstore介绍

书店里有很多书架，每个书架有自己的主题和大小，分别表示摆放的图书的主题和数量

## 要点
1. 数据库
2. proto
3. 写业务逻辑  
    数据库操作  
    grpc逻辑

## proto文件
proto/bookstore.proto

## 生成命令:
```shell
protoc -I=proto --proto_path=proto --go_out=proto --go_opt=paths=source_relative --go-gr
pc_out=proto --go-grpc_opt=paths=source_relative --grpc-gateway_out=proto --grpc-gateway_opt=paths=source_relative  proto/hello.proto
```

