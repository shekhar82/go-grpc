package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/shekhar82/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't connect %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	var wg sync.WaitGroup

	wg.Add(4)
	go func() {
		defer wg.Done()
		doSum(c)
	}()

	//wg.Add(1)
	go func() {
		defer wg.Done()
		doSubtraction(c)
	}()

	//wg.Add(1)
	go func() {
		defer wg.Done()
		doMultiplication(c)
	}()

	//wg.Add(1)
	go func() {
		defer wg.Done()
		doDivision(c)
	}()

	wg.Wait()
}

func doSum(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do unary RPC Sum...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 4,
	}

	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("Errro while calling Sum RPC %v", err)
	}

	log.Printf("Response from Sum  %v", res.Result)
}

func doSubtraction(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do unary RPC Subtraction...")
	req := &calculatorpb.SubtractRequest{
		FirstNumber:  10,
		SecondNumber: 4,
	}

	res, err := c.Subtract(context.Background(), req)

	if err != nil {
		log.Fatalf("Errro while calling Subtrat RPC %v", err)
	}

	log.Printf("Response from Subtract  %v", res.Result)
}

func doMultiplication(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do unary RPC Multiplication...")
	req := &calculatorpb.MultiplyRequest{
		FirstNumber:  10,
		SecondNumber: 4,
	}

	res, err := c.Multiply(context.Background(), req)

	if err != nil {
		log.Fatalf("Errro while calling Multiply RPC %v", err)
	}

	log.Printf("Response from Multiply  %v", res.Result)
}

func doDivision(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do unary RPC Division...")

	req := &calculatorpb.DivisonRequest{
		FirstNumber:  10,
		SecondNumber: 2,
	}

	res, err := c.Divide(context.Background(), req)

	if err != nil {
		log.Fatalf("Errro while calling Divison RPC %v", err)
	}

	log.Printf("Response from Divison  %v", res.Result)
}
