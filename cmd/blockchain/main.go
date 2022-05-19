package main

import (
	"os"

	"github.com/MellKam/blockchain-golang/internal/blockchain"
	"github.com/MellKam/blockchain-golang/internal/cli"
)

func main() {
	defer os.Exit(0)

	b := blockchain.NewBlockchain()
	defer b.Database.Close()

	cli := cli.BlockchainCli{Blockchain: b}
	cli.Run()
}
