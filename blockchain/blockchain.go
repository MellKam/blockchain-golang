package blockchain

import (
	"fmt"

	excepion "github.com/MellKam/blockchain-golang/pkg/exception"
	badger "github.com/dgraph-io/badger/v3"
)

type Blockchain struct {
	LastHash HashType
	Database *badger.DB
}

func NewBlockchain() *Blockchain {
	options := badger.DefaultOptions(DB_DATA_PATH)

	// create badger database
	db, err := badger.Open(options)
	excepion.HandleError(err)

	lastHash, err := initDatabase(db)
	excepion.HandleError(err)

	return &Blockchain{lastHash, db}
}

func initDatabase(db *badger.DB) (HashType, error) {
	var lastHash HashType

	err := db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(LAST_HASH_KEY))

		// if not found blockchain in database,
		// then create new and set it to db
		if err == badger.ErrKeyNotFound {
			fmt.Println("Not found existing blockchain")

			genesis := createGenesisBlock()
			err := SetBlockToDb(txn, genesis)
			lastHash = genesis.Hash

			return err
		}

		excepion.HandleError(err)

		fmt.Println("Found existing blockchain")

		// if successfully found lastHash, then blockchain
		// already initialized in database and we wand
		// to get lastHash:[32]byte from this item
		data, err := item.ValueCopy(nil)
		excepion.HandleError(err)

		copy(lastHash[:], data[:32])

		return err
	})

	return lastHash, err
}

func (b *Blockchain) AddBlock(data string) {
	block := NewBlock(b.LastHash, data)
	mineBlockWithPOW(block)

	err := b.Database.Update(func(txn *badger.Txn) error {
		return SetBlockToDb(txn, block)
	})
	excepion.HandleError(err)

	b.LastHash = block.Hash
}

func mineBlockWithPOW(block *Block) {
	pow := NewProofOfWork(block)
	nonce, hash := pow.MineBlock()

	block.Hash = hash
	block.Nonce = nonce
}

func createGenesisBlock() *Block {
	genesis := NewBlock([32]byte{}, "Genesis Block")
	mineBlockWithPOW(genesis)

	return genesis
}

func (b *Blockchain) NewBlockchainIterator() *BlockchainIterator {
	return &BlockchainIterator{b.LastHash, b.Database}
}
