package main

import (
	"context"
	"fmt"
	"github.com/DuC-cnZj/order_protos"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func main() {
	dial, err := grpc.Dial("localhost:6666", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := order_protos.NewOrderClient(dial)
	list, err := client.List(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
}
