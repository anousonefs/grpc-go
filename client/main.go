package main

import (
	"flag"
	"fmt"
	"log"

	"client/services"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	status "google.golang.org/grpc/status"
)

func main() {

	var cc *grpc.ClientConn
	var err error
	var creds credentials.TransportCredentials

	host := flag.String("host", "localhost:50051", "gRPC server host")
	tls := flag.Bool("tls", false, "Connection uses TLS if true, else plain")
	flag.Parse()
	if *tls {
		certFile := "../tls/ca.crt"
		creds, err = credentials.NewClientTLSFromFile(certFile, "")
	} else {
		creds = insecure.NewCredentials()
	}

	cc, err = grpc.Dial(*host, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	calculcatorClient := services.NewCalculatorClient(cc)
	calculatorService := services.NewCalculatorService(calculcatorClient)
	err = calculatorService.Hello("anousone")
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
