package blockchain

type Blockchain struct {
	Blocks []*Block
}

func (chain *Blockchain) AddBlock(data string) {
	lastBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := NewBlock(lastBlock.Hash, data)

	pow := NewProofOfWork(newBlock)
	nonce, hash := pow.MineBlock()
	newBlock.Hash = hash
	newBlock.Nonce = nonce

	chain.Blocks = append(chain.Blocks, newBlock)
}

func (chain *Blockchain) createGenesisBlock() *Block {
	return NewBlock([32]byte{}, "Genesis Block")
}

func NewBlockchain() *Blockchain {
	blockchain := &Blockchain{}
	blockchain.Blocks = []*Block{blockchain.createGenesisBlock()}

	return blockchain
}
