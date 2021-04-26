package gomyblockchain

//TODO: AddTransactionの所持金がなければ送金できない処理をコメントアウトから外す

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/haltingstate/secp256k1-go"
)

/*
InitBlockChain BlockChainを初期化しBlockChainを返します
*/
func InitBlockChain(blockChainAddress string, port int) *BlockChain {
	BCLogger("(BC *BlockChain)", "InitBlockChain()", "Initialize BlockChain")
	blockchain := new(BlockChain)
	blockchain.Port = port
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

func (BC *BlockChain) setNeighbours() {
	BC.Neighbours = FindNeighbours(LocalHost, BC.Port, NeighboursIPStartRange, NeighboursIPEndRange, BlockChainPortStartRange, BlockChainPortEndRange)
}

/*
SyncNeighbours 近くにいるノードを探索します
*/
func (BC *BlockChain) SyncNeighbours() {
	for {
		BC.setNeighbours()
		time.Sleep(BlockChainNeighboursSyncTymeSec * time.Second)
	}
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

	// neighborsのpoolも削除
	client := &http.Client{}
	for _, node := range BC.Neighbours {
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%s", node)+"/transactions", nil)
		_, err = client.Do(req)
		if err != nil {
			panic(err)
		}
	}
	return block
}

/*
AddTransaction 与えられたトランザクションを検証しトランザクションプールに格納します
*/
func (BC *BlockChain) AddTransaction(transaction Transaction, senderPublicKey []byte, signature []byte) bool {
	// 所有しているbicoinより少ない場合送れない
	// if BC.CalculateTotalAmount(transaction.SenderAddress) < transaction.Value {
	// 	fmt.Println("AddTransaction:Sender Address don't have enough amount.")
	// 	return false
	// }

	result := BC.VerifyTransactionSignature(senderPublicKey, signature, transaction)
	if result != true {
		fmt.Println("VerifyTransactionSignature: faild")
		return false

	}
	BC.TransactionPool = append(BC.TransactionPool, transaction)
	return true
}

/*
AddTransactionSync トランザクションをaddして外部のノードに送信します．AddTransactionをラップしてます
*/
func (BC *BlockChain) AddTransactionSync(transaction Transaction, senderPublicKey []byte, signature []byte) bool {
	isAddTransaction := BC.AddTransaction(transaction, senderPublicKey, signature)
	transactionWithSig := TransactionWithSig{
		RecipientAddress: transaction.RecipientAddress,
		SenderAddress:    transaction.SenderAddress,
		Value:            transaction.Value,
		PublicKey:        hex.EncodeToString(senderPublicKey),
		Signature:        hex.EncodeToString(signature),
	}
	client := &http.Client{}
	if isAddTransaction {
		for _, node := range BC.Neighbours {
			fmt.Println("node: " + node)
			jsonT, _ := json.Marshal(transactionWithSig)
			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%s", node)+"/transactions", bytes.NewBuffer(jsonT))
			if err != nil {
				return false
			}
			req.Header.Set("Content-Type", "application/json")
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			fmt.Println(resp.StatusCode)
		}
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
	// transactionPoolにトランザクションがない場合
	if len(BC.TransactionPool) == 0 {
		return false
	}
	BCLogger("(BC *BlockChain)", "Mining()", "Run Mining")
	BC.addMiningTransaction(CreateTransaction(MiningSender, minerAddress, MiningReward))
	nonce := BC.proofOfWork()
	BC.CreateBlock(nonce)
	BCLogger("(BC *BlockChain)", "Mining()", "Mining Success")
	client := &http.Client{}
	for _, node := range BC.Neighbours {
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://%s", node)+"/consensus", nil)
		if err != nil {
			return false
		}
		_, err = client.Do(req)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	return true
}

/*
BackgroundMining マイニングをバックグラウンドで実行する関数
close(miningStatusCh)すると停止します
*/
func (BC *BlockChain) BackgroundMining(miningStatusCh chan bool, minerAddress string) {
	for {
		select {
		case <-miningStatusCh:
			BCLogger("(BC *BlockChain)", "BackgroundMining()", "Stop")
			return
		default:
			BC.Mining(minerAddress)
			time.Sleep(MiningWaitTimeSec * time.Second)
		}
	}
}

/*
VerifyTransactionSignature トランザクションが有効であるかを検証します
*/
func (BC *BlockChain) VerifyTransactionSignature(senderPublicKey []byte, signature []byte, transaction Transaction) bool {
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

/*
ValidChain chainを検証します
*/
func (BC *BlockChain) ValidChain(chain []Block) bool {
	preBlock := chain[0]
	currentIndex := 1
	for currentIndex < len(chain) {
		block := chain[currentIndex]
		if block.PreHash != BC.blockHash(preBlock) {
			return false
		}
		if !BC.validProof(block.Transactions, block.PreHash, block.Nonce) {
			fmt.Println("validProof false")
			return false
		}
		preBlock = block
		currentIndex++
	}
	return true
}

/*
ResolveConflicts 他ノードとのチェーンを比較します
*/
func (BC *BlockChain) ResolveConflicts() bool {
	longestChain := []Block{}
	maxLength := len(BC.Chain)
	for _, node := range BC.Neighbours {
		res, err := http.Get("http://" + node + "/chain")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		chain := []Block{}
		if err := json.Unmarshal(body, &chain); err != nil {
			log.Fatal(err)
		}
		chainLength := len(chain)
		fmt.Println(BC.ValidChain(chain))
		if chainLength > maxLength && BC.ValidChain(chain) {
			maxLength = chainLength
			longestChain = chain
		}
	}
	if longestChain != nil {
		BC.Chain = longestChain
		return true
	}
	return false
}
