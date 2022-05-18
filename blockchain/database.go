package blockchain

import (
	excepion "github.com/MellKam/blockchain-golang/pkg/exception"
	badger "github.com/dgraph-io/badger/v3"
)

const DB_DATA_PATH = "../../tmp/chain"
const LAST_HASH_KEY = "LastHash"

func SetBlockToDb(txn *badger.Txn, block *Block) error {
	err := txn.Set(block.Hash[:], block.SerializeToBytes())
	excepion.HandleError(err)

	err = txn.Set([]byte(LAST_HASH_KEY), block.Hash[:])

	return err
}

func GetBlockFromDb(txn *badger.Txn, hash HashType) (*Block, error) {
	item, err := txn.Get(hash[:])
	excepion.HandleError(err)

	byteBlock, err := item.ValueCopy(nil)
	block := DeserializeBytesToBlock(byteBlock)

	return block, err
}
