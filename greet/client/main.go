package main

import (
	"log"

	constants "github.com/MegalLink/grpc-go-1.18/greet"
	pb "github.com/MegalLink/grpc-go-1.18/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(constants.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect : %v", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
}
