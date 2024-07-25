package db

import (
	"sync"
	"time"
)

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

type Record struct {
	key       string
	value     string
	createdAt int
}

type DB struct {
	mu              sync.RWMutex
	dbSize          int
	arrayOfPointers [][]Record
}

func NewDB() DB {
	return DB{
		mu:              sync.RWMutex{},
		dbSize:          100,
		arrayOfPointers: [][]Record{},
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
	db.arrayOfPointers[ind] = append(db.arrayOfPointers[ind], Record{
		key:       key,
		value:     value,
		createdAt: int(time.Now().Unix()),
	})
	return nil
}

// TODO: добавить передачу по ссылке переменной для записи value
func (db *DB) Get(key string) (string, error) {
	switch {
	case (len(key) < 1):
		return "", ErrInvalidKey
	}
	ind, err := db.Hash(key)
	if err != nil {
		return "", ErrInvalidKey
	}
	for _, rec := range db.arrayOfPointers[ind] {
		if rec.key == key {
			return rec.value, nil
		}
	}
	return "", nil
}

func (db *DB) Delete(key string) error {
	switch {
	case (len(key) < 1):
		return ErrInvalidKey
	}
	ind, err := db.Hash(key)
	if err != nil {
		return ErrInvalidKey
	}
	for _, rec := range db.arrayOfPointers[ind] {
		if rec.key == key {
			rec = Record{}
		}
	}
	return nil
}
