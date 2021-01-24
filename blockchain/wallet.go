package blockchain

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"

	"github.com/haltingstate/secp256k1-go"
	"github.com/shengdoushi/base58"
)

/*
GenerateWallet walletを作成し、walletを返します
*/
func GenerateWallet() Wallet {
	var wallet Wallet
	pubKey, priKey := secp256k1.GenerateKeyPair()
	wallet.PrivateKey = priKey
	wallet.PublicKey = pubKey
	wallet.Address = generateBlockChainAddress(wallet.PublicKey)
	return wallet
}

/*
Sign PrivateKeyでmsgの署名を行い、signatureを返します
*/
func (wallet *Wallet) Sign(msg []byte) []byte {
	signature := secp256k1.Sign(msg, wallet.PrivateKey)
	return signature
}

/*
VerifySignature 公開鍵と署名を使用して署名の有効性を検証します
*/
func VerifySignature(msg []byte, signature []byte, pubKey []byte) bool {
	result := secp256k1.VerifySignature(msg, signature, pubKey)
	if result == 0 {
		return false //Invalid Signature
	}
	return true // Varid Signature
}

func generateBlockChainAddress(publicKey []byte) string {
	// 1. 公開鍵を SHA-256 にかけ その後 RIPEMD-160 をかけて PublicKeyHash を作成
	ripemd := ripemd160.New()
	pubS1 := sha256.Sum256(publicKey)
	ripemd.Write(pubS1[:])
	publicKeyHash := ripemd.Sum(nil)

	// 2. PublicKeyHash の先頭に 0x00 (ネットワークバイト)を付与
	networkBytes := []byte{00}
	publicKeyHashWithNetworkBytes := append(networkBytes, publicKeyHash...)

	// 3. 2 に SHA-256 をかける
	pubS2 := sha256.Sum256(publicKeyHashWithNetworkBytes)
	// 4. 3 に SHA-256 をかける
	pubS3 := sha256.Sum256(pubS2[:])

	// 5. 4 の先頭 4bytes をチェックサムとして切り取る
	checkSum := pubS3[:4]

	// 6. チェックサム を 2 の最後に連結
	publicKeyHashWithCheckSum := append(publicKeyHashWithNetworkBytes, checkSum...)

	// 7. 6 を BASE58 でエンコードして アドレスを生成
	bitcoinAlphabet := base58.BitcoinAlphabet
	walletAddress := base58.Encode(publicKeyHashWithCheckSum, bitcoinAlphabet)
	return walletAddress
}
