package helper

import (
    "crypto/tls"
    "crypto/x509"
    "google.golang.org/grpc/credentials"
    "io/ioutil"
)

// 加载服务端证书
func GetServerCert() credentials.TransportCredentials {

    cert, _ := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
    certPool := x509.NewCertPool()
    ca, _ := ioutil.ReadFile("cert/ca.pem")
    certPool.AppendCertsFromPEM(ca)

    creds := credentials.NewTLS(&tls.Config{
        Certificates: []tls.Certificate{cert},
        ClientAuth:   tls.RequireAndVerifyClientCert,
        ClientCAs:    certPool,
    })

    return creds
}

// 获取客户端证书
func GetClientCert() credentials.TransportCredentials {
    cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
    certPool := x509.NewCertPool()
    ca, _ := ioutil.ReadFile("cert/ca.pem")
    certPool.AppendCertsFromPEM(ca)

    creds := credentials.NewTLS(&tls.Config{
        Certificates: []tls.Certificate{cert},
        ServerName:   "localhost",
        RootCAs:      certPool,
    })

    return creds
}
