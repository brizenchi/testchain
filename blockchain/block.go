package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash    []byte
	Data    []byte
	PreHash []byte
	Nonce   int
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.Hash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}
func CreateBlock(data string, preHash []byte) *Block {
	block := &Block{
		Hash:    []byte{},
		Data:    []byte(data),
		PreHash: preHash,
		Nonce:   0,
	}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

type BlockChain struct {
	Blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) {
	preBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, preBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
func InitBlockChain() *BlockChain {
	return &BlockChain{Blocks: []*Block{Genesis()}}
}
