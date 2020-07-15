#!/bin/sh
cd protofiles
protoc --go_out=plugins=grpc:../services prod.proto
protoc --go_out=plugins=grpc:../services ProdModel.proto
protoc --go_out=plugins=grpc:../services Order.proto

protoc --grpc-gateway_out=logtostderr=true:../services prod.proto
protoc --grpc-gateway_out=logtostderr=true:../services Order.proto

