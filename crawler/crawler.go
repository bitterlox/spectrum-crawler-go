package crawler

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Bitterlox/spectrum-crawler-go/models"
	"github.com/Bitterlox/spectrum-crawler-go/storage"
)

type Crawler struct {
	backend *storage.MongoDB
}

func New(db *storage.MongoDB) *Crawler {
	return &Crawler{db}
}

func (c *Crawler) Start() {
	log.Println("Starting block Crawler")

	if c.backend.IsFirstRun() {
		c.Init()
	}

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
