package main

import (
	"context"
	"log"
	"time"

	pb "github.com/chkda/go-grpc-course/ecommerce/ecommerce"
	"google.golang.org/grpc"
)

const (
	address = "localhost:25000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect")
	}
	defer conn.Close()

	c := pb.NewProductInfoClient(conn)

	name := "Apple Iphone 11"
	description := "Meet Apple Iphone 11."
	price := float32(699.00)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})

	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", resp.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: resp.Value})
	if err != nil {
		log.Fatalf("Could not fetch product: %v", err)
	}
	log.Printf("Product: %v", product.String())
}
