package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/shekhar82/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.Result, error) {
	fmt.Printf("Sum method was invoked for %v\n", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber

	sum := firstNumber + secondNumber

	result := &calculatorpb.Result{
		Result: sum,
	}

	return result, nil
}

func (*server) Subtract(ctx context.Context, req *calculatorpb.SubtractRequest) (*calculatorpb.Result, error) {
	fmt.Printf("Subtract method was invoked for %v\n", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber

	subtract := firstNumber - secondNumber

	result := &calculatorpb.Result{
		Result: subtract,
	}

	return result, nil
}

func (*server) Multiply(ctx context.Context, req *calculatorpb.MultiplyRequest) (*calculatorpb.Result, error) {
	fmt.Printf("Multiply method was invoked for %v\n", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber

	multiplication := firstNumber * secondNumber

	result := &calculatorpb.Result{
		Result: multiplication,
	}

	return result, nil
}

func (*server) Divide(ctx context.Context, req *calculatorpb.DivisonRequest) (*calculatorpb.Result, error) {
	fmt.Printf("Divide method was invoked for %v\n", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber

	if secondNumber == 0 {
		return nil, fmt.Errorf("you can't use zero in denominator %v", req)
	}

	div := firstNumber / secondNumber

	result := &calculatorpb.Result{
		Result: div,
	}

	return result, nil
}

func main() {
	fmt.Println("Hello Calculator Server....")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to start %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
