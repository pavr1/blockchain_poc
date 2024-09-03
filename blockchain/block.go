package blockchain

type BlockChain struct {
	Blocks []*Block
}
type Block struct {
	Hash     []byte
	PrevHash []byte
	Data     []byte
	Nounce   int
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, prevHash, []byte(data), 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nounce = nonce

	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// create the very first block in the chain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// function to init the block chain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
