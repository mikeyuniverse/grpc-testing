package main

import (
	"context"
	"log"
	"sync"
	"time"

	increment "github.com/mikeyuniverse/grpc-testing/bidistreaming/pkg/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverAddr = "localhost:8000"
)

func main() {
	conn, err := grpc.NewClient(
		serverAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	client := increment.NewIncrementAPIClient(conn)

	stream, err := client.Inc(ctx)
	if err != nil {
		log.Fatalf("prepare stream: %v", err)
	}
	defer stream.CloseSend()

	var value int64 = 14

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		t := time.NewTicker(time.Second)
		defer t.Stop()

		for range t.C {
			err = stream.Send(&increment.IncValue{Value: value})
			if err != nil {
				log.Println("send: ", err)
				return
			}

			resp, err := stream.Recv()
			if err != nil {
				log.Println("recv: ", err)
				return
			}

			value = resp.Value

			log.Println("new value:", value)
		}
	}()

	wg.Wait()
}
