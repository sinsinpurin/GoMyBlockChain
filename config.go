package gomyblockchain

// MiningDifficulty マイニング難易度
const MiningDifficulty = 3

// MiningSender マイニング報酬の送信者
const MiningSender = "THE BROCKCHAIN"

// MiningReward マイニングの報酬額
const MiningReward = 100

// MiningWaitTimeSec バックグラウンドで実行するminingの間隔(secound)
const MiningWaitTimeSec = 15

/// Search Node ///
// Localのtest用の設定
const LocalHost = "192.168.1.2"

// BlockChainPortStartRange 探索するノードのポート番号
const BlockChainPortStartRange = 8085

// BlockChainPortEndRange 探索するノードのポート番号
const BlockChainPortEndRange = 8088

// NeighboursIPStartRange  探索するノードのipの幅
const NeighboursIPStartRange = 0

// NeighboursIPEndRange 探索するノードのipの幅
const NeighboursIPEndRange = 1

// BlockChainNeighboursSyncTymeSec 周辺のノードを探索する間隔
const BlockChainNeighboursSyncTymeSec = 5
