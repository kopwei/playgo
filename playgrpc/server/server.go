package server

import (
	"context"
	"github.com/kopwei/playgo/playgrpc/proto"
)

type PlayGrpcServer struct{}

func GetSimpleResponce(context.Context, *proto.SimpleMsg) (*proto.SimpleMsg, error) {
	return nil, nil
}
