package server

import (
	"log"

	colorUtil "github.com/lucasb-eyer/go-colorful"
	pb "github.com/willcj33/hello-kube/color-generator/protos"
	context "golang.org/x/net/context"
)

var placement = "routes.go"

// ColorGeneratorServer defines the struct for our grpc server implementation
type ColorGeneratorServer struct{}

// GetColor gets a random color
func (s *ColorGeneratorServer) GetColor(ctx context.Context, empty *pb.Empty) (*pb.Color, error) {
	currentColor := colorUtil.FastHappyColor()
	log.Printf("Color: %v\n", currentColor.Hex())
	return &pb.Color{
		Hex: currentColor.Hex(),
	}, nil
}

// NewServer gets a new instance of ColorGeneratorServer
func NewServer() *ColorGeneratorServer {
	s := new(ColorGeneratorServer)
	return s
}
