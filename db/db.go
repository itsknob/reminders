package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dgraph-io/badger"
)

const (
	badgerDiscardRatio = 0.5
	badgerGCInterval   = 10 * time.Minute
)

type (
	DB interface {
		Get(namespace, key []byte) (value []byte, err error)
		Set(namespace, key []byte) error
		Has(namespace, key []byte) (bool, error)
		Close() error
	}
	BadgerDB struct {
		db         *badger.DB
		ctx        context.Context
		cancelFunc context.CancelFunc
		logger     badger.Logger
	}
)

func InitDb(dataDir string, logger badger.Logger) (DB, error) {
	// is that the right permissions?
	if err := os.MkdirAll(dataDir, 0774); err != nil {
		return nil, err
	}

	badgerDB, err := badger.Open(badger.DefaultOptions(dataDir))
	if err != nil {
		return nil, err
	}

	bdb := &BadgerDB{
		db:     badgerDB,
		logger: logger,
	}
	bdb.ctx, bdb.cancelFunc = context.WithCancel(context.Background())

	go bdb.RunGC()
	return bdb, nil
}

func (bdb *BadgerDB) Get(namespace, key []byte) (value []byte, err error) {
	panic("Not Implemented")
}

// todo change function signature to match interface types
func (bdb *BadgerDB) Set(namespace, key []byte) (err error) {
	panic("Not Implemented")
}

func (bdb *BadgerDB) Has(namespace, key []byte) (ok bool, err error) {
	_, err = bdb.Get(namespace, key)
	switch err {
	case badger.ErrKeyNotFound:
		ok, err = false, nil
	case nil:
		ok, err = true, nil
	}
	return //implicit due to named return values
}
func (bdb *BadgerDB) Close() error {
	bdb.cancelFunc()
	return bdb.db.Close()
}

func (bdb *BadgerDB) RunGC() {
	ticker := time.NewTicker(badgerGCInterval)
	for {
		select {
		case <-ticker.C:
			err := bdb.db.RunValueLogGC(badgerDiscardRatio)
			if err != nil {
				if err == badger.ErrNoRewrite {
					bdb.logger.Debugf("no BadgerDB GC occurred: %v", err)
				} else {
					bdb.logger.Errorf("failed to GC BadgerDB: %v", err)
				}
			}
		case <-bdb.ctx.Done():
			return
		}
	}
}

func badgerNamespaceKey(namespace, key []byte) []byte {
	prefix := []byte(fmt.Sprintf("%s/", namespace))
	return append(prefix, key...)
}
