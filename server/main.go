package main

import (
	"flag"
	"fmt"
	"grpc-learn/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {

	var s *grpc.Server
	tls := flag.Bool("tls", false, "Connection uses TLS if true, else plain")
	flag.Parse()
	if *tls {
		certFile := "../tls/server.crt"
		keyFile := "../tls/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatal(err)
		}
		s = grpc.NewServer(grpc.Creds(creds))
	} else {
		s = grpc.NewServer()
	}

	services.RegisterCalculatorServer(s, services.NewCalculatorServer())

	reflection.Register(s)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Listening on :50051")
	if *tls {
		fmt.Println(" with TLS")
	} else {
		fmt.Println()
	}
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
