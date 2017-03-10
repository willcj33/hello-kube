package main

// Import packages
import (
	"log"
	"time"

	"golang.org/x/net/context"

	pb "github.com/willcj33/hello-kube/color-requestor/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var forever = make(chan struct{})

func main() {
	conn, err := grpc.Dial("color-generator:50051", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewColorGeneratorServiceClient(conn)
	tickerLoad := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-tickerLoad.C:
				color, err := client.GetColor(context.Background(), &pb.Empty{})
				if err != nil {
					log.Println(err)
					return
				}
				log.Printf("Color: %v\n", color.Hex)
			case <-quit:
				tickerLoad.Stop()
			}
		}
	}()
	<-forever
}