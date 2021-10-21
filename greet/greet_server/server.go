package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/shekhar82/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.Greeting.FirstName

	result := "Hello " + firstName + "!"

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Hello Server")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to start %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}

}
