package main

import (
	"fmt"
	"grpc-learn/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()
	services.RegisterCalculatorServer(s, services.NewCalculatorServer())

	reflection.Register(s)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on :50051")
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
