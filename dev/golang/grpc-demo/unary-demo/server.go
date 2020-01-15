package main

import (
	"context"
	"log"
	"net"

	"github.com/tabakazu/grpc-unary-demo/hellopb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	firstName := req.GetHello().GetFirstName()
	result := "Hello " + firstName
	res := &hellopb.HelloResponse{
		Result: result,
	}
	log.Printf("Received: %v", result)
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hellopb.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
