package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Bitterlox/spectrum-crawler-go/config"
	// "github.com/Bitterlox/spectrum-crawler-go/crawler"
	"github.com/Bitterlox/spectrum-crawler-go/storage"
)

var cfg config.Config

func init() {

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, TimestampFormat: time.RFC822})

	v, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	if v {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func readConfig(cfg *config.Config) {

	if len(os.Args) == 1 {
		log.Fatalln("Please specify config")
	}

	conf := os.Args[1]
	conf, _ = filepath.Abs(conf)

	log.Printf("Loading config: %v", conf)

	configFile, err := os.Open(conf)
	if err != nil {
		log.Fatal("File error: ", err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&cfg); err != nil {
		log.Fatal("Config error: ", err.Error())
	}
}

func main() {
	readConfig(&cfg)
	rand.Seed(time.Now().UnixNano())

	mongo, err := storage.NewConnection(&cfg.Mongo)

	if err != nil {
		log.Fatalf("Can't establish connection to mongo: %v", err)
	} else {
		log.Fatalf("Successfully connected to mongo")
	}

	err = mongo.Ping()

	if err != nil {
		log.Printf("Can't establish connection to mongo: %v", err)
	} else {
		log.Println("PING")
	}
	//
	// if cfg.Crawler.Enabled {
	// 	go startBlockUnlocker()
	// }

	quit := make(chan bool)
	<-quit
}
