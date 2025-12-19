package main

import (
	"context"
	"log"

	status "github.com/mikeyuniverse/grpc-testing/wrong-enum-value/pkg"
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

	statusClient := status.NewStatusAPIClient(conn)

	_, err = statusClient.CheckStatus(context.Background(), &status.CheckStatusRequest{
		Status: status.Status_STATUS_SECOND + 20,
	})
	if err != nil {
		log.Fatalf("check status: %v\n", err)
	}
}
