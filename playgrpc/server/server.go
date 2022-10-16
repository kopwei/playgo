package server

import (
	"context"
	"fmt"
	pb "github.com/kopwei/playgo/playgrpc/proto"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
	"net"
)

type playGrpcServer struct{}

func (pgs *playGrpcServer) GetSimpleResponce(context.Context, *pb.SimpleMsg) (*pb.SimpleMsg, error) {
	return nil, nil
}

func newServer() *playGrpcServer {
	s := &playGrpcServer{}
	return s
}

func StartServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		klog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterPlayGrpcServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
