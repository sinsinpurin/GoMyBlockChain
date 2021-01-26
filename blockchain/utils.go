package blockchain

import "log"

/*
BCLogger BlockChainのlogの表示
*/
func BCLogger(functionType string, function string, msg string) {
	log.Printf("[BLOCKCHAIN] %s %s : %s", functionType, function, msg)
}
