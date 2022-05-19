package blockchain

import (
	"fmt"

	"github.com/MellKam/blockchain-golang/pkg/handler"
	"github.com/dgraph-io/badger/v3"
)

type Blockchain struct {
	LastHash HashType
	Database *badger.DB
}

func NewBlockchain() *Blockchain {
	options := badger.DefaultOptions(DB_DATA_PATH)
	options.Logger = nil

	// create badger database
	db, err := badger.Open(options)
	handler.HandlePossibleError(err)

	lastHash, err := initDatabase(db)
	handler.HandlePossibleError(err)

	return &Blockchain{lastHash, db}
}

// Create struct that are responsible for iteration throught chain blocks
func (b *Blockchain) NewBlockchainIterator() *BlockchainIterator {
	return &BlockchainIterator{b.LastHash, b.Database}
}

func initDatabase(db *badger.DB) (HashType, error) {
	var lastHash HashType

	err := db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(LAST_HASH_KEY))

		// if not found blockchain in database,
		// then create new and set it to db
		if err == badger.ErrKeyNotFound {
			fmt.Printf("Not found existing blockchain\n\n")

			genesis := createGenesisBlock()
			err := SetBlockToDb(txn, genesis)
			lastHash = genesis.Hash

			return err
		}

		handler.HandlePossibleError(err)

		fmt.Printf("Found existing blockchain\n\n")

		// if successfully found lastHash, then blockchain
		// already initialized in database and we wand
		// to get lastHash:[32]byte from this item
		data, err := item.ValueCopy(nil)
		handler.HandlePossibleError(err)

		copy(lastHash[:], data[:32])

		return err
	})

	return lastHash, err
}

func (b *Blockchain) AddBlock(data string) (*Block, error) {
	block := NewBlock(b.LastHash, data)
	mineBlockWithPOW(block)

	err := b.Database.Update(func(txn *badger.Txn) error {
		return SetBlockToDb(txn, block)
	})
	handler.HandlePossibleError(err)

	b.LastHash = block.Hash

	return block, err
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
