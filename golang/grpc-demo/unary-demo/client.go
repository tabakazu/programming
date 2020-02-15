package main

import (
	"context"
	"log"

	"github.com/tabakazu/grpc-unary-demo/hellopb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := hellopb.NewHelloServiceClient(conn)
	doUnary(c)
}

func doUnary(c hellopb.HelloServiceClient) {
	req := &hellopb.HelloRequest{
		Hello: &hellopb.Hello{
			FirstName: "Taro",
			LastName:  "Tanaka",
		},
	}
	res, err := c.Hello(context.Background(), req)
	if err != nil {
		log.Fatalf("could not hello: %v", err)
	}
	log.Printf("Hello: %s", res.Result)
}
