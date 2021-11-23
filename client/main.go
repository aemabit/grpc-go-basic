package main

import (
	"blockchain-grpc/pb/blockpb"
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
)

var client blockpb.BlockchainClient

func main() {
	addFlag := flag.Bool("add", false, "add new block")
	listFlat := flag.Bool("list", false, "get the block")
	flag.Parse()

	conn, err := grpc.Dial("localhost: 50052", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}

	client = blockpb.NewBlockchainClient(conn)

	if *addFlag {
		addBlock()
	}

	if *listFlat {
		getBlockchain()
	}
}

func addBlock() {
	block, err := client.AddBlock(context.Background(), &blockpb.AddBlockRequest{
		Data: time.Now().String(),
	})

	if err != nil {
		log.Fatalf("unable to add block")
	}

	log.Printf("new block hash %s \n", block.Hash)
}

func getBlockchain() {
	bc, err := client.GetBlockchain(context.Background(), &blockpb.GetBlockchainRequest{})

	if err != nil {
		log.Fatalf("unable to get blockchain", err)
	}

	log.Println("blocks:")
	for _, b := range bc.Blocks {
		log.Printf("hash: %s, prev block hash: %s, data: %s", b.Hash, b.PrevBlockHash, b.Data)
	}

}
