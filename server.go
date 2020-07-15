package main

import (
    "api.com/rpc/helper"
    services "api.com/rpc/services"
    "fmt"
    "google.golang.org/grpc"
    "net"
)

const (
    GrpcPort = "8091"
)

func main() {

    // 初始化证书版的grpc 服务
    rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCert()))

    services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
    services.RegisterOrderServerServer(rpcServer, new(services.OrderService))

    listen, err := net.Listen("tcp", fmt.Sprintf(":%s", GrpcPort))

    if err != nil {
        panic(err)
    }
    rpcServer.Serve(listen)
}
