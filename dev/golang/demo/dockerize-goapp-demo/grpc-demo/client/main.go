package main

import (
	"log"

	pb "app/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("server:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Hello"})
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err,
			})
			log.Fatalf("could not greet: %v", err)
			return
		}

		ctx.JSON(200, gin.H{
			"message": r.Message,
		})
		log.Printf("Greeting: %s", r.Message)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
