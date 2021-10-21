package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/shekhar82/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't connect %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//doUnary(c)
	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Shekhar",
			LastName:  "Gupta",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Errro while calling Greet RPC %v", err)
	}

	log.Printf("Response from greet  %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Server streaming RPC...")

	req := &greetpb.GreetManyTimeRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Shekhar",
			LastName:  "Gupta",
		},
	}
	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling server stream Greet Many times RPC %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// We have reached at the end of stream
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}

		log.Printf("Response from Greet Many times %v\n", msg.GetResult())
	}
}
