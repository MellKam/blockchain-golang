package blockchain

import (
	"bytes"
	"encoding/gob"

	"github.com/MellKam/blockchain-golang/pkg/exception"
)

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

func (b *Block) SerializeToBytes() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(b)
	excepion.HandleError(err)

	return buffer.Bytes()
}

func DeserializeBytesToBlock(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	excepion.HandleError(err)

	return &block
}
