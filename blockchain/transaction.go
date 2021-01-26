package blockchain

///   Transaction   ///

/*
CreateTransaction トランザクションを作成します
*/
func CreateTransaction(senderAddress string, recipientAddress string, value uint64) Transaction {
	transaction := Transaction{
		SenderAddress:    senderAddress,
		RecipientAddress: recipientAddress,
		Value:            value,
	}
	return transaction
}
