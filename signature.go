package gomyblockchain

import (
	"crypto/sha256"
	"encoding/json"
)

/*
GenerateSignature トランザクションの署名を行います
*/
func GenerateSignature(wallet Wallet, transaction Transaction) []byte {
	transactionJSON, _ := json.Marshal(transaction)
	message := sha256.Sum256([]byte(string(transactionJSON)))
	signature := wallet.Sign(message[:])
	return signature
}
