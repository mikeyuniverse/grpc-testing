package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	status "github.com/mikeyuniverse/grpc-testing/wrong-enum-value/pkg"
	"google.golang.org/grpc"
)

const (
	serverPort = "8000"
	serverHost = "localhost"
)

func main() {
	listener, err := net.Listen("tcp", serverHost+":"+serverPort)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	status.RegisterStatusAPIServer(server, &Server{})

	log.Printf("grpc server at %v\n", listener.Addr())
	err = server.Serve(listener) // blocks

	onExitMsg := "finished: no errors"
	exitCode := 0
	if err != nil {
		onExitMsg = fmt.Sprintf("finished with error: %s", err)
		exitCode = 1
	}

	log.Println(onExitMsg)
	os.Exit(exitCode)
}

type Server struct {
	status.UnimplementedStatusAPIServer
}

func (s *Server) CheckStatus(ctx context.Context, r *status.CheckStatusRequest) (*status.CheckStatusResponse, error) {
	gotStatus := r.GetStatus()

	log.Println("Raw status: ", gotStatus)
	log.Println("String: ", gotStatus.String())
	log.Println("Number: ", gotStatus.Number())

	return nil, nil
}
