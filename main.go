package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	blocks []*Block
}
type Block struct {
	Hash     []byte
	PrevHash []byte
	Data     []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.PrevHash, b.Data}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, prevHash, []byte(data)}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// create the very first block in the chain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// function to init the block chain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x \n", block.PrevHash)
		fmt.Printf("Data in block: %s \n", block.Data)
		fmt.Printf("Hash: %x \n", block.Hash)
	}
}
