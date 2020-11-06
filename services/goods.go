package goods

import (
	"context"
	"github.com/DuC-cnZj/micro-goods/protos"
	"log"
)

type Goods struct {
	goods_proto.UnimplementedBBQServer
}

func (g *Goods) Eat(ctx context.Context, request *goods_proto.Request) (*goods_proto.Response, error) {
	log.Println("goods bbq svc eat().")

	return &goods_proto.Response{
		Code: "200",
		Msg:  "goods bbq svc: eat called",
	}, nil
}

func (g *Goods) mustEmbedUnimplementedBBQServer() {
	return
}
