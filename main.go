package main

// Block
// struct - prevHash, data, hash, nonce
// func - CreateBlock, Genesis

// Blockchain
// struct - []Block
// func - InitBlockchain
// method - AddBlock

// ProofOfWork
// struct - block, target
// func - NewProof, ToHex
// method - InitData, RunProof, Validate

import (
	"fmt"
	"strconv"

	bc "github.com/bhanupbalusu/gobc_practice/blockchain"
)

func main() {
	chain := bc.InitBlockchain()
	chain.AddBlock("First Block after Genesis Block.")
	chain.AddBlock("Second Block after Genesis Block.")
	chain.AddBlock("Third Block after Genesis Block.")

	for _, block := range chain.Blocks {
		fmt.Printf("\nPrevious Hash : %x\n", block.PrevHash)
		fmt.Printf("Block Data : %s\n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Printf("Nonce : %d\n", block.Nonce)
		pow := bc.NewProofOfWork(block)
		fmt.Printf("Block is Validated : %s\n", strconv.FormatBool(pow.Validate()))
	}
}
