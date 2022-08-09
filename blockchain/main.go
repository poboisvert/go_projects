package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type blockChain struct {
	blocks[]*Block
}

type Block struct {
	Hash[]byte
	Data[]byte
	PrevHash[]byte
}

func(b *Block) deriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func createBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.deriveHash()
	return block
}

func (chain *blockChain) addBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := createBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func genesis() *Block {
	return createBlock("Genesis", []byte{})
}

func initChain() *blockChain {
	return &blockChain{[]*Block{genesis()}}
}

func main() {
	version := "1.0"
	fmt.Println("Go version: ", version)

	chain := initChain()
	chain.addBlock("init after genesis")
	chain.addBlock("2nd after genesis")
	chain.addBlock("3nd after genesis")
	chain.addBlock("4nd after genesis")
	chain.addBlock("5nd after genesis")

	for _, block := range chain.blocks {
		fmt.Printf("previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data contain: %x\n", block.Data)
	}

}