package blockchain

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

// create the very first block in the chain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
