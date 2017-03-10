package main

import (
	"net"

	serverImp "github.com/willcj33/hello-kube/color-generator/grpc"
	pb "github.com/willcj33/hello-kube/color-generator/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// Import packages

var (
	forever = make(chan struct{})
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterColorGeneratorServiceServer(grpcServer, serverImp.NewServer())
	grpcServer.Serve(lis)
	<-forever
}
