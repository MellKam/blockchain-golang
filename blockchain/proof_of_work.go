package blockchain

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"

	"github.com/MellKam/blockchain-golang/pkg/converter"
)

// Ratio 4 to 1
// 00xxxx... = 8 Difficulty
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
			converter.NumberToByteSlice(nonce),
			converter.NumberToByteSlice(Difficulty),
		},
		[]byte{},
	)
}

func (pow ProofOfWork) MineBlock() (uint, HashType) {
	var (
		intHash *big.Int = big.NewInt(0)
		hash    HashType = [32]byte{}
		nonce   uint     = 0
	)

	for nonce < math.MaxInt64 {
		isBlockValid := pow.validateBlock(nonce, &hash, intHash)

		if isBlockValid {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash
}

func (pow *ProofOfWork) validateBlock(
	nonce uint,
	hash *HashType,
	intHash *big.Int,
) bool {
	data := pow.createBlockData(nonce)

	*hash = sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	// if our intHash less then target
	// then we fount right hash
	return intHash.Cmp(pow.Target) == -1
}
