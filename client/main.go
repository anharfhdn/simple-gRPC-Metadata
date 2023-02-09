package main

import (
	pb "anharfhdn/learn/grpc-metadata/proto/echo"
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func CallerUnaryEcho(c pb.EchoClient, message string) {
	fmt.Printf("----CallerUnaryEcho----\n")

	md := metadata.Pairs("Timestamps", time.Now().Format(time.StampNano))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	response, err := c.UnaryEcho(ctx, &pb.EchoRequest{Message: message})
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Response:\n")
	fmt.Printf("-> %s", response.Message)
}

func main() {
	conn, err := grpc.Dial(":1800", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)
	CallerUnaryEcho(client, "message-1\n")
	CallerUnaryEcho(client, "anharfhdn\n")

}
