package grpcHandler

import (
	"context"

	"github.com/sercand/kuberesolver"
	pb "github.com/willcj33/hello-kube/color-requestor/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var client pb.ColorGeneratorServiceClient

func Init() {
	balancer := kuberesolver.New()
	conn, err := balancer.Dial("kubernetes://color-generator:50051", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client = pb.NewColorGeneratorServiceClient(conn)
}

func GetColor() (*pb.Color, error) {
	return client.GetColor(context.Background(), &pb.Empty{})
}
