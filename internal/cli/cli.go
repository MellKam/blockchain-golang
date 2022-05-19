package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/MellKam/blockchain-golang/internal/blockchain"
	"github.com/MellKam/blockchain-golang/pkg/handler"
)

type BlockchainCli struct {
	Blockchain *blockchain.Blockchain
}

const cliUsage = `BlockchainCLI usage:
	- print
		Print the whole chain of blocks
	- add -block "BLOCK_DATA"
		Create new block and add it to the chain
`

func (cli *BlockchainCli) printUsage() {
	fmt.Printf(cliUsage)
}

func (cli *BlockchainCli) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *BlockchainCli) addBlock(data string) {
	block, err := cli.Blockchain.AddBlock(data)
	if err != nil {
		fmt.Println("Error during block creation!")
		return
	}

	pow := blockchain.NewProofOfWork(block)
	isBlockValid := strconv.FormatBool(pow.ValidateBlockHash())

	cli.pringBlock(block, isBlockValid)
}

func (cli *BlockchainCli) pringBlock(block *blockchain.Block, isBlockValid string) {
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Hash: %x\n", block.Hash)
	fmt.Printf("PreviousHash: %x\n", block.PreviousHash)
	fmt.Printf("Validity: %s\n", isBlockValid)

	fmt.Println()
}

func (cli *BlockchainCli) printChain() {
	var block *blockchain.Block
	var pow *blockchain.ProofOfWork
	iterator := cli.Blockchain.NewBlockchainIterator()

	for {
		block = iterator.Next()
		pow = blockchain.NewProofOfWork(block)

		cli.pringBlock(block, strconv.FormatBool(pow.ValidateBlockHash()))

		if block.PreviousHash == [32]byte{} {
			break
		}
	}
}

func (cli *BlockchainCli) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		handler.HandlePossibleError(err)

		if addBlockCmd.Parsed() {
			fmt.Println(*addBlockData)

			if *addBlockData == "" {
				addBlockCmd.Usage()
				runtime.Goexit()
			}

			cli.addBlock(*addBlockData)
		}

	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		handler.HandlePossibleError(err)

		if printChainCmd.Parsed() {
			cli.printChain()
		}

	default:
		cli.printUsage()
		runtime.Goexit()
	}
}
