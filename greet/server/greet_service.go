package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/MegalLink/grpc-go-1.18/greet/proto"
)

func (ss *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{Result: fmt.Sprintf("Hello %s", in.FirstName)}, nil
}
