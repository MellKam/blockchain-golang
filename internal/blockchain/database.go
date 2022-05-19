package blockchain

import (
	"github.com/MellKam/blockchain-golang/pkg/handler"
	"github.com/dgraph-io/badger/v3"
)

const DB_DATA_PATH = "../../tmp/chain"
const LAST_HASH_KEY = "LastHash"

func SetBlockToDb(txn *badger.Txn, block *Block) error {
	err := txn.Set(block.Hash[:], block.SerializeToBytes())
	handler.HandlePossibleError(err)

	// we nead to reset lastHash because we create new block
	// and now lastHasn will be hash of this block
	err = txn.Set([]byte(LAST_HASH_KEY), block.Hash[:])

	return err
}

func GetBlockFromDb(txn *badger.Txn, hash HashType) (*Block, error) {
	item, err := txn.Get(hash[:])
	handler.HandlePossibleError(err)

	byteBlock, err := item.ValueCopy(nil)
	block := DeserializeBytesToBlock(byteBlock)

	return block, err
}
