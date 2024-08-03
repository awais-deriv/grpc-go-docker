package main

import (
	"context"
	hello "grpc-proto-test/client/generated"

	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "server:50051" // Use the service name 'server'
)

func main() {

	conn, err := grpc.NewClient(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := hello.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.SayHello(ctx, &hello.HelloRequest{Name: "world"})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())

}
