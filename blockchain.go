package main

import "time"

type Block struct {
	index         uint
	timestamp     time.Time
	proof         uint
	previous_hash string
}

type Blockchain struct {
	chain []Block
}

func NewBlockchain() *Blockchain {
	b := &Blockchain{}

	b.CreateBlock(1, "0")

	return b
}

func (b *Blockchain) CreateBlock(proof uint, previous_hash string) {
	block := Block{
		index:         uint(len(b.chain) + 1),
		timestamp:     time.Now(),
		proof:         proof,
		previous_hash: previous_hash,
	}

	b.chain = append(b.chain, block)
}
