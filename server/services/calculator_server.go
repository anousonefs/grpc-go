package services

import (
	context "context"
	"fmt"
	"io"
	"time"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type calculatorServer struct {
}

func NewCalculatorServer() CalculatorServer {
	return calculatorServer{}
}

func (calculatorServer) mustEmbedUnimplementedCalculatorServer() {}

func (calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}
	result := fmt.Sprintf("Hello %s", req.Name)
	res := HelloResponse{
		Result: result,
	}
	return &res, nil
}

func (calculatorServer) Fibonacci(req *FibonacciRequest, stream Calculator_FibonacciServer) error {
	for n := uint32(0); n <= req.N; n++ {
		result := fib(n)
		res := FibonacciResponse{
			Result: result,
		}
		stream.Send(&res)
		time.Sleep(time.Second)
	}
	return nil
}

func fib(n uint32) uint32 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}

func (calculatorServer) Average(stream Calculator_AverageServer) error {
	sum := 0.0
	count := 0.0
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += res.Number
		count++
	}
	res := AverageResponse{
		Result: sum / count,
	}
	return stream.SendAndClose(&res)
}

func (calculatorServer) Sum(stream Calculator_SumServer) error {
	sum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		res := SumResponse{
			Result: sum,
		}
		if err := stream.Send(&res); err != nil {
			return err
		}
	}
	return nil
}
