package models

import (
	"github.com/Bitterlox/spectrum-crawler-go/util"
)

type RawBlockDetails struct {
	Number string `bson:"number" json:"number"`
	Hash   string `bson:"hash" json:"hash"`
}

func (rbn *RawBlockDetails) Convert() (uint64, string) {
	return util.DecodeHex(rbn.Number), rbn.Hash
}

type RawBlock struct {
	Number          string           `bson:"number" json:"number"`
	Timestamp       string           `bson:"timestamp" json:"timestamp"`
	Transactions    []RawTransaction `bson:"transactions" json:"transactions"`
	Hash            string           `bson:"hash" json:"hash"`
	ParentHash      string           `bson:"parentHash" json:"parentHash"`
	Sha3Uncles      string           `bson:"sha3Uncles" json:"sha3Uncles"`
	Miner           string           `bson:"miner" json:"miner"`
	Difficulty      string           `bson:"difficulty" json:"difficulty"`
	TotalDifficulty string           `bson:"totalDifficulty" json:"totalDifficulty"`
	Size            string           `bson:"size" json:"size"`
	GasUsed         string           `bson:"gasUsed" json:"gasUsed"`
	GasLimit        string           `bson:"gasLimit" json:"gasLimit"`
	Nonce           string           `bson:"nonce" json:"nonce"`
	Uncles          []string         `bson:"uncles" json:"uncles"`
	//
	BlockReward  string `bson:"blockReward" json:"blockReward"`
	UnclesReward string `bson:"unclesReward" json:"unclesReward"`
	AvgGasPrice  string `bson:"avgGasPrice" json:"avgGasPrice"`
	TxFees       string `bson:"txFees" json:"txFees"`
	//
	ExtraData string `bson:"extraData" json:"extraData"`
}

func (b *RawBlock) Convert() *Block {
	//log.Debugf("raw block: %+v", b)
	return &Block{
		Number:          util.DecodeHex(b.Number),
		Timestamp:       util.DecodeHex(b.Timestamp),
		Transactions:    b.Transactions,
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
		Uncles:          b.Uncles,
		// Empty
		BlockReward:  0,
		UnclesReward: 0,
		AvgGasPrice:  0,
		TxFees:       0,
		//
		ExtraData: b.ExtraData,
	}
}

type Block struct {
	Number          uint64           `bson:"number" json:"number"`
	Timestamp       uint64           `bson:"timestamp" json:"timestamp"`
	Transactions    []RawTransaction `bson:"transactions" json:"transactions"`
	Hash            string           `bson:"hash" json:"hash"`
	ParentHash      string           `bson:"parentHash" json:"parentHash"`
	Sha3Uncles      string           `bson:"sha3Uncles" json:"sha3Uncles"`
	Miner           string           `bson:"miner" json:"miner"`
	Difficulty      string           `bson:"difficulty" json:"difficulty"`
	TotalDifficulty string           `bson:"totalDifficulty" json:"totalDifficulty"`
	Size            uint64           `bson:"size" json:"size"`
	GasUsed         uint64           `bson:"gasUsed" json:"gasUsed"`
	GasLimit        uint64           `bson:"gasLimit" json:"gasLimit"`
	Nonce           string           `bson:"nonce" json:"nonce"`
	Uncles          []string         `bson:"uncles" json:"uncles"`
	//
	BlockReward  uint64 `bson:"blockReward" json:"blockReward"`
	UnclesReward uint64 `bson:"unclesReward" json:"unclesReward"`
	AvgGasPrice  uint64 `bson:"avgGasPrice" json:"avgGasPrice"`
	TxFees       uint64 `bson:"txFees" json:"txFees"`
	//
	ExtraData string `bson:"extraData" json:"extraData"`
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
