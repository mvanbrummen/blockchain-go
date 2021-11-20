package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	blockchain := NewBlockchain()

	r := gin.Default()

	r.GET("/mine_block", func(c *gin.Context) {
		previousBlock := blockchain.GetPreviousBlock()
		previousProof := previousBlock.proof
		proof := blockchain.ProofOfWork(previousProof)
		previousHash := hash(previousBlock)
		block := blockchain.CreateBlock(proof, previousHash)

		c.JSON(200, gin.H{
			"message":       "A block is mined",
			"index":         block.index,
			"timestamp":     block.timestamp,
			"proof":         block.proof,
			"previous_hash": block.previous_hash,
		})
	})

	r.GET("/get_chain", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"chain":  blockchain.chain,
			"length": len(blockchain.chain),
		})
	})

	r.GET("/valid", func(c *gin.Context) {
		valid := blockchain.ChainValid()

		c.JSON(200, gin.H{
			"valid": valid,
		})
	})

	r.Run(":8080")
}
