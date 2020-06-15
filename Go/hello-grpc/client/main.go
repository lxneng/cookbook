package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pb "hello.grpc/hello"
	"os"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewHelloServiceClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Greeting(ctx, &pb.HelloReq{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
