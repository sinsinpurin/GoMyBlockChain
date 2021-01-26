package main

import (
	"github.com/sinsinpurin/GoMyBlockchain/blockchain"
)

func main() {
	MyBlockChainAddress := "test recipient address"
	BC := blockchain.InitBlockChain(MyBlockChainAddress)
	wallet := blockchain.GenerateWallet()
	wallet.PrintWalletInfo()
	// block1
	transaction1 := blockchain.CreateTransaction(wallet.Address, MyBlockChainAddress, 100000000)
	BC.AddTransaction(transaction1, wallet.PublicKey, blockchain.GenerateSignature(wallet, transaction1))
	BC.Mining()

	BC.PrintChain()

	transaction2 := blockchain.CreateTransaction(wallet.Address, MyBlockChainAddress, 50000000)
	BC.AddTransaction(transaction2, wallet.PublicKey, blockchain.GenerateSignature(wallet, transaction2))
	transaction3 := blockchain.CreateTransaction(wallet.Address, MyBlockChainAddress, 600000)
	BC.AddTransaction(transaction3, wallet.PublicKey, blockchain.GenerateSignature(wallet, transaction3))
	BC.Mining()
	BC.PrintChain()

	BC.Mining()
	BC.PrintChain()

}
