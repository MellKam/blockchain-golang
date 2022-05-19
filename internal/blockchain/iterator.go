package blockchain

import (
	"github.com/MellKam/blockchain-golang/pkg/handler"
	"github.com/dgraph-io/badger/v3"
)

type BlockchainIterator struct {
	currentHash HashType
	database    *badger.DB
}

func (iter *BlockchainIterator) Next() *Block {
	var block *Block

	err := iter.database.View(func(txn *badger.Txn) error {
		var err error
		block, err = GetBlockFromDb(txn, iter.currentHash)

		return err
	})
	handler.HandlePossibleError(err)

	iter.currentHash = block.PreviousHash

	return block
}
