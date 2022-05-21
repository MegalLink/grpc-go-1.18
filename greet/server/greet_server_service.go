package main

import (
	"context"
	"fmt"
	"io"
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

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked with \n")

	res := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{Result: fmt.Sprint("Greetings:" + res)})
		}
		if err != nil {
			log.Fatalf("stream reading error: %v\n", err)
		}
		greeting := LanguajeGreet[req.Languaje]

		res += fmt.Sprintf("greeting: %s ", greeting)
	}
}

func (s *Server) GreetEveryOne(stream pb.GreetService_GreetEveryOneServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream reading error: %v\n", err)
		}
		greeting := LanguajeGreet[req.Languaje]

		if err := stream.Send(&pb.GreetResponse{Result: fmt.Sprintf("%s->%s", req.FirstName, greeting)}); err != nil {
			log.Fatalf("stream sending error: %v\n", err)
		}
	}
	return nil
}
