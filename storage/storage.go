package storage

import (
	"math/big"
	"time"

	"github.com/Bitterlox/spectrum-crawler-go/models"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Address  string `json:"address"`
}

type MongoDB struct {
	session *mgo.Session
	db      *mgo.Database
}

func NewConnection(cfg *Config) (*MongoDB, error) {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{cfg.Address},
		Database: cfg.Database,
		Username: cfg.User,
		Password: cfg.Password,
	})
	if err != nil {
		return nil, err
	}
	return &MongoDB{session, session.DB("")}, nil
}

func (m *MongoDB) Init() {
	store := &models.Store{
		Timestamp: time.Now().Unix(),
		Symbol:    "UBQ",
		Supply:    "36108073197716300000000000",
	}

	ss := m.db.C(models.STORE)

	if err := ss.Insert(store); err != nil {
		log.Fatalf("Could not init sysStore", err)
	}
	log.Warnf("Initialized sysStore")

}

func (m *MongoDB) Ping() error {
	return m.session.Ping()
}

func (m *MongoDB) DB() *mgo.Database {
	return m.db
}

func (m *MongoDB) IsFirstRun() bool {
	var store models.Store
	err := m.DB().C(models.STORE).Find(&bson.M{}).Limit(1).One(&store)

	if err != nil {
		if err.Error() == "not found" {
			return true
		} else {
			log.Fatalf("Error during initialization: %v", err)
		}
	}

	return false
}

func (m *MongoDB) IsPresent(height uint64) bool {
	var rbn models.RawBlockDetails
	err := m.DB().C(models.BLOCKS).Find(&bson.M{"number": height}).Limit(1).One(&rbn)

	if err != nil {
		if err.Error() == "not found" {
			return false
		} else {
			log.Errorf("Error checking for block in db: %v", err)
		}
	}

	if number, _ := rbn.Convert(); number == height {
		return false
	}

	return true
}

func (m *MongoDB) IsForkedBlock(height uint64, hash string) bool {
	var rbn models.RawBlockDetails
	err := m.DB().C(models.BLOCKS).Find(&bson.M{"number": height}).Limit(1).One(&rbn)

	if err != nil {
		if err.Error() == "not found" {
			return false
		} else {
			log.Errorf("Error checking for block in db: %v", err)
		}
	}

	if bn, contendentHash := rbn.Convert(); bn == height && contendentHash != hash {
		return true
	}

	return false
}

func (m *MongoDB) AddTransaction(tx *models.Transaction) error {
	ss := m.db.C(models.TXNS)

	if err := ss.Insert(tx); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) AddTokenTransfer(tt *models.TokenTransfer) error {
	ss := m.db.C(models.TRANSFERS)

	if err := ss.Insert(tt); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) AddUncle(u *models.Uncle) error {
	ss := m.db.C(models.UNCLES)

	if err := ss.Insert(u); err != nil {
		return err
	}
	return nil
}
func (m *MongoDB) UpdateSupply(minted *big.Int) {

}
func (m *MongoDB) AddBlock(block models.Block) {

}
func (m *MongoDB) AddForkedBlock() {

}
