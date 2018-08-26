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

	genesis := &models.Block{
		Number:          0,
		Timestamp:       1485633600,
		Transactions:    nil,
		Hash:            "0x406f1b7dd39fca54d8c702141851ed8b755463ab5b560e6f19b963b4047418af",
		ParentHash:      "0x0000000000000000000000000000000000000000000000000000000000000000",
		Sha3Uncles:      "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
		Miner:           "0x3333333333333333333333333333333333333333",
		Difficulty:      "80000000000",
		TotalDifficulty: "80000000000",
		Size:            524,
		GasUsed:         0,
		GasLimit:        134217728,
		Nonce:           "0x0000000000000888",
		Uncles:          nil,
		// Empty
		BlockReward:  0,
		UnclesReward: 0,
		AvgGasPrice:  0,
		TxFees:       0,
		//
		ExtraData: "0x4a756d6275636b734545",
	}

	gb := m.db.C(models.BLOCKS)

	if err := gb.Insert(genesis); err != nil {
		log.Fatalf("Could not init genesis block", err)
	}

	log.Warnf("Initialized sysStore, genesis")

}

func (m *MongoDB) UpdateStore(supply uint64, head *models.Block, price string) error {

	x := big.NewInt(0)
	x.SetUint64(supply)

	new_supply, err := m.GetSupply()

	if err != nil {
		return err
	}

	new_supply.Add(new_supply, x)

	err = m.DB().C(models.STORE).Update(&bson.M{}, &bson.M{"supply": new_supply.String(), "price": price, "head": head})

	if err != nil {
		return err
	}

	return nil
}

func (m *MongoDB) GetSupply() (*big.Int, error) {
	var store models.Store

	err := m.DB().C(models.STORE).Find(&bson.M{}).Limit(1).One(&store)

	if err != nil {
		return big.NewInt(0), err
	}

	x := big.NewInt(0)

	x.SetString(store.Supply, 10)

	return x, nil
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

	if height == 0 {
		return true
	}

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

func (m *MongoDB) AddBlock(b *models.Block) error {
	ss := m.db.C(models.BLOCKS)

	if err := ss.Insert(b); err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) UpdateSupply(minted *big.Int) {

}

func (m *MongoDB) AddForkedBlock() {

}
