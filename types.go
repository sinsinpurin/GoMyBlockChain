package gomyblockchain

import "time"

/*
BlockChainServer の構成
*/
type BlockChainServer struct {
	BlockChain BlockChain `json:"BlockChain"`
	Wallet     Wallet     `json:"Wallet"`
}

/*
BlockChain の構成
TransactionPool : []Transaction ( マイニングされていないトランザクションを一時的に貯める )
Chain           : Block         ( 生成されたブロックBlock )
BlockChainAddress: string       ( ブロックチェーンを所有するアドレス )
*/
type BlockChain struct {
	TransactionPool   []Transaction `json:"TransactionPool"`
	Chain             []Block       `json:"Chain"`
	BlockChainAddress string        `json:"BlockChainAddress"`
	Port              int           `json:"Port"`
	Neighbours        []string      `json:"Neighbours"`
}

/*
Block の構成
PreHash        :  Hash?    ( 以前のブロックのハッシュ)
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

/*
TransactionWithSig の構成
*/
type TransactionWithSig struct {
	RecipientAddress string `json:"RecipientAddress"`
	SenderAddress    string `json:"SenderAddress"`
	Value            uint64 `json:"Value"`
	PublicKey        string `json:"SenderPublicKey"`
	Signature        string `json:"Signature"`
}

/*
Wallet の構成
*/
type Wallet struct {
	PrivateKey []byte `json:"PrivateKey"`
	PublicKey  []byte `json:"PublicKey"`
	Address    string `json:"Address"`
}

type ResChain struct {
	Chain []Block
}
