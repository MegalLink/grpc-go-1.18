package main

import (
	"context"
	"io"
	"log"
	"time"

	constants "github.com/MegalLink/grpc-go-1.18/greet"
	pb "github.com/MegalLink/grpc-go-1.18/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Jeferson",
		Languaje:  constants.RU,
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting %s\n", res.Result)
}

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")
	req := &pb.GreetRequest{
		FirstName: "Jeferson",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream reading error: %v\n", err)
		}

		log.Printf("GreetingManyTimes %s\n", msg.Result)
	}
}

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	requests := []*pb.GreetRequest{
		{Languaje: constants.EN},
		{Languaje: constants.JP},
		{Languaje: constants.HI},
		{Languaje: constants.IT},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Could not LongGreet times: %v\n", err)
	}

	for _, req := range requests {
		log.Printf("Request: %v", req)
		stream.SendMsg(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error getting response from LongGreat: %v\n", err)
	}

	log.Printf("LongGreet %s\n", res.Result)
}
