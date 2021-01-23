package main

import (
	"github.com/sinsinpurin/GoMyBlockchain/blockchain"
)

func main() {
	MyBlockChainAddress := "My Address"
	BC := blockchain.InitBlockChain(MyBlockChainAddress)

	// block1
	BC.AddTransaction(BC.CreateTransaction("masaki", MyBlockChainAddress, 100000000))
	BC.Mining()

	// BC.PrintChain()

	// block2
	BC.AddTransaction(BC.CreateTransaction(MyBlockChainAddress, "coco", 1000000000))
	BC.AddTransaction(BC.CreateTransaction("masaki", "michi", 100000000))
	BC.Mining()

	BC.PrintChain()
	BC.PrintAddressAmount(MyBlockChainAddress)
}
