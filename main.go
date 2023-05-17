package main

import (
	"fmt"
	"strconv"
	"testchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("first")
	chain.AddBlock("second")
	chain.AddBlock("third")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x \n", block.PreHash)
		fmt.Printf("Data in Block: %s  \n", block.Data)
		fmt.Printf("Hash: %x \n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	}
}
