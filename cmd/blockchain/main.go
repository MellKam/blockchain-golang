package main

import (
	"fmt"

	"github.com/MellKam/blockchain-golang/blockchain"
)

func main() {
	b := blockchain.NewBlockchain()

	b.AddBlock("First block after genesis")
	b.AddBlock("Second block after genesis")
	b.AddBlock("Third block after genesis")

	for i, block := range b.Blocks {
		fmt.Printf(`
			--- Block #%d
			PreviousHash: %x
			Hash: %x
			Data: %s
			Nonce: %d
		`, i, block.PreviousHash, block.Hash, block.Data, block.Nonce)
	}
}
