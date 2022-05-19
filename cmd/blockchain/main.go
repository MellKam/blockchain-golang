package main

import (
	"fmt"
	"strconv"

	"github.com/MellKam/blockchain-golang/blockchain"
)

func main() {
	b := blockchain.NewBlockchain()

	b.AddBlock("First block after genesis")
	b.AddBlock("Second block after genesis")
	b.AddBlock("Third block after genesis")

	var block *blockchain.Block

	iterator := b.NewBlockchainIterator()
	for {
		block = iterator.Next()

		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("PreviousHash: %x\n", block.PreviousHash)
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("Validity: %s\n", strconv.FormatBool(pow.ValidateBlockHash()))

		fmt.Println()

		if block.PreviousHash == [32]byte{} {
			break
		}
	}
}
