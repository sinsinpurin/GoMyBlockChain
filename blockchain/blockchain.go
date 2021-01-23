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
)

// MiningDifficulty マイニング難易度
const MiningDifficulty = 3

// MiningSender マイニング報酬の送信者
const MiningSender = "THE BROCKCHAIN"

// MiningReward マイニングの報酬額
const MiningReward = 1.0

/*
InitBlockChain BlockChainを初期化しBlockChainを返します
*/
func InitBlockChain(blockChainAddress string) *BlockChain {
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
func (BC *BlockChain) Mining() bool {
	BC.AddTransaction(BC.CreateTransaction(MiningSender, BC.BlockChainAddress, MiningReward))
	nonce := BC.proofOfWork()
	BC.CreateBlock(nonce)
	log.Println("(BC *BlockChain) Mining(): Success")
	return true
}

/*
CalculateTotalAmount
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
