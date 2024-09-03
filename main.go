package main

import "bytes"

type Block struct {
	Hash     []byte
	PrevHash []byte
	Data     []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.PrevHash, b.Data}, []byte{})
}

func main() {

}
