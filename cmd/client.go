package main

import (
	"context"
	"fmt"
	goods_proto "github.com/DuC-cnZj/micro-goods/protos"
	"github.com/DuC-cnZj/order_protos"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func main() {
	dialGoods, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	dialOrder, err := grpc.Dial("localhost:6666", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	defer dialGoods.Close()
	defer dialOrder.Close()
	goodsClient := goods_proto.NewBBQClient(dialGoods)
	orderClient := order_protos.NewOrderClient(dialOrder)
	eat, err := goodsClient.Eat(context.Background(), &goods_proto.Request{})
	if err != nil {
		panic(err)
	}
	list, err := orderClient.List(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}

	fmt.Println("goods eat: ", eat)
	fmt.Println("order list: ", list)
}
