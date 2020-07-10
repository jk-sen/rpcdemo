package main

import (
    "api.com/rpc/helper"
    "api.com/rpc/services"
    "context"
    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    "google.golang.org/grpc"
    "log"
    "net/http"
)


func main() {

    //Register gRPC server endpoint
    // 注册http 服务路由
    mux := runtime.NewServeMux()

    // 上下文
    ctx := context.Background()

    //option
    option := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCert())}

    err := services.RegisterProdServiceHandlerFromEndpoint(ctx, mux, "localhost:8091", option)

    if err != nil {
        log.Fatal(err)
    }

    httpServer := &http.Server{
        Addr: ":82",
        Handler: mux,
    }

    err = httpServer.ListenAndServe()

    if err != nil {
        log.Fatal(err)
    }

    log.Println("http 服务启动成功")
}
