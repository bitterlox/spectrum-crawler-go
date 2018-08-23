package config

import (
	// "github.com/Bitterlox/spectrum-crawler-go/crawler"
	"github.com/Bitterlox/spectrum-crawler-go/storage"
)

type Config struct {
	Crawler struct {
		Enabled bool `json:"enabled"`
	} `json:"crawler"`
	Mongo storage.Config
}

// {
//   "mongodb": {
//     "user": "spectrum",
//     "password": "UBQ4Lyfe",
//     "database": "spectrumdb",
//     "address": "localhost",
//     "port": 27017
//   },
//
//   "mongodbtest": {
//     "user": "spectrum",
//     "password": "UBQ4Lyfe",
//     "database": "spectrum-test",
//     "address": "localhost",
//     "port": 27017
//   }
// }
