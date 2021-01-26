package main

import (
	"testing"

	"github.com/sinsinpurin/GoMyBlockchain/blockchain"
)

/*
TestTransactionVerifyByBlockChain BlockChainが有効なトランザクションの検証のテスト
*/
func TestTransactionVerifyByBlockChain(t *testing.T) {
	MyBlockChainAddress := "My Address"
	BC := blockchain.InitBlockChain(MyBlockChainAddress)
	wallet := blockchain.GenerateWallet()
	transaction := blockchain.CreateTransaction("masaki", MyBlockChainAddress, 1000000)
	transactionSignature := blockchain.GenerateSignature(wallet, transaction)
	if result := BC.VerifyTransactionSignature(wallet.PublicKey, transactionSignature, transaction); result == true {
		// fmt.Println("Blockchain verify transaction")
	} else {
		t.Error("BlockChain deny")
	}
}
