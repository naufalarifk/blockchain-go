package main

import (
	"blockchain-go/block"
)



func main() {
	bc := block.NewBlockchain()
	defer bc.Db.Close()

	cli := block.CLI{Bc:bc}
	cli.Run()
}