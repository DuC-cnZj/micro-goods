package main

import (
	goods_proto "github.com/DuC-cnZj/micro-goods/protos"
	goods "github.com/DuC-cnZj/micro-goods/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	goods_proto.RegisterBBQServer(server, &goods.Goods{})
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
