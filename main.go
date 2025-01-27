package main

import (
	"blockchain-go/block"
	"fmt"
)



func main() {
	bc := block.NewBlockchain()
	bc.AddBlock("Send 1 BTC to Naufal")
	bc.AddBlock("Send 2 more BTC to Naufal")

	for _,block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}