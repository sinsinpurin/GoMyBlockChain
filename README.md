# GoMyBlockchain

Go Document

```zsh
% go doc -http=:8080
```

# Wallet の作成

## 秘密鍵と公開鍵の生成

1. 秘密鍵(32bytes)をランダムに生成
2. 秘密鍵から secp256k1 を使用して公開鍵を生成

## アドレスの生成

Bitcoin を採用

1. 公開鍵を SHA-256 にかけ その後 RIPEMD-160 をかけて PublicKeyHash を作成
2. PublicKeyHash の先頭に 0x00 (ネットワークバイト)を付与
3. 2 に SHA-256 をかける
4. 3 に SHA-256 をかける
5. 4 の先頭 4bytes をチェックサムとして切り取る
6. チェックサム を 2 の最後に連結
7. 6 を BASE58 でエンコードして アドレスを生成

# 参考文献

[ビットコインウォレットを Javascript で作ってみよう](https://note.com/strictlyes/n/n5432a4c5bd36)
