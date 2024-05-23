package main

import (
	// system
	"context"
	"io"
	"log"
	"math/rand"
	"time"

	// bidirectional streaming
	mathpb "bidirectional_streaming/proto/mathpb"

	// external
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	host string = "0.0.0.0"
	port string = "50050"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)

	conn, err := grpc.NewClient(host+":"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("[client] Failed to connect to server: %v", err)
		return
	}

	client := mathpb.NewMathClient(conn)

	stream, err := client.Max(context.Background())
	if err != nil {
		log.Fatalf("[client] Failed to create stream %v", err)
	}

	var max int32
	ctx := stream.Context()

	done := make(chan bool)

	go func() {
		for i := 1; i <= 10; i++ {
			rnd := int32(rand.Intn(i))

			req := mathpb.Request{Num: rnd}

			if err := stream.Send(&req); err != nil {
				log.Fatalf("[client] Failed to send data to server: %v", err)
				continue
			}

			log.Printf("[client] Sent %d to server", req.Num)
			time.Sleep(time.Millisecond * 200)
		}

		if err := stream.CloseSend(); err != nil {
			log.Fatalf("[client] Failed to get response from server: %v", err)
			return
		}
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					close(done)
					return
				}

				log.Fatalf("[client] Failed to receive message from server: %v", err)
				continue
			}

			max = resp.Result
			log.Printf("[client] (from server) Received new max=%d", max)
		}
	}()

	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done

	log.Printf("[client] Finished processing data from server with max=%d", max)
}
