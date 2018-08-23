package storage

import (
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

func (m *MongoDB) Ping() error {
	return m.session.Ping()
}

func (m *MongoDB) DB() *mgo.Database {
	return m.db
}

func (m *MongoDB) IsFirstRun() bool {
	var store *models.Store
	err := m.db.C(models.STORE).Find(&bson.M{}).One(&store)
	if err != nil {
		log.Debugf("err: %v", err)
		return true
	}
	log.Debugf("Store: %v", store)
	return false
}
