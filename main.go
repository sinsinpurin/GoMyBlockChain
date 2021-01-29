package main

import (
	"github.com/sinsinpurin/GoMyBlockchain/blockchain"
)

func main() {
	MyBlockChainAddress := "test recipient address"
	BC := blockchain.InitBlockChain(MyBlockChainAddress)
	wallet := blockchain.GenerateWallet()
	BC.PrintWalletInfo(wallet)

	// block1
	BC.Mining(wallet.Address)

	transaction2 := blockchain.CreateTransaction(wallet.Address, MyBlockChainAddress, 1)
	BC.AddTransaction(transaction2, wallet.PublicKey, blockchain.GenerateSignature(wallet, transaction2))
	BC.Mining(wallet.Address)
	BC.PrintChain()
	BC.PrintWalletInfo(wallet)
}
