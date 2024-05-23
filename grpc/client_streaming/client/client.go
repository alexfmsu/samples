package main

import (
	// system
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	// client_streaming
	orderspb "client_streaming/proto/orderspb"

	// external
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	host string = "0.0.0.0"
	port string = "50050"
)

func main() {
	conn, err := grpc.NewClient(host+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("[client] Failed to connect to server: %v", err)
		return
	}
	defer conn.Close()

	client := orderspb.NewOrdersServiceClient(conn)

	stream, err := client.PostOrder(context.Background())
	if err != nil {
		log.Fatalf("[client] Failed to post order: %v", err)
		return
	}

	for i := 0; i < 10; i++ {
		order := orderspb.OrderRequest{
			Price:    float64(rand.Intn(1000)),
			Quantity: int64(rand.Intn(10)),
		}

		stream.Send(&order)
		time.Sleep(time.Second)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("[client] Failed to get response from server: %v", err)
		return
	}

	fmt.Println("[client] (from server) Executed orders:", resp.ExecutedOrders)
}
