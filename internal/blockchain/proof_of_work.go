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

	// 1 to the power of (256 - Difficulty)
	target.Lsh(target, HashBitsNumber-Difficulty)

	return &ProofOfWork{block, target}
}

func (pow ProofOfWork) createBlockData(nonce uint) []byte {
	return bytes.Join(
		[][]byte{
			pow.Block.PreviousHash[:],
			pow.Block.Data,
			converter.IntegerToByteSlice(nonce),
			converter.IntegerToByteSlice(Difficulty),
		},
		[]byte{},
	)
}

func (pow ProofOfWork) MineBlock() (uint, HashType) {
	var (
		intHash *big.Int = big.NewInt(0)
		hash    HashType = [32]byte{}
		nonce   uint     = 0
		data    []byte
	)

	for nonce < math.MaxInt64 {
		data = pow.createBlockData(nonce)
		hash = sha256.Sum256(data)

		intHash.SetBytes(hash[:])

		// if intHash < target so we found
		// right hash and we can exit from cycle
		if intHash.Cmp(pow.Target) == -1 {
			break
		}

		// otherwise increment nonce and rerun cycle
		// until we found correct hash
		nonce++
	}

	return nonce, hash
}

func (pow *ProofOfWork) ValidateBlockHash() bool {
	data := pow.createBlockData(pow.Block.Nonce)
	hash := sha256.Sum256(data)

	if hash == pow.Block.Hash {
		return true
	}

	return false
}
