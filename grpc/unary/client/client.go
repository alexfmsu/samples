package main

import (
	// system
	"context"
	"fmt"
	"log"

	// unary
	loginpb "unary/proto/loginpb"

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

	grpcClient := loginpb.NewLoginServiceClient(conn)

	req := loginpb.LoginRequest{
		Login:    "login",
		Password: "passwd",
	}

	resp, err := grpcClient.Login(context.Background(), &req)
	if err != nil {
		log.Fatalf("[client] Failed to login: %v", err)
		return
	}

	fmt.Println("[client] (from server)", resp)
}
