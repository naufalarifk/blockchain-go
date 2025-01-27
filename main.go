package main

import (
	"blockchain-go/block"
	"fmt"
	"strconv"
)



func main() {
	bc := block.NewBlockchain()
	bc.AddBlock("Send 1 BTC to Naufal")
	bc.AddBlock("Send 2 more BTC to Naufal")

	for _,b := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", b.PrevBlockHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		pow := block.NewProofOfWork(b)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}