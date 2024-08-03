package main

import (
	"context"
	"fmt"
	"log"
	"net"

	hello "grpc-proto-test/server/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	hello.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &hello.HelloReply{Message: fmt.Sprintf("Hello %s", in.GetName())}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	reflection.Register(s)
	hello.RegisterGreeterServer(s, &server{})
	fmt.Println("Server started at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
