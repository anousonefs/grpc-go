package services

import (
	context "context"
	"fmt"
	"io"
	"time"
)

type CalculatorService interface {
	Hello(name string) error
	Fibonacci(n uint32) error
}

type calculatorService struct {
	calculatorClient CalculatorClient
}

func NewCalculatorService(calculatorClient CalculatorClient) CalculatorService {
	return calculatorService{calculatorClient}
}

func (base calculatorService) Hello(name string) error {
	req := HelloRequest{
		Name: name,
	}
	res, err := base.calculatorClient.Hello(context.Background(), &req)
	if err != nil {
		return err
	}
	fmt.Printf("service: hello \n")
	fmt.Printf("Request : %v\n", req.Name)
	fmt.Printf("Response : %v\n", res.Result)
	return nil
}

func (base calculatorService) Fibonacci(n uint32) error {
	req := FibonacciRequest{
		N: n,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stream, err := base.calculatorClient.Fibonacci(ctx, &req)
	if err != nil {
		return err
	}
	fmt.Printf("service: fibonacci \n")
	fmt.Printf("Request : %v\n", req.N)
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("Response : %v\n", res.Result)
	}
	return nil
}
