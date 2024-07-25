package db

import "sync"

type Store interface {
	Put(string, string) error
	Get(string) (string, error)
	Delete(string) error
	Hash(string) (int, error)
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

func (db *DB) Hash(key string) (int, error) {
	switch {
	case (len(key) < 1):
		return -1, ErrInvalidKey
	case (len(key) > 100):
		return -1, ErrKeyTooLong
	}
	var sum int = 0
	for _, j := range []byte(key) {
		sum += int(j)
	}
	return sum % db.dbSize, nil
}

func (db *DB) Lock() {
	db.mu.Lock()
}

func (db *DB) Unlock() {
	db.mu.Unlock()
}

func (db *DB) RLock() {
	db.mu.RLock()
}

func (db *DB) RUnlock() {
	db.mu.RUnlock()
}

func (db *DB) Put(key string, value string) error {
	switch {
	case (len(key) < 1):
		return ErrInvalidKey
	case (len(value) < 1):
		return ErrInvalidValue
	}
	ind, err := db.Hash(key)
	if err != nil {
		return ErrInvalidKey
	}
	db.arrayOfPointers[ind][len(db.arrayOfPointers)] = value
	return nil
}
