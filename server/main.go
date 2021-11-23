package main

import (
	"blockchain-grpc/pb/blockpb"
	services "blockchain-grpc/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50052")

	if err != nil {
		log.Fatalf("unable to listen on 50052 port: %v", err)
	}

	s := grpc.NewServer()
	blockpb.RegisterBlockchainServer(s, &services.BlockchainService{})

	s.Serve(listener)
}
