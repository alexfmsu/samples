package main

import (
	// system
	"context"
	"fmt"
	"log"
	"net"

	// unary
	loginpb "unary/proto/loginpb"

	// external
	"google.golang.org/grpc"
)

type server struct{}

var (
	port string = "50050"
)

func (s *server) Login(ctx context.Context, req *loginpb.LoginRequest) (*loginpb.LoginResponse, error) {
	if req.Login == "login" && req.Password == "passwd" {
		res := loginpb.LoginResponse{Result: "Ok"}
		fmt.Println("[server] (from client)", req)
		return &res, nil
	}

	res := loginpb.LoginResponse{Result: "Error"}
	return &res, nil
}

func main() {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("[server] Failed to listen port %s: %v", port, err)
		return
	}
	fmt.Println("[server] Listening port " + port + "...")

	grpcServer := grpc.NewServer()

	loginpb.RegisterLoginServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("[server] Failed to register LoginServiceServer: %v", err)
	}
}
