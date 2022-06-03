package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/chkda/go-grpc-course/orderMgmnt/orderMgmnt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedOrderManagementServer
}

const (
	port = ":25000"
)

var orderMap = make(map[string]*pb.Order)

func (s *server) GetOrder(ctx context.Context, in *pb.OrderID) (*pb.Order, error) {
	_, exists := orderMap[in.Value]
	if exists {
		return orderMap[in.Value], nil
	}
	return nil, errors.New("order does not exist")
}

func (s *server) AddOrder(ctx context.Context, in *pb.Order) (*pb.OrderID, error) {
	in.Id = uuid.New().String()
	orderMap[in.Id] = in
	log.Printf("Order Added: %v", in.Id)
	return &pb.OrderID{Value: in.Id}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Unable to listen to %v", port)
	}
	log.Printf("Listening on port %v", port)
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Unable to serve :%v", err)
	}
}
