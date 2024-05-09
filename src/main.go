package main

import (
	"fmt"
	"strconv"

	"github.com/labstack/gommon/log"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Jonas")
	bc.AddBlock("Send 3 BTC to Jay")

	for _, block := range bc.blocks {
		log.Infof("Prev. hash: %x", block.PrevBlockHash)
		log.Infof("Data: %s", block.Data)
		log.Infof("Hash: %x", block.Hash)

		pow := NewProofOfWork(block)
		log.Infof("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
