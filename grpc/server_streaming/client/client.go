package main

import (
	// system
	"context"
	"fmt"
	"io"
	"log"

	// server_streaming
	lotspb "server_streaming/proto/lotspb"

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
		log.Fatalf("Failed to connect to server: %v", err)
		return
	}
	defer conn.Close()

	client := lotspb.NewLotsServiceClient(conn)

	req := lotspb.LotsRequest{
		Limit: 3,
	}

	resp, err := client.ActiveLots(context.Background(), &req)
	if err != nil {
		log.Fatalf("Failed to get active lots: %v", err)
		return
	}

	for {
		lots, err := resp.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("Failed to receive active lots from server: %v", err)
		}

		fmt.Println("[from server] ActiveLots:", lots.Lot)
	}
}
