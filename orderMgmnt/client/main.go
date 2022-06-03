package main

import (
	"context"
	"log"
	"time"

	pb "github.com/chkda/go-grpc-course/orderMgmnt/orderMgmnt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:25000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewOrderManagementClient(conn)
	orders := []pb.Order{
		pb.Order{Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA", Price: 1800.00},
		pb.Order{Items: []string{"Apple Watch S4"}, Destination: "San Jose, CA", Price: 400.00},
		pb.Order{Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "Mountain View, CA", Price: 400.00},
		pb.Order{Items: []string{"Amazon Echo"}, Destination: "San Jose, CA", Price: 30.00},
		pb.Order{Items: []string{"Amazon Echo", "Apple iPhone XS"}, Destination: "Mountain View, CA", Price: 300.00},
	}
	orderIDs := make([]string, 0, 5)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, order := range orders {
		resp, err := c.AddOrder(ctx, &order)
		if err != nil {
			log.Fatalf("Unable to add order: %v", err)
		}
		log.Println("OrderID:", resp.Value)
		orderIDs = append(orderIDs, resp.Value)
	}

	for _, orderID := range orderIDs {
		log.Println("OrderID:", orderID)
		resp, err := c.GetOrder(ctx, &pb.OrderID{Value: orderID})
		if err != nil {
			log.Printf("Unable to fetch order: %v/n", err)
		}
		log.Printf("OrderId %v: %v", orderID, resp)
	}
}
