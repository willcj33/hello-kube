package grpcHandler

import (
	"github.com/sercand/kuberesolver"
	pb "github.com/willcj33/hello-kube/color-requestor/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func NewClient() (pb.ColorGeneratorServiceClient, *grpc.ClientConn) {
	balancer := kuberesolver.New()
	conn, err := balancer.Dial("kubernetes://color-generator:50051", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	client := pb.NewColorGeneratorServiceClient(conn)
	return client, conn
}
