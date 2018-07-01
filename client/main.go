package main

import (
	"context"
	"flag"
	"grpc-bockchain/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

var client proto.BlockchainClient

func main() {
	addFlag := flag.Bool("add", false, "Add new Block")
	listFlag := flag.Bool("list", false, "Get all Blocks in Blockchain")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannt dial server: %v", err)
	}

	client = proto.NewBlockchainClient(conn)
	if *addFlag {
		addBlock()
	}

	if *listFlag {
		getBlockchain()
	}
}

func addBlock() {
	bl, err := client.AddBlock(
		context.Background(),
		&proto.AddBlockReq{
			Data: time.Now().String(),
		},
	)
	if err != nil {
		log.Fatalf("Cannt add new Block. Err: %v", err)
	}
	log.Printf("Added new Block. Hash is %s\n", bl.Hash)
}

func getBlockchain() {
	bc, err := client.GetBlockchain(
		context.Background(),
		&proto.GetBlockchainReq{},
	)
	if err != nil {
		log.Fatalf("Cannt get Blockchain. Err: %v", err)
	}
	log.Print("Blockchain:\n")
	for _, b := range bc.Blocks {
		log.Printf("%+v \n", b)
	}
}
