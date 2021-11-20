package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math"
	"time"
)

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
	b := &Blockchain{
		[]Block{},
	}

	b.CreateBlock(1, "0")

	return b
}

func (b *Blockchain) CreateBlock(proof uint, previousHash string) Block {
	block := Block{
		index:         uint(len(b.chain) + 1),
		timestamp:     time.Now(),
		proof:         proof,
		previous_hash: previousHash,
	}

	b.chain = append(b.chain, block)

	return block
}

func (b *Blockchain) GetPreviousBlock() Block {
	return b.chain[len(b.chain)-1]
}

func (b *Blockchain) ProofOfWork(previousProof uint) uint {
	newProof := uint(1)
	checkProof := false

	for !checkProof {
		fmt.Printf("Trying proof %d\n", newProof)
		i := math.Pow(float64(newProof), 2) - math.Pow(float64(previousProof), 2)

		hashOperation := sha256.Sum256([]byte(fmt.Sprintf("%f", i)))

		hash := fmt.Sprintf("%x", hashOperation[:])

		if hash[:4] == "0000" {
			checkProof = true
		} else {
			newProof++
		}
	}
	return newProof
}

func (b *Blockchain) ChainValid() bool {
	previousBlock := b.chain[0]
	blockIdx := 0

	for blockIdx < len(b.chain) {
		block := b.chain[blockIdx]
		if block.previous_hash != hash(previousBlock) {
			return false
		}

		previousProof := previousBlock.proof
		proof := block.proof

		i := math.Pow(float64(proof), 2) - math.Pow(float64(previousProof), 2)

		hashOperation := sha256.Sum256([]byte(fmt.Sprintf("%f", i)))

		hash := fmt.Sprintf("%x", hashOperation[:])

		if hash[:4] == "0000" {
			return false
		}

		previousBlock = block
		blockIdx++
	}
	return true
}

func hash(block Block) string {
	json, err := json.Marshal(block)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", sha256.Sum256(json))
}
