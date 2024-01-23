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
	err = calculatorService.Fibonacci(10)
	if err != nil {
		log.Fatal(err)
	}
}
