package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pb "hello.grpc/hello"
	"net"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) Greeting(ctx context.Context, in *pb.HelloReq) (*pb.HelloResp, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResp{Message: "Hello " + in.GetName()}, nil
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
