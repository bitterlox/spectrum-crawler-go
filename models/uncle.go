package models

import (
	"math/big"

	"github.com/Bitterlox/spectrum-crawler-go/util"
)

type RawUncle struct {
	Number      string `bson:"number" json:"number"`
	Position    string `bson:"position" json:"position"`
	BlockNumber string `bson:"blockNumber" json:"blockNumber"`
	Hash        string `bson:"hash" json:"hash"`
	ParentHash  string `bson:"parentHash" json:"parentHash"`
	Sha3Uncles  string `bson:"sha3Uncles" json:"sha3Uncles"`
	Miner       string `bson:"miner" json:"miner"`
	Difficulty  string `bson:"difficulty" json:"difficulty"`
	GasUsed     string `bson:"gasUsed" json:"gasUsed"`
	GasLimit    string `bson:"gasLimit" json:"gasLimit"`
	Timestamp   string `bson:"timestamp" json:"timestamp"`
	Reward      string `bson:"reward" json:"reward"`
}

func (rw *RawUncle) Convert() *Uncle {
	return &Uncle{
		Number:     util.DecodeHex(rw.Number).Int64(),
		Position:   util.DecodeHex(rw.Position).Int64(),
		Hash:       rw.Hash,
		ParentHash: rw.ParentHash,
		Sha3Uncles: rw.Sha3Uncles,
		Miner:      rw.Miner,
		Difficulty: rw.Difficulty,
		GasUsed:    util.DecodeHex(rw.GasUsed),
		GasLimit:   util.DecodeHex(rw.GasLimit),
		Timestamp:  util.DecodeHex(rw.Timestamp).Int64(),
	}
}

type Uncle struct {
	Number      int64    `bson:"number" json:"number"`
	Position    int64    `bson:"position" json:"position"`
	BlockNumber int64    `bson:"blockNumber" json:"blockNumber"`
	Hash        string   `bson:"hash" json:"hash"`
	ParentHash  string   `bson:"parentHash" json:"parentHash"`
	Sha3Uncles  string   `bson:"sha3Uncles" json:"sha3Uncles"`
	Miner       string   `bson:"miner" json:"miner"`
	Difficulty  string   `bson:"difficulty" json:"difficulty"`
	GasUsed     *big.Int `bson:"gasUsed" json:"gasUsed"`
	GasLimit    *big.Int `bson:"gasLimit" json:"gasLimit"`
	Timestamp   int64    `bson:"timestamp" json:"timestamp"`
	Reward      *big.Int `bson:"reward" json:"reward"`
}
