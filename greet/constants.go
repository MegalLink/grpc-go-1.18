package greet

import (
	pb "github.com/MegalLink/grpc-go-1.18/greet/proto"
)

// constants
const ServerAddress = "localhost:50051"

//types
type Server struct {
	pb.GreetServiceServer
}
