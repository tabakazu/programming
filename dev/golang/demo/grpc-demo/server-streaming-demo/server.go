package main

import (
	"log"
	"net"
	"strconv"
	"time"

	"github.com/tabakazu/grpc-server-streaming-demo/chatpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) PostChatMessage(req *chatpb.ChatRequest, stream chatpb.ChatService_PostChatMessageServer) error {
	content := req.GetChatMessage().GetContent()
	for i := 0; i < 10; i++ {
		result := content + " number " + strconv.Itoa(i)
		res := &chatpb.ChatResponse{
			Result: result,
		}
		stream.Send(res)
		log.Printf("Received: %v", result)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	chatpb.RegisterChatServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
