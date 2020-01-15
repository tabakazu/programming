package main

import (
	"context"
	"io"
	"log"

	"github.com/tabakazu/grpc-server-streaming-demo/chatpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := chatpb.NewChatServiceClient(conn)
	doServerStreaming(c)
}

func doServerStreaming(c chatpb.ChatServiceClient) {
	req := &chatpb.ChatRequest{
		ChatMessage: &chatpb.ChatMessage{
			Content: "Hello",
		},
	}

	res, err := c.PostChatMessage(context.Background(), req)

	if err != nil {
		log.Fatalf("could not hello: %v", err)
	}

	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		log.Printf("Hello: %s", msg.GetResult())
	}
}
