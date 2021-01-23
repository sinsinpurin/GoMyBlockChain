package blockchain

///   Transaction   ///

/*
AddTransaction 与えられたトランザクションをトランザクションプールに格納します
*/
func (BC *BlockChain) AddTransaction(transaction Transaction) bool {
	BC.TransactionPool = append(BC.TransactionPool, transaction)
	return true
}

/*
CreateTransaction トランザクションを作成します
*/
func (BC *BlockChain) CreateTransaction(senderAddress string, recipientAddress string, value uint64) Transaction {
	transaction := Transaction{
		SenderAddress:    senderAddress,
		RecipientAddress: recipientAddress,
		Value:            value,
	}
	return transaction
}
