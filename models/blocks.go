package models

import (
	"github.com/Bitterlox/spectrum-crawler-go/util"
	log "github.com/sirupsen/logrus"
)

type RawBlock struct {
	Number          string `bson:"number" json:"number"`
	Timestamp       string `bson:"timestamp" json:"timestamp"`
	Transactions    string `bson:"transactions" json:"transactions"`
	Hash            string `bson:"hash" json:"hash"`
	ParentHash      string `bson:"parentHash" json:"parentHash"`
	Sha3Uncles      string `bson:"sha3Uncles" json:"sha3Uncles"`
	Miner           string `bson:"miner" json:"miner"`
	Difficulty      string `bson:"difficulty" json:"difficulty"`
	TotalDifficulty string `bson:"totalDifficulty" json:"totalDifficulty"`
	Size            string `bson:"size" json:"size"`
	GasUsed         string `bson:"gasUsed" json:"gasUsed"`
	GasLimit        string `bson:"gasLimit" json:"gasLimit"`
	Nonce           string `bson:"nonce" json:"nonce"`
	Uncles          string `bson:"uncles" json:"uncles"`
	BlockReward     string `bson:"blockReward" json:"blockReward"`
	UnclesReward    string `bson:"unclesReward" json:"unclesReward"`
	AvgGasPrice     string `bson:"avgGasPrice" json:"avgGasPrice"`
	TxFees          string `bson:"txFees" json:"txFees"`
	ExtraData       string `bson:"extraData" json:"extraData"`
}

type Block struct {
	Number          int64  `bson:"number" json:"number"`
	Timestamp       int64  `bson:"timestamp" json:"timestamp"`
	Transactions    int64  `bson:"transactions" json:"transactions"`
	Hash            string `bson:"hash" json:"hash"`
	ParentHash      string `bson:"parentHash" json:"parentHash"`
	Sha3Uncles      string `bson:"sha3Uncles" json:"sha3Uncles"`
	Miner           string `bson:"miner" json:"miner"`
	Difficulty      string `bson:"difficulty" json:"difficulty"`
	TotalDifficulty string `bson:"totalDifficulty" json:"totalDifficulty"`
	Size            int64  `bson:"size" json:"size"`
	GasUsed         int64  `bson:"gasUsed" json:"gasUsed"`
	GasLimit        int64  `bson:"gasLimit" json:"gasLimit"`
	Nonce           string `bson:"nonce" json:"nonce"`
	Uncles          int64  `bson:"uncles" json:"uncles"`
	BlockReward     string `bson:"blockReward" json:"blockReward"`
	UnclesReward    string `bson:"unclesReward" json:"unclesReward"`
	AvgGasPrice     string `bson:"avgGasPrice" json:"avgGasPrice"`
	TxFees          string `bson:"txFees" json:"txFees"`
	ExtraData       string `bson:"extraData" json:"extraData"`
}

func (b *RawBlock) Convert() *Block {
	log.Debugf("raw block: %+v", b)
	height := util.DecodeHex(b.Number)
	return &Block{
		Number:          height,
		Timestamp:       util.DecodeHex(b.Timestamp),
		Transactions:    util.DecodeHex(b.Transactions),
		Hash:            b.Hash,
		ParentHash:      b.ParentHash,
		Sha3Uncles:      b.Sha3Uncles,
		Miner:           b.Miner,
		Difficulty:      b.Difficulty,
		TotalDifficulty: b.TotalDifficulty,
		Size:            util.DecodeHex(b.Size),
		GasUsed:         util.DecodeHex(b.GasUsed),
		GasLimit:        util.DecodeHex(b.GasLimit),
		Nonce:           b.Nonce,
		Uncles:          util.DecodeHex(b.Uncles),
		BlockReward:     util.BaseBlockReward(height).String(),
		// Fix these
		UnclesReward: b.UnclesReward,
		AvgGasPrice:  b.AvgGasPrice,
		TxFees:       b.TxFees,
		ExtraData:    b.ExtraData,
	}
}

// totalDifficulty:
// 	0x9423e4f040bcf071
// gasUsed:
// 	0x0
// number:
// 	0x8b0c7
// timestamp:
// 	0x5b7ef868
// difficulty:
// 	0xee9d8d30d76
// size:
// 	0x223
// gasLimit:
// 	0x79d60c
