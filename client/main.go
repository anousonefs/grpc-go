package main

import (
	"log"

	"client/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	/* err = calculatorService.Hello("John") */
	/* err = calculatorService.Fibonacci(10) */
	/* err = calculatorService.Average(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) */
	err = calculatorService.Sum(1, 2, 3, 4, 5, 6)
	if err != nil {
		log.Fatal(err)
	}
}
