package goods

import (
	"context"
	goods_proto "github.com/DuC-cnZj/micro-goods/protos"
	"log"
	"strconv"
)

type Iphone struct {
	goods_proto.UnimplementedIphoneServer
}

func (g *Iphone) GetOneByType(ctx context.Context, request *goods_proto.Request) (*goods_proto.Response, error) {
	log.Println("goods bbq svc eat().")

	return &goods_proto.Response{
		Code: "200",
		Msg:  "goods svc: get iphone " + strconv.Itoa(int(request.Type)),
	}, nil
}

func (g *Iphone) mustEmbedUnimplementedBBQServer() {
	return
}
