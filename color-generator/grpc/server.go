package server

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	pb "github.com/willcj33/hello-kube/color-generator/protos"
	context "golang.org/x/net/context"
)

var placement = "routes.go"

// ColorGeneratorServer defines the struct for our grpc server implementation
type ColorGeneratorServer struct{}

// GetColor gets a random color
func (s *ColorGeneratorServer) GetColor(ctx context.Context, empty *pb.Empty) (*pb.Color, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	h := fmt.Sprintf("%v%v%v", strconv.FormatInt(rand.Int63n(256), 16), strconv.FormatInt(rand.Int63n(256), 16), strconv.FormatInt(rand.Int63n(256), 16))
	log.Printf("Color: %v\n", h)
	return &pb.Color{
		Hex: h,
	}, nil
}

// NewServer gets a new instance of ColorGeneratorServer
func NewServer() *ColorGeneratorServer {
	s := new(ColorGeneratorServer)
	return s
}
