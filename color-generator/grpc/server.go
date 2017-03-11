package server

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	pb "github.com/willcj33/hello-kube/color-generator/protos"
	context "golang.org/x/net/context"
)

// ColorGeneratorServer defines the struct for our grpc server implementation
type ColorGeneratorServer struct{}

// GetColor gets a random color
func (s *ColorGeneratorServer) GetColor(ctx context.Context, empty *pb.Empty) (*pb.Color, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	return &pb.Color{
		Hex: fmt.Sprintf("%v%v%v",
			strconv.FormatInt(rand.Int63n(256), 16),
			strconv.FormatInt(rand.Int63n(256), 16),
			strconv.FormatInt(rand.Int63n(256), 16)),
	}, nil
}

// NewServer gets a new instance of ColorGeneratorServer
func NewServer() *ColorGeneratorServer {
	s := new(ColorGeneratorServer)
	return s
}
