package main

import (
	// system
	"io"
	"log"
	"net"

	// bidirectional streaming
	mathpb "bidirectional_streaming/proto/mathpb"

	// external
	"google.golang.org/grpc"
)

type server struct{}

var (
	port string = "50050"
)

func (s *server) Max(srv mathpb.Math_MaxServer) error {
	log.Println("[server] Started processing data from client")

	var max int32

	ctx := srv.Context()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("[server] Finished processing data from client")
				return nil
			}

			log.Fatalf("[server] Failed to receive message from client: %v", err)
			continue
		}

		if req.Num <= max {
			continue
		}

		max = req.Num

		resp := mathpb.Response{Result: max}
		if err := srv.Send(&resp); err != nil {
			log.Fatalf("[server] Failed to send message to client: %v", err)
			continue
		}

		log.Printf("[server] Send new max=%d to client", max)
	}
}

func main() {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("[server] Failed to listen port %s: %v", port, err)
		return
	}

	grpcServer := grpc.NewServer()

	mathpb.RegisterMathServer(grpcServer, &server{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("[server] Failed to register MathServer: %v", err)
		return
	}
}
