package config

import (
	"github.com/Bitterlox/spectrum-crawler-go/crawler"
	"github.com/Bitterlox/spectrum-crawler-go/rpc"
	"github.com/Bitterlox/spectrum-crawler-go/storage"
)

type Config struct {
	Threads int            `json:"threads"`
	Crawler crawler.Config `json:"crawler"`
	Mongo   storage.Config `json:"mongo"`
	Rpc     rpc.Config     `json:"rpc"`
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
