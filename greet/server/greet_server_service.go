package main

import (
	"context"
	"fmt"
	"log"

	constants "github.com/MegalLink/grpc-go-1.18/greet"
	pb "github.com/MegalLink/grpc-go-1.18/greet/proto"
)

var LanguajeGreet = map[string]string{
	constants.ES: "Hola",
	constants.FR: "Bonjour ",
	constants.EN: "Hello",
	constants.IT: "Ciao",
	constants.RU: "Priviet",
	constants.HI: "Namaste",
	constants.JP: "Konnichi wa",
}

func (s *Server) Greet(ctx context.Context, request *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", request)
	greetResponse := LanguajeGreet[request.Languaje]

	return &pb.GreetResponse{Result: fmt.Sprintf("%s %s", greetResponse, request.FirstName)}, nil
}

func (s *Server) GreetManyTimes(request *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", request)
	// for maps is key,value
	for _, value := range LanguajeGreet {
		res := fmt.Sprintf("%s ,%s", value, request.FirstName)
		stream.Send(&pb.GreetResponse{Result: res})
	}

	return nil
}
