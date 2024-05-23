package main

import (
	// system
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"

	// client_streaming
	orderspb "client_streaming/proto/orderspb"

	// external
	"google.golang.org/grpc"
)

type server struct{}

var (
	port string = "50050"
)

func (s *server) PostOrder(stream orderspb.OrdersService_PostOrderServer) error {
	var executedOrders int

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				resp := orderspb.OrdersResponse{
					ExecutedOrders: int64(executedOrders),
				}

				stream.SendAndClose(&resp)
				break
			}

			log.Fatalf("[server] Failed to receive message from client: %v", err)
			continue
		}

		if rand.Intn(1000)%2 == 0 {
			fmt.Printf("[server] (from client) Executed order with price %.2f and quantity %d\n", req.Price, req.Quantity)
			executedOrders++
		}
	}

	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("[server] Failed to listen port %s: %v", port, err)
		return
	}
	fmt.Println("[server] Listening port " + port + "...")

	grpcServer := grpc.NewServer()

	orderspb.RegisterOrdersServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("[server] Failed to register OrdersServiceServer: %v", err)
		return
	}
}
