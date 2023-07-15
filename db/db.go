package db

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/dgraph-io/badger/v3"
)

var (
	instance *badger.DB
	once     sync.Once
)

func Instance() *badger.DB {
	once.Do(func() {
		var err error
		instance, err = initDB()
		if err != nil {
			log.Fatalf("Error initializing database: %v", err)
		}
	})
	return instance
}

func initDB() (*badger.DB, error) {
	homeDir, _ := os.UserHomeDir()
	dbPath := filepath.Join(homeDir, ".shelf", "shelf.db")
	opts := badger.DefaultOptions(dbPath)
	opts.Logger = nil
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return db, nil
}
