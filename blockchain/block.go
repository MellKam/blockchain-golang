package blockchain

const HashBitsNumber uint = 256

type HashType [32]byte

type Block struct {
	Hash         HashType
	PreviousHash HashType
	Data         []byte
	Nonce        uint
}

func NewBlock(previousHash HashType, data string) *Block {
	return &Block{PreviousHash: previousHash, Data: []byte(data), Nonce: 0}
}
