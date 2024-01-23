package main

import (
	"fmt"
	"log"

	"client/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	status "google.golang.org/grpc/status"
)

func main() {

	creds := insecure.NewCredentials()
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()
	calculcatorClient := services.NewCalculatorClient(cc)
	calculatorService := services.NewCalculatorService(calculcatorClient)
	err = calculatorService.Hello("")
	/* err = calculatorService.Fibonacci(10) */
	/* err = calculatorService.Average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) */
	/* err = calculatorService.Sum(1, 2, 3, 4, 5, 6) */
	if err != nil {
		if grpcErr, ok := status.FromError(err); ok {
			fmt.Printf("[%v] %v\n", grpcErr.Code(), grpcErr.Message())
		} else {
			log.Fatal(err)
		}
	}
}
