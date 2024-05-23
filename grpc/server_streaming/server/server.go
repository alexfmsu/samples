package main

import (
	// system
	"log"
	"math/rand"
	"net"
	"time"

	// server_streaming
	lotspb "server_streaming/proto/lotspb"

	// external
	"google.golang.org/grpc"
)

type server struct{}

var (
	port string = "50050"
)

func (s *server) ActiveLots(req *lotspb.LotsRequest, resp lotspb.LotsService_ActiveLotsServer) error {
	startPrice := rand.Intn(100)

	for i := 1; i < 11; i++ {
		res := lotspb.LotsResponse{
			Lot: &lotspb.Lot{
				ID:    int64(i),
				Desc:  "Description",
				Price: float64(startPrice * i),
			},
		}

		resp.Send(&res)
		time.Sleep(time.Second)

		if req.Limit < int64(i)+1 {
			break
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

	grpcServer := grpc.NewServer()

	lotspb.RegisterLotsServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("[server] Failed to register LotsServiceServer: %v", err)
		return
	}
}
