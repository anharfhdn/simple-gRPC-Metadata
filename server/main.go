package main

import (
	pb "anharfhdn/learn/grpc-metadata/proto/echo"
	"context"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) UnaryEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	fmt.Printf("-----UnaryEcho-----\n")
	fmt.Printf("-----Incoming request-----\n")
	fmt.Printf("Message: %s -> send echo\n", in.Message)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.DataLoss, "error UnaryEcho")
	}
	if v, ok := md["Timestamp"]; ok {
		fmt.Printf("Metadata Timestamp")
		for i, e := range v {
			fmt.Printf("%d.%s", i, e)
		}
	}

	return &pb.EchoResponse{Message: in.Message}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":1800")
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEchoServer(grpcServer, &server{})

	grpcServer.Serve(lis)

}
