package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math"
	"time"
)

type Block struct {
	Index        uint      `json:"index,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
	Proof        uint      `json:"proof,omitempty"`
	PreviousHash string    `json:"previous_hash,omitempty"`
}

type Blockchain struct {
	Chain []Block `json:"chain,omitempty"`
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
		Index:        uint(len(b.Chain) + 1),
		Timestamp:    time.Now(),
		Proof:        proof,
		PreviousHash: previousHash,
	}

	b.Chain = append(b.Chain, block)

	return block
}

func (b *Blockchain) GetPreviousBlock() Block {
	return b.Chain[len(b.Chain)-1]
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
	previousBlock := b.Chain[0]
	blockIdx := 0

	for blockIdx < len(b.Chain) {
		block := b.Chain[blockIdx]
		if block.PreviousHash != hash(previousBlock) {
			return false
		}

		previousProof := previousBlock.Proof
		proof := block.Proof

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
