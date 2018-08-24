package crawler

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Bitterlox/spectrum-crawler-go/models"
	"github.com/Bitterlox/spectrum-crawler-go/rpc"
	"github.com/Bitterlox/spectrum-crawler-go/storage"
)

type Config struct {
	Enabled     bool   `json:"enabled"`
	Interval    string `json:"interval"`
	maxRoutines int    `json:"routines"`
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
		c.Init()
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
				timer.Reset(interval)
			}
		}
	}()

}

func (c *Crawler) SyncLoop() {
	var wg sync.WaitGroup
	var startBlock int64

	block, err := c.rpc.GetLatestBlock()

	if err != nil {
		log.Errorf("Error getting blockno: %v", err)
	}

	wg.Add(c.cfg.maxRoutines)

	for startBlock = block.Number; c.backend.IsPresent(startBlock); startBlock-- {
		go c.Sync(block, &wg)
	}

	wg.Wait()
}

func (c *Crawler) Sync(block *models.Block, wg *sync.WaitGroup) {

	log.Printf("block: %+v", block.Number)

	time.Sleep(30 * time.Second)

	wg.Done()
}

func (c *Crawler) Init() {
	store := &models.Store{
		Timestamp: time.Now().Unix(),
		Symbol:    "UBQ",
		Supply:    "36108073197716300000000000",
	}

	ss := c.backend.DB().C(models.STORE)

	if err := ss.Insert(store); err != nil {
		log.Fatalf("Could not init sysStore", err)
	}
	log.Warnf("Initialized systore")

}
