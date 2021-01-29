package blockchain

// ToDo: アドレス型を作りたい
// ToDo: Hash型を作りたい

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/haltingstate/secp256k1-go"
)

/*
InitBlockChain BlockChainを初期化しBlockChainを返します
*/
func InitBlockChain(blockChainAddress string) *BlockChain {
	BCLogger("(BC *BlockChain)", "InitBlockChain()", "Initialize BlockChain")
	blockchain := new(BlockChain)
	blockchain.BlockChainAddress = blockChainAddress
	initBlock := Block{
		PreHash:      "init hash",
		Nonce:        0,
		Timestamp:    time.Now(),
		Transactions: nil,
	}
	blockchain.Chain = append(blockchain.Chain, initBlock)
	return blockchain
}

/*
CreateBlock ブロックを作成します
*/
func (BC *BlockChain) CreateBlock(nonce uint) Block {
	preBlock := BC.Chain[len(BC.Chain)-1]
	preHash := BC.blockHash(preBlock)
	block := Block{
		PreHash:      preHash,
		Nonce:        nonce,
		Timestamp:    time.Now(),
		Transactions: BC.TransactionPool,
	}
	BC.Chain = append(BC.Chain, block)
	BC.TransactionPool = nil
	return block
}

/*
AddTransaction 与えられたトランザクションを検証しトランザクションプールに格納します
*/
func (BC *BlockChain) AddTransaction(transaction Transaction, senderPublicKey []byte, signature []byte) bool {
	if BC.CalculateTotalAmount(transaction.SenderAddress) < transaction.Value {
		log.Fatalln("AddTransaction:Sender Address don't have enough amount.")
		return false
	}
	if result := BC.VerifyTransactionSignature(senderPublicKey, signature, transaction); result == true {
		BC.TransactionPool = append(BC.TransactionPool, transaction)
		return true
	}
	return false
}

/*
addMiningTransaction マイニング報酬のトランザクションをトランザクションプールに格納します
*/
func (BC *BlockChain) addMiningTransaction(transaction Transaction) bool {
	BC.TransactionPool = append(BC.TransactionPool, transaction)
	return true
}

/*
hash ブロックのhash(sha256)を出力
*/
func (BC *BlockChain) blockHash(block Block) string {
	blockHashByte, _ := json.Marshal(block)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(string(blockHashByte))))
}

func (BC *BlockChain) validProof(transactions []Transaction, preHash string, nonce uint) bool {
	guessBlock := Block{
		PreHash:      preHash,
		Nonce:        nonce,
		Timestamp:    time.Now(),
		Transactions: transactions,
	}
	guessHash := BC.blockHash(guessBlock)
	return guessHash[:MiningDifficulty] == strings.Repeat("0", MiningDifficulty)
}

func (BC *BlockChain) proofOfWork() uint {
	transactions := make([]Transaction, len(BC.TransactionPool))
	copy(transactions, BC.TransactionPool)
	preHash := BC.blockHash(BC.Chain[len(BC.Chain)-1])
	nonce := uint(0)
	for BC.validProof(transactions, preHash, nonce) == false {
		nonce++
	}
	return nonce
}

/*
Mining マイニングを行います
*/
func (BC *BlockChain) Mining(minerAddress string) bool {
	BCLogger("(BC *BlockChain)", "Mining()", "Run Mining")
	BC.addMiningTransaction(CreateTransaction(MiningSender, minerAddress, MiningReward))
	nonce := BC.proofOfWork()
	BC.CreateBlock(nonce)
	return true
}

/*
VerifyTransactionSignature トランザクションが有効であるかを検証します
*/
func (BC *BlockChain) VerifyTransactionSignature(senderPublicKey []byte, signature []byte, transaction Transaction) bool {
	BCLogger("(BC *BlockChain)", "VerifyTransactionSignature", "Run Verify Transaction by BlockChain")
	transactionJSON, _ := json.Marshal(transaction)
	message := sha256.Sum256([]byte(string(transactionJSON)))
	result := verifySignature(message[:], signature, senderPublicKey)
	return result
}

/*
verifySignature 公開鍵と署名を使用して署名の有効性を検証します
*/
func verifySignature(msg []byte, signature []byte, pubKey []byte) bool {
	result := secp256k1.VerifySignature(msg, signature, pubKey)
	if result == 0 {
		return false //Invalid Signature
	}
	return true // Varid Signature
}

/*
CalculateTotalAmount アドレスの保有量を返します
*/
func (BC *BlockChain) CalculateTotalAmount(blockChainAddress string) uint64 {
	var totalAmount uint64
	for _, block := range BC.Chain {
		for _, transaction := range block.Transactions {
			value := transaction.Value
			if blockChainAddress == transaction.RecipientAddress {
				totalAmount += value
			}
			if blockChainAddress == transaction.SenderAddress {
				totalAmount -= value
			}
		}
	}
	return totalAmount
}
