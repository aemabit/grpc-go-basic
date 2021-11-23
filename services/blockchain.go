package services

import (
	"blockchain-grpc/pb/blockpb"
	"context"
	"crypto/sha256"
	"encoding/hex"
)

type BlockchainService struct {
	blockpb.UnimplementedBlockchainServer
}

type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

type Blockchain struct {
	Blocks []*Block
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

func NewBlock(data string, prevBlockHash string) *Block {
	block := &Block{
		Data:          data,
		PrevBlockHash: prevBlockHash,
	}

	return block
}

func (bc *Blockchain) AddBlockFactory(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)

	return newBlock
}

func NewBlockchain() *Blockchain {

	return &Blockchain{
		Blocks: []*Block{NewGenesisBlock()},
	}
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "")
}

func (s *BlockchainService) AddBlock(ctx context.Context, in *blockpb.AddBlockRequest) (*blockpb.AddBlockResponse, error) {

	block := NewBlockchain().AddBlockFactory(in.Data)

	return &blockpb.AddBlockResponse{
		Hash: block.Hash,
	}, nil
}

func (s *BlockchainService) GetBlockchain(ctx context.Context, in *blockpb.GetBlockchainRequest) (*blockpb.GetBlockchainResponse, error) {
	resp := new(blockpb.GetBlockchainResponse)

	for _, b := range NewBlockchain().Blocks {
		resp.Blocks = append(resp.Blocks, &blockpb.Block{
			PrevBlockHash: b.PrevBlockHash,
			Hash:          b.Hash,
			Data:          b.Data,
		})
	}
	return resp, nil
}
