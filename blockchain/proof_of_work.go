package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"

	"github.com/MellKam/blockchain-golang/pkg/converter"
)

const Difficulty uint = 8

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, HashBitsNumber-Difficulty)

	return &ProofOfWork{block, target}
}

func (pow ProofOfWork) createBlockData(nonce uint) []byte {
	return bytes.Join(
		[][]byte{
			pow.Block.PreviousHash[:],
			pow.Block.Data,
			converter.UintToByteSlice(nonce),
			converter.UintToByteSlice(Difficulty),
		},
		[]byte{},
	)
}

func (pow ProofOfWork) MineBlock() (uint, HashType) {
	var intHash big.Int
	var hash HashType
	var nonce uint = 0

	for nonce < math.MaxInt64 {
		data := pow.createBlockData(nonce)
		hash = sha256.Sum256(data)

		fmt.Println(hash)

		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash
}

func (pow *ProofOfWork) ValidateBlock() bool {
	var intHash big.Int

	data := pow.createBlockData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}
