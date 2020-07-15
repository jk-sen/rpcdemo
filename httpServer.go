package main

import (
    "api.com/rpc/helper"
    "api.com/rpc/services"
    "context"
    "fmt"
    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    "google.golang.org/grpc"
    "log"
    "net/http"
)

var (
    HttpPort = "8081"
)

func main() {

    //Register gRPC server endpoint
    mux := runtime.NewServeMux()

    ctx := context.Background()

    grpcEndPoint := "localhost:8091"

    option := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCert())}

    // 注册商品服务
    err := services.RegisterProdServiceHandlerFromEndpoint(ctx, mux, grpcEndPoint, option)

    // 注册订单服务
    err = services.RegisterOrderServerHandlerFromEndpoint(ctx, mux, grpcEndPoint, option)

    if err != nil {
        log.Fatal(err)
    }

    httpServer := &http.Server{
        Addr:    fmt.Sprintf(":%s", HttpPort),
        Handler: mux,
    }

    err = httpServer.ListenAndServe()

    if err != nil {
        log.Fatal(err)
    }

    log.Println("http 服务启动成功")
}
