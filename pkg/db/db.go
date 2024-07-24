package db

import "sync"

type Store interface {
	Put(string, string) error
	Get(string) (string, error)
	Delete(string) error
	Lock() error
	Unlock() error
	RLock() error
	RUnlock() error
}

type DB struct {
	mu              sync.RWMutex
	dbSize          int
	arrayOfPointers [][]string
}

func NewDB() DB {
	return DB{
		mu:              sync.RWMutex{},
		dbSize:          100,
		arrayOfPointers: [][]string{},
	}
}

func (db *DB) Hash(key string) int {
	var sum int = 0
	for _, j := range []byte(key) {
		sum += int(j)
	}
	return sum % db.dbSize
}
