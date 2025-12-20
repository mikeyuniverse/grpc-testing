package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	increment "github.com/mikeyuniverse/grpc-testing/bidistreaming/pkg/api/v1"
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

	increment.RegisterIncrementAPIServer(server, &Server{})

	log.Printf("grpc server at %v\n", listener.Addr())
	err = server.Serve(listener)

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
	increment.UnimplementedIncrementAPIServer
}

func (s *Server) Inc(r grpc.BidiStreamingServer[increment.IncValue, increment.IncValue]) error {
	for {
		req, err := r.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}

			return fmt.Errorf("recv: %w", err)
		}

		log.Println("got value:", req.Value)

		req.Value++

		log.Println("send value:", req.Value)

		err = r.Send(req)
		if err != nil {
			return fmt.Errorf("send: %w", err)
		}
	}
}
