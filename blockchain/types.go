package blockchain

import "time"

/*
BlockChain の構成
TransactionPool : []Transaction ( マイニングされていないトランザクションを一時的に貯める )
Chain           : Block         ( 生成されたブロックBlock )
*/
type BlockChain struct {
	TransactionPool   []Transaction `json:"TransactionPool"`
	Chain             []Block       `json:"Chain"`
	BlockChainAddress string        `json:"BlockChainAddress"`
}

/*
Block の構成
PreHash  :　Hash?          ( 以前のブロックのハッシュ)
Nonce          :  int      ( nonceの値 )
Timestamp      :  Time     ( 作成時刻 )
Transactions   :  Transaction[]
*/
type Block struct {
	PreHash      string        `json:"PreHash"`
	Nonce        uint          `json:"Nonce"`
	Timestamp    time.Time     `json:"Timestamp"`
	Transactions []Transaction `json:"Transactions"`
}

/*
Transaction の構成

*/
type Transaction struct {
	RecipientAddress string `json:"RecipientAddress"`
	SenderAddress    string `json:"SenderAddress"`
	Value            uint64 `json:"Value"`
}
