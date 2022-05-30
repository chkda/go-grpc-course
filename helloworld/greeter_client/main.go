package main

import (
	"context"
	pb "github.com/chkda/go-grpc-course/helloworld/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	address = "localhost:25000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := &pb.HelloRequest{
		Name: "Chhaya",
	}
	response, err := c.SayHello(ctx, request)

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Println("Greeting: ", response.GetMessage())
}
