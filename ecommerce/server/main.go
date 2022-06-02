package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/chkda/go-grpc-course/ecommerce/ecommerce"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const (
	port = ":25000"
)

type server struct {
	pb.UnimplementedProductInfoServer
	productMap map[string]*pb.Product
}

func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	uuid := uuid.New().String()
	in.Id = uuid

	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	log.Printf("Product %v : %v - Added", in.Id, in.Name)
	return &pb.ProductID{Value: in.Id}, nil
}

func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	_, exists := s.productMap[in.Value]

	if exists {
		return s.productMap[in.Value], nil
	}
	return nil, errors.New("product does not exist")
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Listening on port: ", port)
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
