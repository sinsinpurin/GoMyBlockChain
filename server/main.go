package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sinsinpurin/GoMyBlockchain/blockchain"
)

// Port ポート番号の設定
var Port string

var blockChainCache blockchain.BlockChainServer

func main() {
	flag.StringVar(&Port, "p", ":8080", "Set Port Number")
	flag.Parse()
	fmt.Println(Port)
	e := echo.New()
	e.GET("/chain", getChain)
	e.Logger.Fatal(e.Start(Port))

}

func getChain(c echo.Context) error {
	bc := getBlockChain()
	return c.JSON(http.StatusOK, bc.Chain)
}

func getBlockChain() blockchain.BlockChain {
	if blockChainCache.BlockChain.BlockChainAddress == "" {
		blockChainCache.Wallet = blockchain.GenerateWallet()
		blockChainCache.BlockChain = *blockchain.InitBlockChain(blockChainCache.Wallet.Address)
	}
	return blockChainCache.BlockChain
}
