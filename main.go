package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash    []byte
	Data    []byte
	PreHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.Hash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}
func CreateBlock(data string, preHash []byte) *Block {
	block := &Block{
		Data:    []byte(data),
		PreHash: preHash,
	}
	block.DeriveHash()
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

func main() {
	chain := InitBlockChain()
	chain.AddBlock("first")
	chain.AddBlock("second")
	chain.AddBlock("third")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x \n", block.PreHash)
		fmt.Printf("Data in Block: %s  \n", block.Data)
		fmt.Printf("Hash: %x \n", block.Hash)
	}
}
