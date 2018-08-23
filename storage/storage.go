package storage

import (
	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
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

func (m *MongoDB) Close() {
	log.Debugln("Bye bye")
	m.session.Close()
}
