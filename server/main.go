package main

import (
	"grpc-bockchain/proto"
	"grpc-bockchain/server/blockchain"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Cann't listen on port 8080. Because: %v", err)
	}

	server := grpc.NewServer()
	proto.RegisterBlockchainServer(server, &Server{
		Blockchain: blockchain.NewBlockchain(),
	})
	server.Serve(listener)
}

// Server is struct to implement func interface
type Server struct {
	Blockchain *blockchain.Blockchain
}

// AddBlock is a simpe RPC to create new BC
func (s *Server) AddBlock(ctx context.Context, in *proto.AddBlockReq) (*proto.AddBlockRes, error) {
	block := s.Blockchain.AddBlock(in.Data)
	return &proto.AddBlockRes{
		Hash: block.Hash,
	}, nil
}

// GetBlockchain is a simpe RPC to get a BC
func (s *Server) GetBlockchain(ctx context.Context, in *proto.GetBlockchainReq) (*proto.GetBlockchainRes, error) {
	res := new(proto.GetBlockchainRes)
	for _, b := range s.Blockchain.Blocks {
		res.Blocks = append(res.Blocks, &proto.Block{
			PrevBlockHash: b.PrevBlockHash,
			Hash:          b.Hash,
			Data:          b.Data,
		})
	}
	return res, nil
}
