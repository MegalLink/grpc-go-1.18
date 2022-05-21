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
		if err := stream.SendMsg(req); err != nil {
			log.Fatalf("stream sending error: %v\n", err)
		}
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error getting response from LongGreat: %v\n", err)
	}

	log.Printf("LongGreet %s\n", res.Result)
}

func doGreatGreetEveryOne(c pb.GreetServiceClient, wait bool) {
	requests := []pb.GreetRequest{
		{Languaje: constants.EN, FirstName: "Jeff"},
		{Languaje: constants.JP, FirstName: "Lizz"},
		{Languaje: constants.HI, FirstName: "Dome"},
		{Languaje: constants.IT, FirstName: "Henry"},
	}
	stream, err := c.GreetEveryOne(context.Background())
	if err != nil {
		log.Fatalf("Could not LongGreet times: %v\n", err)
	}
	if wait {
		sendAndReceivaAll(stream, requests)
	} else {
		sendAndReceiveOneByOne(stream, requests)
	}
}

func sendAndReceivaAll(stream pb.GreetService_GreetEveryOneClient, requests []pb.GreetRequest) {
	for _, req := range requests {
		log.Printf("Request: %v", &req)
		if err := stream.SendMsg(&req); err != nil {
			log.Fatalf("stream sending error: %v\n", err)
		}
		time.Sleep(1 * time.Second)
	}
	stream.CloseSend()
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream reading error: %v\n", err)
			break
		}

		log.Printf("GreetEveryOne %s\n", msg.Result)
	}
}

func sendAndReceiveOneByOne(stream pb.GreetService_GreetEveryOneClient, requests []pb.GreetRequest) {
	waitc := make(chan struct{})
	go func() {
		for _, req := range requests {
			log.Printf("Request: %v", &req)
			if err := stream.SendMsg(&req); err != nil {
				log.Fatalf("stream sending error: %v\n", err)
			}
			time.Sleep(2 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("stream reading error: %v\n", err)
				break
			}

			log.Printf("GreetEveryOne %s\n", msg.Result)
		}

		close(waitc)
	}()

	<-waitc
}
