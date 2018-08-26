package crawler

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Bitterlox/spectrum-crawler-go/models"
	"github.com/Bitterlox/spectrum-crawler-go/rpc"
	"github.com/Bitterlox/spectrum-crawler-go/storage"
	"github.com/Bitterlox/spectrum-crawler-go/util"
)

type Config struct {
	Enabled     bool   `json:"enabled"`
	Interval    string `json:"interval"`
	MaxRoutines int    `json:"routines"`
}

type Crawler struct {
	backend *storage.MongoDB
	rpc     *rpc.RPCClient
	cfg     *Config
}

func New(db *storage.MongoDB, rpc *rpc.RPCClient, cfg *Config) *Crawler {
	return &Crawler{db, rpc, cfg}
}

func (c *Crawler) Start() {
	log.Println("Starting block Crawler")

	if c.backend.IsFirstRun() {
		c.backend.Init()
	}

	interval, err := time.ParseDuration(c.cfg.Interval)
	if err != nil {
		log.Fatalf("Crawler: can't parse duration: %v", err)
	}

	timer := time.NewTimer(interval)

	log.Printf("Block refresh interval: %v", interval)
	go func() {
		for {
			select {
			case <-timer.C:
				go c.SyncLoop()
				//timer.Reset(interval)
			}
		}
	}()

}

func (c *Crawler) SyncLoop() {
	var wg sync.WaitGroup
	var currentBlock uint64
	var routines int

	startBlock, err := c.rpc.LatestBlockNumber()

	if err != nil {
		log.Errorf("Error getting blockNo: %v", err)
	}

	for currentBlock = startBlock; !c.backend.IsPresent(currentBlock); currentBlock-- {

		block, err := c.rpc.GetBlockByHeight(currentBlock)

		if err != nil {
			log.Errorf("Error getting block: %v", err)
		}

		// TODO: try to test this

		if c.backend.IsPresent(currentBlock) && c.backend.IsForkedBlock(currentBlock, block.Hash) {
			go c.SyncForkedBlock(block, &wg)
		} else {
			go c.Sync(block, &wg)
		}

		wg.Add(1)
		routines++

		if routines == c.cfg.MaxRoutines {
			wg.Wait()
			routines = 0
		}
	}
}

func (c *Crawler) SyncForkedBlock(block *models.Block, wg *sync.WaitGroup) {
	defer wg.Done()

	height := block.Number

	dbblock, err := c.backend.GetBlock(height)

	if err != nil {
		log.Errorf("Error getting forked block: %v", err)
	}

	c.backend.AddForkedBlock(dbblock)

	c.backend.UpdateStore(dbblock.BlockReward, block, "1", true)

	c.backend.ReorgPurge(height)

	c.Sync(block, wg)

}

func (c *Crawler) Sync(block *models.Block, wg *sync.WaitGroup) {
	defer wg.Done()

	var avgGasPrice, txFees, uncleRewards, minted uint64

	blockReward := util.CaculateBlockReward(block.Number, len(block.Uncles))

	if len(block.Transactions) > 0 {
		avgGasPrice, txFees = c.ProcessTransactions(block.Transactions, block.Timestamp)
	}

	if len(block.Uncles) > 0 {
		uncleRewards = c.ProcessUncles(block.Uncles, block.Number)
	}

	minted = blockReward + uncleRewards

	block.BlockReward = minted
	block.AvgGasPrice = avgGasPrice
	block.TxFees = txFees
	block.UnclesReward = uncleRewards

	c.backend.UpdateStore(minted, block, "1", false)
	c.backend.AddBlock(block)

	log.Printf("Block (%v): added %v transactions, %v uncles", block.Number, len(block.Transactions), len(block.Uncles))

}

func (c *Crawler) ProcessUncles(uncles []string, height uint64) uint64 {

	var uncleRewards uint64

	for k, _ := range uncles {

		uncle, err := c.rpc.GetUncleByBlockNumberAndIndex(height, k)

		if err != nil {
			log.Errorf("Error getting uncle: %v", err)
			return 0

		}

		// TODO: broken func
		uncleReward := util.CaculateUncleReward(height, uncle.Number)

		uncleRewards += uncleReward

		uncle.BlockNumber = height
		uncle.Reward = uncleReward

		err = c.backend.AddUncle(uncle)

		if err != nil {
			log.Errorf("Error inserting tx into backend: %v", err)
			return 0
		}

	}
	return uncleRewards
}

func (c *Crawler) ProcessTransactions(txs []models.RawTransaction, timestamp uint64) (uint64, uint64) {

	var avgGasPrice, txFees uint64

	for _, v := range txs {

		v := v.Convert()

		receipt, err := c.rpc.GetTxReceipt(v.Hash)

		if err != nil {
			log.Errorf("Error getting tx receipt: %v", err)
		}

		avgGasPrice += v.Gas

		txFees += (v.GasPrice * receipt.GasUsed)

		v.Timestamp = timestamp
		v.GasUsed = receipt.GasUsed
		v.ContractAddress = receipt.ContractAddress
		v.Logs = receipt.Logs

		if v.IsTokenTransfer() {

			tktx := v.GetTokenTransfer()

			tktx.BlockNumber = v.BlockNumber
			tktx.Hash = v.Hash
			tktx.Timestamp = v.Timestamp

			c.backend.AddTokenTransfer(tktx)
		}

		err = c.backend.AddTransaction(v)

		if err != nil {
			log.Errorf("Error inserting tx into backend: %v", err)
		}

	}
	return avgGasPrice / uint64(len(txs)), txFees
}
