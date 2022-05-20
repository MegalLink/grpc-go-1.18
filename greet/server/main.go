package main

import (
	"log"
	"net"

	constants "github.com/MegalLink/grpc-go-1.18/greet"
	pb "github.com/MegalLink/grpc-go-1.18/greet/proto"
	"google.golang.org/grpc"
)

//types
type Server struct {
	pb.GreetServiceServer
}

func main() {
	listener, err := net.Listen("tcp", constants.ServerAddress)
	if err != nil {
		log.Fatalf("Failed to listen on : %v", err)
	}

	log.Printf("Listening on %s\n", constants.ServerAddress)
	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
