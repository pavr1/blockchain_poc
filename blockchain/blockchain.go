package blockchain

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

// function to init the block chain
func InitBlockChain() *BlockChain {
	opts := badger.DefaultOptions(dbPath)

	//this is being done in the DefaultOptions function
	// //db metadata storage path
	// opts.Dir = dbPath
	// //db value storage path
	// opts.ValueDir = dbPath

	lastHash := []byte{}

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(txn *badger.Txn) error {
		// check if the table is empty. "lh" stands for last hash
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("Genesis created")

			genesis := Genesis()
			err = txn.Set(genesis.Hash, genesis.Serialize())
			if err != nil {
				return err
			}
			err = txn.Set([]byte("lh"), genesis.Hash)
			if err != nil {
				return err
			}

			lastHash = genesis.Hash

			return nil
		} else {
			//if there is data in the db then get the last hash
			item, err := txn.Get([]byte("lh"))
			if err != nil {
				return err
			}
			err = item.Value(func(val []byte) error {
				lastHash = val

				return nil
			})

			return nil
		}

	})

	if err != nil {
		log.Fatal(err)
	}

	chain := BlockChain{lastHash, db}

	return &chain
}

func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			lastHash = val

			return nil
		})

		return err
	})

	newBlock := CreateBlock(data, lastHash)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err = txn.Set(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			return err
		}

		err = txn.Set([]byte("lh"), newBlock.Hash)
		if err != nil {
			return err
		}

		chain.LastHash = newBlock.Hash

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (chain *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{chain.LastHash, chain.Database}
}

func (iter *BlockChainIterator) Next() *Block {
	var block *Block
	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			block = Deserialize(val)

			return nil
		})

		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	iter.CurrentHash = block.PrevHash

	return block
}
