package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// Block stores basic infomation
type Block struct {
	Hash          string
	PrevBlockHash string
	Data          string
}

// Blockchain is list of block
type Blockchain struct {
	Blocks []*Block
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

// NewBlock create new block
func NewBlock(data string, prevBlockHash string) *Block {
	bl := &Block{
		PrevBlockHash: prevBlockHash,
		Data:          data,
	}
	bl.setHash()
	return bl
}

// AddBlock creates a new block and add it to blockchain
func (bc *Blockchain) AddBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
	return newBlock
}

// NewBlockchain init new blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewBlock("Genesis Block", "")}}
}
