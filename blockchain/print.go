package blockchain

import (
	"fmt"
	"strings"
)

/*
PrintChain Chainに格納されているブロックを表示します
*/
func (BC BlockChain) PrintChain() {
	for chainIndex, block := range BC.Chain {
		fmt.Printf("%s Chain %d %s \n", strings.Repeat("=", 40), chainIndex, strings.Repeat("=", 40))
		fmt.Printf("Previous hash: %s \n", block.PreHash)
		fmt.Printf("Nonce: %d \n", block.Nonce)
		fmt.Printf("Time Stamp: %s \n", block.Timestamp)
		for _, transaction := range block.Transactions {
			fmt.Printf("%s \n", strings.Repeat("-", 89))
			fmt.Printf("SenderAddress   : %s \n", transaction.SenderAddress)
			fmt.Printf("RecipientAddress: %s \n", transaction.RecipientAddress)
			fmt.Printf("Value           : %d \n", transaction.Value)
		}
		fmt.Printf("%s \n \n \n", strings.Repeat("=", 89))
	}
}

/*
PrintTransactionPool TransactionPoolに格納されているトランザクションを表示します
*/
func (BC *BlockChain) PrintTransactionPool() {
	for transactionPoolIndex, transaction := range BC.TransactionPool {
		fmt.Printf("%s Transaction Pool %d %s \n", strings.Repeat("=", 40), transactionPoolIndex, strings.Repeat("=", 40))
		fmt.Printf("Sender Address: %s \n", transaction.SenderAddress)
		fmt.Printf("Recipient Address: %s \n", transaction.RecipientAddress)
		fmt.Printf("Value:  %d satoshi \n", transaction.Value)
	}
}

/*
PrintAddressAmount アドレスの合計値を表示します
*/
func (BC *BlockChain) PrintAddressAmount(blockChainAddress string) {
	fmt.Printf("[%s]: %d satoshi \n", blockChainAddress, BC.CalculateTotalAmount(blockChainAddress))
}

/*
PrintWalletInfo Walletの情報を表示します
*/
func (wallet *Wallet) PrintWalletInfo() {
	fmt.Printf("%s Wallet Info %s \n", strings.Repeat("=", 40), strings.Repeat("=", 40))
	fmt.Printf("PriKey: %x \n", wallet.PrivateKey)
	fmt.Printf("PubKey: %x \n", wallet.PublicKey)
	fmt.Printf("Address: %s \n", wallet.Address)
	fmt.Printf("%s \n \n \n", strings.Repeat("=", 93))
}
